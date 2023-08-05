package maker

import (
	"strconv"
)

func UnorderedList(text []string) string {
	list := Break

	for _, t := range text {
		list += "* " + t + Break
	}

	return list
}

func OrderedList(text []string) string {
	list := Break

	for i, t := range text {
		num := strconv.Itoa(i + 1)
		list += num + ". " + t + Break
	}
	return list
}
