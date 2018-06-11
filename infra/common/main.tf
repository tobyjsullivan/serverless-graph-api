variable "lambda_package" {}
variable "env" {}

resource "aws_iam_role" "lambda_role" {
  name_prefix = "graph-api-${var.env}"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_lambda_function" "handler" {
  filename         = "${var.lambda_package}"
  source_code_hash = "${base64sha256(file(var.lambda_package))}"
  function_name    = "graph-api-${var.env}"
  handler          = "lambda_handler"
  runtime          = "go1.x"
  role             = "${aws_iam_role.lambda_role.arn}"
}

output "lambda_func_arn" {
  value = "${aws_lambda_function.handler.arn}"
}

output "lambda_func_name" {
  value = "${aws_lambda_function.handler.function_name}"
}
