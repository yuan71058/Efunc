package main

import (
	"fmt"
	"time"

	"github.com/yuan71058/Efunc/en/utils"
)

type User struct {
	Id   int64  `xorm:"pk autoincr"`
	Name string `xorm:"varchar(50)"`
	Age  int    `xorm:"int"`
}

func main() {
	fmt.Println("=== Database Demo ===")

	// SQLite (in-memory for demo)
	fmt.Println("\n--- SQLite ---")
	engine, err := utils.Database_ConnectSQLite(":memory:?cache=shared")
	if err != nil {
		fmt.Println("Connect error:", err)
		return
	}
	defer utils.Database_Close(engine)

	utils.Database_Ping(engine)
	fmt.Println("Connected to SQLite")

	// Sync tables
	err = utils.Database_SyncTables(engine, new(User))
	if err != nil {
		fmt.Println("Sync error:", err)
		return
	}
	fmt.Println("Tables synced")

	// Insert
	users := []User{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}
	affected, err := utils.Database_Insert(engine, &users)
	if err != nil {
		fmt.Println("Insert error:", err)
		return
	}
	fmt.Printf("Inserted %d records\n", affected)

	// Query all
	var allUsers []User
	err = utils.Database_Find(engine, &allUsers)
	if err != nil {
		fmt.Println("Find error:", err)
		return
	}
	for _, u := range allUsers {
		fmt.Printf("  User: id=%d name=%s age=%d\n", u.Id, u.Name, u.Age)
	}

	// Query with condition
	var youngUsers []User
	err = utils.Database_FindWhere(engine, &youngUsers, "age < ?", 30)
	if err != nil {
		fmt.Println("FindWhere error:", err)
		return
	}
	fmt.Println("\nYoung users (age < 30):")
	for _, u := range youngUsers {
		fmt.Printf("  User: id=%d name=%s age=%d\n", u.Id, u.Name, u.Age)
	}

	// Count
	count, err := utils.Database_Count(engine, new(User))
	if err != nil {
		fmt.Println("Count error:", err)
		return
	}
	fmt.Printf("\nTotal users: %d\n", count)

	// Update
	affected, err = utils.Database_UpdateWhere(engine, &User{Age: 26}, "name = ?", "Alice")
	if err != nil {
		fmt.Println("Update error:", err)
		return
	}
	fmt.Printf("Updated %d records\n", affected)

	// Get single record
	var alice User
	found, err := utils.Database_Get(engine.Where("name = ?", "Alice"), &alice)
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	if found {
		fmt.Printf("Alice's new age: %d\n", alice.Age)
	}

	// Transaction
	fmt.Println("\n--- Transaction ---")
	result, err := utils.Database_Transaction(engine, func(session interface{}) (interface{}, error) {
		return nil, nil
	})
	_ = result
	if err != nil {
		fmt.Println("Transaction error:", err)
	} else {
		fmt.Println("Transaction completed")
	}

	// BuntDB key-value store
	fmt.Println("\n--- BuntDB ---")
	bdb, err := utils.BuntDB_Open(":memory:")
	if err != nil {
		fmt.Println("BuntDB open error:", err)
		return
	}
	defer utils.BuntDB_Close(bdb)

	utils.BuntDB_Set(bdb, "greeting", "Hello from BuntDB")
	val, _ := utils.BuntDB_Get(bdb, "greeting")
	fmt.Printf("BuntDB value: %s\n", val)

	utils.BuntDB_SetJSON(bdb, "user:1", map[string]interface{}{"name": "Alice", "age": 25})
	jsonVal, _ := utils.BuntDB_Get(bdb, "user:1")
	fmt.Printf("BuntDB JSON: %s\n", jsonVal)

	_ = time.Now

	fmt.Println("\nDone!")
}