terraform {
  backend "s3" {
    bucket = "terraform-states.tobyjsullivan.com"
    key    = "states/stateless-graph/staging.tfstate"
    region = "us-east-1"
  }
}

provider "aws" {
  region = "us-east-1"
}
