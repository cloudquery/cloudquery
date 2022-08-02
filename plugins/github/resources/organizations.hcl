service          = "github"
output_directory = "."
add_generate     = true

resource "github" "" "organizations" {
  path = "github.com/google/go-github/v45/github.Installation"

  column "text_matches" {
    type              = "json"
    generate_resolver = true
  }

  options {
    primary_keys = ["id"]
  }

}

