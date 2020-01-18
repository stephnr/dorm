package main

import "github.com/stephnr/dorm"

type TodoItem struct {
	dorm.Entity
	ID          string `dorm:"HASH"`
	UserID      string `dorm:"RANGE" dormav:"user_id"`
	Todo        string
	Description string `dormav:"-"`
}

func main() {
	todoList, _ := dorm.LoadTable("todo-lists", TodoItem, &dorm.TableOptions{
		ForceCreate: true,
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
