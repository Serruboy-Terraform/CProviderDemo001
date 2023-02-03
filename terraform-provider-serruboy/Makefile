
all: build test

build:
		go get -v
		go build -o terraform-provider-serruboy

test:
		terraform init 
		terraform plan 
		terraform apply 