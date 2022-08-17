FROM golang:1.19
LABEL description="Two port app."
WORKDIR / 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o two-port-app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /two-port-app ./
CMD ["./two-port-app"]  