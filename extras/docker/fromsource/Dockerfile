# Dockerfile to build a container image from the current working tree. This
# includes any non-committed changes.
#
# Build from the root of the repository with:
#
#  $ buildah bud --network=host --file=extras/docker/fromsource/Dockerfile .
#


# need to use fedora 27 until glide is fixed in fedora 28
# (or we get an alternative glide)
FROM fedora:27 AS build
MAINTAINER Heketi Developers <heketi-devel@gluster.org>

LABEL version="1.3.1"
LABEL description="Development build"

# let's setup all the necessary environment variables
ENV BUILD_HOME=/build
ENV GOPATH=$BUILD_HOME/golang
ENV PATH=$GOPATH/bin:$PATH

COPY . $GOPATH/src/github.com/heketi/heketi/

# install dependencies, build and cleanup
RUN dnf -y install glide golang git make mercurial && \
    dnf -y clean all && \
    cd $GOPATH/src/github.com/heketi && \
    cd $GOPATH/src/github.com/heketi/heketi && \
    glide install -v && \
    make && \
    cp heketi /usr/bin/heketi && \
    cp client/cli/go/heketi-cli /usr/bin/heketi-cli


# the final container can be based on the latest Fedora version
FROM fedora:latest

MAINTAINER Heketi Developers <heketi-devel@gluster.org>

LABEL version="1.3.1"
LABEL description="Development build"

# install latest updates
RUN true \
    && dnf -y update \
    && dnf clean all \
    && true

# copy the compiled binaries from the build container
COPY --from=build /usr/bin/heketi /usr/bin/heketi
COPY --from=build /usr/bin/heketi-cli /usr/bin/heketi-cli

# post install config and volume setup
ADD extras/docker/fromsource/heketi.json /etc/heketi/heketi.json
ADD extras/docker/fromsource/heketi-start.sh /usr/bin/heketi-start.sh

VOLUME /etc/heketi

RUN mkdir /var/lib/heketi
VOLUME /var/lib/heketi

# expose port, set user and set entrypoint with config option
ENTRYPOINT ["/usr/bin/heketi-start.sh"]
EXPOSE 8080
