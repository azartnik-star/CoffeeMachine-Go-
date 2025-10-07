package main
import f "fmt"
type Recipe struct {
	water, milk, beans, price, cups int
}

var espresso = Recipe{water: 250, milk: 0, beans: 16, cups: 1, price: 4}
var latte = Recipe{water: 350, milk: 75, beans: 20, cups: 1, price: 7}
var capp = Recipe{water: 200, milk: 100, beans: 12, cups: 1, price: 6}

type MachineState struct {
	water, milk, beans, money, cups int
}

func buy (stte *MachineState, recipe Recipe) {
	state.water -= recipe.water
	state.milk -= recipe.milk
	state.beans -= recipe.beans
	state.cups -= recipe.cups
	state.money += recipe.price
}

func fill (state *MachineState) {
	var water, milk, beans, cups int
	f.Println("Write how many ml of water you want to add:")
	f.Scan(&water)
	f.Println("Write how many ml of milk you want to add:")
	f.Scan(&milk)
	f.Println("Write how many grams of coffee beans you want to add:")
	f.Scan(&beans)
	f.Println("Write how many disposable cups you want to add:")
	f.Scan(&cups)
	state.water += water
	state.milk += milk
	state.beans += beans
	state.cups += cups
}

func take (state *MachineState) {
	f.Printf ("I gave you $%d\n\n", state.money)
	state.money = 0
}

func main () {
	var state = MachineState{water: 400, milk: 540, beans: 120, cups: 9, money: 550}
	printState(state)
	f.Println("\nWrite action (buy, fill, take):")
	var input string
	f.Scan(&input)
	switch input {
	case "buy":
		f.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuchino: ")
		var drink int
		f.Scan(&drink)
		switch drink {
		case 1:
			buy(&state, espresso)
		case 2:
			buy(&state, latte)
		case 3:
			buy(%state, capp)
		}
	case "fill":
		fill(&state)
	case "take":
		take(&state)
	}
	printState(state)
}