
provider "example" {
  api_key = "key"
  api_url = "http://127.0.0.1:8080"
}

resource "example_server" "web" {
    name = "web"
    image = "ubutu_18"
    region = "eu"
    size = "small"
}
