FROM golang:1.16.13-bullseye AS build

ADD . /go/src/github.com/kubeflow/mpi-operator
WORKDIR /go/src/github.com/kubeflow/mpi-operator
RUN make mpi-operator

FROM debian:bullseye-20211220
COPY --from=build /go/src/github.com/kubeflow/mpi-operator/_output/cmd/bin/* /opt/

ENTRYPOINT ["/opt/mpi-operator"]
CMD ["--help"]
