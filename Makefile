
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
BINARY_NAME_TIMER=timersrv
BINARY_NAME_PROJECT=projectsrv
BINARY_NAME_CLIENT=clientsrv

VERSION ?= latest

all: test buildgql builduser buildorder buildstore buildfrontend buildtimer buildproject buildclient docker
buildgql: 
	cd gqlsrv/cli && $(GOBUILD) -o $(BINARY_NAME_GQL) -v 

builduser: 
	cd usersrv/cli && $(GOBUILD) -o $(BINARY_NAME_USER) -v 

buildtimer: 
	cd timersrv/cli && $(GOBUILD) -o $(BINARY_NAME_TIMER) -v 

buildproject: 
	cd projectsrv/cli && $(GOBUILD) -o $(BINARY_NAME_PROJECT) -v 

buildclient: 
	cd clientsrv/cli && $(GOBUILD) -o $(BINARY_NAME_CLIENT) -v 

buildstore: 
	cd eventstore/cli && $(GOBUILD) -o $(BINARY_NAME_STORE) -v 

buildorder: 
	cd ordersrv/cli && $(GOBUILD) -o $(BINARY_NAME_ORDER) -v 

buildfrontend: 
	cd frontend && npm i && npm run build

codegen:
	go get -u github.com/DealTap/graphql-gen-go
	# cd gqlsrv && graphql-gen-go schema/schema.graphql --out_dir api --pkg types
	cd gqlsrv && graphql-gen-go schema/schema.graphql --out_dir api --pkg gql
	

test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME_GQL)
	rm -f $(BINARY_NAME_USER)

runall:
	make rungql & make runuser & make runtimer & make runproject & make runclient & wait

rungql:
	cd gqlsrv/cli && go run main.go -config config.conf -usersrv localhost:10000 -timersrv localhost:10001 -projectsrv localhost:10002 -clientsrv localhost:10003 -zipkin http://localhost:9411/api/v1/spans | sed -e 's/^/[Command1] /'

runuser:
	cd usersrv/cli && go run main.go --port 10000 --zipkin http://localhost:9411/api/v1/spans | sed -e 's/^/[Command2] /'

timergen:
	cd timersrv/api/storage && go generate ./ent

runtimer:
	cd timersrv/cli && go run main.go -config config.conf --port 10001 --zipkin http://localhost:9411/api/v1/spans


runproject:
	cd projectsrv/cli && go run main.go --port 10002 --zipkin http://localhost:9411/api/v1/spans

projectgen:
	cd projectsrv/api/storage && go generate ./ent

runclient:
	cd clientsrv/cli && go run main.go --port 10003 --zipkin http://localhost:9411/api/v1/spans

clientgen:
	cd clientsrv/api/storage && go generate ./ent

runfrontend:
	cd frontend && GRAPHQL="http://localhost:8090" && npm run start

runorder:
	cd ordersrv/cli && go run main.go -config config.conf --port 10001 --zipkin http://localhost:9411/api/v1/spans

runstore:
	cd eventstore/cli && go run main.go --port 10002 --zipkin http://localhost:9411/api/v1/spans

docker: docker-build docker-push

docker-build: builddockerfrontend builddockergqlsrv builddockerusersrv builddockertimersrv builddockerprojectsrv builddockerclientsrv
	
docker-push: pushdockerusersrv 	pushdockergqlsrv pushdockerfrontend pushdockertimersrv pushdockerprojectsrv pushdockerclientsrv

dockerfrontend: builddockerfrontend pushdockerfrontend

dockergqlsrv: builddockergqlsrv pushdockergqlsrv

dockerusersrv: builddockerusersrv pushdockerusersrv

dockertimersrv: builddockertimersrv pushdockertimersrv

dockerprojectsrv: builddockerprojectsrv pushdockerprojectsrv

dockerclientsrv: builddockerclientsrv pushdockerclientsrv

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

builddockertimersrv:
	docker build -t obiwan007/timersrv:${VERSION} -f ./Dockerfile_timer . 

pushdockertimersrv:
	docker push obiwan007/timersrv:${VERSION}

builddockerprojectsrv:
	docker build -t obiwan007/projectsrv:${VERSION} -f ./Dockerfile_project . 

pushdockerprojectsrv:
	docker push obiwan007/projectsrv:${VERSION}

builddockerclientsrv:
	docker build -t obiwan007/clientsrv:${VERSION} -f ./Dockerfile_client . 

pushdockerclientsrv:
	docker push obiwan007/clientsrv:${VERSION}


compose-build:
	docker-compose build

compose-up:
	docker-compose up

kapplyzipkin:
	kubectl apply -f config/zipkin.yaml

#################################################################################### KUBERNETES 

kapply:
	kubectl apply -f config/frontendservice.yaml
	kubectl apply -f config/gqlservice.yaml
	kubectl apply -f config/userservice.yaml
	kubectl apply -f config/timerservice.yaml
	kubectl apply -f config/projectservice.yaml
	kubectl apply -f config/clientservice.yaml
	kubectl apply -f config/zipkin.yaml
	kubectl apply -f config/secret.yaml

kfrontend:
	kubectl rollout restart deployment frontend-deployment

kgql:
	kubectl rollout restart deployment api-deployment

kredeploy:
	kubectl rollout restart deployment api-deployment
	kubectl rollout restart deployment user-deployment
	kubectl rollout restart deployment timer-deployment
	kubectl rollout restart deployment project-deployment
	kubectl rollout restart deployment client-deployment
	kubectl rollout restart deployment frontend-deployment

kdashboard:
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0/aio/deploy/recommended.yaml

okteto:
	export KUBECONFIG="${HOME}/Downloads/okteto-kube.config:${KUBECONFIG:-$HOME/.kube/config}"
#	k get deployment api-deployment -o yaml | sed "s/\(image: obiwan007\/gqlsrv\):.*$/\1:VERSION/" | grep image

#################################################################################### KUBERNETES 

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

postgres:
	brew services start postgresql