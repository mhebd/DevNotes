package main

import (
	"fmt"
	"net/http"
	"os"

	"devnotes.com/db"
	"devnotes.com/router"
)

func main() {
	db.Start()

	fmt.Printf("Server is running on http://127.0.0.1:%v \n", os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), router.Router())
}
