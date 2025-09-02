package main

import f "fmt"

const (
	waterPerCupMl = 200
	milkPerCupMl  = 50
	coffeePerCupG = 15
)

func minN(nums ...int) int {
	// нахождение минимального числа
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]
	for _, v := range nums[1:] {
		if min > v {
			min = v
		}
	}
	return min
}

func howManyCups(w, m, c int) int {
	// функция считает количество чашек которое можно приготовить из заданного кол-ва ингредиентов
	water := w / waterPerCupMl
	milk := m / milkPerCupMl
	coffee := c / coffeePerCupG
	return minN(water, milk, coffee)
}

func main() {
	var waterMl, milkMl, coffeeG int
	f.Println("Write how many ml of water the coffee machine has:")
	f.Scan(&waterMl)
	f.Println("Write how many ml of milk the coffee machine has:")
	f.Scan(&milkMl)
	f.Println("Write how many grams of coffee the coffee machine has:")
	f.Scan(&coffeeG)
	f.Println("Write how many cups of coffee you will need:")
	var cupsNeeded int
	f.Scan(&cupsNeeded)

	howMany := howManyCups(waterMl, milkMl, coffeeG)

	switch {
	case howMany == cupsNeeded:
		f.Println("Yes, I can make that amount of coffee")
	case howMany > cupsNeeded:
		dif := howMany - cupsNeeded
		f.Printf("Yes, I can make that amount of coffee (and even %v more than that) \n", dif)
	case howMany < cupsNeeded:
		f.Printf("No, I can make only %d cups of coffee \n", howMany)
	}
}
