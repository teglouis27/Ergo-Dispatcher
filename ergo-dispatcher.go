package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/ajstarks/svgo"
)

const (
	width       = 600
	height      = 600
	numRed      = 5  // Warehouses
	numYellow   = 10 // Factories
	numBlue     = 7  // Trucks/Ships/Airplanes
	circleR     = 250
	centerX     = width / 2
	centerY     = height / 2
	dotRadius   = 6
	moveSpeed   = 3
	steps       = 100 // Number of movement steps
)

type Entity struct {
	X, Y   int
	Color  string
	Items  []string
	IsMoving bool
}

func randomItems() []string {
	items := []string{"Iron", "Copper", "Gold", "Textiles", "Electronics"}
	n := rand.Intn(3) + 1
	var selected []string
	for i := 0; i < n; i++ {
		selected = append(selected, items[rand.Intn(len(items))])
	}
	return selected
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create SVG file
	file, err := os.Create("ergo_dispatcher.svg")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	svg := svg.New(file)
	svg.Start(width, height)
	svg.Circle(centerX, centerY, circleR, "fill:white;stroke:black;stroke-width:2")

	// Initialize entities
	warehouses := placeEntities(numRed, "red", false)
	factories := placeEntities(numYellow, "yellow", false)
	vehicles := placeEntities(numBlue, "blue", true)

	// Draw static entities
	for _, e := range append(warehouses, factories...) {
		drawEntity(svg, e)
	}

	// Move vehicles
	for step := 0; step < steps; step++ {
		for i := range vehicles {
			moveEntity(&vehicles[i])
			for _, e := range append(warehouses, factories...) {
				if checkCollision(vehicles[i], e) {
					exchangeItems(&vehicles[i], &e)
				}
			}
		}
	}

	// Draw moving entities
	for _, e := range vehicles {
		drawEntity(svg, e)
	}

	svg.End()
	fmt.Println("SVG file 'ergo_dispatcher.svg' created successfully!")
}

func placeEntities(count int, color string, isMoving bool) []Entity {
	var entities []Entity
	for i := 0; i < count; i++ {
		for {
			x := rand.Intn(2*circleR) + (centerX - circleR)
			y := rand.Intn(2*circleR) + (centerY - circleR)
			if (x-centerX)*(x-centerX)+(y-centerY)*(y-centerY) <= circleR*circleR {
				entities = append(entities, Entity{X: x, Y: y, Color: color, Items: randomItems(), IsMoving: isMoving})
				break
			}
		}
	}
	return entities
}

func drawEntity(svg *svg.SVG, e Entity) {
	tooltip := fmt.Sprintf("Items: %v", e.Items)
	svg.Circle(e.X, e.Y, dotRadius, fmt.Sprintf("fill:%s", e.Color))
	svg.Title(tooltip)
}

func moveEntity(e *Entity) {
	if e.IsMoving {
		e.X += rand.Intn(2*moveSpeed) - moveSpeed
		e.Y += rand.Intn(2*moveSpeed) - moveSpeed
	}
}

func checkCollision(a, b Entity) bool {
	dist := (a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y)
	return dist < (2 * dotRadius * 2 * dotRadius)
}

func exchangeItems(a, b *Entity) {
	if len(a.Items) > 0 && len(b.Items) > 0 {
		idxA, idxB := rand.Intn(len(a.Items)), rand.Intn(len(b.Items))
		a.Items[idxA], b.Items[idxB] = b.Items[idxB], a.Items[idxA]
	}
}
