package basics

import "fmt"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR STRUCTS HERE

// item struktúra: alap adatok
type item struct {
	id    int
	name  string
	price int
}

// game struktúra: beágyazza az item-et (nincs neve a mezőnek, csak a típusa!)
type game struct {
	item
	genre string
}

// newGame returns a new game struct.
func newGame(id int, name string, price int, genre string) game {
	return game{
		item:  item{id: id, name: name, price: price},
		genre: genre,
	}
}

// String stringifies an item.
// Formátum: <id>: <name> costs <price>
func (i item) String() string {
	return fmt.Sprintf("%d: %s costs %d", i.id, i.name, i.price)
}

// String stringifies a game.
// Formátum: Game <id>: <name> costs <price> of genre <genre>
// Mivel az item be van ágyazva, közvetlenül elérjük a g.id-t, g.name-et!
func (g game) String() string {
	return fmt.Sprintf("Game %d: %s costs %d of genre %s", g.id, g.name, g.price, g.genre)
}

// newGameList creates a game store.
// Adatok: (1, god of war, 50, action adventure), (3, minecraft, 20, sandbox), (4, warcraft, 40, strategy)
func newGameList() []game {
	return []game{
		newGame(1, "god of war", 50, "action adventure"),
		newGame(3, "minecraft", 20, "sandbox"),
		newGame(4, "warcraft", 40, "strategy"),
	}
}

// queryById returns the game in the specified store with the given id or returns a "no such game" error.
func queryById(games []game, id int) (game, error) {
	for _, g := range games {
		if g.id == id {
			return g, nil
		}
	}
	// Ha nem találtuk meg, üres game struct és hiba
	return game{}, fmt.Errorf("no such game")
}

// listNameByPrice returns the name of the game(s) with price equal or smaller than a given price.
func listNameByPrice(games []game, price int) []string {
	var result []string // Üres slice
	for _, g := range games {
		if g.price <= price {
			result = append(result, g.name)
		}
	}
	return result
}
