FROM golang:1.14-alpine
# golang:1.14-alpine with Git

# Install git
#RUN set -ex; \
#    apk update; \
#    apk add --no-cache git

# Set working directory
WORKDIR /mnt/e/Project/gomod/docker_compose_test/

# Run tests
CMD go test -v ./...