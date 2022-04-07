FROM golang as build-stage

WORKDIR /tinyurl

COPY . .

RUN go build -o tinyurl

FROM alpine
RUN apk --no-cache add ca-certificates libc6-compat
WORKDIR /tinyurl/
COPY --from=build-stage /tinyurl/tinyurl .
EXPOSE 8080
CMD ["/tinyurl"]