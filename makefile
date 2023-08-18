# Variables
GOLANG          := golang:1.20
ALPINE          := alpine:3.18
KIND            := kindest/node:v1.27.1

KIND_CLUSTER    := gopher-cluster
NAMESPACE       := models-system
APP             := models
REPO_NAME 		:= jessemolina
SERVICE_NAME    := models-api
VERSION         := 0.0.1
SERVICE_IMAGE   := $(REPO_NAME)/$(APP):$(VERSION)

# Developer
dev-up: k8s-up \
		dkr-build \
		k8s-load

dev-down: k8s-down

# Docker
dkr-build:
	docker build \
	-f deploy/docker/dockerfile.models \
	-t $(SERVICE_IMAGE) \
	--build-arg BUILD_REF=$(VERSION) \
	--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
	.

dkr-logs:
	docker logs $(SERVICE_NAME)

dkr-run:
	docker run -d --name $(SERVICE_NAME) $(SERVICE_IMAGE)

dkr-status:
	docker ps -a

dkr-sh:
	docker exec -it $(SERVICE_NAME) /bin/sh

dkr-prune:
	docker stop $(SERVICE_NAME)
	docker rm $(SERVICE_NAME)
	docker image rm $(SERVICE_IMAGE)

# Go
go-tidy:
	go mod tidy

go-run:
	go run cmd/services/models-api/main.go

# Helm
helm-install:
	helm install $(SERVICE_NAME) deploy/helm/models \
	-f deploy/k8s/dev/models/values.yaml \
	--create-namespace \
	--namespace $(NAMESPACE)

helm-upgrade:
	helm upgrade $(SERVICE_NAME) deploy/helm/models \
	--namespace $(NAMESPACE)

# Kubernetes
k8s-up:
	kind create cluster \
		--image $(KIND) \
		--name $(KIND_CLUSTER) \
		--config deploy/k8s/dev/kind-config.yaml

	kubectl wait --timeout=120s --namespace=local-path-storage --for=condition=Available deployment/local-path-provisioner

k8s-down:
	kind delete cluster --name $(KIND_CLUSTER)

k8s-load:
	kind load docker-image $(SERVICE_IMAGE) --name $(KIND_CLUSTER)

k8s-status:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces
