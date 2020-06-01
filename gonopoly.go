package gonopoly

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type property struct {
	Name  string
	Price int
	Color string
}

type cardSet struct {
	Props []property
	Total int
}

// CardCount Change this if you want to start with less properties by default
const CardCount = 4

// Randomize returns random cards based on input
func Randomize(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	cardCount, cErr := strconv.Atoi(params.Get("cardCount"))
	playerCount, pErr := strconv.Atoi(params.Get("playerCount"))

	if cErr != nil || pErr != nil || cardCount <= 0 || playerCount <= 0 {
		http.Error(w, "Bad Card/Player Count", 422)
		return
	}
	cardSet, err := run(playerCount, cardCount)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	js, err := json.Marshal(cardSet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	cardSet, err := run(4, 4)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Enjoy!", cardSet)
}

func (c *cardSet) AddCard(p property) {
	c.Props = append(c.Props, p)
	c.Total += p.Price
}

func run(playersNum int, cardCount int) (cards map[int]*cardSet, err error) {
	cards = make(map[int]*cardSet)
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
		var cardGroup cardSet
		(cards)[i] = &cardGroup
		for c := 0; c < cardCount; c++ {
			randomIndex := rand.Intn(len(props))
			(cards)[i].AddCard(props[randomIndex])
			// Remove the element at index i from a.
			props[randomIndex] = props[len(props)-1] // Copy last element to index i.
			props[len(props)-1] = property{}         // Erase last element (write zero value).
			props = props[:len(props)-1]             // Truncate slice.
		}
	}
	return cards, nil
}
