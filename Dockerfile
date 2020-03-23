FROM golang:alpine AS build

WORKDIR /go/src/routine.sh/nested-service
COPY . /go/src/routine.sh/nested-service
RUN go build -o /app

FROM alpine:latest
COPY --from=build /app /app
ENTRYPOINT [ "/app" ]
