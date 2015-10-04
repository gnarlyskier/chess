package core

import (
	"unicode"
)

type Player interface {
	RenderIcon(icon rune) rune
}

type white string
type black string

var (
	// Players
	WHITE white = "white"
	BLACK black = "black"
)

func (_ white) RenderIcon(icon rune) rune {
	return unicode.ToLower(icon)
}

func (_ black) RenderIcon(icon rune) rune {
	return unicode.ToUpper(icon)
}
