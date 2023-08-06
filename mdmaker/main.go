package mdmaker

import (
	"log"
	
	"github.com/malutkah/markdowncreator/settings"
)

func main() {
	
	list := []string{"1st", "2nd", "3rd"}
	
	data := [][]string{
		{"Name", "Age", "City"},
		{"John Doe", "23", "New York"},
		{"Jane Doe", "27", "San Francisco"},
	}
	
	text := Header(settings.Title, "Wasmachensachen")
	text += Header(settings.Sub, "Hi")
	text += Header(settings.Header1, "Header 1")
	text += Line
	text += Text("dayumm", settings.Strike)
	text += Line
	text += OrderedList(list)
	text += UnorderedList([]string{"some", "stupid", "shit"})
	text += Link("Some Link", "https://google.com")
	text += Line + Break
	text += Table(data)
	
	err := CreateMarkdown(text, "new")
	if err != nil {
		log.Fatalln(err)
	}
}
