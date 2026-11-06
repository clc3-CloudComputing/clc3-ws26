# Pull in the layer of the base image: alpine:3.20
FROM alpine:3.22

# Copy binary demo to the folder `/bin/`
COPY demo /bin/demo

# Run the service demo when a container is launched
CMD ["/bin/demo"]