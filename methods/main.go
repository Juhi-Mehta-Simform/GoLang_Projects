package main

func main() {
	var (
		mobydick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 15}
		rubik     = puzzle{title: "rubik's cube", price: 5}
		yoda      = toy{title: "Yoda", price: 30}
	)

	var store list
	store = append(store, &minecraft, &tetris, mobydick, rubik, &yoda)
	store.discount(.5)
	store.print()

	// interface values are comparable
	// fmt.Println(store[0] == &minecraft)
	// fmt.Println(store[3] == mobydick)

	// var items []*game
	// items = append(items, &minecraft, &tetris)

	// my := list(items)
	// my = nil
	// my.print()
	// (minecraft).discount(0.5)

	// mobydick.print()
	// minecraft.print()
	// tetris.print()

	// book.print(mobydick)
	// game.print(minecraft)
	// game.print(tetris)

}
