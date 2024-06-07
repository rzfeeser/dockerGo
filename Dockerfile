FROM docker.io/golang:1.16-alpine

# Set destination for COPY
# /app is a common working directory to create
WORKDIR /app

# Move our copies of go.mod and go.sum
# into our container space
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./

# Perform a Go build with the -o switch to control the name of the output 
RUN go build -o /dockergo

# When we launch the container we will specify the ports to open
# via the docker command. This is more for documentation
EXPOSE 9876

# (Optional) environment variable that our dockerised
# application can make use of. The value of environment
# variables can also be set via parameters supplied
# to the docker command on the command line.
#ENV HTTP_PORT=5009

# Each Docker file may have one CMD
# This is what we want to "do" to kick things off
CMD [ "/dockergo" ]
