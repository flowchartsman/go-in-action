package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

//strings

func main() {
	basicStr := "Hello, Gophers!"
	fmt.Println(basicStr)

	// Unicode characters are supported natively.
	unicodeStr := "你好，地鼠!"
	fmt.Println(unicodeStr)

	strWithEscapes := "Hello\n\u5730\u9F20"
	fmt.Println(strWithEscapes) // output: Hello<newline>地鼠

	rawString := `Hello\n\u5730\u9F20`
	fmt.Println(rawString) // output: Hello\n\u5730\u9F20
	rawStrWithNewlines := `I can
span विभिन्न lines`
	fmt.Println(rawStrWithNewlines)

	// string operations
	fmt.Println("hello" == "hello") // true

	fmt.Println("地鼠" == "\u5730\u9F20") // true, these escape codes are the
	// same characters

	fmt.Println("地鼠" == `\u5730\u9F20`) // false, escapes don't apply in
	// raw strings

	fmt.Println("abc" < "cba")   // true
	fmt.Println("andy" > "bill") // false

	strHello := "hello"
	strGophers := " Gophers!"
	strHello += strGophers
	fmt.Println(strHello) // output: "hello Gopher!"

	asciiCharStr := "easy, right?"
	fmt.Println(len(asciiCharStr)) // output: 12

	unicodeCharStr := "地鼠"
	fmt.Println(len(unicodeCharStr)) // output: 6

	// iterating an ASCII string and printing the bytes as strings
	// output: e a s y ,   r i g h t ?
	for i := 0; i < len(asciiCharStr); i++ {
		fmt.Print(string(asciiCharStr[i]) + " ")
	}
	fmt.Println()

	// iterating a Unicode string and attempting to print the bytes as
	// strings
	// output: å  ° é ¼
	for i := 0; i < len(unicodeCharStr); i++ {
		fmt.Print(string(unicodeCharStr[i]) + " ")
	}
	fmt.Println()

	// several single characters of various actual byte lengths
	fmt.Printf("% x\n", "A") // output: 41
	fmt.Printf("% x\n", "À") // output: c3 80
	fmt.Printf("% x\n", "Ᾰ") // output: e1 be b8
	fmt.Printf("% x\n", "🅰") // output: f0 9f 85 b0

	for i, r := range unicodeCharStr {
		fmt.Printf("%d:%s ", i, string(r))
	}
	//output: 0:地 3:鼠
	fmt.Println()

	characters := []rune(unicodeCharStr)
	for i, r := range characters {
		fmt.Printf("%d:%s ", i, string(r))
	}
	//output: 0:地 1:鼠
	fmt.Println()

	interestingCharacters := "À🅰"
	for _, r := range interestingCharacters {
		fmt.Printf("rune: %s byte length: %d\n", string(r), utf8.RuneLen(r))
		if unicode.IsLetter(r) && unicode.IsUpper(r) {
			fmt.Println("rune is an uppercase letter")
			fmt.Println("lowercase is:", string(unicode.ToLower(r)))
		}
		if unicode.IsSymbol(r) {
			fmt.Println("rune is a symbolic character")
		}
	}
}
