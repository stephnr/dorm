package dorm_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/stephnr/dorm"
)

type Todo struct {
	dorm.Table
	ID          string `dorm:"HASH"`
	UserID      string `dorm:"RANGE" dormav:"user_ID"`
	Todo        string
	Description string `dormav:"-"`
}

func TestEntity_Key(t *testing.T) {
	awsConfig := &aws.Config{Region: aws.String("eu-west-1")}

	dorm.LoadTable(dorm.TableLoadOptions{
		AwsConfig: awsConfig,
		TableName: "todo-list",
		Type:      Todo{},
		Options: &dorm.TableOptions{
			ForceCreate: true,
		},
	})

	// item := &TodoItem{
	// 	ID:     "123",
	// 	UserID: "456",
	// }

	// dorm.Key(item)
}
