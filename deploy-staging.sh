#!/bin/bash

#fool's man CI/CD
pwd

echo "LOCAL: Build and save image"
docker build . -f Dockerfile-staging --build-arg enviro=staging  -t api-ecommerce-golang
#docker build . -t  api-ecommerce-golang

# menjadikan file ke tar
docker save -o image.tar  api-ecommerce-golang

echo "LOCAL: Copy image over scp"

if ! scp -P 30004 image.tar dev@codethusiast.my.id:
then
  echo "scp failed"
  exit
fi


echo "REMOTE STAGING"

echo "STAGING: Stop and remove container  api-ecommerce-golang & image  api-ecommerce-golang"
ssh dev@codethusiast.my.id -p 30004 'sudo -S docker container stop  api-ecommerce-golang'
ssh dev@codethusiast.my.id -p 30004 'sudo -S docker container rm  api-ecommerce-golang'
ssh dev@codethusiast.my.id -p 30004 'sudo -S docker image rm  api-ecommerce-golang'

echo "STAGING: Run  api-ecommerce-golang image as  api-ecommerce-golang"
echo "load docker image"
ssh dev@codethusiast.my.id -p 30004 'sudo -S docker load -i image.tar'
echo "end load docker image"
# ssh dev@codethusiast.my.id -p 30004 'sudo -S docker container run --name  api-ecommerce-golang --rm -e PORT=8282 -p 38282:8282  api-ecommerce-golang'
# ssh dev@codethusiast.my.id -p 30004 'sudo -S docker run  api-ecommerce-golang --name  api-ecommerce-golang'

# # create container
# echo "create docker container"
# ssh dev@codethusiast.my.id -p 30004 'sudo -S docker container create --name  api-ecommerce-golang -e PORT=8282 -e INSTANCE_ID="my first instance" -p 38889:8080  api-ecommerce-golang --network="mysql_staging_apps"'
# echo "end create docker container"
# # run  container
# echo "run docker container"
# ssh dev@codethusiast.my.id -p 30004 'sudo -S docker container start  api-ecommerce-golang'
# echo "end run docker container"
# # 

echo "start run docker container"


echo "pwd"
ssh dev@codethusiast.my.id -p 30004 'pwd'
echo "ls -l"
ssh dev@codethusiast.my.id -p 30004 'ls -l'

ssh dev@codethusiast.my.id -p 30004 'sudo -S docker run -d --net mysql_staging_apps -e PORT=9797 -p 38989:9797 --name  api-ecommerce-golang -v /home/contoh-upload-files:/app/api-ecommerce/upload-files  api-ecommerce-golang'
echo "end run docker container"

#ssh dev@codethusiast.my.id -p 30004 'docker run -d --net docker_job2go_test_db  api-ecommerce-golang --name  api-ecommerce-golang'

echo "STAGING: Cleanup"
ssh dev@codethusiast.my.id -p 30004 'rm image.tar'

echo "LOCAL: Cleanup"
docker image rm  api-ecommerce-golang

rm image.tar

echo "ok"
