version := 1.0.2

all:
	rm ./../terraform.* -f
	rm ./../.terraform.* -f
	rm ./../.terraform -rf
	rm go.mod -f
	rm go.sum -f
	rm terraform-provider-demo001 -f
	go get -v
	go mod init terraform-provider-demo001
	go fmt

	go mod tidy
	go build -o terraform-provider-demo001

	mkdir -p ~/.terraform.d/plugins/terraform-example.com/exampleprovider/demo001/$(version)/linux_amd64
	cp terraform-provider-demo001 ~/.terraform.d/plugins/terraform-example.com/exampleprovider/demo001/$(version)/linux_amd64
	#rm go.mod -f
	#rm go.sum -f
	#rm ./../terraform.* -f
	#rm ./../.terraform.* -f
	#rm ./../.terraform -rf
	#rm terraform-provider-demo001 -f



