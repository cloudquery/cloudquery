resource "azurerm_resource_group" "frontdoor" {
  name     = "${var.prefix}-frontdoor"
  location = "East US"
}

resource "azurerm_frontdoor" "test" {
  name = "${var.prefix}-frontdoor"
  tags = var.tags

  resource_group_name = azurerm_resource_group.frontdoor.name

  enforce_backend_pools_certificate_name_check = false

  routing_rule {
    name               = "exampleRoutingRule1"
    accepted_protocols = ["Http", "Https"]
    patterns_to_match  = ["/*"]
    frontend_endpoints = ["${var.prefix}-frontdoor"]
    forwarding_configuration {
      forwarding_protocol = "MatchRequest"
      backend_pool_name   = "exampleBackendBing"
    }
  }

  backend_pool_load_balancing {
    name = "exampleLoadBalancingSettings1"
  }

  backend_pool_health_probe {
    name = "exampleHealthProbeSetting1"
  }

  backend_pool {
    name = "exampleBackendBing"
    backend {
      host_header = "www.bing.com"
      address     = "www.bing.com"
      http_port   = 80
      https_port  = 443
    }

    load_balancing_name = "exampleLoadBalancingSettings1"
    health_probe_name   = "exampleHealthProbeSetting1"
  }

  frontend_endpoint {
    name      = "${var.prefix}-frontdoor"
    host_name = "${var.prefix}-frontdoor.azurefd.net"
  }
}
