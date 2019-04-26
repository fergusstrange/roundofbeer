data "aws_caller_identity" "current" {}

variable "aws_region" {
  type = "string"
}

provider "aws" {
  region = "eu-west-1"
  version = "2.7.0"
}

terraform {
  backend "s3" {
    bucket = "roundofbeerinfra"
    key = "terraform/terraform.tfstate"
    region = "eu-west-1"
  }

  required_version = ">= 0.9.3"
}

resource "aws_s3_bucket" "roundofbeer_static_content" {
  bucket = "roundof.beer"
  region = "${var.aws_region}"
  versioning {
    enabled = true
  }
}

# Lambda Permissions
resource "aws_iam_policy" "roundofbeer_aws_iam_policy_lambda_execution" {
  name = "roundofbeer_aws_iam_policy_lambda_execution"
  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "dynamodb:*"
            ],
            "Resource": [
                "arn:aws:dynamodb:${var.aws_region}:${data.aws_caller_identity.current.account_id}:table/*",
                "arn:aws:dynamodb:${var.aws_region}:${data.aws_caller_identity.current.account_id}:table/*/*"
            ]
        }
    ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "roundofbeer_aws_iam_role_policy_attachment_lambda_execution" {
  role = "roundofbeer_lambda_function"
  policy_arn = "${aws_iam_policy.roundofbeer_aws_iam_policy_lambda_execution.arn}"
}

# API Gateway permissions
resource "aws_iam_role" "roundofbeer_gateway_aws_iam_role" {
  name = "roundofbeer_gateway_aws_iam_role"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "apigateway.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "roundofbeer_aws_iam_role_policy_invoke_lambda" {
  name = "roundofbeer_aws_iam_role_policy_invoke_lambda"
  role = "${aws_iam_role.roundofbeer_gateway_aws_iam_role.id}"
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Resource": [
        "*"
      ],
      "Action": [
        "lambda:InvokeFunction"
      ]
    },
    {
      "Effect": "Allow",
      "Resource": [
        "${aws_s3_bucket.roundofbeer_static_content.arn}",
        "${aws_s3_bucket.roundofbeer_static_content.arn}/*",
        "*"
      ],
      "Action": [
        "s3:GetObject",
        "s3:HeadObject",
        "s3:*"
      ]
    }
  ]
}
EOF
}

resource "aws_lambda_permission" "roundofbeer_aws_lambda_permission" {
  statement_id = "AllowAPIGatewayInvoke"
  action = "lambda:InvokeFunction"
  function_name = "${var.apex_function_hello}"
  principal = "apigateway.amazonaws.com"
  source_arn = "${aws_api_gateway_deployment.roundofbeer_aws_api_gateway_deployment_prod.execution_arn}/*/*"
}

# API Gateway
resource "aws_api_gateway_rest_api" "roundofbeer_aws_api_gateway_rest_api" {
  name = "roundofbeer_aws_api_gateway_rest_api"
  endpoint_configuration {
    types = [
      "REGIONAL"
    ]
  }
}

# API Gateway Root Resource
resource "aws_api_gateway_method" "roundofbeer_aws_api_gateway_method_root" {
  rest_api_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.id}"
  resource_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.root_resource_id}"
  http_method = "ANY"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "roundofbeer_aws_api_gateway_integration_root" {
  rest_api_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.id}"
  resource_id = "${aws_api_gateway_method.roundofbeer_aws_api_gateway_method_root.resource_id}"
  http_method = "${aws_api_gateway_method.roundofbeer_aws_api_gateway_method_root.http_method}"
  credentials = "${aws_iam_role.roundofbeer_gateway_aws_iam_role.arn}"
  integration_http_method = "GET"
  type = "AWS"
  uri = "arn:aws:apigateway:${var.aws_region}:s3:path/${aws_s3_bucket.roundofbeer_static_content.bucket}/index.html"
}

resource "aws_api_gateway_method_response" "roundofbeer_aws_api_gateway_method_response_200_root" {
  rest_api_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.id}"
  resource_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.root_resource_id}"
  http_method = "${aws_api_gateway_method.roundofbeer_aws_api_gateway_method_root.http_method}"
  status_code = "200"
  response_parameters = {
    "method.response.header.Timestamp" = true
    "method.response.header.Content-Length" = true
    "method.response.header.Content-Type" = true
  }
}

resource "aws_api_gateway_integration_response" "roundofbeer_aws_api_gateway_integration_response_root" {
  rest_api_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.id}"
  resource_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.root_resource_id}"
  http_method = "${aws_api_gateway_method.roundofbeer_aws_api_gateway_method_root.http_method}"
  status_code = "${aws_api_gateway_method_response.roundofbeer_aws_api_gateway_method_response_200_root.status_code}"
  selection_pattern = "-"
  response_parameters = {
    "method.response.header.Timestamp" = "integration.response.header.Date"
    "method.response.header.Content-Length" = "integration.response.header.Content-Length"
    "method.response.header.Content-Type" = "integration.response.header.Content-Type"
  }
}

# API Gateway Resource
resource "aws_api_gateway_resource" "roundofbeer_aws_api_gateway_resource" {
  rest_api_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.id}"
  parent_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.root_resource_id}"
  path_part = "{proxy+}"
}

resource "aws_api_gateway_method" "roundofbeer_aws_api_gateway_method" {
  rest_api_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.id}"
  resource_id = "${aws_api_gateway_resource.roundofbeer_aws_api_gateway_resource.id}"
  http_method = "ANY"
  authorization = "NONE"
  request_parameters {
    "method.request.path.proxy" = false
  }
}

