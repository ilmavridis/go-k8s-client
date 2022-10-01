run:
	@kubectl create -f .\go-k8s-client.yaml

show:
	@kubectl wait --for=condition=Complete job/$(shell kubectl get jobs --no-headers -o custom-columns=":metadata.name" -n report) -n report
	@kubectl logs $(shell kubectl get pods --no-headers -o custom-columns=":metadata.name" -n report) -n report

clear:
	@kubectl delete  -f .\go-k8s-client.yaml

# Application runs inside the K8s cluster
runinside:
	@printf "\n * Deploy the associated Kubernetes resources and run the application as job. *\n"
	@make -s run  >/dev/null && make show
	@make -s clear >/dev/null &
	@printf "\n * Report produced, associated Kubernetes resources cleaned up. *\n" 

# Application runs outside the K8s cluster
runoutside: 
	@go run main.go $(NS)