FROM golang:1.21.1-alpine as builder

RUN apk --no-cache add ca-certificates git

FROM scratch
ENV GOTIFY_SERVER_PORT="80"

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app
ADD gotify-app /app/
EXPOSE 80
ENTRYPOINT ["./gotify-app"]
