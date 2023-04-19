#Building stage
FROM golang:1.19.0-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go 

#Running stages
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/app.env .

# EXPOSE 9090/tcp
EXPOSE 9090
ENTRYPOINT ["bash", "-c","/app/main"]