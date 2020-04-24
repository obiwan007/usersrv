
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME_GQL=gqlsrv
BINARY_NAME_USER=usersrv

VERSION ?= latest

all: test buildgql builduser
buildgql: 
	cd gqlsrv/cli && $(GOBUILD) -o $(BINARY_NAME_GQL) -v 
builduser: 
	cd usersrv/cli && $(GOBUILD) -o $(BINARY_NAME_USER) -v 
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
	cd frontend && GRAPHQL="http://localhost:8090" && npm run dev:static 
	cd usersrv/cli && ./$(BINARY_NAME_USER) --port 10000 --zipkin http://localhost:9411/api/v1/spans


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

kapply:
	kubectl apply -f config/frontendservice.yaml
	kubectl apply -f config/gqlservice.yaml
	kubectl apply -f config/userservice.yaml
#	kubectl apply -f config/zipkin.yaml

kfrontend:
	kubectl rollout restart deployment frontend-deployment

kgql:
	kubectl rollout restart deployment api-deployment

redeploy:
	kubectl rollout restart deployment api-deployment
	kubectl rollout restart deployment user-deployment
	kubectl rollout restart deployment frontend-deployment

okteto:
	export KUBECONFIG=$HOME/Downloads/okteto-kube.config:${KUBECONFIG:-$HOME/.kube/config}
#	k get deployment api-deployment -o yaml | sed "s/\(image: obiwan007\/gqlsrv\):.*$/\1:VERSION/" | grep image