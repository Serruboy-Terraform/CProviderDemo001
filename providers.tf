terraform {
  required_providers {
    myfile = {
      version = "=0.1.0"
      source = "myorg.com/custom/serruboy"
    }
  }
}

provider "myfile" {
  encoding = "utf8"
}