#ARG GO_BUILDER_IMAGE_URI=golang
#ARG GO_BUILDER_IMAGE_VERSION=1.10
#FROM ${GO_BUILDER_IMAGE_URI}:${GO_BUILDER_IMAGE_VERSION} AS builder
FROM golang:1.10 AS builder
WORKDIR /
COPY ./ /
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM busybox:latest
WORKDIR /
COPY --from=builder ./server /
COPY lab.csv /
COPY student.csv /
CMD ["./server"]