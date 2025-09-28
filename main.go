package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func learnBoolean() {

	var knowKannada bool = true
	var knowTelugu bool = true

	if knowTelugu && knowKannada {
		fmt.Println("Knows kannada and telugu")
	} else if knowTelugu {
		fmt.Println("Know only telugu")
	}
}

func learnIntegerTypes() {
	var largeNumber int = 420000          // Signed int (default size)
	var smallNumber uint8 = 250           // Max value is 255
	sum := largeNumber + int(smallNumber) // Type conversion is mandatory
	fmt.Println(sum)

	var unicodeChar rune = '😊' // Rune: Represents a single character (Unicode code point)
	fmt.Printf("Rune: %c, Value: %d\n", unicodeChar, unicodeChar)

	area := 3.14159                       // Inferred as float64
	var distance float32 = 100.5          // Explicitly declared
	finalArea := area + float64(distance) // Cannot directly add float64 and float32!
	fmt.Println("Final Area:", finalArea)

	val := 16.0
	root := math.Sqrt(val)
	fmt.Println("Sqrt:", root)

	ceiling := math.Ceil(4.3)
	fmt.Println("Ceiling:", ceiling)

	numStr := "123"
	i, err := strconv.Atoi(numStr)

	if err == nil {
		fmt.Printf("Converted int: %d\n", i+1) // Output: 124
	}
	newStr := strconv.Itoa(456)
	fmt.Printf("Converted string: %s\n", newStr) // Output: 456
}

func learnStringTypes() {
	name := "chandrasekhar"
	msg := "welcome"
	fmt.Println(msg + " " + name)
	for j, eachChar := range name {
		fmt.Printf("String at byte index %d: %c\n", j, eachChar)
	}

	sUnicode := "こんにちは"
	for i, char := range sUnicode {
		fmt.Printf("Rune at byte index %d: %c\n", i, char)
	}

	fmt.Println("length of name:", len(name))
	fmt.Println("length of sUnicode:", len(sUnicode))

	data := "apple,banana,cherry"

	parts := strings.Split(data, ",") // Split the string
	fmt.Println("Split:", parts)

	joined := strings.Join(parts, " and ") // Join the slice back into a string with a different separator
	fmt.Println("Joined:", joined)

	exists := strings.Contains(joined, "banana")
	fmt.Println("Contains banana:", exists)

	numStr := "12345"
	i, err := strconv.Atoi(numStr)

	if err == nil {
		fmt.Printf("Converted int: %d\n", i)
	}

	s := strconv.Itoa(99)
	fmt.Println("Converted string:", s)
	fmt.Printf("Variable type of s is: %T\n", s)
}

func learnSliceTypes() {

	numbers := []int{10, 20, 30, 40, 50}
	subNumbers := numbers[1:3]

	fmt.Printf("numbers: %v\n", numbers)
	fmt.Println(subNumbers)

	subNumbers[0] = 100
	fmt.Println(subNumbers)

	fmt.Println("length - ", len(numbers))
	fmt.Println("capacity - ", cap(numbers))

	newVar := make([]int, 2, 5) // Create a slice with length 2, capacity 5
	fmt.Printf("length: %d, capacity: %d\n", len(newVar), cap(newVar))

	makeExample := make([]int, 0, 10) // Create a slice with length 0, capacity 5
	fmt.Printf("length: %d, capacity: %d\n", len(makeExample), cap(makeExample))

	appendExample := []int{1, 2, 3, 4, 5}
	fmt.Printf("Before - length: %d, capacity: %d\n", len(appendExample), cap(appendExample))
	appendExample = append(appendExample, 6, 7) // The capacity may have increased (e.g., doubled)
	fmt.Printf("After - length: %d, capacity: %d\n", len(appendExample), cap(appendExample))

	copyExample := []int{100, 200, 300, 400, 500}
	newCopy := make([]int, len(copyExample))
	copy(newCopy, copyExample)
	newCopy[0] = 1
	fmt.Println("newCopy ", newCopy)
}

func main() {
	learnBoolean()
	learnIntegerTypes()
	learnStringTypes()
	learnSliceTypes()
}
