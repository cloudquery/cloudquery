cloudquery {
  connection {
    dsn =  "host=localhost user=postgres password=pass DB.name=postgres port=5432"
  }
  provider "test" {
    source = "cloudquery"
    version = "v0.0.0"
  }
}


provider "aws" {
  configuration {
    account "1" {
      regions: ["asdas"]
      resources: ["ab", "c"]
    }

    regions: ["adsa"]
  }
  resources: ["slow_resource"]
}