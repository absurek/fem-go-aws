package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FemGoAwsStackProps struct {
	awscdk.StackProps
}

func NewFemGoAwsStack(scope constructs.Construct, id string, props *FemGoAwsStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	table := awsdynamodb.NewTable(stack, jsii.String("MyUserTable"), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("username"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		TableName:     jsii.String("UserTable"),
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})

	lambda := awslambda.NewFunction(stack, jsii.String("MyLambdaFunction"), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Code:         awslambda.AssetCode_FromAsset(jsii.String("../dist/lambda"), nil),
		Handler:      jsii.String("handler"),
		Architecture: awslambda.Architecture_ARM_64(),
	})

	table.GrantReadWriteData(lambda)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewFemGoAwsStack(app, "FemGoAwsStack", &FemGoAwsStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil
}
