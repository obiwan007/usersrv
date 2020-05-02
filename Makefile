
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME_GQL=gqlsrv
BINARY_NAME_USER=usersrv
BINARY_NAME_ORDER=ordersrv
BINARY_NAME_STORE=storesrv

VERSION ?= latest

all: test buildgql builduser buildorder buildstore buildfrontend docker
buildgql: 
	cd gqlsrv/cli && $(GOBUILD) -o $(BINARY_NAME_GQL) -v 

builduser: 
	cd usersrv/cli && $(GOBUILD) -o $(BINARY_NAME_USER) -v 

buildstore: 
	cd eventstore/cli && $(GOBUILD) -o $(BINARY_NAME_STORE) -v 

buildorder: 
	cd ordersrv/cli && $(GOBUILD) -o $(BINARY_NAME_ORDER) -v 

buildfrontend: 
	cd frontend && npm i && npm run build



test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME_GQL)
	rm -f $(BINARY_NAME_USER)
rungql:
	cd gqlsrv/cli && $(GOBUILD) -o $(BINARY_NAME_GQL) -v 
	cd gqlsrv/cli && ./$(BINARY_NAME_GQL) -config config.conf --server_addr localhost:10000 --zipkin http://localhost:9411/api/v1/spans

runuser:
	cd usersrv/cli && $(GOBUILD) -o $(BINARY_NAME_USER) -v 
	cd usersrv/cli && ./$(BINARY_NAME_USER) --port 10000 --zipkin http://localhost:9411/api/v1/spans

runfrontend:
	cd frontend && GRAPHQL="http://localhost:8090" && npm run start

runorder:
	cd ordersrv/cli && $(GOBUILD) -o $(BINARY_NAME_ORDER) -v 
	cd ordersrv/cli && ./$(BINARY_NAME_ORDER) -config config.conf --port 10001 --zipkin http://localhost:9411/api/v1/spans

runstore:
	cd eventstore/cli && $(GOBUILD) -o $(BINARY_NAME_STORE) -v 
	cd eventstore/cli && ./$(BINARY_NAME_STORE) --port 10002 --zipkin http://localhost:9411/api/v1/spans

docker: docker-build docker-push

docker-build: builddockerfrontend builddockergqlsrv builddockerusersrv
	# docker build -t obiwan007/gqlsrv:${VERSION} -f ./gql_Dockerfile . 
	# docker build -t obiwan007/usersrv:${VERSION} -f ./user_Dockerfile . 
	# cd frontend && docker build -t obiwan007/frontend:${VERSION} -f ./nginxDockerfile .
	
docker-push: pushdockerusersrv 	pushdockergqlsrv pushdockerfrontend
	# docker push obiwan007/frontend:${VERSION}	
	# docker push obiwan007/gqlsrv:${VERSION}
	# docker push obiwan007/usersrv:${VERSION}

dockerfrontend: builddockerfrontend pushdockerfrontend

dockergqlsrv: builddockergqlsrv pushdockergqlsrv

dockerusersrv: builddockerusersrv pushdockerusersrv

builddockerfrontend:
	cd frontend && docker build -t obiwan007/frontend:${VERSION} -f ./nginxDockerfile .

pushdockerfrontend:
	docker push obiwan007/frontend:${VERSION}

builddockergqlsrv:
	docker build -t obiwan007/gqlsrv:${VERSION} -f ./gql_Dockerfile . 

pushdockergqlsrv:
	docker push obiwan007/gqlsrv:${VERSION}

builddockerusersrv:
	docker build -t obiwan007/usersrv:${VERSION} -f ./user_Dockerfile . 

pushdockerusersrv:
	docker push obiwan007/usersrv:${VERSION}



compose-build:
	docker-compose build

compose-up:
	docker-compose up

kapplyzipkin:
	kubectl apply -f config/zipkin.yaml

kapply:
	kubectl apply -f config/frontendservice.yaml
	kubectl apply -f config/gqlservice.yaml
	kubectl apply -f config/userservice.yaml
	kubectl apply -f config/zipkin.yaml

kfrontend:
	kubectl rollout restart deployment frontend-deployment

kgql:
	kubectl rollout restart deployment api-deployment

kredeploy:
	kubectl rollout restart deployment api-deployment
	kubectl rollout restart deployment user-deployment
	kubectl rollout restart deployment frontend-deployment

kdashboard:
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0-beta8/aio/deploy/recommended.yaml

okteto:
	export KUBECONFIG="${HOME}/Downloads/okteto-kube.config:${KUBECONFIG:-$HOME/.kube/config}"
#	k get deployment api-deployment -o yaml | sed "s/\(image: obiwan007\/gqlsrv\):.*$/\1:VERSION/" | grep image

protobuf:
	protoc -I proto/ proto/*.proto --go_out=plugins=grpc:proto

dbinstall:
	brew install cockroachdb/tap/cockroach

dbcreate:
	brew install cockroachdb/tap/cockroach
	cockroach user set markus --insecure
	cockroach sql --insecure -e 'CREATE DATABASE ordersdb'
	cockroach sql --insecure -e 'GRANT ALL ON DATABASE ordersdb TO markus'

dbcluster:
	cockroach start --insecure --store=ordersdb-1 --host=localhost --background
	cockroach start --insecure --store=ordersdb-2 --host=localhost --port=26258 --http-port=8081 --join=localhost:26257 --background
	cockroach start --insecure --store=ordersdb-3 --host=localhost --port=26259 --http-port=8082 --join=localhost:26257 --background

dbshell:
	cockroach sql --url="postgresql://markus@localhost:26257/ordersdb?sslmode=disable";

runnats:
	nats-streaming-server --store file --dir ./data --max_msgs 0 --max_bytes 0