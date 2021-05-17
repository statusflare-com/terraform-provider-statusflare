TEST?=$$(go list ./...)
HOSTNAME=github.com
NAMESPACE=statusflare-com
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

testacc: 
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m -parallel 1

clean:
	rm -rf ${BINARY}
	rm -rf ./examples/.terraform
	rm -rf ./examples/terraform.*