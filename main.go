package main

// import (
// 	. "./models"
// )

func main() {
	go Connect("config.json")

	ListenAndServe(8080)
}
