FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/ContainerSolutions/locust_exporter/
COPY . .
RUN GO111MODULE=on go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/locust_exporter

FROM scratch
COPY --from=builder /go/bin/locust_exporter /go/bin/locust_exporter
ENTRYPOINT ["/go/bin/locust_exporter"]