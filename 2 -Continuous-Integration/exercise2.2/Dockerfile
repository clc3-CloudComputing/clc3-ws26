# Pull in the base image
FROM alpine:3.24
# Copy binary demo to the folder `/bin/`
COPY demo /bin/demo
# Run the service demo when a container is launched
ENTRYPOINT ["/bin/demo"]