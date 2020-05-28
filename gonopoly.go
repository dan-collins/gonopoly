package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gookit/color"
)

type property struct {
	Name  string
	Price int
	Color string
}

// CardCount Change this if you want to start with less properties by default
const CardCount = 4

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Enjoy!")
}

func run() (err error) {
	fmt.Print("How many players?:")
	var playersNum int
	_, err = fmt.Scanf("%d", &playersNum)
	if err != nil {
		return err
	}

	props := make([]property, 0, 28)
	props = append(props, property{Name: "Mediterranean Ave.", Price: 60, Color: "mgb"},
		property{Name: "Baltic Ave.", Price: 60, Color: "mgb"},
		property{Name: "Oriental Ave.", Price: 100, Color: "cyan"},
		property{Name: "Vermont Ave.", Price: 100, Color: "cyan"},
		property{Name: "Connecticut Ave.", Price: 120, Color: "cyan"},
		property{Name: "St. Charles Place", Price: 140, Color: "magenta"},
		property{Name: "States Ave.", Price: 140, Color: "magenta"},
		property{Name: "Virginia Ave.", Price: 160, Color: "magenta"},
		property{Name: "St. James Place", Price: 180, Color: "lightRedEx"},
		property{Name: "Tennessee Ave.", Price: 180, Color: "lightRedEx"},
		property{Name: "New York Ave.", Price: 200, Color: "lightRedEx"},
		property{Name: "Kentucky Ave.", Price: 220, Color: "red"},
		property{Name: "Indiana Ave.", Price: 220, Color: "red"},
		property{Name: "Illinois Ave.", Price: 240, Color: "red"},
		property{Name: "Atlantic Ave.", Price: 260, Color: "lightYellow"},
		property{Name: "Ventnor Ave.", Price: 260, Color: "lightYellow"},
		property{Name: "Marvin Gardens", Price: 280, Color: "lightYellow"},
		property{Name: "Pacific Ave.", Price: 300, Color: "green"},
		property{Name: "North Carolina Ave.", Price: 300, Color: "green"},
		property{Name: "Pennsylvania Ave.", Price: 320, Color: "green"},
		property{Name: "Park Place", Price: 350, Color: "blue"},
		property{Name: "Boardwalk", Price: 400, Color: "blue"},
		property{Name: "Electric Company", Price: 150, Color: "gray"},
		property{Name: "Water Works", Price: 150, Color: "gray"},
		property{Name: "Reading Railroad", Price: 200, Color: "darkGray"},
		property{Name: "Pennsylvania Railroad", Price: 200, Color: "darkGray"},
		property{Name: "B. & O. Railroad", Price: 200, Color: "darkGray"},
		property{Name: "Short Line Railroad", Price: 200, Color: "darkGray"})

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	for i := 1; i <= playersNum; i++ {
		cost := 0
		color.Notice.Println("Player", i, "starting cards:")
		for c := 0; c < CardCount; c++ {
			randomIndex := rand.Intn(len(props))
			prop := props[randomIndex]
			cost += prop.Price
			color.Printf("<%s>%s</> - Price: %d\n", prop.Color, prop.Name, prop.Price)
			// Remove the element at index i from a.
			props[randomIndex] = props[len(props)-1] // Copy last element to index i.
			props[len(props)-1] = property{}         // Erase last element (write zero value).
			props = props[:len(props)-1]             // Truncate slice.
		}
		color.Notice.Println("Player", i, "total cost: $", cost, "\n")
	}
	return nil
}
