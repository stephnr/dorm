package dorm_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/stephnr/dorm"
)

type Todo struct {
	dorm.Table
	ID          string `dorm:"HASH"`
	UserID      int64  `dorm:"RANGE" dormav:"user_ID"`
	Todo        string
	Description string `dormav:"-"`
}

func TestEntity_Key(t *testing.T) {
	awsConfig := &aws.Config{Region: aws.String("eu-west-1")}

	db := dorm.Open(awsConfig)

	db.LoadTable("todo-list", Todo{}, &dorm.TableOptions{
		ForceCreate: true,
	})

	// item := &TodoItem{
	// 	ID:     "123",
	// 	UserID: "456",
	// }

	// dorm.Key(item)
}
