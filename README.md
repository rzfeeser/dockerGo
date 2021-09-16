# dockerGo

The purpose of this repository is to provide a simple Go program for "Dockerization". Builds may be accomplished using `Dockerfile` and `Dockerfile.multistage`.

#### Getting started

git clone https://github.com/rzfeeser/dockergo

#### Running Dockerfile (building a container)

sudo docker build --tag dockergo .

#### Running Dockerfile.multistage (improving container builds)

(sudo) docker build -t dockergo:multistage -f Dockerfile.multistage .

#### Run the image

(sudo) docker run -d  -p 8085:8085 dockergo:latest
(sudo) docker run -d  -p 8085:8085 dockergo:multistage

#### Stop all containers

sudo docker stop $(sudo docker ps -aq)

#### Delete all running containers

sudo docker container prune
