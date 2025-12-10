package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	fmt.Println(CheckType(numDecimal))
	fmt.Println(CheckType(numOctal))
	fmt.Println(CheckType(numHexadecimal))
	fmt.Println(CheckType(pi))
	fmt.Println(CheckType(name))
	fmt.Println(CheckType(isActive))
	fmt.Println(CheckType(complexNum))

	str := ConvertToStr(numDecimal) + ConvertToStr(numOctal) + ConvertToStr(numHexadecimal) + ConvertToStr(pi) +
		ConvertToStr(name) + ConvertToStr(isActive) + ConvertToStr(complexNum)

	fmt.Println(str)

	strRune := ConvertToRune(str)

	fmt.Println(strRune)

	strRune = append(strRune[:len(strRune)/2], append([]rune("go-2024"), strRune[len(strRune):]...)...)

	fmt.Println(Hash(strRune))

}

func CheckType(val any) reflect.Type {
	return reflect.TypeOf(val)
}

func ConvertToStr(val any) string {

	switch v := val.(type) {
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case string:
		return v
	case bool:
		return strconv.FormatBool(v)
	case complex64:
		return fmt.Sprintf("%v", v)
	default:
		return ""
	}

}

func ConvertToRune(val string) []rune {

	return []rune(val)

}

func Hash(val []rune) [32]byte {

	return sha256.Sum256([]byte(string(val)))

}
