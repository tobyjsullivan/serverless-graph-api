# Serverless GraphQL API

This is a prototype implementation of GraphQL API implemented on serverless architecture on AWS.

## Setup

You'll have to alter the backend configuration in `infra/staging/providers.tf` to something that works for your own AWS account.

```bash
make infra/staging/init
make infra/staging/apply
```

## Running

```bash
curl -X POST "$(cd infra/staging && terraform output api_invoke_url)/query" --data '{ "query": "{ hello }" }'
```

