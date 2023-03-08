#!/bin/bash

#fool's man CI/CD
pwd

echo "DEPLOY LOCAL"

# stop container
docker container stop api-ecommerce-golang 
# delete container
docker container rm api-ecommerce-golang 
# delete images
docker image rm api-ecommerce-golang-images

echo "LOCAL: Build and save image"
docker build . -f Dockerfile --build-arg enviro=local  -t api-ecommerce-golang-images
# docker build . -f Dockerfile -t api-ecommerce-golang-images

echo "LOCAL: docker run images"

# untuk folder upload di dalam /app
# docker run -d -e PORT=9797  -p 38989:9797 --name api-ecommerce-golang -v D:/interview/backend/transvision/coba-upload-file:/app/upload-file  api-ecommerce-golang-images

# unutk folder di dalam aplikasi
docker run -d -e PORT=9797  -p 38989:9797 --name api-ecommerce-golang -v D:/interview/backend/transvision/coba-upload-file:/app/api-ecommerce/upload-files/images  api-ecommerce-golang-images

echo "ok"
