---
title: Set up a prometheus by a prometheus operator in a kubernetes cluster
date: 2023-02-26T00:00:00
tags:
  - kubernetes
  - prometheus
  - prometheus-operator
---

This document is written by following [this document](https://prometheus-operator.dev/docs/user-guides/getting-started/).

# Install a prometheus operator

First, create CRDs and resources in the `default` namespace.

```fish
set LATEST (curl -s https://api.github.com/repos/prometheus-operator/prometheus-operator/releases/latest | jq -cr .tag_name)
curl -sLO https://github.com/prometheus-operator/prometheus-operator/releases/download/$LATEST/bundle.yaml
kubectl apply -f bundle.yaml
```

Note that I failed to create a prometheus resource because the CRD definition was too big like next.

```bash
The CustomResourceDefinition "prometheuses.monitoring.coreos.com" is invalid: metadata.annotations: Too long: must have at most 262144 bytes
```

Hence, I created the resource by the next command.

```
kubectl apply --server-side=true -f https://raw.githubusercontent.com/prometheus-operator/kube-prometheus/main/manifests/setup/0prometheusCustomResourceDefinition.yaml
```

## Create a Prometheus resource

Create a prometheus server with exposing its admin API
```yml
apiVersion: monitoring.coreos.com/v1
kind: Prometheus
metadata:
  name: prometheus
spec:
  serviceAccountName: prometheus
  podMonitorSelector:
    matchLabels:
      team: frontend
  resources:
    requests:
      memory: 400Mi
  enableAdminAPI: true
```

```yml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: prometheus
```

```yml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: prometheus
rules:
- apiGroups: [""]
  resources:
  - nodes
  - nodes/metrics
  - services
  - endpoints
  - pods
  verbs: ["get", "list", "watch"]
- apiGroups: [""]
  resources:
  - configmaps
  verbs: ["get"]
- apiGroups:
  - networking.k8s.io
  resources:
  - ingresses
  verbs: ["get", "list", "watch"]
- nonResourceURLs: ["/metrics"]
  verbs: ["get"]
```

```yml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: prometheus
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: prometheus
subjects:
- kind: ServiceAccount
  name: prometheus
  namespace: default
```

## An application monitor

On an application to be monitored, create a ServiceMonitor CRD.
For example, if there is the next example app,

```yml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: example-app
spec:
  replicas: 3
  selector:
    matchLabels:
      app: example-app
  template:
    metadata:
      labels:
        app: example-app
    spec:
      containers:
      - name: example-app
        image: fabxc/instrumented_app
        ports:
        - name: web
          containerPort: 8080
---
kind: Service
apiVersion: v1
metadata:
  name: example-app
  labels:
    app: example-app
spec:
  selector:
    app: example-app
  ports:
  - name: web
    port: 8080
```

Then deploy the ServiceMonitor like next

```yml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: example-app
  labels:
    team: frontend
spec:
  selector:
    matchLabels:
      app: example-app
  endpoints:
  - port: web
```
