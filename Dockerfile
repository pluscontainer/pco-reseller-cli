FROM golang:1.26@sha256:c83e68f3ebb6943a2904fa66348867d108119890a2c6a2e6f07b38d0eb6c25c5 AS build
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
