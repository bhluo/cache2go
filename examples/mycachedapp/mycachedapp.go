package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/heytom7/cache2go"
)

var text string

func main() {
	cache := cache2go.Cache("myCache")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		printMenu()

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			fmt.Print("Enter key: ")
			scanner.Scan()
			key := scanner.Text()

			fmt.Print("Enter value: ")
			scanner.Scan()
			value := scanner.Text()

			cache.Add(key, 60*time.Second, value)

			fmt.Println("Value added to cache.")

		case "2":
			fmt.Print("Enter key to retrieve value: ")
			scanner.Scan()
			key := scanner.Text()

			res, err := cache.Value(key)
			if err == nil {
				fmt.Println("Found value in cache:", res.Data())
			} else {
				fmt.Println("Error retrieving value from cache:", err)
			}

		case "3":
			fmt.Print("Enter key to delete value: ")
			scanner.Scan()
			key := scanner.Text()

			cache.Delete(key)

			fmt.Println("Value deleted from cache.")

		case "4":
			fmt.Println("Flushing the cache...")
			cache.Flush()
			fmt.Println("Cache flushed.")

		case "5":
			fmt.Println("Data of CacheItems in CacheTable")
			items := cache.GetItems()
			if len(items) == 0 {
				fmt.Println("nil!")
				break
			}
			for key, value := range items {
				fmt.Printf("Key: %v, Value: %v\n", key, value.Data())
			}

		case "6":
			fmt.Println("Exiting program...")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func printMenu() {
	fmt.Println("Welcome to Cache System!")
	fmt.Println("1. Add value to cache")
	fmt.Println("2. Retrieve value from cache")
	fmt.Println("3. Delete value from cache")
	fmt.Println("4. Flush cache")
	fmt.Println("5. Print all CacheItems")
	fmt.Println("6. Exit")
	fmt.Print("Enter option: ")
}
