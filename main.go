package main

import (
	"github.com/novitaekasari/JWToken/database"
	"github.com/novitaekasari/JWToken/router"
)

// func main() {
// 	database.StartDB()
// 	r := router.StartApp()
// 	r.Run(":8080")
// }

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}