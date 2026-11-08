FROM golang:1.27-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 go build -o /demo .

FROM alpine:3.24
COPY --from=build /demo /bin/demo
ENTRYPOINT ["/bin/demo"]
