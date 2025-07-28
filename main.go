package main

import f "fmt"

func main() {
	f.Println("Write how many cups of coffee you will need:")
	numbersCups := 0
	f.Scan(&numbersCups)
	water := 200 * numbersCups
	milk := 50 * numbersCups
	coffeBeans := 15 * numbersCups
	f.Printf("For %v cups of coffee you will need:", numbersCups)
	f.Println("Starting to make a coffee")
	f.Printf("Grinding coffee beans %v g \n", coffeBeans)
	f.Printf("Boiling water %v ml \n", water)
	f.Println("Mixing boiled water with crushed coffee beans")
	f.Println("Pouring coffee into the cup")
	f.Printf("Pouring some milk into the cup %v ml \n", milk)
	f.Println("Coffee is ready!")
}
