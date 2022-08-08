#/bin/bash

if [[ -z "${TF_BACKEND_BUCKET}" ]]; then
  echo "Missing env=TF_BACKEND_BUCKET"
  exit 1;
else
  TF_BACKEND_BUCKET="${TF_BACKEND_BUCKET}"
fi

if [[ -z "${TF_BACKEND_KEY}" ]]; then
  echo "Missing env=TF_BACKEND_KEY"
  exit 1;
else
  TF_BACKEND_KEY="${TF_BACKEND_KEY}"
fi

if [[ -z "${TF_BACKEND_REGION}" ]]; then
  TF_BACKEND_REGION="us-east-1"
else
  TF_BACKEND_REGION="${TF_BACKEND_REGION}"
fi

touch terraform-override.tf

echo "
terraform {
  backend "s3" {}
}
" > terraform-override.tf

terraform init -backend-config="bucket=$TF_BACKEND_BUCKET" \
   -backend-config="key=$TF_BACKEND_KEY" \
   -backend-config="region=$TF_BACKEND_REGION"