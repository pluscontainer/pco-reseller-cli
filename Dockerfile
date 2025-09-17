FROM golang:1.25@sha256:8305f5fa8ea63c7b5bc85bd223ccc62941f852318ebfbd22f53bbd0b358c07e1 AS build
RUN mkdir /src
COPY ./ /src/
WORKDIR /src
RUN go build -o /ps-openstack-client .

FROM ubuntu:noble@sha256:353675e2a41babd526e2b837d7ec780c2a05bca0164f7ea5dbbd433d21d166fc
RUN mkdir /app
RUN apt update
RUN apt install -y ca-certificates

COPY --from=build /ps-openstack-client /app/ps-openstack-client
ENTRYPOINT [ "/app/ps-openstack-client" ]
