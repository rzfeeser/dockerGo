# dockerGo
Alta3 Research | https://alta3.com
@RZFeeser      | A simple Go Webservice

The purpose of this repository is to provide a simple Go program for container testing within Docker and Kubernetes.

This service mimics the Simple Python Flask Service available @:
https://gitlab.com/alta3research/simpleflaskservice


The following endpoints are currently available:
- /endpoint - Description of what is returned
- /        - HTTP + HTML
- /ping    - HTTP + JSON
- /spock   - HTTP + JSON 
- /env     - HTTP + JSON
- /alta3 - HTTP + JSON
- /health - HTTP + JSON (response may be deplayed by setting HEALTH_DELAY)
   

Some behaviors of the webservice may be modified by setting the following environmental variables.

*Default Vaules:*
- ENV_VAR_TO_SET - Description
- PORT - The default port to run the webservice on. The default is 9876
- HEALTH_DELAY - The time in seconds to delay an HTTP response sent to /health. The default is 0. The only reason to change this might be mimicing a failure within a Kubernetes health check.
- VERSION - The version returned by the server. The default is 0.1.


Builds may be accomplished using the`Dockerfile`and Dockerfile.multistage`.

#### Getting started

git clone https://github.com/rzfeeser/dockergo

#### Running Dockerfile (building a container)

sudo docker build --tag dockergo .

#### Running Dockerfile.multistage (improving container builds)

(sudo) docker build -t dockergo:multistage -f Dockerfile.multistage .

#### Run the image

(sudo) docker run -d  -p 9876:9876 dockergo:latest
(sudo) docker run -d  -p 9876:9876 dockergo:multistage

#### Stop all containers

sudo docker stop $(sudo docker ps -aq)

#### Delete all running containers

sudo docker container prune


#### Author(s)
Russell Zachary Feeser - @RZFeeser
