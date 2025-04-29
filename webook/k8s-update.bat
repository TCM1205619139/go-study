@echo off

kubectl delete deployment webook
kubectl delete service webook
docker rmi -f flycash/webook:v0.0.1

call ./build-image.bat

kubectl apply -f k8s-webook-deployment.yaml
kubectl apply -f k8s-webook-service.yaml