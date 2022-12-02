package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func solve_part_1(input []string) {
	calories := 0
	max_elf_calories := 0
	for _, calory := range input {
		if calory == "" {
			if max_elf_calories < calories {
				max_elf_calories = calories
			}

			calories = 0
			continue
		}

		calory_int, err := strconv.Atoi(calory)
		if err != nil {
			log.Fatal(err)
		}

		calories += calory_int
	}
	fmt.Printf("Solve part 1: %d\n", max_elf_calories)
}

func solve_part_2(input []string) {
	calories := 0
	first_elf := 0
	second_elf := 0
	third_elf := 0
	for _, calory := range input {
		if calory == "" {
			if first_elf < calories {
				third_elf = second_elf
				second_elf = first_elf
				first_elf = calories
			} else if second_elf < calories {
				third_elf = second_elf
				second_elf = calories
			} else if third_elf < calories {
				third_elf = calories
			}

			calories = 0
			continue
		}

		calory_int, err := strconv.Atoi(calory)
		if err != nil {
			log.Fatal(err)
		}

		calories += calory_int
	}
	fmt.Printf("Solve part 2: %d\n", first_elf+second_elf+third_elf)
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Missing file name argument")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	input := []string{}

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	solve_part_1(input)
	solve_part_2(input)
}
