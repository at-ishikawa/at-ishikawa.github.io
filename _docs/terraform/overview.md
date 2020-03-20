---
title: Terraform overview
---
Wrote in November, 2019.

Basic concepts
===
There are some basic components for terraform.

1. Resource: one infrastructure element, like a virtual machine. Resource type is defined in a _provider_.
1. Provider: a collection of resource types. Most provider offer resource types on a platform. The example of a provider is AWS, GCP, or Kubernetes.
1. Module: a set of resources. In order to use modules, following elements may be used.
    1. Input variables:, defined as _variable_.
    1. Output variables: defined as _output_.
    1. Local variables: defined in _locals_.
1. Data Sources: Data defined and used externally from current terraform. This may be a different terraform state.


State
---
The current infrastructure configuration is stored in _state_ by terraform.
Terraform stores the state in __terraform.tfstate__ localy by default, but using _backend_, state can be stored remotely like AWS S3.
The state can contain sensitive data such as a database password.
To avoid this, by some backends, encryption or other features are provided.


External web services or tools
===

Terraform Cloud
---
[https://app.terraform.io/app]()

The managed service for terraform.
See [official page](https://www.terraform.io/docs/cloud/index.html) for more details.


 Terraform registry
 ---
[https://registry.terraform.io/]()

Reusable modules published on registry.


GitHub Actions
---
[https://www.terraform.io/docs/github-actions/index.html]()


Circle CI Orbs
---
[https://circleci.com/orbs/registry/orb/ovotech/terraform]()

There are many unofficial terraform orbs, but this has check, plan, and apply features with workspace, parallelism, and github comments.


Further readings
===
See [official pages](https://www.terraform.io/docs/configuration/index.html) for more details.
