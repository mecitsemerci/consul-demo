# Consul Demo

## Docker compose on local

    docker-compose up

## Consul installation on kubernetes

    helm repo add hashicorp https://helm.releases.hashicorp.com

    helm install --values helm-consul-values.yaml consul hashicorp/consul --create-namespace --namespace consul --version "1.0.0"


Consul UI

    kubectl port-forward svc/consul-ui --namespace consul 8500:8500

Test

    kubectl apply -f counting.yaml && kubectl apply -f dashboard.yaml

    kubectl apply --filename intentions.yaml

    kubectl port-forward svc/dashboard --namespace default 9002:9002

## APP deployment
Cache-API deploy

    kubectl apply -f go-cache-api.yaml  


### References
https://developer.hashicorp.com/consul/tutorials/kubernetes/kubernetes-minikube