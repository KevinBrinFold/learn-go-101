package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Name      string
	Age       uint8
	Class     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {

	dsn := "root:MuhF23#@!!@tcp(127.0.0.1:3306)/dd_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("db connection error")
	}

	fmt.Println("success connect to database")

	db.AutoMigrate(&User{})

	// data1 := []User{
	// 	{
	// 		Name: "Budi Fredie",
	// 		Age: 9,
	// 		Class: "V",
	// 	},
	// 	{
	// 		Name: "Boni Angkasa",
	// 		Age: 8,
	// 		Class: "V",
	// 	},
	// 	{
	// 		Name: "Buzz Budiman",
	// 		Age: 7,
	// 		Class: "IV",
	// 	},
	// 	{
	// 		Name: "Budi Baikhati",
	// 		Age: 11,
	// 		Class: "VI",
	// 	},
	// 	{
	// 		Name: "Budi Beraksi",
	// 		Age: 12,
	// 		Class: "VI",
	// 	},
	// 	{
	// 		Name: "Steven Brin",
	// 		Age: 17,
	// 		Class: "XI",
	// 	},
	// 	{
	// 		Name: "Fumi Redford",
	// 		Age: 17,
	// 		Class: "XI",
	// 	},
	// }

	// db.Create(&data1)

	data := []User{}

	// db.Raw("SELECT name, age, class FROM users WHERE age > ?", "10").Scan(&data)
	db.Raw("SELECT name, age, class FROM users WHERE class = ?", "XI").Scan(&data)

	for _, i := range data {

		fmt.Println("Nama :", i.Name)
		fmt.Println("Umur :", i.Age)
		fmt.Println("Kelas :", i.Class)
	}

}