package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/ajstarks/svgo"
)

const (
	width       = 500
	height      = 500
	circleX     = width / 2
	circleY     = height / 2
	circleR     = 200
	numRed      = 5  // Warehouses
	numYellow   = 10 // Factories
	numBlue     = 15 // Trucks/Ships/Airplanes
	dotRadius   = 5
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create SVG file
	file, err := os.Create("map.svg")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	svg := svg.New(file)
	svg.Start(width, height)

	// Draw background circle
	svg.Circle(circleX, circleY, circleR, "fill:white;stroke:black;stroke-width:2")

	// Place dots
	placeDots(svg, numRed, "red")
	placeDots(svg, numYellow, "yellow")
	placeDots(svg, numBlue, "blue")

	svg.End()
	fmt.Println("SVG file 'map.svg' created successfully!")
}

func placeDots(s *svg.SVG, count int, color string) {
	for i := 0; i < count; i++ {
		// Generate a random point inside the circle
		for {
			x := rand.Intn(2*circleR) + (circleX - circleR)
			y := rand.Intn(2*circleR) + (circleY - circleR)

			// Ensure point is within the circle
			if (x-circleX)*(x-circleX)+(y-circleY)*(y-circleY) <= circleR*circleR {
				s.Circle(x, y, dotRadius, fmt.Sprintf("fill:%s", color))
				break
			}
		}
	}
}
