# DevOps-GO-APP
DevOps Go APP

## GO-APP
mkdir go-app && cd $_
go mod init go-app
touch main.go

### Run locally
go run main.go
http://localhost:8080

### Build and run
go build
 ./go-app 

cd ../

## Docker
docker build -t go-app:latest .
docker run -p 8080:8080 go-app:latest

## Local regitry
### Run docker registry
docker run -d -p 5000:5000 --restart=always --name registry registry:2

### Push to local registry
docker tag go-app:v1 localhost:5000/go-app:v1
docker push localhost:5000/go-app:v1

## Push to kind
kind load docker-image go-app:latest --name kind

### check images in kind
docker exec -it kind-control-plane crictl images

## K8s
kubectl apply -f k8s/manifests/deployment.yaml
kubectl apply -f k8s/manifests/service.yaml 
kubectl apply -f k8s/manifests/ingress.yaml 

kubectl port-forward deployment/go-app 8080:8080 &
kubectl port-forward svc/go-app 8080:80

http://localhost:8080

kubectl get svc
docker exec -it kind-control-plane curl http://10.96.130.232

## Ingress controller
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
kubectl get pods -n ingress-nginx

kind delete cluster
kind create cluster --config kind-config.yaml

kubectl rollout restart deployment go-app
kubectl wait --for=condition=Ready pod -l app=go-app


## HELM
mkdir helm && cd $_
helm create go-app-chart
cd go-app-chart

rm -r Charts
rm -r templates/*

cp ../../k8s/manifests/* templates/

### Values
add values in values.yaml

helm install go-app ./go-app-chart
helm list

helm uninstall go-app
helm upgrade go-app ./go-app-chart

## CI 

### Github Workflows
https://github.com/marketplace

### Checkout
https://github.com/marketplace/actions/checkout

## CD

### Argo CD


