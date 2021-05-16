HOSTNAME=statusflare.com
NAMESPACE=statusflare
NAME=statusflare
BINARY=terraform-provider-${NAME}
VERSION=1.0
OS_ARCH=darwin_amd64

default: install

build:
	go build -o ${BINARY} 

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

clean:
	rm -rf ${BINARY}
	rm -rf ./examples/.terraform
	rm -rf ./examples/terraform.*
