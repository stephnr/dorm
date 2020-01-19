package dorm

import "github.com/aws/aws-sdk-go/aws"

// DB contains current information about the AWS connection
type DB struct {
	// AWS Config details
	awsConfig *aws.Config
}

func Open(cfg *aws.Config) *DB {
	return &DB{awsConfig: cfg}
}
