package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/stephnr/dorm"
)

type Todo struct {
	ID          string `dorm:"HASH"`
	UserID      string `dorm:"RANGE" dormav:"user_id"`
	Todo        string
	Description string `dormav:"-"`
}

func main() {
	awsConfig := &aws.Config{Region: aws.String("eu-west-1")}

	// table error?
	todoList := dorm.LoadTable(dorm.TableLoadOptions{
		AwsConfig: awsConfig,
		TableName: "todo-list",
		Type:      Todo{},
		Options: &dorm.TableOptions{
			ForceCreate: true,
		},
	})

	todoList.Scan( /* ... */ )
	todoList.Query( /* ... */ )
	todoList.Get( /* ... */ )
	todoList.Reload( /* ... */ )
	todoList.Put( /* ... */ )
	todoList.Upsert( /* ... */ )
	todoList.Update( /* ... */ )
	todoList.Save( /* ... */ )
	todoList.Delete( /* ... */ )
}
