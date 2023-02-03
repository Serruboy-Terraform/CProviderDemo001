resource "example_server" "name" {
  uuid_count="1"
}

resource "example_server_2" "name" {
  country="ES"
}

output "example_server_2" {
  value = example_server_2.name.country
}