FROM golang:1.18.1-alpine3.14 AS stage-one
WORKDIR /app
COPY . .
RUN go build -o app main.go


FROM alpine:3.14
WORKDIR /mopoen-remake
COPY --from=stage-one /app .
COPY --from=stage-one /app/config.env .
CMD [ "./app" ]

