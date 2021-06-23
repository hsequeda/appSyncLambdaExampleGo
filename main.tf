variable "AWS_REGION" {
  type = string
}

provider "aws" {
  region = var.AWS_REGION
}

data "archive_file" "lambda_zip" {
  type        = "zip"
  source_file  = "./function/main"
  output_path = "app_sync_lambda_example.zip"
}

resource "aws_lambda_function" "app_sync_lambda_example" {
  filename         = "app_sync_lambda_example.zip"
  function_name    = "app_sync_lambda_example"
  handler          = "main"
  role             = aws_iam_role.iam_for_lambda_tf.arn
  source_code_hash = data.archive_file.lambda_zip.output_base64sha256
  runtime          = "go1.x"
}

resource "aws_iam_role" "iam_for_lambda_tf" {
  name               = "iam_for_lambda_tf"
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
