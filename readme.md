# UserSrv a testing ground for Microservice experiments

This project evolved from a simple testing ground to a more full-fletched time tracker solution. My intend is to make this a viable tool for all kind of time tracking/reporting - i.E. for freelance workers working on different projects and trying to create reports out of ther work based on hours.

## Tech used

- golang as language for backend
- ReactJS/TypeScript/Node as frontend system
- GraphQL as API service (frontend2backend)
- GRPC for calling backend2backend
- zipkin/jaeger as tracing app
- etcd for discovery (not used currently)
- docker-compose for bringing up everything together on a dev system.
- Kubernetes as a system to orchestrate and make it available. Using a Minikube instance with apache as a router on a LinuxVM for now.

Some libraries to mention:
- Reactjs, materialui, create-react-app, typescript
- ent as entity framework https://github.com/facebookincubator/ent
- graphql backend github.com/graph-gophers/graphql-go, code generator used to precreate some basic model classes (github.com/DealTap/graphql-gen-go).
- opentracing (grpc, zipkin) to support tracing functionality.


## Services (excluding external)

### Frontend

 - Little React project using create-react-app with Typescript
 - material-ui as UI/UX lib.
 - nginx hosting page
 - nginx proxy graphql/rest requests to API service
 - Authorization uses Microsoft, Google OAuth2 services. Token will be evaluated in backend. New JWT Token will be created with clains.

### API Service

 - handling of login/auth via JWT token
 - providing GraphQL service/discovery
 - Forwarding requests to other GRPC based services.
 - JWT token will be retrieved from cookie/authorization handler and forwarded to grpc based services.
 - Each service could decide which data it will give based on the decoded JWT token.
 - GraphQL gateway could also filter data based on JWT.

### User Service

 - sample service storing username, email etc. is serialized to ENT orm to Postgresql
 - interfaced via GRPC protobuf services.
 - planned Teams support to make it easier to work on bigger projects and keep track on each team member.

### Timer Service

 - Handling timer entries. is serialized to ENT orm to Postgresql
 - interfaced via GRPC protobuf services.

### Client Service

 - Handling clients for one user. is serialized to ENT orm to Postgresql
 - Planned on extending that so that multiple user could work on one client and group users to teams.
 - interfaced via GRPC protobuf services.

### Project Service

 - Handling projects for one user. is serialized to ENT orm to Postgresql
 - Planned on extending that so that multiple user could work on one project via a team teams.
 - interfaced via GRPC protobuf services.

## Makefile for handling all the stuff

- make rungql runs the API GraphQL service
- make runuser/runclient/runtimer/runproject runs corresponding service
- make runfrontend runs the frontendservice

- make compose-build builds all the docker images for usage in docker-compose
- make compose-up runs the composed app completely

- make docker builds all images and push them to docker hub

- make kapply applies all configs to your kubernetes cluster. Be sure to deploy some services/deployments like jaeger seperately. Also all the secrets need to me manually changed.
- make kreapply restarts the deployment to make sure all versions are up to date.

Currently all docker images will be tagged with "latest". Github actions will be triggered to build everything via docker - but not pushed. On my linux staging machine a github hook is listending and rebuilding everything including push and redeploy. So the app will be updateded in a few minutes automatically.

Plans on using helm. This might help to deploy in a more secure way as using versioned image names - and be able to rollback.

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

`docker run -it --rm -p 8080:80 custom-nginx`

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

