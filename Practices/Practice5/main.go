package main

import "fmt"

func main() {
	db := ConnectDB()
	defer db.Close()

	fmt.Println("✅ Connected to PostgreSQL!")

	users, _ := GetAllUsers(db)
	fmt.Println("Before transfer:", users)

	err := TransferBalance(db, 1, 2, 100)
	if err != nil {
		fmt.Println("❌ Transfer error:", err)
	} else {
		fmt.Println("💸 Transfer success!")
	}

	users, _ = GetAllUsers(db)
	fmt.Println("After transfer:", users)
}
