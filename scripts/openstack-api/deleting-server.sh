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

server_ip=$(curl -H "X-Auth-Project-Id: Kubernetes" -H "$header" -X GET ${NOVA}/servers | jq .servers[0].id)

#removing double quotes for further reasons
temp="${server_ip%\"}"
temp="${temp#\"}"

curl -H "X-Auth-Project-Id: Kubernetes" -H "$header" -X DELETE ${NOVA}/servers/"$temp"


