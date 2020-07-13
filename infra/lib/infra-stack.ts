import * as cdk from '@aws-cdk/core';
import * as ecr from '@aws-cdk/aws-ecr';
import { DockerImageAsset } from '@aws-cdk/aws-ecr-assets';
import * as ec2 from '@aws-cdk/aws-ec2';
import * as ecs from '@aws-cdk/aws-ecs';
import * as ecs_patterns from '@aws-cdk/aws-ecs-patterns'
import * as path from "path";

/**
 * The port range to open up for dynamic port mapping
 */
const EPHEMERAL_PORT_RANGE = ec2.Port.tcpRange(32768, 65535);

export class InfraStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    // @ts-ignore
    const cliqueupApiRepository = new ecr.Repository(this, 'cliqueup-service-api');

    // @ts-ignore
    const asset = new DockerImageAsset(this, 'cliqueup-service-api-image', {
      directory: path.join(__dirname, '../../')
    });

    // For better iteration speed, it might make sense to put this VPC into
    // a separate stack and import it here. We then have two stacks to
    // deploy, but VPC creation is slow so we'll only have to do that once
    // and can iterate quickly on consuming stacks. Not doing that for now.
    /*const vpc = new ec2.Vpc(this, 'MyVpc', { maxAzs: 2 });

    // @ts-ignore
    const cluster = new ecs.Cluster(this, 'Ec2Cluster', { vpc });
    cluster.addCapacity('DefaultAutoScalingGroup', {
      instanceType: ec2.InstanceType.of(ec2.InstanceClass.T2, ec2.InstanceSize.MICRO)
    });

    // Instantiate ECS Service with just cluster and image
    // @ts-ignore
    const ecsService = new ecs_patterns.NetworkLoadBalancedEc2Service(this, "Ec2Service", {
      cluster,
      memoryLimitMiB: 512,
      taskImageOptions: {
        image: ecs.ContainerImage.fromRegistry(cliqueupApiRepository.repositoryName),
      }
    });

    // Need target security group to allow all inbound traffic for
    // ephemeral port range (when host port is 0).
    ecsService.service.connections.allowFromAnyIpv4(EPHEMERAL_PORT_RANGE);*/
  }
}
