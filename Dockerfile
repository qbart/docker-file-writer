FROM golang:1.17.9-alpine3.15 as build
RUN mkdir /build
ADD . /build
WORKDIR /build
RUN go build -o job .

FROM alpine:latest

RUN addgroup user 
RUN adduser -S -G user user
USER user

WORKDIR /app
COPY --from=build /build/job .
CMD ["/app/job"]

