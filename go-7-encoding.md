# Hexadecimal, octal, ASCII, UTF8, Unicode, Runes

## 1. Base 16: the hexadecimal representation

To represent a binary number, you need to align many zeros and one. This notation is verbose. To represent the decimal number 1324, we needed to use 11 characters in binary. That’s why we need to have a numbering system more convenient to express large numbers.

Hexadecimal is also a positional numeral system that uses 16 characters to represent a number.

- The prefix Hexa means 6 in Latin

- Decimal is coming from the Latin word Decem which means 10

Those characters are numbers and letters. We use the numbers from 0 to 9 (10 characters) and the letters from A to F (6 characters).

Let’s take an example: 1324 in base ten is equivalent to 52C in base 16

![image](./images/hexa_decimal.png)

The digits from 0 to 9 correspond to the same value in the decimal system. The letters A correspond to 10, the letter B to 11 ...etc. This is a specificity of the hexadecimal numeral system; we use letters to represent numeric values.

Usually, this specificity generate confusion and questions to my students, what I typically reply is that you have to admit it; we needed more characters so we took letters...

![image](./images/hexa_decimal_representation.png)

You can see that we introduced in this notation letters. That’s because from 0 to 9, you have ten characters, ten digits, but with a base-16 numbering system, we need six more characters. That’s why we have taken the first six letters of the alphabet. This is a historical choice; other characters could have replaced letters, the system would have been still the same.

The method you can use to convert a hexadecimal number to a decimal number is similar to the previous one. We take the rightmost character we find its decimal equivalent, then we multiply it by 16 at the power 0. In our example, we have the letter C.The equivalent of C is 12.

To print the hexadecimal representation of a number, you can use fmt functions. :

```go
// hexadecimal-octal-ascii-utf8-unicode-runes/hexa-lower/main.go
package main

import "fmt"

func main() {
    n := 2548
    fmt.Printf("%x", n)
}
```
This program will output: 9f4 (which is the hexadecimal representation of the decimal number 2548)."%x" is the formatting verb for hexadecimal (with letters lowercase).

**If you want to represent a number in hexadecimal in your code, add 0x before the numeral :**

## 2. Base 8: the octal representation

I have almost forgotten another numeral system! The octal!

It uses a base 8, which means eight different characters. The numbers from 0 to 7 were chosen. The conversion from decimal to octal is similar to the methods that I have presented before. Let’s take an example :

![image](./images/hexa_decimal_representation.png)

We begin by the rightmost character, and we multiply it by eight at the power 0, which is 1. Then we take the next character: 5 to multiply it by eight at the power 1, which is 8...

The octal system is notably used to represent permissions on a file for Unix operating systems.

In the same fashion as hexadecimal, the fmt package defines two formatting verbs for octal :

```go
// /hexadecimal-octal-ascii-utf8-unicode-runes/octal/main.go
package main

import "fmt"

func main() {
    n2 := 0x9F4
    fmt.Printf("Decimal : %d\n", n2)

    // n3 is represented using the octal numeral system
    n3 := 02454
    // alternative : n3 := 0o2454

    // convert in decimal
    fmt.Printf("decimal: %d\n", n3)

    // n4 is represented using the decimal numeral system
    n4 := 1324
    // output n4 (decimal) in octal
    fmt.Printf("octal: %o\n", n4)
    // output n4 (decimal) in octal (with a 0o prefix)
    fmt.Printf("octal with prefix : %O\n", n4)

}
```

Output :

```go
Decimal : 2548
decimal: 1324
octal: 2454
octal with prefix : 0o2454
```

**"%o"** allow you to print the number in octal

**"%O"** allow you to print the number in octal with a **"0o"** prefix

## 3. Data representation bits, nibble, bytes, and words

Bit is an abbreviation for Binary digit .For instance 10100101100 is made of 11 binary digits, in other words, 11 bits. It’s very usual to group bits together. Groups exist in various sizes:

- A nibble is composed of 4 bits

- A byte is composed of 8 bits (two nibbles)

- A word is composed of 16 bits (two bytes)

- A doubleword is composed of 32 bits (two words)

- A quadword is composed of 16×4=64 bits (four words)

