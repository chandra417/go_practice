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
	fmt.Printf("Variable type of s is: %T", s)
}

func main() {
	learnBoolean()
	learnIntegerTypes()
	learnStringTypes()
}
