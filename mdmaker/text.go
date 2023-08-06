package mdmaker

import (
	"strings"
)

// create markdown header

// # Title
// ## Subtitle
// ### Header 1
// #### Header 2
// ##### Header 3
// ###### Header 4

func Header(setting int, text string) string {
	return Break + strings.Repeat("#", setting) + " " + text + Break
}

func Link(text string, url string) string {
	return Break + "[" + text + "](" + url + ")" + Break
}

func Image(alt string, src string) string {
	return Break + "![" + alt + "](" + src + ")" + Break
}

func Text(text string, style int) string {
	switch style {
	case Paragraph:
		return Break + text + Break
	case Bold:
		return Break + "**" + text + "**" + Break
	case Italic:
		return Break + "_" + text + "_" + Break
	case Strike:
		return Break + "~~" + text + "~~" + Break
	default:
		return Break + text + Break
	}
}
