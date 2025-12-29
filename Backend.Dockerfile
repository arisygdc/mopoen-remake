FROM golang:1.18.1-alpine3.14 AS stage-one
WORKDIR /app
COPY . .
RUN go build -o app main.go


FROM alpine:3.14
RUN addgroup -S appgroup && adduser -S goapp -G appgroup
WORKDIR /home/goapp/app

# Copy aplikasi saja (config akan dibuat dari env vars)
COPY --from=stage-one /app/app .

# Copy entrypoint script
COPY docker-entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/docker-entrypoint.sh

# Ganti ownership ke user goapp
RUN chown -R goapp:appgroup /home/goapp/app

USER goapp

ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["./app"]      
