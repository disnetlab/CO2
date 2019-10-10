#! /bin/bash

export NOVA=http://192.168.0.1:8774/v2.1/e8475c6dc63f4587af4603a23d7bf624 
export OS_TOKEN=$(curl -i   -H "Content-Type: application/json"   -d '
{ "auth": {
    "identity": {
      "methods": ["password"],
      "password": {
        "user": {
          "name": "mkahani",
          "domain": { "id": "default" },
          "password": "kirwu4-xYqwit-cizqic",
          "tenantName": "Kubernetes"
        }
      }
    },
    "scope": {
      "project": {
        "name": "Kubernetes",
        "domain": { "id": "default" }
      }
    }
  }
}'   http://192.168.0.1:5000/v3/auth/tokens |  grep X-Subject-Token: | sed -e "s/X-Subject-Token: //")

header='X-Auth-Token: '$OS_TOKEN

<<COMMENT1
echo "IMAGES"
curl -H "X-Auth-Project-Id: Kubernetes" -H "$header" -X GET ${NOVA}/images | tail -n 1 | python -m json.tool
echo "FLAVORS"
curl -H "X-Auth-Project-Id: Kubernetes" -H "$header" -X GET ${NOVA}/flavors | python -m json.tool
echo "SERVERS"
curl -H "X-Auth-Project-Id: Kubernetes" -H "$header" -X GET ${NOVA}/servers | python -m json.tool
echo "NETWORKS
curl -H "Content-Type: application/json" -H "$header" -X GET http://192.168.0.1:9696/v2.0/networks | python -m json.tool
COMMENT1


echo "CREATING VMS"


#creating kubeadm token

touch start.sh
token=$(kubeadm token create --print-join-command)
printf "#! /bin/bash\n\n" >> start.sh
echo $token >> start.sh
token=$(base64 start.sh)
rm start.sh
token=$(echo "${token}" | tr -d '[:space:]')



#creating random string to name worker-nodes

random_string=$(sudo head /dev/urandom | tr -dc A-Za-z0-9 | head -c 13 ; echo '')
worker_name="worker-""$random_string"
echo $worker_name

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -H "X-Auth-Project-Id: Kubernetes" -H "$header" -X POST ${NOVA}/servers -d '{"server": {"OS-DCF:diskConfig": "AUTO", "name": "'"$worker_name"'", "imageRef": "3d00ca1a-1bd1-4a24-80df-538d514ae98d", "availability_zone": "nova", "flavorRef": "2", "max_count": 1, "min_count": 1, "key_name": "mohammad",
 "user_data" : "'"$token"'"
  , "networks": [{"uuid": "976655bc-34bc-4fc8-a76b-63f21243ae9a"}]}}'
