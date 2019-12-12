
provider "example" {
  api_key = "key"
  api_url = "http://requestbin.net/r/1jdf7yi1"
}

resource "example_server" "web" {
    name = "web"
    image = "ubutu_18"
    region = "eu"
    size = "small"
}