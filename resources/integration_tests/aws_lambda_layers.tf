
//dont know how to add layer version policy
resource "aws_lambda_layer_version" "lambda_layer" {
  filename = data.archive_file.lambda_zip_inline.output_path
  layer_name = "lambda_layer${var.test_prefix}${var.test_suffix}"

  compatible_runtimes = [
    "nodejs12.x"]
}

data "archive_file" "lambda_zip_inline" {
  type = "zip"
  output_path = "./tmp/lambda_zip_inline.zip"
  source {
    content = <<EOF
module.exports.handler = async (event, context, callback) => {
	const what = "world";
	const response = `Hello $${what}!`;
	callback(null, response);
};
EOF
    filename = "main.js"
  }
}
