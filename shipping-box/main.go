package main

import "fmt"

//Product struct
type Product struct {
	Name   string
	Length int //millimeters
	Width  int
	Height int
}

//Box struct
type Box struct {
	Length int
	Width  int
	Height int
}

/*
	Use and retrieve the best box based on the products you have on your cart
	So the problem is haw you can organize the boxes of the products to use the smaller possible box available
	TODO: Use concurrency!
*/

var (
	totalLength, totalWidth, totalHeight int = 0, 0, 0
	bestBox                              Box
)

func getBestBox(availableBoxes []Box, products []Product) Box {
	//TODO: Complete!
	for _, product := range products {
		totalLength += product.Length
		totalWidth += product.Width
		totalHeight += product.Height
	}
	for _, box := range availableBoxes {
		if box.Length >= totalLength && box.Width >= totalWidth && box.Height >= totalHeight {
			bestBox = box
			break
		}
	}
	return bestBox
}

func main() {
	products := []Product{
		{
			Name:   "Product1",
			Length: 5,
			Width:  2,
			Height: 2,
		},
		{
			Name:   "Product2",
			Length: 8,
			Width:  5,
			Height: 5,
		},
		{
			Name:   "Product3",
			Length: 6,
			Width:  5,
			Height: 5,
		},
	}
	availableBoxes := []Box{
		{
			Length: 20,
			Width:  10,
			Height: 10,
		},
		{
			Length: 20,
			Width:  10,
			Height: 20,
		},
		{
			Length: 25,
			Width:  15,
			Height: 20,
		},
		{
			Length: 50,
			Width:  15,
			Height: 20,
		},
		{
			Length: 60,
			Width:  30,
			Height: 30,
		},
	}

	fmt.Println(getBestBox(availableBoxes, products))

}
