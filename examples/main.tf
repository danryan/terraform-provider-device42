provider "device42" {
  api_endpoint = "https://device42.cudaops.com/api/1.0"
}

data "device42_devices" "first_10" {
  limit = 10
}

resource "device42_device" "new_shit" {
  name = "new thingy"
}
