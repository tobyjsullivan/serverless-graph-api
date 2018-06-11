.PHONY: build package infra/staging/apply infra/staging/init infra/staging/destroy infra/format

build/lambda_handler: lambda_handler.go
	GOOS=linux GOARCH=amd64 go build -o build/lambda_handler lambda_handler.go

build: build/lambda_handler

build/package.zip: build/lambda_handler
	zip -j build/package.zip build/lambda_handler

package: build/package.zip

infra/staging/.terraform/terraform.tfstate: infra/staging/providers.tf
	cd infra/staging && terraform init

infra/staging/init: infra/staging/.terraform/terraform.tfstate
	terraform fmt infra/

infra/staging/apply: infra/staging/init build/package.zip infra/common/main.tf infra/staging/main.tf
	terraform fmt infra/
	cd infra/staging && terraform apply -var 'lambda_package=../../build/package.zip'

infra/staging/destroy: infra/staging/init
	terraform fmt infra/
	cd infra/staging && terraform destroy -var 'lambda_package=../../build/package.zip'
