package main

import (
    "fmt"
    "unicode/utf8"
)

//func UpperCase(x string) string { //大文字があるか判定

//}

//func LowerCase(x string) string { //小文字があるか判定

//}

func LengthEight(x string) bool { //8文字以上か判定

    if(utf8.RuneCountInString(ja) >= 8) {
    	return true
    }
    else {
    	return false
    }
}

//func Number(x string) bool { //数字が入っているか判定
	// for _, r := range x {
	// 	if '0' <= r && r <= '9' {
	// 		return true
	// 	}
	// }
	// return false
//}

//func Symbol(x string) string { //記号が入っているか判定

//}

func main() {
	var s string

	fmt.Print("Input your password: ")
	fmt.Scanf("%d", &s)

	//UpperCase(s)
	//LowerCase(s)
	LengthEight(s)
	//Number(s)
	//Symbol(s)
}
