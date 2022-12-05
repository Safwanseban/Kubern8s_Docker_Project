
FROM golang:alpine as builder
WORKDIR /Kubern8s_Docker_Project

COPY . .

RUN go build -o main .


FROM alpine:latest

COPY --from=builder /Kubern8s_Docker_Project/main .

COPY . .
CMD ["./main"]



