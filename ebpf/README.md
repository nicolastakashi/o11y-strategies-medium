## How to run:

First thing you need to do is to run `make setup` to prepare you environment, that command you create a minikube cluster and install all the required dependency.

Right after this you should have all the pods up and running, you can check this by the output provided or running `kubectl -n kube-system get pods`.

## Port forward Rely and UI

To access the [Hubble UI](http://localhost:12000) and also execute the connectivity test, you must run the following commands:

```bash
make port-forward-relay
make port-forward-ui
```

## Running connectivity test

Now you already have acess to Hubble UI you can use the following command to generate traffic and test Hubble.

```bash
make test-connectivity
```

## Requirements
- [Cilium CLI](https://formulae.brew.sh/formula/cilium-cli)
- [Hubble CLI](https://formulae.brew.sh/formula/hubble)

## References
- [Getting Started](https://docs.cilium.io/en/stable/gettingstarted/)
- [Hubble UI](https://docs.cilium.io/en/stable/gettingstarted/hubble/)