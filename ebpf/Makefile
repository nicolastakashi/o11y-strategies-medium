start-minikube:
	@minikube start --network-plugin=cni --cni=false

start-kind:
	@kind create cluster --config ./kind.yaml

install-cilium:
	@helm upgrade -i cilium cilium/cilium --namespace kube-system \
	--set=operator.replicas=1 --reuse-values \
   	--set hubble.relay.enabled=true \
   	--set hubble.ui.enabled=true \
	--set prometheus.enabled=true \
  	--set operator.prometheus.enabled=true \
	--set hubble.metrics.enableOpenMetrics=true \
  	--set hubble.metrics.enabled="{dns,drop,tcp,flow,port-distribution,icmp,httpV2:exemplars=true;labelsContext=source_ip\,source_namespace\,source_workload\,destination_ip\,destination_namespace\,destination_workload\,traffic_direction}"

install-monitoring:
	@kubectl apply -f https://raw.githubusercontent.com/cilium/cilium/v1.13/examples/kubernetes/addons/prometheus/monitoring-example.yaml

install-starwars:
	@kubectl apply -f ./http-sw-app.yaml

restart-unmanaged-pods:
	@kubectl get pods --all-namespaces -o custom-columns=NAMESPACE:.metadata.namespace,NAME:.metadata.name,HOSTNETWORK:.spec.hostNetwork --no-headers=true | grep '<none>' | awk '{print "-n "$$1" "$$2}' | xargs -L 1 -r kubectl delete pod --force --grace-period 0

sleep:
	sleep 60

setup: start-minikube install-cilium sleep restart-unmanaged-pods install-monitoring
	@kubectl -n kube-system get pods

port-forward-relay:
	@kubectl -n kube-system port-forward svc/hubble-relay 4245:80

port-forward-ui:
	@kubectl -n kube-system port-forward svc/hubble-ui 12000:80

port-forward-grafana:
	@kubectl -n cilium-monitoring port-forward service/grafana --address 0.0.0.0 --address :: 3000:3000

test-connectivity:
	@while true; do cilium connectivity test; done

xwing-request-landing:
	@while true; do kubectl exec xwing -- curl -s -XPOST deathstar.default.svc.cluster.local/v1/request-landing; done

tiefighter-request-landing:
	@while true; do kubectl exec tiefighter -- curl -s -XPOST deathstar.default.svc.cluster.local/v1/request-landing; done

destroy:
	@minikube delete