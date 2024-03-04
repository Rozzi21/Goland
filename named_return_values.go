package main

import "fmt"

func main() {
	title, author, chapter := getCompleteManga()
	fmt.Println(title, author, chapter)
}

func getCompleteManga() (title, author, chapter string) {
	title = "Naruto"
	author = "Mashashi kishimoto"
	chapter = "400"

	return title, author, chapter

}
