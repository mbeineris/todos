package model

import (
	"encoding/binary"
	"encoding/json"
	"github.com/boltdb/bolt"
	"todos/database"
)

func UpdateTodo(bucketName string, TodoListID int, dataStruct Todo) error {
	// Handle DB changes
	err := database.DBCon.Update(func(tx *bolt.Tx) error {

		// Get TodoBucket instance
		b := tx.Bucket([]byte(bucketName))

		// Get TodoList JSON by id from TodoBucket
		v := b.Get([]byte(itob(TodoListID)))

		var todo_data_item TodoList

		// Unserialize JSON to Tododata instance
		if err := json.Unmarshal(v, &todo_data_item); err != nil {
			panic(err)
		}

		// Find next available id for autoincrementing
		if len(todo_data_item.Todos) > 0 {
			dataStruct.ID = todo_data_item.Todos[len(todo_data_item.Todos)-1].ID + 1
		} else {
			dataStruct.ID = 0
		}

		// Append Todo into Todos of the TodoList instance
		todo_data_item.Todos = append(todo_data_item.Todos, dataStruct)

		// Serialize TodoList instance
		buf, err := json.Marshal(todo_data_item)
		if err != nil {
			return err
		}

		// Update TodoList record in DB
		return b.Put(itob(todo_data_item.ID), buf)
	})

	return err
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
