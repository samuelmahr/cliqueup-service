#!/usr/bin/env bash

# this was for early development, not recommended to use
# must run from project root directory
go test -count=1 -p=1 ./...

cd infra || exit
cdk synth || exit
cdk bootstrap || exit
cdk deploy --require-approval never || exit

migrate -database "${CLIQUEUP_POSTGRESQL_URL}" -path migrations up