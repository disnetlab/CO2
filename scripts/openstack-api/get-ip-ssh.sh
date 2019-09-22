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

slave_ip=$(curl -H "X-Auth-Project-Id: Kubernetes" -H "$header" -X GET ${NOVA}/servers/detail | jq .servers[0].addresses.private_network[0].addr)
temp="${slave_ip%\"}"
temp="${temp#\"}"
echo "$temp"

sudo ssh ubuntu@"$temp" -i /home/ubuntu/private-key
