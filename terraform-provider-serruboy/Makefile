
all: init build install

init:
		go get -v
		go mod init terraform-provider-serruboy
		go fmt

build:
		go mod tidy
		go build -o terraform-provider-serruboy

install:

		mkdir -p ~/.terraform.d/plugins/terraform-example.com/exampleprovider/serruboy/1.0.0/linux_amd64
		cp terraform-provider-serruboy ~/.terraform.d/plugins/terraform-example.com/exampleprovider/serruboy/1.0.0/linux_amd64
		rm go.mod -f
		rm go.sum -f
		rm terraform-provider-serruboy -f
    	
		

