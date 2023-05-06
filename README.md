# eqemu-kube

EQEmulator on Kubernetes

## Requirements
- [Docker](https://docs.docker.com/engine/install/)
- [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
- [Helm](https://helm.sh/docs/intro/install/)

## Steps
- `make up` spin up a local k8s kind instance
- `make install` install CRDS for kind instance
- `make run` run the operator as a standalone binary that connects


## Developer Requirements
- [Kubebuilder](https://book.kubebuilder.io/quick-start.html#installation)