With Go, you can create a slice of bytes. Lots of common standard package functions and methods are taking as arguments slice of bytes. Let’s see how we can create a slice of byte.

```go
// /hexadecimal-octal-ascii-utf8-unicode-runes/slice-of-byte/main.go
package main

import "fmt"

func main() {
    b := make([]byte, 0)
    b = append(b, 255)
    b = append(b, 10)
    fmt.Println(b)
}
```

In the previous snippet, we created a slice of bytes (with the builtin make) then we appended to the slice two numbers.

Golang byte type is an alias of uint8. Uint8 means that we can store unsigned (without any signs, so no negative numbers) integers on 8 bits (a byte) of data. The minimum value is 0 (the binary digit 000000~2~) the maximum value is 255 (11111111~2~​ which is equivalent to the decimal number 2^7^ + 2^6^ + 2^5^ + 2^4^ + 2^3^ + 2^2^ + 2^1^ + 2^0^)

That’s why we can only append to a byte slice numbers from 0 to 255. If you try to append a number greater than 255, you will get the following error :

> constant 256 overflows byte

**To print the binary representation of a number, you can use the "%b" formatting verb :**

```go
// /hexadecimal-octal-ascii-utf8-unicode-runes/decimal-binary/main.go
package main

import "fmt"

func main() {
    n2 := 0x9F4
    fmt.Printf("Decimal : %d\n", n2)
    fmt.Printf("Binary : %b\n", n2)
}
```
Output:

```go
Decimal : 2548
Binary : 100111110100
```

## 4. What about other characters? 

What if you want to store something other than numbers? For instance how could we store this Haiku from Masaoki Shiki :

```bash
spring rain:
browsing under an umbrella
at the picture-book store
```

Is the byte type appropriate? A byte is nothing more than an unsigned integer stored on 8 bits. This Haiku is composed of letters and special characters. We have an“:” and a “-” we also have line breaks... How can we store those characters?

We have to find a way to give each letter and even special characters an unique code. You have maybe heard about UTF-8, ASCII, Unicode? This section will explain what they are and how they work. Once I started programming (that was not in Go), character encoding was something obscure, and I did not find it interesting. I think that character encoding could be essential because I have spent nights at work on problems that could have been resolved with a basic understanding of character encoding.

The history of character encoding is very rich. With the development of the telegraph, we needed a way to encode messages in a way that could be transportable on an electrical wire. One of the earliest attempts was the Morse code. It is composed of four symbols: short signal, long signal, short space, long space (Wikipedia). Each letter of the alphabet could be encoded in morse. For instance, A was encoded as a short signal followed by a long signal. The plus sign “+” was encoded with “short long short long short”.

## 5. Vocabulary

We need to define a common vocabulary to understand character encoding :

- Character This can be written by our hand. It conveys a signification. For instance, the sign “+” is a character. It means adding something to something else. A character can be a letter, a sign, or an ideogram.

- Character set: this a collection of distinct characters. Often you will see or hear the abbreviation “charset”.

- Code point : each character from a character set has an equivalent numeric value that uniquely identify this character. This numeric value is a code point.

## 6. Character sets and encoding

There is one character set that you want to know : Unicode. It is a standard that lists the vast majority of characters from living languages that are used today on computers

It is composed of 137,374 characters for it’s version 11.0. Unicode is like an enormous table that maps a character to a code point. For instance, the character “A” is mapped to the code point “0041”.

With Unicode, we have our basis, our table of characters, now the next challenge is to find a way to encode those characters, to put those code point into bytes of data. This is precisely what ASCII and UTF-8 do.

![image](./images/unicodepoint.png)
![image](./images/compare_unicode_character.png)

## 7. How ASCII works?

- ASCII means American Standard Code for Information Interchange. It has been developed during the sixties. The objective was to find a way to encode characters used to transmit messages.

ASCII encode characters on seven binary digits. Another binary digit is a parity bit. A parity bit is used to detect transmission errors. It’s added after the seven first bits, and its value is 0. If the number of ones is odd, then the parity bit is 1; if the number is even, it’s set to 0.

A byte of data can store each character (8 bits). How many integers can you create with only 7 bits ? With one single bit, we can encode two values, 0 and 1, with 2 bits, we can encode four distinct values. When you add a bit, you multiply by two the number of values you can encode. With 7 bits, you can encode 128 integers. More generally, the number of unsigned integers you can encode with n binary digits is two at the power n.

