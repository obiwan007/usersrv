# UserSrv a testing ground for Microservice experiments

## Tech used

- golang as language for backend
- ReactJS/TypeScript/Node as frontend system
- GraphQL as API service (frontend2backend)
- GRPC for calling backend2backend
- zipkin/jaeger as tracing app
- etcd for discovery (not used currently)
- docker-compose for bringing up everything together.

## Services (excluding external)

### Frontend

 - Little ReactGQ sample project
 - nodejs serving via koa server
 - proxy graphql/rest requests to API service

### API Service

 - handling of login/auth via JWT token
 - providing GraphQL service/discovery
 - Forwarding requests to user service

### User Service

 - sample service storing username, email etc. in some on disk file
 - interfaced via GRPC protobuf services.

## Makefile for handling all the stuff

- make rungql runs the API GraphQL service
- make runuser runs user service
- make runfronted runs the frontendservice

- make compose-build builds all the docker images for usage in docker-compose
- make compose-up runs the composed app completely

- make docker builds all images and push them to docker hub

# some tooling and instructions (basically to remember stuff)
## Generating proto models
`sh protogen.sh`

## Testing the gRPC interface

### gRPC Testing App 
`brew cask install bloomrpc`

Download via https://github.com/uw-labs/bloomrpc


## Service discovery with etcd
`etcdctl put my-service/1.2.3.4 '{"Addr":"localhost:10000","Metadata":"..."}'`
### fixing some strange dependency (probably not necessary anymore)
`rm -r $GOPATH/src/go.etcd.io/etcd/vendor/golang.org/x/net/trace`

# Tracing with zipkin:
`docker run -d -p 9411:9411 openzipkin/zipkin`

# Cert generation
`openssl req -newkey rsa:2048 \
  -new -nodes -x509 \
  -days 3650 \
  -out cert.pem \
  -keyout key.pem \
  -subj "/C=US/ST=California/L=Mountain View/O=Your Organization/OU=Your Unit/CN=localhost"`

# Running the App
## run server

`go run server.go -tls -cert_file ../certs/cert.pem -key_file ../certs/key.pem`

## run client
`go run main.go -tls -ca_file ../certs/cert.pem`

# Docker Builds
## Buidl container

`docker build -t obiwan007/usersrv .`
`docker push obiwan007/usersrv`

## Run service

`docker run -it --rm -p 10000:10000 --name usersrv obiwan007/usersrv`

# Clustering with Kubernetis
## Kubectl

https://kubernetes.io/docs/tasks/tools/install-kubectl/

`kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0-beta8/aio/deploy/recommended.yaml`
`kubectl proxy`

https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-multiple-clusters/

## Local Development
### Minicube Installation

https://minikube.sigs.k8s.io/docs/start/

`brew install minikube`

Show Dashboard

`minikube dashboard`


### Deploy to minikube

`kubectl create -f userservice.yaml`

### Usefull commands

`k kubectl apply -f config/nginx.yaml`

`k get pods`

`kubectl expose deployment nginx-deployment --port=80 --type=LoadBalancer`

