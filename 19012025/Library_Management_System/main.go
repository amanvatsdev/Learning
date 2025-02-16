package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Library struct {
	Id     int
	Title  string
	Author string
	Genre  string
}

func main() {
	libraryData := map[string]Library{
		"George": {
			Id:     12,
			Title:  "1984",
			Author: "George Orwell",
			Genre:  "Fictionn",
		},
		"Jonr": {
			Id:     25,
			Title:  "Harry Potter and the Philosopher`s Stone",
			Author: "John",
			Genre:  "Fantasy",
		},
		"Rupi": {
			Id:     65,
			Title:  "Milk and Honey",
			Author: "Rupi Kaur",
			Genre:  "Poetry",
		},
	}
	//asking to the user what he wants
	fmt.Println("Choose any one option")
	fmt.Println(`	1.Adding a book
	2.Removing a book
	3.Searching for books by title`)
	//Scanning the outpuy
	var option int
	fmt.Scanln(&option)
	switch option {
	case 1:
		fmt.Println("type the things as asked")

		fmt.Println("Type Name of Author")
		var author string
		fmt.Scan(&author)
		fmt.Println("Type Id number")
		var id int
		fmt.Scan(&id)
		fmt.Println("type the title name")
		var title string
		fmt.Scan(&title)
		fmt.Println("type Genre name ")
		var genre string
		fmt.Scan(&genre)
		libraryData[author] = Library{
			Id:     id,
			Title:  title,
			Author: author,
			Genre:  genre,
		}
	case 2:
		fmt.Println(`		Removing a book 
		Type the name Of Author`)
		var author string
		fmt.Scan(&author)
		authour := strings.TrimSpace(author)
		_, Exist := libraryData[authour]
		if Exist {
			delete(libraryData, authour)
		} else {
			fmt.Println("This Name was not in the list")
		}
	case 3:
		fmt.Println("You have chosen the third option.")
		fmt.Println("Type the title you want to print:")
		title := bufio.NewScanner(os.Stdin)
		if title.Scan() {
			titleName := title.Text()
			fmt.Println("hello", titleName)
		}
		fmt.Println(libraryData)
	}
}
