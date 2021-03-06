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
	Player string
	Props  []property
	Total  int
}

// CardCount Change this if you want to start with less properties by default
const CardCount = 4

// Randomize returns random cards based on input
func Randomize(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	playerCount, err := strconv.Atoi(params.Get("playerCount"))
	if err != nil || playerCount > 8 || playerCount <= 0 {
		http.Error(w, "Bad Player Count", 422)
		return
	}
	utilities, err := strconv.ParseBool(params.Get("utilities"))
	if err != nil {
		http.Error(w, "Utilities parameter must be true or false", 422)
		return
	}
	railroads, err := strconv.ParseBool(params.Get("railroads"))
	if err != nil {
		http.Error(w, "Railroads parameter must be true or false", 422)
		return
	}
	maxCards := 22
	if utilities {
		maxCards += 2
	}
	if railroads {
		maxCards += 4
	}
	cardCount, err := strconv.Atoi(params.Get("cardCount"))
	if err != nil || cardCount <= 0 || cardCount*playerCount > maxCards {
		http.Error(w, "Bad Card Count", 422)
		return
	}

	cardSet, err := getCards(playerCount, cardCount, utilities, railroads)
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
	cardSet, err := getCards(4, 4, true, true)
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

func getCards(playersNum int, cardCount int, utilities bool, railroads bool) (cards []*cardSet, err error) {
	props := createPropertySlice(utilities, railroads)
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

	for i := 1; i <= playersNum; i++ {
		var cardGroup cardSet
		cardGroup.Player = fmt.Sprintf("Player %d", i)
		for c := 0; c < cardCount; c++ {
			randomIndex := rand.Intn(len(props))
			cardGroup.AddCard(props[randomIndex])
			// Remove the element at index i from a.
			props[randomIndex] = props[len(props)-1] // Copy last element to index i.
			props[len(props)-1] = property{}         // Erase last element (write zero value).
			props = props[:len(props)-1]             // Truncate slice.
		}
		cards = append(cards, &cardGroup)
	}
	return cards, nil
}

func createPropertySlice(utilities bool, railroads bool) (props []property) {

	props = append(props, property{Name: "Mediterranean Ave.", Price: 60, Color: "#864C38"},
		property{Name: "Baltic Ave.", Price: 60, Color: "#864C38"},
		property{Name: "Oriental Ave.", Price: 100, Color: "#AADDF2"},
		property{Name: "Vermont Ave.", Price: 100, Color: "#AADDF2"},
		property{Name: "Connecticut Ave.", Price: 120, Color: "#AADDF2"},
		property{Name: "St. Charles Place", Price: 140, Color: "#C53884"},
		property{Name: "States Ave.", Price: 140, Color: "#C53884"},
		property{Name: "Virginia Ave.", Price: 160, Color: "#C53884"},
		property{Name: "St. James Place", Price: 180, Color: "#EC8B2C"},
		property{Name: "Tennessee Ave.", Price: 180, Color: "#EC8B2C"},
		property{Name: "New York Ave.", Price: 200, Color: "#EC8B2C"},
		property{Name: "Kentucky Ave.", Price: 220, Color: "#DB2428"},
		property{Name: "Indiana Ave.", Price: 220, Color: "#DB2428"},
		property{Name: "Illinois Ave.", Price: 240, Color: "#DB2428"},
		property{Name: "Atlantic Ave.", Price: 260, Color: "#FFF004"},
		property{Name: "Ventnor Ave.", Price: 260, Color: "#FFF004"},
		property{Name: "Marvin Gardens", Price: 280, Color: "#FFF004"},
		property{Name: "Pacific Ave.", Price: 300, Color: "#13A857"},
		property{Name: "North Carolina Ave.", Price: 300, Color: "#13A857"},
		property{Name: "Pennsylvania Ave.", Price: 320, Color: "#13A857"},
		property{Name: "Park Place", Price: 350, Color: "#0066A4"},
		property{Name: "Boardwalk", Price: 400, Color: "#0066A4"})
	if utilities {
		props = append(props, property{Name: "Electric Company", Price: 150, Color: "#a5a5a5"},
			property{Name: "Water Works", Price: 150, Color: "#a5a5a5"})
	}
	if railroads {
		props = append(props, property{Name: "Reading Railroad", Price: 200, Color: "#6b6b6b"},
			property{Name: "Pennsylvania Railroad", Price: 200, Color: "#6b6b6b"},
			property{Name: "B. & O. Railroad", Price: 200, Color: "#6b6b6b"},
			property{Name: "Short Line Railroad", Price: 200, Color: "#6b6b6b"})
	}
	return
}
