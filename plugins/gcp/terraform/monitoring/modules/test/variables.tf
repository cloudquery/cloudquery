variable "prefix" {
  description = "Prefix to use for all name resources"
  type        = string
  validation {
    condition     = length(var.prefix) == 2
    error_message = "The prefix should be exactly two characters."
  }
}

variable "project_id" {
  description = "Project where test resources will be deployed."
}

variable "region" {
  description = "Region where test resources will be deployed."
  default    = "us-east1"
}

variable "labels" {
  type = map
  default = {
    "environment" = "cq-provider-gcp"
  }
}