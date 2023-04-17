SHELL := /bin/bash

# ================================================================
# GO

go-run:
	go run cmd/services/api/main.go

go-tidy:
	go mod tidy
	go mod vendor

# ================================================================
# DOCKER

IMAGE = gopher
VERSION := 1.0

docker-build:
	docker build \
	-f deploy/docker/Dockerfile \
	-t $(IMAGE):$(VERSION) \
	--build-arg BUILD_REF=$(VERSION) \
	--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
	.

docker-run:
	docker run 	\
	-p 3000:3000 -p 4000:4000 \
	-e DEBUG_PORT=3000 -e API_PORT=4000 \
	$(IMAGE):$(VERSION)


docker-sh:
	docker run -it $(IMAGE):$(VERSION) sh


# ================================================================
# KIND

KIND_CLUSTER = developer

kind-up:
	kind create cluster \
		--image kindest/node:v1.24.0@sha256:0866296e693efe1fed79d5e6c7af8df71fc73ae45e3679af05342239cdc5bc8e \
		--name $(KIND_CLUSTER) \
		--config deploy/k8s/kind/kind-config.yaml

kind-down:
	kind delete cluster --name $(KIND_CLUSTER)

kind-load:
	kind load docker-image gopher:$(VERSION) --name $(KIND_CLUSTER)

# ================================================================
# KUBE

kube-run:
	kubectl run gopher --image=gopher:$(VERSION) --port=4000

kube-apply:
	kubectl apply -f deploy/k8s/base/namespace.yaml
	kubectl apply -f deploy/k8s/base
	kubectl config set-context --current=true --namespace=gopher-ns

kube-images:
	kubectl get pods --all-namespaces -o jsonpath="{..image}" | tr -s '[[:space:]]' '\n' | sort | uniq
