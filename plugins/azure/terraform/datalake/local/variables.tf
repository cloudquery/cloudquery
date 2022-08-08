variable "prefix" {
  description = "Prefix to use for all name resources"
  type        = string
  validation {
    condition     = length(var.prefix) == 2
    error_message = "The prefix should be exactly two characters."
  }
}
