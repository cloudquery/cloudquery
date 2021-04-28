cloudquery {

  provider "aws" {
    source = "cloudquery"
    version = "v0.3.8"
  }
  provider "gcp" {
    source = "cloudquery"
    version = "v0.2.1"
  }

}