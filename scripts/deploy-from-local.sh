#!/usr/bin/env bash

# must run from project root directory
cd infra || exit
cdk synth || exit
cdk bootstrap || exit
cdk deploy || exit