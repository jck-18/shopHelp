package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jck-18/shophelp/shopHelp"
)

func main() {
	// Create a new cart
	cart := shopHelp.NewCart()

	// Set up a reader to get user input from the console
	reader := bufio.NewReader(os.Stdin)

	for {
		// Prompt for an item name
		fmt.Print("Enter item name (or 'done' to finish): ")
		nameInput, _ := reader.ReadString('\n')
		name := strings.TrimSpace(nameInput)

		// Check for exit condition
		if strings.ToLower(name) == "done" {
			break
		}

		// Prompt for an item price
		fmt.Print("Enter item price: ")
		priceInput, _ := reader.ReadString('\n')
		priceStr := strings.TrimSpace(priceInput)

		// Convert price string to float64
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			fmt.Println("Invalid price. Please try again.")
			continue
		}

		// Add the item to the cart
		cart.AddItem(name, price)
		fmt.Printf("Added \"%s\" with price $%.2f\n\n", name, price)
	}

	// Display all items in the cart
	fmt.Println("\nItems in your shopping cart:")
	for _, item := range cart.Items() {
		fmt.Printf("- %s: $%.2f\n", item.Name, item.Price)
	}

	// Display the total price
	fmt.Printf("Total: $%.2f\n", cart.Total())
}
