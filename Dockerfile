FROM alpine:latest
COPY nested-service /service
ENTRYPOINT ["/service"]
