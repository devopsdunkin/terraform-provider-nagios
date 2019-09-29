# Go env vars
GOCMD=go
GOBUILD=${GOCMD} build
GOTEST= ${GOCMD} test -v

# Terraform env vars
TF_PROVIDER_BIN=./terraform-provider-nagios

all: setup build test run

build:
	${GOBUILD} -o ${TF_PROVIDER_BIN}

cleanup: nagios-cleanup build


docker-start:
	docker start nagiosxi

init:
	docker pull tgoetheyn/docker-nagiosxi:latest

log-cleanup:
	

nagios-cleanup:
	newman run terraform-provider-nagios.postman_collection.json

run: docker-start tf-init tf-plan tf-apply

setup:
	${GOCMD} run env/setup.go

test:
	${GOTEST} ./nagios

tf-apply:
	terraform apply -auto-approve the_plan

tf-destroy:
	terraform destroy -auto-approve

tf-init:
	terraform init

tf-plan:
	terraform plan -out the_plan