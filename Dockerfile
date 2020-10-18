# Dockerfile
# start with an alpine image for small footprint
FROM golang:1.10-alpine
# install required tools
RUN apk update && apk add bash git curl
# install dep system-wide
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
# copy project to container go src and set the work dir to it
ADD . /go/src/github.com/mahendrafathan/registration
WORKDIR /go/src/github.com/mahendrafathan/registration
# build and install the application
RUN go install
# run the application
CMD ["/go/bin/registration"]