| Number of bits | Number of values |
| ----------- | ----------- |
| 1      | 2                |
| 2      | 4                |
| 3      | 8                |
| 4      | 16               |
| 5      | 32               |
| 6      | 64               |
| 7      | 128              |

ASCII allows you to encode 128 different characters. For each character, we have a specific code point. Unsigned integer values represent code points.

![image](./images/ascii_code.png)

On the previous figure1, you can see the USASCII code chart. This table allows you to convert a byte into a character. For instance the letter B is equivalent to 1000010 (binary) (column 4, row 2)

## 8. How UTF-8 works?

- UTF-8 means Universal Character Set Transformation Format1 - 8 bits. It has been invented by two people that are also the creators of Go: Rob Pike and Ken Thompson! The design of this type of encoding is very ingenious. I will try to explain it briefly :

UTF-8 is a variable width encoding system. It means that characters are encoded using one to four bytes (a byte represents eight binary digits).

![image](./images/utf8_system.png)

On the figure 5 you can see the encoding rules of UTF-8. A character can be encoded on 1 to 4 bytes.

The code points that can be encoded using only one byte are from U+0000 to U+007F (included). This range is composed of 128 characters. (from 0 to 127, there are 128 numbers2

But more characters need to be encoded! That’s why the creators of UTF-8 had the idea of adding bytes to the system. The first additional byte begins with a one and a 0; those are fixed. It signals to decoders that we are now using 2 bytes to encode our characters we simply add the bits “110”. It says to UTF-8 decoders, “be careful; we are 2 !”.

If we use 2 bytes, we have 11 bits free (8 * 2 - 5 (fixed bits) =11). We can encode the characters which have the Unicode code point from U+0080 to U+07FF included. How many characters does that represent?

- 0080 in hex = 128 in decimal

- 07FF in hex = 2047 in decimal

- from 0080 to 07FF there are 2047-128+1=1920

You might ask why do we add a one to the count... That’s because characters are indexed from the code point 0.

If you use 3 bytes, then the first byte will start with the fixed bits 1110. This will signal to decoders that the character is encoded using 3 bytes. In other words, the next characters will begin after the third byte. The two additional bytes are beginning with 10. With three encoding bytes, you have 16 bits free (8 * 3 - 8 (fixed bits) =16). You can encode characters from U+0800 to U+FFFF.

If you have understood how it works for 3 bytes, then you should have no problem to know how the system works with 4 bytes. Inside our first byte, we fix the five first bits (11110). Then we have three additional bytes. If we subtract the fixed bits from the total number of bits, we have 21 bits available. It means that we can encode code points from U+10000 to U+10FFFF.

## 9. Strings

A string is “a sequence of characters”. For instance "Test" is a string composed of 4 different characters: T, e, s, and t. Strings are prevalent; we use them to store raw text inside our program. They are generally readable by humans. For instance, the first name and the last name of an application user are two strings.

Characters can come from different character sets. If you use the character set ASCII, you have to choose from 128 characters available.

Each character has a corresponding code point in the character set. As we have seen before, the code point is an unsigned integer arbitrarily chosen. Strings are stored using bytes. Let’s take the example of a string composed only of ASCII characters :

```bash
Hello
```

A single byte can store each character. This string can be stored with the following bits :

```bash
01001000 01100101 01101100 01101100 01101111
```

![image](./images/string_to_binary.png)

In Go strings are immutables, meaning that they cannot be modified once created.

## 10. String literals

There are two “types” of strings literals :

- raw string literals. They are defined between back quotes.

    - Forbidden characters are

        - back quotes
    - Discarded characters are

        - Carriage returns (\r)
- interpreted string literals. They are defined between double-quotes.
    - Forbidden characters are
        - new lines

        - unescaped double quotes

```go
// /hexadecimal-octal-ascii-utf8-unicode-runes/string-literals/main.go
package main

import "fmt"

func main() {

    raw := `spring rain:
browsing under an umbrella
at the picture-book store`
    fmt.Println(raw)

    interpreted := "i love spring"
    fmt.Println(interpreted)
}
```

You can note that inside this snippet of code, we did not say to Go which character set we use. This is because string literals are **implicitly encoded using UTF-8**.

## 11. Runes

Behind the scene, a string is a collection of bytes. We can iterate over the bytes of a string with a for loop:

```go
package main
import "fmt"

func main() {
	s := "ぶん Golang"
	for _, v := range s {
		// v is of type rune
		fmt.Printf("Unicode code point : %U - character '%c' - binary %b - hex %X - Decimal %d\n ", v, v, v, v, v)
	}
}
```

This program will iterate over each character of the string. Inside the for loop v is of type *rune*.*rune* is a built-in type that is defined as follow :

```go
// rune is an alias for int32 and is equivalent to int32 in all ways. It is
// used, by convention, to distinguish character values from integer values.
type rune = int32
```

A **rune** represent a Unicode code point.

- Unicode code points are numeric values.

- By convention, they are always noted with the following format: **"U+X"** where X is the hexadecimal representation of the code point. *X* should have four characters.

If X has less than four characters, we add zeros.

Ex: The character **"o"** has a code point equal to **111** (in decimal). 111 in hexadecimal is written 6F. The decimal code point is **U+006F**

To print the code point in the conventional format, you can use the format verb "**%U"**.

![image](./images/unicode_point_string.png)

Note that you can create a rune by using simple quotes :

```go
// /hexadecimal-octal-ascii-utf8-unicode-runes/rune/main.go
package main

import "fmt"

func main(){
    var aRune rune = 'Z'
    fmt.Printf("Unicode Code point of &#39;%c&#39;: %U\n", aRune, aRune)
}
```

## 12. Test yourself

1. True or false : “785G” is an hexadecimal numeral


- False

- The letter G cannot be part of hexadecimal numbers.

- However, the letters A to F can be part of a hexadecimal number.

2. True or false : “785f” and “785F” represent the same quantity

- This is true

- The fact that a letter is capitalized does not change its signification.

3. What is the formatting verb to represent a hexadecimal number (with a capitalized letter)?

- %X

4. What is the formatting verb to represent a number in decimal?

- %d

5. What is a code point?

- A code point is a numeric value that identifies a character in a character set

6. Fill the blanks. _______ is a character set, ______ is an encoding standard.

- **Unicode** is a character set, **UTF-8** is an encoding standard.

7. True or false: UTF-8 allows you to encode fewer characters than ASCII.

- False

8. How many bytes can I use to encode a character using the UTF-8 encoding system?

- From 1 to 4 bytes

- It depends on the character

## 13. Key takeaways

- Hexadecimal is a numeration system like decimal and binary

- With hexadecimal, a number is represented using 16 characters :

    - 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, A, B, C, D, E, F
- With fmt functions (fmt.Sprintf and fmt.Printf) you can use “formatting verbs” to represent a number using a specific numeral system

    - %b for binary

    - %X and%x for hexadecimal

    - %d for decimal

    - %o for octal

- **Character** This is something that can be written by our hand, which conveys a signification. Ex: “-”, “A” , “a”

- **Character set**: this a collection of distinct characters. Often you will see or hear the abbreviation “charset”.

- **Code point** : each character from a character set as an equivalent numeric value that uniquely identify this character. This numeric value is a code point.

- Unicode is a character set that is composed of 137.000 + characters.

- Each character has a code point. For instance"A" character is equivalent to the code point U+0041

- ASCII is an encoding technique that can encode only 128 characters.

- UTF-8 is an encoding technique that can encode more than 1 million characters

- With UTF-8, any character is encoded using 1 to 4 bytes.

- **rune** is a builtin type

- A rune represents the Unicode code point of a character.

- To create a rune, you can use simple quotes :

```bash
var aRune rune = 'Z'
```

- When you iterate over a string, you will iterate over runes.

```go
// /hexadecimal-octal-ascii-utf8-unicode-runes/iterate-over-string/main.go
package main

import "fmt"

func main() {
    b := "hello"
    for i := 0; i < len(b); i++ {
        fmt.Println(b[i])
    }
    // will output :
    // 104
    // 101
    // 108
    // 108
    // 111
    // and NOT :
    // h
    // e
    // l
    // l
    // o
}
```

- In Go, strings are immutable, meaning that they cannot be changed once created.

