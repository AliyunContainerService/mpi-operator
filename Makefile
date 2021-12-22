BIN_DIR=_output/cmd/bin
IMG_BUILDER=docker
IMAGE_NAME?=kubeflow/mpi-operator
RELEASE_VERSION=v20211222
GIT_COMMIT:=$(shell git rev-parse --short HEAD)

build: all

all: init fmt mpi-operator kubectl-delivery

mpi-operator:
	go build -o ${BIN_DIR}/mpi-operator ./cmd/mpi-operator/

kubectl-delivery:
	go build -o ${BIN_DIR}/kubectl-delivery ./cmd/kubectl-delivery/

init:
	mkdir -p ${BIN_DIR}

fmt:
	go fmt ./...

clean:
	rm -fr ${BIN_DIR}

images:
	@echo "version: ${RELEASE_VERSION}"
	${IMG_BUILDER} build --no-cache -t ${IMAGE_NAME}:${RELEASE_VERSION}-${GIT_COMMIT} .

.PHONY: clean