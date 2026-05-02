FROM golang:1.25@sha256:0d8c14c93010e9eab98599f1ddd37e80b8fd39e9c662d670c4e4d9d0b101831d AS build
RUN mkdir /src
COPY ./ /src/
WORKDIR /src
RUN go build -o /ps-openstack-client .

FROM ubuntu:noble@sha256:c4a8d5503dfb2a3eb8ab5f807da5bc69a85730fb49b5cfca2330194ebcc41c7b
RUN mkdir /app
RUN apt update
RUN apt install -y ca-certificates

COPY --from=build /ps-openstack-client /app/ps-openstack-client
ENTRYPOINT [ "/app/ps-openstack-client" ]
