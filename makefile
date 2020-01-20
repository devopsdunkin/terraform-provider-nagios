# Go env vars
GOCMD=go
GOBUILD=${GOCMD} build
GOTEST= ${GOCMD} test -v

# Terraform env vars
TF_PROVIDER_BIN=./terraform-provider-nagios

build:
	${GOBUILD} -o ${TF_PROVIDER_BIN}

test:
	${GOTEST} ./nagios

release:
	gox -osarch="linux/amd64" -output="./bin/terraform-provider-nagios_linux_amd64"
	gox -osarch="windows/amd64" -output="./bin/terraform-provider-nagios_windows_amd64.exe"
	gox -osarch="darwin/amd64" -output="./bin/terraform-provider-nagios_darwin_amd64"

resource:
	touch nagios/${name}.go
	touch nagios/resource_${name}.go
	touch nagios/resource_${name}_test.go
	touch docs/resources/resource_${name}.md

data_source:
	touch nagios/data_source_${name}.go
	touch nagios/data_source_${name}_test.go
	touch docs/data_sources/data_source_${name}.md