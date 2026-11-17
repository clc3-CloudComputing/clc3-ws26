FROM --platform=$BUILDPLATFORM golang:1.27-alpine AS build
ARG TARGETOS
ARG TARGETARCH
WORKDIR /src
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /demo .

FROM alpine:3.24
COPY --from=build /demo /bin/demo
ENTRYPOINT ["/bin/demo"]
