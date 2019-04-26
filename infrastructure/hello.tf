variable "apex_function_hello" {
  type = "string"
}

# API Gateway Resource
resource "aws_api_gateway_resource" "roundofbeer_aws_api_gateway_resource_hello" {
  rest_api_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.id}"
  parent_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.root_resource_id}"
  path_part = "/api/round"
}

resource "aws_api_gateway_method" "roundofbeer_aws_api_gateway_method_hello" {
  rest_api_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.id}"
  resource_id = "${aws_api_gateway_resource.roundofbeer_aws_api_gateway_resource_hello.id}"
  http_method = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "roundofbeer_aws_api_gateway_integration_hello" {
  rest_api_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.id}"
  resource_id = "${aws_api_gateway_resource.roundofbeer_aws_api_gateway_resource_hello.id}"
  http_method = "${aws_api_gateway_method.roundofbeer_aws_api_gateway_method_hello.http_method}"
  credentials = "${aws_iam_role.roundofbeer_gateway_aws_iam_role.arn}"
  integration_http_method = "POST"
  type = "AWS_PROXY"
  uri = "arn:aws:apigateway:${var.aws_region}:lambda:path/2015-03-31/functions/${var.apex_function_hello}/invocations"
}

resource "aws_api_gateway_method_response" "roundofbeer_aws_api_gateway_method_response_200_hello" {
  rest_api_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.id}"
  resource_id = "${aws_api_gateway_resource.roundofbeer_aws_api_gateway_resource_hello.id}"
  http_method = "${aws_api_gateway_method.roundofbeer_aws_api_gateway_method_hello.http_method}"
  status_code = "200"
  response_parameters = {
    "method.response.header.Timestamp" = true
    "method.response.header.Content-Length" = true
    "method.response.header.Content-Type" = true
  }
  response_models {
    "application/json" = "Empty"
  }
}

resource "aws_api_gateway_integration_response" "roundofbeer_aws_api_gateway_integration_response_200_hello" {
  rest_api_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.id}"
  resource_id = "${aws_api_gateway_resource.roundofbeer_aws_api_gateway_resource_hello.id}"
  http_method = "${aws_api_gateway_method.roundofbeer_aws_api_gateway_method_hello.http_method}"
  status_code = "${aws_api_gateway_method_response.roundofbeer_aws_api_gateway_method_response_200_hello.status_code}"
  selection_pattern = "-"
  response_parameters = {
    "method.response.header.Timestamp" = "integration.response.header.Date"
    "method.response.header.Content-Length" = "integration.response.header.Content-Length"
    "method.response.header.Content-Type" = "integration.response.header.Content-Type"
  }
  response_templates {
    "application/json" = "Empty"
  }
}

resource "aws_api_gateway_method_response" "roundofbeer_aws_api_gateway_method_response_400_hello" {
  rest_api_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.id}"
  resource_id = "${aws_api_gateway_resource.roundofbeer_aws_api_gateway_resource_hello.id}"
  http_method = "${aws_api_gateway_method.roundofbeer_aws_api_gateway_method_hello.http_method}"
  status_code = "400"
}

resource "aws_api_gateway_integration_response" "roundofbeer_aws_api_gateway_integration_response_400_hello" {
  rest_api_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.id}"
  resource_id = "${aws_api_gateway_resource.roundofbeer_aws_api_gateway_resource_hello.id}"
  http_method = "${aws_api_gateway_method.roundofbeer_aws_api_gateway_method_hello.http_method}"
  status_code = "${aws_api_gateway_method_response.roundofbeer_aws_api_gateway_method_response_400.status_code}"
  selection_pattern = "4\\d{2}"
}

resource "aws_api_gateway_method_response" "roundofbeer_aws_api_gateway_method_response_500_hello" {
  rest_api_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.id}"
  resource_id = "${aws_api_gateway_resource.roundofbeer_aws_api_gateway_resource_hello.id}"
  http_method = "${aws_api_gateway_method.roundofbeer_aws_api_gateway_method_hello.http_method}"
  status_code = "500"
}

resource "aws_api_gateway_integration_response" "roundofbeer_aws_api_gateway_integration_response_500_hello" {
  rest_api_id = "${aws_api_gateway_rest_api.roundofbeer_aws_api_gateway_rest_api.id}"
  resource_id = "${aws_api_gateway_resource.roundofbeer_aws_api_gateway_resource_hello.id}"
  http_method = "${aws_api_gateway_method.roundofbeer_aws_api_gateway_method_hello.http_method}"
  status_code = "${aws_api_gateway_method_response.roundofbeer_aws_api_gateway_method_response_500_hello.status_code}"
  selection_pattern = "5\\d{2}"
}