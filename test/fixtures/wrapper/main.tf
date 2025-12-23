module "projects" {
  source = "../../modules/project"
}

module "storage-buckets" {
  source = "./storage-buckets"
}