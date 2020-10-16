FROM golang:1.15.3 as goartifact
WORKDIR /app
COPY src/main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
WORKDIR /root/
COPY --from=goartifact /app/app .
CMD ["./app"]

