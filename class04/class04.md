## class04

### strings

- need to think about strings in a "logical way" and "physical way"
- strings in go are all unicode
- technique to represent international characters
- old days ascii: 
    - represent a character with 7 bits
    - represent american english characters
    - basically, all characters fit into 1 byte
- international languages accent marks, non-roman languages like chinese and arabic
- we need technique to represent those
- unicode uses numbers that are bigger what fits into a byte
- a rune is the go equivalent of a 'character', sometimes called 'wide character'
- rune is a synonym for a int32 (that 4 byte is big enough to represent any unicode code point)
- a unicode code point is a numerical value that represents a specific character within the Unicode standard. It's essentially a unique identifier for each character
- but for effiency we don't wanna represent every character all the times with 4 bytes, a lot of programs are just gonna have a ascii characters
- theres a technique for encoding unicode called UTF-8
- a short way of representing unicode in bytes (invented by some who worked on go)
- Bell labs developed UTF-8 as an efficient way to encode unicode
- "physically" they are the UTF-8 encoding of unicode character
- theres a type called the byte (uint8)
- A string is physically a sequence of bytes required to encode the unicode characters that are there logically (the runes)
- the runes is a synonim for a 32 bit int.

Runes (characters) are enclosed in a single quotes `'a'`

**A STRING IS AN INMUTABLE SEQUENCE OF "CHARACTERS"**
- *physically* a sequence of bytes (UTF-8 encoding)
- *logically* a sequence of (unicode) runes

"Raw" strings use backktick quotes `string with "quotes"`
They don't evaluate escape characters such as \n

**THE LENGTH OF A STRING IS THE LENGTH OF BYTE STRING NECESARY TO ENCODE THE STRING IN UTF-8 **
- LOGICALLY/VISUALLY 5 CHARS
- PHYSICALLY IS 5 PRINTABLE CHARS