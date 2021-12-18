service = "zoom"
output_directory = "../resources"

resource "zoom" "users" {
    path = "github.com/zoom-lib-golang/zoom-lib-golang/zoom.User"
}