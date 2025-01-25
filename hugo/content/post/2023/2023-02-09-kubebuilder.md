---
date: "2023-02-09T00:00:00Z"
tags:
- kubernetes
- kubebuilder
title: Getting Started with Kubernetes API by Kubebuilder
---


This document just follows [a quick tutorial](https://book.kubebuilder.io/quick-start.html#installation) for kubebuilder and learn its behavior.

# Install

Install `kubebuilder`.

```bash
curl -L -o kubebuilder https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)
chmod +x kubebuilder && mv kubebuilder /usr/local/bin/
```

# Workaround a first project
At first, create a project and a Guestbook API, and make manifests

```bash
kubebuilder init --domain at-ishikawa.github.io --repo at-ishikawa.github.io/at-ishikawa.github.io
kubebuilder create api --group webapp --version v1 --kind Guestbook
# Create Resource [y/n] y
# Create Controller [y/n] y
make manifests
```

Then you can see its manifest file under `config/crd/bases`.

## Deploy it into a kubernetes cluster

At first, create the namespace on the cluster.

```bash
kubectl create namespace kubebuilder
```

Then install CRD.

```bash
make install
```
Now you can see its CRD

```fish
> kubectl get crd | grep guestbook
guestbooks.webapp.at-ishikawa.github.io          2023-02-10T04:41:32Z
```

Then run a controller in a foreground job

```bash
make run
```

Create a sample resource for CRD
```bash
kubectl apply -f config/samples/webapp_v1_guestbook.yaml
```

Then deploy an image on the cluster
```bash
make docker-build docker-push IMG=$registry/$project:$tag
make deploy IMG=$registry/$project:$tag
```

Then you can see a deployment and pod in the namespace `tutorial-system`.

```bash
> kubectl get deploy -n tutorial-system
NAME                          READY   UP-TO-DATE   AVAILABLE   AGE
tutorial-controller-manager   1/1     1            1           9m57s
> kubectl get pods -n tutorial-system
NAME                                           READY   STATUS    RESTARTS   AGE
tutorial-controller-manager-649cdb6fb5-62nwb   2/2     Running   0          12m
```

This controller never handle a guestbook crd.

```
> kubectl get guestbook -n kubebuilder
NAME               AGE
guestbook-sample   2m6s
```

## Architecture

See [this page](https://book.kubebuilder.io/architecture.html) for the overview of the architecture.

* Manager
  * Controller: Handles an event. Event is first filtered by the predicate and then handled by a reconciler
    * Reconciler contains a logic of a controller for the resource
  * Webhook: Receives an admission request and set the default fields by Defaulter and rejects an invalid object by a Validator
    * They're mutating admission webhook and validating admission webhook
  * k8s API Client
  * Cache

The example of the reconciling of a controller is as follows, from [this page](https://book.kubebuilder.io/cronjob-tutorial/controller-implementation.html):

1. Load the named CronJob
1. List all active jobs, and update the status
1. Clean up old jobs according to the history limits
1. Check if we’re suspended (and don’t do anything else if we are)
1. Get the next scheduled run
1. Run a new job if it’s on schedule, not past the deadline, and not blocked by our concurrency policy
1. Requeue when we either see a running job (done automatically) or it’s time for the next scheduled run.
