#!/bin/zsh
##VPCs
aws --region eu-west-1 cloudformation create-stack --stack-name VPCAPI --template-body file://cloudformation/vpc/vpc.yaml --parameters file://cloudformation/vpc/vpc_param_api.json
aws --region eu-west-1 cloudformation update-stack --stack-name VPCAPI --template-body file://cloudformation/vpc/vpc.yaml --parameters file://cloudformation/vpc/vpc_param_api.json
aws --region eu-west-1 cloudformation create-stack --stack-name VPCConsumer --template-body file://cloudformation/vpc/vpc.yaml --parameters file://cloudformation/vpc/vpc_param_consumer.json
aws --region eu-west-1 cloudformation update-stack --stack-name VPCConsumer --template-body file://cloudformation/vpc/vpc.yaml --parameters file://cloudformation/vpc/vpc_param_consumer.json

# 
