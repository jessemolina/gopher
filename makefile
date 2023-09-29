# Variables
GOLANG          := golang:1.20
ALPINE          := alpine:3.18
KIND            := kindest/node:v1.27.1

KIND_CLUSTER    := gopher-cluster
NAMESPACE       := models-system
REPO_NAME 		:= jessemolina
APP             := models
VERSION         := 0.0.1
SERVICE_IMAGE   := $(REPO_NAME)/$(APP):$(VERSION)
RELEASE_NAME    := models-api

# ================================================================
# Developer

dev-up: kind-up \
		docker-build \
		kind-load \
		helm-install

dev-update: docker-build \
			helm-uninstall \
			kind-load \
			helm-install

dev-upgrade: docker-build \
			 kind-load \
			 helm-upgrade

dev-down: kind-down

# ================================================================
# Docker

docker-build:
	docker build \
	-f deploy/docker/dockerfile.models \
	-t $(SERVICE_IMAGE) \
	--build-arg BUILD_REF=$(VERSION) \
	--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
	.

docker-logs:
	docker logs $(RELEASE_NAME)

docker-run:
	docker run -d --name $(RELEASE_NAME) $(SERVICE_IMAGE)

docker-status:
	docker ps -a

docker-sh:
	docker exec -it $(RELEASE_NAME) /bin/sh

docker-prune:
	docker stop $(RELEASE_NAME)
	docker rm $(RELEASE_NAME)
	docker image rm $(SERVICE_IMAGE)

# ================================================================
# Go

go-tidy:
	go mod tidy

go-run:
	go run cmd/services/models-api/main.go

# ================================================================
# Helm

helm-install:
	helm install $(RELEASE_NAME) deploy/helm/models \
	-f deploy/k8s/dev/models/values.yaml \
	--create-namespace \
	--namespace $(NAMESPACE)

helm-status:
	helm list -n $(NAMESPACE)
	@echo "\n"
	helm history -n $(NAMESPACE) $(RELEASE_NAME)

helm-uninstall:
	helm uninstall $(RELEASE_NAME) --namespace $(NAMESPACE)

helm-upgrade:
	helm upgrade $(RELEASE_NAME) deploy/helm/models \
	-f deploy/k8s/dev/models/values.yaml \
	--namespace $(NAMESPACE)

# ================================================================
# Kind

kind-up:
	kind create cluster \
		--image $(KIND) \
		--name $(KIND_CLUSTER) \
		--config deploy/k8s/dev/kind-config.yaml

	kubectl wait --timeout=120s --namespace=local-path-storage --for=condition=Available deployment/local-path-provisioner

kind-down:
	kind delete cluster --name $(KIND_CLUSTER)

kind-load:
	kind load docker-image $(SERVICE_IMAGE) --name $(KIND_CLUSTER)

# ================================================================
# Kubernetes

k8s-logs:
	kubectl logs --namespace=$(NAMESPACE) -l app=$(APP) --all-containers=true -f --tail=100

k8s-status:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces

# ================================================================
# Test

test-local:
	curl -il localhost:3000/test
