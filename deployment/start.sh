#!/bin/bash

PROXY_ENDPOINT=$(terraform output -state=/home/thehuge/projects/profile-db/deployment/proxy/terraform.tfstate -raw ec2_ip)

echo "PROXY_ENDPOINT"
echo $PROXY_ENDPOINT


echo "DB_ENDPOINT"
echo $1

terraform apply -var="redis_proxy_key_dest=/home/thehuge/projects/profile-db/deployment/proxy/.ssh/id_rsa" -var="proxy_access=ec2-user@$PROXY_ENDPOINT" -var="redis_endpoint=$1"
