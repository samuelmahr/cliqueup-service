import * as cdk from '@aws-cdk/core';
import { DockerImageAsset } from '@aws-cdk/aws-ecr-assets';
import * as ec2 from '@aws-cdk/aws-ec2';
import * as ecs from '@aws-cdk/aws-ecs';
import * as ecs_patterns from '@aws-cdk/aws-ecs-patterns'
import * as path from "path";
import * as cognito from '@aws-cdk/aws-cognito'
import * as rds from '@aws-cdk/aws-rds'

/**
 * The port range to open up for dynamic port mapping
 */
const EPHEMERAL_PORT_RANGE = ec2.Port.tcpRange(32768, 65535);

export class InfraStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    // @ts-ignore
    const cliqueupServiceImageAsset = new DockerImageAsset(this, 'cliqueup-service-api-image', {
      directory: path.join(__dirname, '../../'),
      repositoryName: 'cliqueup-service-api'
    });

    // For better iteration speed, it might make sense to put this VPC into
    // a separate stack and import it here. We then have two stacks to
    // deploy, but VPC creation is slow so we'll only have to do that once
    // and can iterate quickly on consuming stacks. Not doing that for now.
    const vpc = new ec2.Vpc(this, 'MyVpc', { maxAzs: 2 });

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
        image: ecs.ContainerImage.fromDockerImageAsset(cliqueupServiceImageAsset),
        containerPort: 8000
      },
      maxHealthyPercent: 200,
      minHealthyPercent: 0
    });

    // ephemeral port range (when host port is 0).
    ecsService.service.connections.allowFromAnyIpv4(EPHEMERAL_PORT_RANGE);

    const cliqueUpUserPool = new cognito.UserPool(this, 'cliqueup-user-pool', {
      userPoolName: 'cliqueup-user-pool',
      selfSignUpEnabled: true,
      userVerification: {
        emailSubject: 'Verify your email for CliqueUp!',
        emailBody: 'Hello {username}, Thanks for signing up to CliqueUp! Your verification code is {####}',
        emailStyle: cognito.VerificationEmailStyle.CODE,
        smsMessage: 'Hello {username}, Thanks for signing up to our CliqueUp! Your verification code is {####}',
      },
      signInAliases: {
        phone: true,
        email: true
      },
      autoVerify: {
        email: true,
        phone: true
      },
      standardAttributes: {
        email: {
          required: true,
          mutable: false,
        },
        phoneNumber: {
          required: false,
          mutable: true,
        },
      },
      passwordPolicy: {
        minLength: 12,
        requireLowercase: true,
        requireUppercase: true,
        requireDigits: true,
        requireSymbols: false,
        tempPasswordValidity: cdk.Duration.days(3),
      },
      accountRecovery: cognito.AccountRecovery.EMAIL_AND_PHONE_WITHOUT_MFA,
    });

    const pgUsername = process.env.POSTGRES_USER || "master";
    const pgPassword = process.env.POSTGRES_PASS || "Passw0rd";
    const dbName = process.env.POSTGRES_DATABASE || "cliqueup";

    // @ts-ignore
    const cliqueupDbInstance = new rds.DatabaseInstance(this, 'Instance', {
      engine: rds.DatabaseInstanceEngine.POSTGRES,
      engineVersion: '11.1',
      instanceType: ec2.InstanceType.of(ec2.InstanceClass.BURSTABLE3, ec2.InstanceSize.MICRO),
      masterUsername: pgUsername,
      masterUserPassword: cdk.SecretValue.plainText(pgPassword),
      databaseName: dbName,
      vpc,
      vpcPlacement: { subnetType: ec2.SubnetType.PUBLIC }
    });

    cliqueupDbInstance.connections.allowDefaultPortFromAnyIpv4()
  }
}
