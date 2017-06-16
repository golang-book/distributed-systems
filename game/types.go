package main

type (
	Color uint32
	Block struct {
		Color Color
	}
)

// Predefined Colors
const (
	ColorBlack Color = 0x000000FF
	ColorWhite Color = 0xFFFFFFFF
)
