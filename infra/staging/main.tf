variable "lambda_package" {}

module "api" {
  source         = "../common"
  lambda_package = "${var.lambda_package}"
  env            = "staging"
}

output "api_invoke_url" {
  value = "${module.api.api_invoke_url}"
}
