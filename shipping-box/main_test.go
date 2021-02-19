package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShippingBox(t *testing.T) {
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
	result := getBestBox(availableBoxes, products)

	assert.NotNil(t, result)
	assert.EqualValues(t, 25, result.Length)
	assert.EqualValues(t, 15, result.Width)
	assert.EqualValues(t, 20, result.Height)

}
