---
date: "2022-05-01T00:00:00Z"
tags:
- tidb
title: Getting Started with TiDB by Kubernetes Operator
---


Getting Started
---
Use minikube by following [this document](https://docs.pingcap.com/tidb-in-kubernetes/stable/get-started#minikube)

### Start minikube
```
minikube start
alias kubectl="minikube kubectl --"
kubectl cluster-info
```

### Apply an operator
Install its CRDs
```
kubectl create -f https://raw.githubusercontent.com/pingcap/tidb-operator/v1.3.2/manifests/crd.yaml
```

Install helm if it's not installed on Ubuntu.
```
curl https://baltocdn.com/helm/signing.asc | sudo apt-key add -
sudo apt-get install apt-transport-https --yes
echo "deb https://baltocdn.com/helm/stable/debian/ all main" | sudo tee /etc/apt/sources.list.d/helm-stable-debian.list
sudo apt-get update
sudo apt-get install helm
```

Install its Operator
```
helm repo add pingcap https://charts.pingcap.org/
kubectl create namespace tidb-admin
helm install --namespace tidb-admin tidb-operator pingcap/tidb-operator --version v1.3.2
kubectl get pods --namespace tidb-admin -l app.kubernetes.io/instance=tidb-operator
```

### Deploy a TiDB cluster

```
kubectl create namespace tidb-cluster && \
    kubectl -n tidb-cluster apply -f https://raw.githubusercontent.com/pingcap/tidb-operator/master/examples/basic/tidb-cluster.yaml
```

Deploy a TiDB monitoring services
```
kubectl -n tidb-cluster apply -f https://raw.githubusercontent.com/pingcap/tidb-operator/master/examples/basic/tidb-monitor.yaml
```

# Monitor a TiDB cluster
Access a Grafana dashboard
```
kubectl port-forward -n tidb-cluster svc/basic-grafana 3000 > pf3000.out &
```
Access `http://127.0.0.1:3000` on a browser.
The default username and password is admin


Access a TiDB dashboard from [this document](https://docs.pingcap.com/tidb-in-kubernetes/dev/access-dashboard#method-1-access-tidb-dashboard-by-port-forward)

```
kubectl port-forward svc/basic-discovery -n tidb-cluster 10262:10262
```
Access `http://localhost:10262/dashboard`.
