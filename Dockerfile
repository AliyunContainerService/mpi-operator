FROM golang:1.15.13 AS build

ADD . /go/src/github.com/kubeflow/mpi-operator
WORKDIR /go/src/github.com/kubeflow/mpi-operator
RUN make mpi-operator

FROM gcr.io/distroless/base-debian10:latest
COPY --from=build /go/src/github.com/kubeflow/mpi-operator/_output/cmd/bin/* /opt/

ENTRYPOINT ["/opt/mpi-operator"]
CMD ["--help"]