package main

type TodoItem struct {
	ID          string `dorm:"HASH"`
	UserID      string `dorm:"RANGE" dormav:"user_id"`
	Todo        string
	Description string `dormav:"-"`
}

func main() {
	todoList, err := dorm.Table("todo-lists", TodoItem, &dorm.TableOptions{
		ForceCreate: true,
	})

	todoList.Scan(/* ... */)
	todoList.Query(/* ... */)
	todoList.Get(/* ... */)
	todoList.Reload(/* ... */)
	todoList.Put(/* ... */)
	todoList.Upsert(/* ... */)
	todoList.Update(/* ... */)
	todoList.Save(/* ... */)
	todoList.Delete(/* ... */)
}
