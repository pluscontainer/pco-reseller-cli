FROM golang:1.25@sha256:0d8c14c93010e9eab98599f1ddd37e80b8fd39e9c662d670c4e4d9d0b101831d AS build
RUN mkdir /src
COPY ./ /src/
WORKDIR /src
RUN go build -o /ps-openstack-client .

FROM ubuntu:noble@sha256:59a458b76b4e8896031cd559576eac7d6cb53a69b38ba819fb26518536368d86
RUN mkdir /app
RUN apt update
RUN apt install -y ca-certificates

COPY --from=build /ps-openstack-client /app/ps-openstack-client
ENTRYPOINT [ "/app/ps-openstack-client" ]
