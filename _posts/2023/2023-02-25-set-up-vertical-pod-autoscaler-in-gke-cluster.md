---
title: Set up a vertical pod autoscaler in a GKE cluster
date: 2023-02-26T19:00:00
tags:
  - kubernetes
  - vertical pod autoscaler
  - multidim pod autoscaler
---

To figure out which kubernetes objects are how much resource, a vertical pod autoscaler might be useful.
It has a feature to either automatically update values, or suggest values.
But it cannot be used with a horizontal pod autoscaler to update values if the horizontal pod autoscaler doesn't use a custom or external metrics. In that case, use a multidimensional pod autoscaling.

In this post, it's to test a vertical pod autoscaler by following [this document](https://cloud.google.com/kubernetes-engine/docs/concepts/verticalpodautoscaler) and [this article](https://softwaremill.com/vertical-pod-autoscaling-with-gcp/).

Besides, there are Pre-GA feature in GCP as of Feb 2023, MultidimPodAutoscaler, written in [this document](https://cloud.google.com/kubernetes-engine/docs/how-to/multidimensional-pod-autoscaling).
And this behaves
- HorizontalPodAutoscaler for CPU
- VerticalPodAutoscaler for memory


# Getting Started

Before using a VPA nor MPA, enable a vertical pod autoscaler on your GKE cluster.
Then create a VPA and MPA for your workload.


## The example of a VPA for a recommendation

For a recommendation configuration, set `.spec.updatePolicy` Off on a VPA.

At first, define a resource like Deployment for a VPA:

```yml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: php-apache
  namespace: test-vpa
spec:
  selector:
    matchLabels:
      run: php-apache
  template:
    metadata:
      labels:
        run: php-apache
    spec:
      containers:
      - name: php-apache
        image: k8s.gcr.io/hpa-example
        ports:
        - containerPort: 80
        resources:
          limits:
            cpu: 500m
            memory: 100M
          requests:
            cpu: 100m
            memory: 10M
```

For the above VPA, define VPA.

```yml
apiVersion: autoscaling.k8s.io/v1
kind: VerticalPodAutoscaler
metadata:
  name: php-apache-vpa
  namespace: test-vpa
spec:
  targetRef:
    apiVersion: "apps/v1"
    kind:       Deployment
    name:       php-apache
  updatePolicy:
    updateMode: "Off" # Only recommendation
  resourcePolicy:
    containerPolicies:
      - containerName: php-apache
        minAllowed:
          cpu: "100m"
          memory: "50Mi"
```

The example output looks like next

```bash
[personal|test-vpa] > kubectl get vpa php-apache-vpa -o yaml | yq .status.recommendation
containerRecommendations:
  - containerName: php-apache
    lowerBound:
      cpu: 100m
      memory: 50Mi
    target:
      cpu: 100m
      memory: 50Mi
    uncappedTarget:
      cpu: 1m
      memory: "11534336"
    upperBound:
      cpu: 100m
      memory: "52428800"
```

Note that we shouldn't use a HPA controlling by CPU.


## The example of a MPA

The target of the deployment is the same as VPA.
Note that there is no recommendation mode currently.

```yml
# https://cloud.google.com/kubernetes-engine/docs/how-to/multidimensional-pod-autoscaling#updatemode
apiVersion: autoscaling.gke.io/v1beta1
kind: MultidimPodAutoscaler
metadata:
  name: php-apache-autoscaler
  namespace: test-vpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: php-apache
  goals:
    metrics:
    - type: Resource
      resource:
      # Define the target CPU utilization request here
        name: cpu
        target:
          type: Utilization
          averageUtilization: 60
  constraints:
    global:
      minReplicas: 1
      maxReplicas: 5
    containerControlledResources: [ memory ]
    container:
    - name: '*'
    # Define boundaries for the memory request here
      requests:
        minAllowed:
          memory: 10Mi
        maxAllowed:
          memory: 100Mi
  policy:
    updateMode: "Auto"
```

Then you can see a HPA and VPA created by this MPA

```bash
[personal|test-vpa] > kubectl get hpa,vpa | grep php-apache-autoscaler
horizontalpodautoscaler.autoscaling/php-apache-autoscaler   Deployment/php-apache   1%/60%    1         5         1          21h
verticalpodautoscaler.autoscaling.k8s.io/php-apache-autoscaler   Auto          11534336   True       21h
```
