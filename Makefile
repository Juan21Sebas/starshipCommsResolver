BUILDPATH=$(CURDIR)
API_NAME=`cat .name-ms`

#VERSION=`cat .version`
#GIT_BRANCH ?= $(shell git name-rev --name-only HEAD)
#GIT_VERSION ?= $(shell git rev-parse HEAD)
#GIT_VERSION_SHORT ?=$(shell git rev-parse --short HEAD)
#GIT_DESCRIPTION ?= $(shell git log -1 HEAD --pretty=format:%s )
#GIT_DATE ?= $(shell git show -s --format=%cd HEAD)
#GIT_COMMIT_TOTAL ?= $(shell git rev-list HEAD --count)
#DATE ?= $(shell date)

#PATH_RESOURCE ?= $(PWD)/resources/
#NAME_FILE ?= microservice-info.json

#JSON_INFO ?= "\"nameApp\": \"$(API_NAME)\", \"version\": \"$(VERSION)\", \"appGroupId\": \"$(API_NAME)\", \"artifact\": \"$(API_NAME)\", \"artifactName\": \"$(API_NAME)\", \"time\": \"$(DATE)\""
#JSON_GIT ?= "\"gitInfo\": { \"branch\":\"$(GIT_BRANCH)\",\"commit\": \"$(GIT_VERSION)\",\"commitAbrev\": \"$(GIT_VERSION_SHORT)\",\"description\": \"$(GIT_DESCRIPTION)\",\"time\": \"$(GIT_DATE)\",\"tags\": \"\",\"commitTotal\": $(GIT_COMMIT_TOTAL)}"
#JSON_MICROSERVICE_INFO ?= "{$(JSON_INFO), $(JSON_GIT)}"

build:
	@echo "Creando Binario ..."
	@go mod vendor
	@go build -mod=vendor -ldflags '-s -w' -o $(BUILDPATH)/build/bin/${API_NAME} cmd/main.go
	#@echo "Binario generado en build/bin/${API_NAME}"
	@echo "Binario generado en build/bin/${API_NAME}"

test:
	@echo "Ejecutando tests..."
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out

coverage:
	@echo "Coverfile..."
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out
	@go tool cover -html=coverfile_out -o coverfile_out.html

docker:
	@echo "Exec Docker Build ...."
	@docker build -t $(API_NAME):$(VERSION) -f $(PWD)/iaas/docker/Dockerfile .

resources:
		@echo "Creando file microservice-info.json"
    ifneq ("$(wildcard $(PATH_RESOURCE)$(NAME_FILE))","")
    	$(shell rm -f  $(PATH_RESOURCE)$(NAME_FILE))
    	$(shell touch $(PATH_RESOURCE)$(NAME_FILE))
    else
    	$(shell rm -f  $(PATH_RESOURCE)$(NAME_FILE))
    	$(shell touch $(PATH_RESOURCE)$(NAME_FILE))
    endif
    	$(shell echo $(JSON_MICROSERVICE_INFO) > $(PATH_RESOURCE)$(NAME_FILE) )

build-all: resources build

start:
	@cp .env ./configs/.env
	@go run ./cmd/main.go

.PHONY: test build