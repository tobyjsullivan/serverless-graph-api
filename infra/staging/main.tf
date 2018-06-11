terraform {
  backend "s3" {
    bucket = "terraform-states.tobyjsullivan.com"
    key    = "states/stateless-graph/staging.tfstate"
    region = "us-east-1"
  }
}

variable "lambda_package" {}

provider "aws" {
  region = "ap-southeast-2"
}

module "api" {
  source         = "../common"
  lambda_package = "${var.lambda_package}"
  env            = "staging"
}

output "lambda_func_arn" {
  value = "${module.api.lambda_func_arn}"
}

output "lambda_func_name" {
  value = "${module.api.lambda_func_name}"
}
