# Provider Garage

`provider-garage` is a [Crossplane](https://crossplane.io/) provider
garage that is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the Garage
API.

## Getting Started

Install using declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-minio
spec:
  package: ghcr.io/eaglesemanation/provider-garage:v0.1.0
EOF
```

You can see the API reference here: https://doc.crds.dev/github.com/eaglesemanation/provider-garage@v0.1.0

## Developing

Run code-generation pipeline:
```console
make generate
```

Test an example against a Kind k8s cluster (provider config is included throug setup.sh):
```console
make e2e UPTEST_EXAMPLE_LIST="examples/namespaced/bucket/bucket.yaml,examples/namespaced/key/key.yaml,examples/namespaced/bucket/permission.yaml"
kind delete cluster -n local-dev
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/eaglesemanation/provider-garage/issues).
