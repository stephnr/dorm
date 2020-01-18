package dorm

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (tbl *Table) validateExistence() {
	svc := dynamodb.New(tbl.aws.session)
	forceCreate := false

	if tbl.options != nil && tbl.options.ForceCreate {
		forceCreate = true
	}

	output, err := svc.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(tbl.model.tableName),
	})

	if err != nil {
		awsErr, ok := err.(awserr.Error)

		if ok && awsErr.Code() == dynamodb.ErrCodeResourceNotFoundException && forceCreate {
			tbl.createTable()
			return
		}

		if ok && awsErr.Code() == dynamodb.ErrCodeResourceNotFoundException && !forceCreate {
			panic(fmt.Sprintf("dorm: Failed to find dynamodb table (%s)", tbl.model.tableName))
		}

		panic(fmt.Sprintf("dorm: Unknown failure occurred when validating if dynamodb table exists\n%+v", err))
	}

	for _, keySchema := range output.Table.KeySchema {
		if *keySchema.KeyType == "HASH" && *keySchema.AttributeName != tbl.model.hashParamName {
			panic(fmt.Sprintf("dorm: Mismatched HASH column names. AWS has (%s) but model uses (%s)", *keySchema.AttributeName, tbl.model.hashParamName))
		}

		if *keySchema.KeyType == "RANGE" && *keySchema.AttributeName != tbl.model.rangeParamName {
			panic(fmt.Sprintf("dorm: Mismatched RANGE column names. AWS has (%s) but model uses (%s)", *keySchema.AttributeName, tbl.model.rangeParamName))
		}
	}
}

func (tbl *Table) createTable() {
	svc := dynamodb.New(tbl.aws.session)

	input := &dynamodb.CreateTableInput{
		TableName: aws.String(tbl.model.tableName),
		KeySchema: []*dynamodb.KeySchemaElement{
			&dynamodb.KeySchemaElement{
				AttributeName: aws.String(tbl.model.hashParamName),
				KeyType:       aws.String("HASH"),
			},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			&dynamodb.AttributeDefinition{
				AttributeName: aws.String(tbl.model.hashParamName),
				AttributeType: aws.String("S"),
			},
		},
	}

	if tbl.model.rangeParamName != "" {
		input.KeySchema = append(input.KeySchema, &dynamodb.KeySchemaElement{
			AttributeName: aws.String(tbl.model.rangeParamName),
			KeyType:       aws.String("RANGE"),
		})

		input.AttributeDefinitions = append(input.AttributeDefinitions, &dynamodb.AttributeDefinition{
			AttributeName: aws.String(tbl.model.rangeParamName),
			AttributeType: aws.String("S"),
		})

		if tbl.options.ReadUnits == nil && tbl.options.WriteUnits == nil {
			input.BillingMode = aws.String(dynamodb.BillingModePayPerRequest)
		}

		_, err := svc.CreateTable(input)

		if err != nil {
			panic(fmt.Sprintf("dorm: Failed to force create dynamodb table: (%s)\n%+v", tbl.model.tableName, err))
		}
	}
}
