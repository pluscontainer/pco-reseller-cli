FROM golang:1.27@sha256:2920bb124f1a2e12a1d39f0cab756ae0068441ee1da7655dec7f5e374944095b AS build
RUN mkdir /src
COPY ./ /src/
WORKDIR /src
RUN go build -o /ps-openstack-client .

FROM ubuntu:resolute@sha256:f3d28607ddd78734bb7f71f117f3c6706c666b8b76cbff7c9ff6e5718d46ff64
RUN mkdir /app
RUN apt update
RUN apt install -y ca-certificates

COPY --from=build /ps-openstack-client /app/ps-openstack-client
ENTRYPOINT [ "/app/ps-openstack-client" ]
