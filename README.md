# Shopping Lab

This repository demonstrates a simple **shopping cart** application in Go. The project contains:
- A **custom Go package** (`shoppingcart`) that manages a shopping cart’s items and calculates totals.
- A **main program** (`main.go`) that interactively prompts users for item names and prices, adds them to the cart, and displays the final total.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

---

## Features

- **Add items** to a cart by typing a name and price interactively.
- **Calculate the total** price of all items in the cart.
- **View a list** of items in the cart.
- **Easy to extend** with additional functionality (discounts, inventory checks, etc.).

---

## Installation

1. **Clone** the repository (HTTPS):
   ```bash
   git clone https://github.com/<YourUsername>/shopping-lab.git
   ```
   Or (SSH):
   ```bash
   git clone git@github.com:<YourUsername>/shopping-lab.git
   ```

2. **Navigate** into the project folder:
   ```bash
   cd shopping-lab
   ```

3. (Optional) **Initialize** or **update** dependencies:
   ```bash
   go mod tidy
   ```

---

## Usage

### Run the Interactive Program

Inside the project folder, run:
```bash
go run main.go
```

You will be prompted to enter items in a loop:

1. **Item Name**: Type the name of the item, then press Enter.
2. **Item Price**: Type a valid price (numeric), then press Enter.
3. Repeat for each item.
4. When done adding items, type `done` in the item-name prompt and press Enter.

**Example session**:

```
Enter item name (or 'done' to finish): Book: Go Programming
Enter item price: 19.99
Added "Book: Go Programming" with price $19.99

Enter item name (or 'done' to finish): Headphones
Enter item price: 59.99
Added "Headphones" with price $59.99

Enter item name (or 'done' to finish): done

Items in your shopping cart:
- Book: Go Programming: $19.99
- Headphones: $59.99
Total: $79.98
```

### Importing the `shoppingcart` Package in Another Project

If you want to **reuse** this package in a separate project:

1. Inside your other Go project, run:
   ```bash
   go get github.com/<YourUsername>/shopping-lab/shoppingcart
   ```
2. In your Go files, import and use the package:
   ```go
   import "github.com/<YourUsername>/shopping-lab/shoppingcart"

   func main() {
       cart := shoppingcart.NewCart()
       cart.AddItem("Laptop", 999.99)
       // ...
   }
   ```

---

## Testing

A basic unit test is provided in `shoppingcart/shoppingcart_test.go`. To run it, navigate to the project’s root directory and execute:

```bash
go test ./...
```

You should see output indicating whether the tests passed.

---

## Project Structure

```
shopping-lab/
├── go.mod
├── main.go                // Interactive CLI for the shopping cart
├── shoppingcart/
│   ├── shoppingcart.go    // Core shopping cart package
│   └── shoppingcart_test.go  // Unit tests for the shopping cart
└── README.md              // Project documentation
```

- **`go.mod`**: Go module definition, including the module path.
- **`main.go`**: The entry point of the application. Prompts for user input, adds items to the cart, and prints results.
- **`shoppingcart` folder**: Contains the `shoppingcart.go` source file, which defines the `Cart` and `Item` structures, and their associated logic.
- **`_test.go` file**: Contains tests for the shopping cart package.

---

## Contributing

1. **Fork** the repository.
2. **Create a new branch** for your feature or bug fix:  
   ```bash
   git checkout -b feature/amazing-feature
   ```
3. **Commit your changes**:  
   ```bash
   git commit -m "Add some amazing feature"
   ```
4. **Push** to the branch:  
   ```bash
   git push origin feature/amazing-feature
   ```
5. **Open a Pull Request** on GitHub.

All contributions are welcome! Feel free to open an issue for suggestions or problems.

---

## License

This project is licensed under the [MIT License](LICENSE). See the [LICENSE](LICENSE) file for details.

---

### Thank You
Thanks for checking out this simple Go-based shopping cart application. If you find it helpful or have feedback, feel free to open an issue or submit a PR!
