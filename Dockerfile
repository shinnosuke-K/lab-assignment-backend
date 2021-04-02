FROM golang:1.16-alpine as builder
WORKDIR /
COPY ./main.go /
COPY ./go.mod /
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM busybox:latest
WORKDIR /
COPY --from=builder ./server /
CMD ["./server"]