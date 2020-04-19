
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME_GQL=gqlsrv
BINARY_NAME_USER=usersrv

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
	cd gqlsrv/cli && ./$(BINARY_NAME_GQL) --server_addr localhost:10000 --zipkin http://localhost:9411/api/v1/spans

runuser:
	cd usersrv/cli && $(GOBUILD) -o $(BINARY_NAME_USER) -v 
	cd usersrv/cli && ./$(BINARY_NAME_USER) --port 10000 --zipkin http://localhost:9411/api/v1/spans

runfrontend:
	cd frontend && GRAPHQL="http://localhost:8090" && npm run dev:static 
	cd usersrv/cli && ./$(BINARY_NAME_USER) --port 10000 --zipkin http://localhost:9411/api/v1/spans


docker: docker-build docker-push

docker-build:
	docker build -t obiwan007/gqlsrv -f ./gql_Dockerfile . 
	docker build -t obiwan007/usersrv -f ./user_Dockerfile . 
	cd frontend && docker build -t obiwan007/frontend -f ./Dockerfile .
	
docker-push:	
	docker push obiwan007/frontend	
	docker push obiwan007/gqlsrv
	docker push obiwan007/usersrv

compose-build:
	docker-compose build

compose-up:
	docker-compose up