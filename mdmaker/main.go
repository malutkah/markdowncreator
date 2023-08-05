package main

import (
	"log"

	"github.com/malutkah/markdowncreator/maker"
	"github.com/malutkah/markdowncreator/settings"
)

func main() {

	list := []string{"1st", "2nd", "3rd"}

	data := [][]string{
		{"Name", "Age", "City"},
		{"John Doe", "23", "New York"},
		{"Jane Doe", "27", "San Francisco"},
	}

	text := maker.Header(settings.Title, "Wasmachensachen")
	text += maker.Header(settings.Sub, "Hi")
	text += maker.Header(settings.Header1, "Header 1")
	text += maker.Line
	text += maker.Text("dayumm", settings.Strike)
	text += maker.Line
	text += maker.OrderedList(list)
	text += maker.UnorderedList([]string{"some", "stupid", "shit"})
	text += maker.Link("Some Link", "https://google.com")
	text += maker.Line
	text += maker.Image("some pic", "1.jpeg")
	text += maker.Line + maker.Break
	text += maker.Table(data)

	err := maker.CreateMarkdown(text, "new")
	if err != nil {
		log.Fatalln(err)
	}
}
