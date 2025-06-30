module "example-vpc" {
  source = "terraform-aws-modules/vpc/aws"
}
module "example-bar" {
  source = "git::https://example.com/bar.git"
}
module "example-storage" {
  source = "../modules/storage"
}
module "example-vpc-2" {
  source = "terraform-aws-modules/vpc/aws"
}