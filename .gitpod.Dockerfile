FROM gitpod/workspace-full:latest

# Install custom tools, runtimes, etc.
# For example "bastet", a command-line tetris clone:
 RUN go get -u github.com/onsi/ginkgo/ginkgo
#
# More information: https://www.gitpod.io/docs/config-docker/
