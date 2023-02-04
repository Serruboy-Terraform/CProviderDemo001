version := 1.0.2

all:
	rm ./../terraform.* -f
	rm ./../.terraform.* -f
	rm ./../.terraform -rf
	rm go.mod -f
	rm go.sum -f
	rm terraform-provider-serruboy -f
	go get -v
	go mod init terraform-provider-serruboy
	go fmt

	go mod tidy
	go build -o terraform-provider-serruboy

	mkdir -p ~/.terraform.d/plugins/terraform-example.com/exampleprovider/serruboy/$(version)/linux_amd64
	cp terraform-provider-serruboy ~/.terraform.d/plugins/terraform-example.com/exampleprovider/serruboy/$(version)/linux_amd64
	#rm go.mod -f
	#rm go.sum -f
	#rm ./../terraform.* -f
	#rm ./../.terraform.* -f
	#rm ./../.terraform -rf
	#rm terraform-provider-serruboy -f



