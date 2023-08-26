package main

import (
	"fmt"
	"strings"
)

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode strings.Builder

	for _, char := range s.name {
		if char >= 'a' && char <= 'z' {
			encodedChar := rune('a' + (int(char-'a')+3)%26) // Shift by 3
			nameEncode.WriteRune(encodedChar)
		} else {
			nameEncode.WriteRune(char)
		}
	}

	s.nameEncode = nameEncode.String()
	return s.nameEncode
}

func (s *student) Decode() string {
	var nameDecode strings.Builder

	for _, char := range s.nameEncode {
		if char >= 'a' && char <= 'z' {
			decodedChar := rune('a' + (int(char-'a')+23)%26) // Shift back by 3
			nameDecode.WriteRune(decodedChar)
		} else {
			nameDecode.WriteRune(char)
		}

	}

	return nameDecode.String()
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of student’s encoded name " + a.nameEncode + " is : " + c.Decode())
	}
}