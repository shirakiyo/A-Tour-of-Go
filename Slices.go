package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			pic[y] = append(pic[y], uint8(y^x))
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
