package mdmaker

import (
	"log"
)

func Test() {
	
	list := []string{"1st", "2nd", "3rd"}
	
	data := [][]string{
		{"Name", "Age", "City"},
		{"John Doe", "23", "New York"},
		{"Jane Doe", "27", "San Francisco"},
	}
	
	text := Header(Title, "Wasmachensachen")
	text += Header(Sub, "Hi")
	text += Header(Header1, "Header 1")
	text += Line
	text += Text("dayumm", Strike)
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
