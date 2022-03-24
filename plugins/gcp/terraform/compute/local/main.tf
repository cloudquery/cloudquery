module "test" {
    source = "../modules/test"
    prefix = var.prefix
    project_id = var.project_id
}