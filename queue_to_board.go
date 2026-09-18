package main

import "fmt"

func main() {
	var input string
	var place int

	var place1 string
	var place2 string
	var place3 string
	var place4 string
	var place5 string

	count := 0

	for {
		if _, err := fmt.Scan(&input); err != nil {
			break
		}

		if input == "очередь" {
			if place1 == "" {
				fmt.Println("1. -")
			} else {
				fmt.Println("1.", place1)
			}

			if place2 == "" {
				fmt.Println("2. -")
			} else {
				fmt.Println("2.", place2)
			}

			if place3 == "" {
				fmt.Println("3. -")
			} else {
				fmt.Println("3.", place3)
			}

			if place4 == "" {
				fmt.Println("4. -")
			} else {
				fmt.Println("4.", place4)
			}

			if place5 == "" {
				fmt.Println("5. -")
			} else {
				fmt.Println("5.", place5)
			}

		} else if input == "количество" {
			fmt.Println("Осталось свободных мест:", 5-count)
			fmt.Println("Всего человек в очереди:", count)

		} else if input == "конец" {
			if place1 == "" {
				fmt.Println("1. -")
			} else {
				fmt.Println("1.", place1)
			}

			if place2 == "" {
				fmt.Println("2. -")
			} else {
				fmt.Println("2.", place2)
			}

			if place3 == "" {
				fmt.Println("3. -")
			} else {
				fmt.Println("3.", place3)
			}

			if place4 == "" {
				fmt.Println("4. -")
			} else {
				fmt.Println("4.", place4)
			}

			if place5 == "" {
				fmt.Println("5. -")
			} else {
				fmt.Println("5.", place5)
			}

			break

		} else {
			fmt.Scan(&place)

			if place < 1 || place > 5 {
				fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", place)

			} else if count == 5 {
				fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", place)

			} else if place == 1 && place1 != "" {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", place)

			} else if place == 2 && place2 != "" {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", place)

			} else if place == 3 && place3 != "" {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", place)

			} else if place == 4 && place4 != "" {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", place)

			} else if place == 5 && place5 != "" {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", place)

			} else {
				if place == 1 {
					place1 = input
				} else if place == 2 {
					place2 = input
				} else if place == 3 {
					place3 = input
				} else if place == 4 {
					place4 = input
				} else if place == 5 {
					place5 = input
				}

				count++
			}
		}
	}
}
