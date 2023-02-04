resource "public_ip" "serru"{
  iptype="v4"
}


data "public_data" "name" {}

data "public_data" "external_ip_from_aws" {
  resolver = "https://checkip.amazonaws.com/"
}

output "external_ips" {
  value = data.public_data.name.result
}

output "external_ip_from_awss" {
  value = data.public_data.external_ip_from_aws
}

output "sssss"{
  value = public_ip.serru.result
}

