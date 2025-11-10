package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("🚀 go Calculator v1.0")
	fmt.Println("=======================")

	reader := bufio.NewReader(os.Stdin)

	for {

		displayMenu()

		fmt.Print("Enter choice (1-5): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if choice == "5" {
			fmt.Println("Exiting the calculator. Goodbye!")
			break
		}

		if choice < "1" || choice > "5" {
			fmt.Println("Invalid choice. Please select a valid operation.")
			continue
		}

		num1, num2 := getNumbers(reader)

		// perform calculation based on user choice
		result, err := performCalculation(choice, num1, num2)

		if err != nil {
			fmt.Printf(" ❌ Error: %v \n", err)
		} else {
			fmt.Printf(" ✅ Result: %.2f\n", result)
		}

		fmt.Println("\n" + strings.Repeat("-", 30))
	}

}

func displayMenu() {
	fmt.Println("\nSelect an operation:")
	fmt.Println("1. Addition (+)")
	fmt.Println("2. Subtraction (-)")
	fmt.Println("3. Multiplication (*)")
	fmt.Println("4. Division (/)")
	fmt.Println("5. Exit")
}

func getNumbers(reader *bufio.Reader) (float64, float64) {

	for {
		fmt.Print("Enter first number: ")
		num1Str, _ := reader.ReadString('\n')
		num1, err1 := strconv.ParseFloat(strings.TrimSpace(num1Str), 64)

		fmt.Print("Enter second number: ")
		num2Str, _ := reader.ReadString('\n')
		num2, err2 := strconv.ParseFloat(strings.TrimSpace(num2Str), 64)

		if err1 != nil || err2 != nil {
			fmt.Println("Invalid input. Please enter numeric values.")
			continue
		}

		return num1, num2
	}

}

func performCalculation(choice string, num1, num2 float64) (float64, error) {

	switch choice {

	case "1":
		return Add(num1, num2), nil
	case "2":
		return Subtract(num1, num2), nil
	case "3":
		return Multiply(num1, num2), nil
	case "4":
		return Divide(num1, num2)
	default:
		return 0, fmt.Errorf("Invalid operation")

	}

}