resource "aws_api_gateway_integration" "roundofbeer_aws_api_gateway_integration" {
  rest_api_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.id}"
  resource_id = "${aws_api_gateway_resource.roundofbeer_aws_api_gateway_resource.id}"
  http_method = "${aws_api_gateway_method.roundofbeer_aws_api_gateway_method.http_method}"
  credentials = "${aws_iam_role.roundofbeer_gateway_aws_iam_role.arn}"
  integration_http_method = "GET"
  type = "AWS"
  request_parameters {
    "integration.request.path.proxy" = "method.request.path.proxy"
  }
  uri = "arn:aws:apigateway:${var.aws_region}:s3:path/${aws_s3_bucket.roundofbeer_static_content.bucket}/{proxy}"
  #integration_http_method = "POST"
  #type = "AWS_PROXY"
  #uri = "arn:aws:apigateway:${var.aws_region}:lambda:path/2015-03-31/functions/${var.apex_function_hello}/invocations"
}

resource "aws_api_gateway_method_response" "roundofbeer_aws_api_gateway_method_response_200" {
  rest_api_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.id}"
  resource_id = "${aws_api_gateway_resource.roundofbeer_aws_api_gateway_resource.id}"
  http_method = "${aws_api_gateway_method.roundofbeer_aws_api_gateway_method.http_method}"
  status_code = "200"
  response_parameters = {
    "method.response.header.Timestamp" = true
    "method.response.header.Content-Length" = true
    "method.response.header.Content-Type" = true
  }
}

resource "aws_api_gateway_integration_response" "roundofbeer_aws_api_gateway_integration_response_200" {
  rest_api_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.id}"
  resource_id = "${aws_api_gateway_resource.roundofbeer_aws_api_gateway_resource.id}"
  http_method = "${aws_api_gateway_method.roundofbeer_aws_api_gateway_method.http_method}"
  status_code = "${aws_api_gateway_method_response.roundofbeer_aws_api_gateway_method_response_200_hello.status_code}"
  selection_pattern = "-"
  response_parameters = {
    "method.response.header.Timestamp" = "integration.response.header.Date"
    "method.response.header.Content-Length" = "integration.response.header.Content-Length"
    "method.response.header.Content-Type" = "integration.response.header.Content-Type"
  }
}

resource "aws_api_gateway_method_response" "roundofbeer_aws_api_gateway_method_response_400" {
  rest_api_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.id}"
  resource_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.root_resource_id}"
  http_method = "${aws_api_gateway_method.roundofbeer_aws_api_gateway_method.http_method}"
  status_code = "400"
}

resource "aws_api_gateway_integration_response" "roundofbeer_aws_api_gateway_integration_response_400" {
  rest_api_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.id}"
  resource_id = "${aws_api_gateway_resource.roundofbeer_aws_api_gateway_resource.id}"
  http_method = "${aws_api_gateway_method.roundofbeer_aws_api_gateway_method.http_method}"
  status_code = "${aws_api_gateway_method_response.roundofbeer_aws_api_gateway_method_response_400.status_code}"
  selection_pattern = "4\\d{2}"
}

resource "aws_api_gateway_method_response" "roundofbeer_aws_api_gateway_method_response_500" {
  rest_api_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.id}"
  resource_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.root_resource_id}"
  http_method = "${aws_api_gateway_method.roundofbeer_aws_api_gateway_method.http_method}"
  status_code = "500"
}

resource "aws_api_gateway_integration_response" "roundofbeer_aws_api_gateway_integration_response_500" {
  rest_api_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.id}"
  resource_id = "${aws_api_gateway_resource.roundofbeer_aws_api_gateway_resource.id}"
  http_method = "${aws_api_gateway_method.roundofbeer_aws_api_gateway_method.http_method}"
  status_code = "${aws_api_gateway_method_response.roundofbeer_aws_api_gateway_method_response_500.status_code}"
  selection_pattern = "5\\d{2}"
}

#API Gateway Custom Domain & Deployment
resource "aws_acm_certificate_validation" "roundofbeer_aws_acm_certificate_validation" {
  certificate_arn = "arn:aws:acm:eu-west-1:673047522944:certificate/4c0c2595-0cd7-419b-bf78-ca3d43d8e931"
}

resource "aws_api_gateway_domain_name" "roundofbeer_aws_api_gateway_domain_name" {
  domain_name = "www.roundof.beer"
  regional_certificate_arn = "${aws_acm_certificate_validation.roundofbeer_aws_acm_certificate_validation.certificate_arn}"

  endpoint_configuration {
    types = [
      "REGIONAL"]
  }
}

resource "aws_route53_record" "roundofbeer_aws_route53_record" {
  name = "${aws_api_gateway_domain_name.roundofbeer_aws_api_gateway_domain_name.domain_name}"
  type = "A"
  zone_id = "Z34NRCFE2Q3TH"

  alias {
    evaluate_target_health = false
    name = "${aws_api_gateway_domain_name.roundofbeer_aws_api_gateway_domain_name.regional_domain_name}"
    zone_id = "${aws_api_gateway_domain_name.roundofbeer_aws_api_gateway_domain_name.regional_zone_id}"
  }
}

resource "aws_api_gateway_base_path_mapping" "roundofbeer_aws_api_gateway_base_path_mapping" {
  api_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.id}"
  domain_name = "${aws_api_gateway_domain_name.roundofbeer_aws_api_gateway_domain_name.domain_name}"
  stage_name = "${aws_api_gateway_deployment.roundofbeer_aws_api_gateway_deployment_prod.stage_name}"
}

# API Gateway deployment
resource "aws_api_gateway_deployment" "roundofbeer_aws_api_gateway_deployment_prod" {
  rest_api_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.id}"
  stage_name = "prod"
  depends_on = [
    "aws_api_gateway_integration.roundofbeer_aws_api_gateway_integration_root",
    "aws_api_gateway_integration.roundofbeer_aws_api_gateway_integration"
  ]
}