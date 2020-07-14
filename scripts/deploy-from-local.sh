#!/usr/bin/env bash

# this was for early development, not recommended to use
# must run from project root directory
cd infra || exit
cdk synth || exit
cdk bootstrap || exit
cdk deploy --require-approval never || exit