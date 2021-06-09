FROM golang:1.16.5 AS builder
WORKDIR /
COPY ./ /
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM busybox
WORKDIR /
COPY --from=builder ./server /
COPY lab.csv /
COPY student.csv /
CMD ["./server"]