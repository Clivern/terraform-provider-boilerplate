
provider "boilerplate" {
  api_key = "key"
  api_url = "http://127.0.0.1:8080"
}

data "boilerplate_image" "image" {
  slug = "UBUNTU_18_04_64BIT"
}

resource "boilerplate_server" "web" {
    name = "web"
    image = data.boilerplate_image.image.id
    region = "eu"
    size = "small"
}

