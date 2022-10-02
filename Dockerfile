FROM golang:1.18.1-alpine3.14 AS stage-one
WORKDIR /app
COPY . .
RUN go build -o app main.go


FROM alpine:3.14
RUN addgroup -S appgroup && adduser -S goapp -G appgroup
USER goapp
WORKDIR /home/goapp/app
COPY --from=stage-one /app .
COPY --from=stage-one /app/config.env .

CMD [ "./app" ]