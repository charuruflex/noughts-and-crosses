# build
FROM golang:alpine as builder
RUN mkdir -p /go/src/app/build /go/src/app
WORKDIR /go/src/app
COPY . .
RUN go get
RUN go build -o /game main.go

# deployment
FROM alpine:latest
ARG PORT
ARG SIZE
ENV PORT ${PORT}
ENV SIZE ${SIZE}
# RUN mkdir /app
WORKDIR /
COPY --from=builder /game .
CMD ["./game"]