FROM golang:1.16-alpine AS build_base


WORKDIR /tmp/app
COPY . .

RUN go mod download
RUN go build -o ./out .

FROM alpine:3.9

COPY --from=build_base /tmp/app/out /app

EXPOSE 80
CMD ["/app"]