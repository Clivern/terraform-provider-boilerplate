
provider "boilerplate" {
  api_key = "key"
  api_url = "http://127.0.0.1:8080"
}

resource "boilerplate_server" "web" {
    name = "web"
    image = "ubuntu_18"
    region = "eu"
    size = "small"
}
