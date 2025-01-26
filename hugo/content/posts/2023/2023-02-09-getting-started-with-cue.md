---
date: "2023-02-09T19:00:00Z"
tags:
- cue
title: Getting Started with cue
---

### Install
```bash
go install cuelang.org/go/cmd/cue@latest
```

### Generate a YAML file from cue

Prepare a sample file `user.cue` by a cue

```cue
import "time"

#User: {
    id: int
    name: string
    created: time.Format("2006-01-02")
}

john: #User
john: {
    id: 1
    name: "John"
    created: "2023-02-09"
}
```

Then run a command to output a yaml file by the `export` subcommand.

```bash
> cue export user.cue --out yaml
john:
  id: 1
  name: John
  created: "2023-02-09"
```

### Use Go packages

Use Go packages like k8s definitions, by following [this ](https://cuelang.org/docs/integrations/go/) and [this](https://cuetorials.com/first-steps/import-configuration/) documents.

```bash
go mod init github.com/at-ishikawa/at-ishikawa.github.com/examples/cue/go
go get k8s.io/api/core/v1
cue get go k8s.io/api/core/v1
```

Now then you can see the files like
```bash
> ls
cue.mod/ go.mod   go.sum
```

Now, let's make a file to create k8s service definitions by cue.
Let's say it's defined in `service.cue`

```cue
import "k8s.io/api/core/v1"

services: [string]: v1.#Service

services: {
    http: {
        // Next is shows an error
        // unknown: "name"
        apiVersion: "v1"
        kind: "Service"
        metadata: {
            name: "http-service"
            namespace: "app"
        }
        spec: {
            selector: {
                "app.kubernetes.io/name": "app"
            }
            ports: [
                {
                    protocol: "TCP"
                    port: 80
                    targetPort: 80
                }
            ]
        }
    }
}
```

Then you can export the yaml file for the kubernetes service `http-service` by

```
cue export service.cue --out yaml -e 'services.http'
```
