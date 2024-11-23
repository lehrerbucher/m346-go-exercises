package main

import "fmt"

func main() {
	var firstName string = "Patrick"
	var lastName string = "Bucher"

	dayOfBirth := 24
	monthOfBirth := 6
	yearOfBirth := 1987

	var numberOfSiblings = 1

	heightInMeters := 1.88

	var zodiacSign rune = '\u264b'

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
