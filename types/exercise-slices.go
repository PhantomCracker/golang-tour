package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	mySlice := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		mySlice[i] = make([]uint8, dx)
	}
	return mySlice
}

func main() {
	pic.Show(Pic)
}