package gobyexample

import (
	"fmt"
	"unicode/utf8"
)

func Strings() {
	// is a slice of bytes
	const name = "Churros"
	fmt.Println("Len: ", len(name))

	for i := 0; i < len(name); i++ {
		// range produces the raw bytes values
		fmt.Printf("%x\n", name[i])
	}

	for _, runeVal := range name {
		fmt.Printf("Rune value: %#U\n", runeVal)
	}

	fmt.Println("Rune count: ", utf8.RuneCountInString(name))

	fmt.Println("HUM...")
	const s = "สวัสดี"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println("Rune count in THAI: ", utf8.RuneCountInString(s));

	fmt.Println()
	for _, r := range s {
		fmt.Printf("RUNA THAI: %c \n", r)
	}
	fmt.Println()

	fmt.Println("HUM...")
	const a = "á";

	// length 2
	fmt.Println("Length: ", len(a));

	for i := 0; i < len(a); i++ {
		fmt.Printf("%x ", a[i]);
	}
	fmt.Println();

	for _, r := range a {
		// t := rune(r);
		fmt.Printf("RUNA: %c\n", r);
	}

	fmt.Println("Rune count: ", utf8.RuneCountInString(a));
}
