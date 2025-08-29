kubectl apply -f k8s/manifests/deployment.yaml
kubectl apply -f k8s/manifests/service.yaml 
kubectl apply -f k8s/manifests/ingress.yaml 

kubectl port-forward deployment/go-app 8080:8080 &
