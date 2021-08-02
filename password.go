package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func LengthEight(x string) bool { //8文字以上か判定

	if utf8.RuneCountInString(x) >= 8 {
		return true
	} else {
		return false
	}

}

func checkCapitalLetter(str string) bool { //大文字のアルファベットが入力されているかチェック
	if "A" <= str && str <= "Z" {
		//fmt.Println("Large Capital OK")
		return true
	} else {
		return false
	}
}

func checkSmallLetter(str string) bool {
	if "a" <= str && str <= "z" {
		//fmt.Println("Large Capital OK")
		return true
	} else {
		return false
	}
}

func checkNumber(str string) bool {
	if "0" <= str && str <= "9" {
		//fmt.Println("Large Capital OK")
		return true
	} else {
		return false
	}
}

func main() {
	var s string

	var checkEight bool = false
	var capitalLetter bool = false
	var smallLetter bool = false
	var number bool = false

	fmt.Print("Input your password: ")
	fmt.Scan(&s)

	//文字列をスライスに一文字ずつ代入
	arr := strings.Split(s, "")

	//スライスの要素をひとつずつ取り出して、最後の要素までfor文を回す
	//largeCap：大文字、小文字、数字が入っているか判断
	for _, str := range arr {
		if capitalLetter == false && checkCapitalLetter(str) {
			capitalLetter = true
		}

		if smallLetter == false && checkSmallLetter(str) {
			smallLetter = true
		}

		if number == false && checkNumber(str) {
			number = true
		}
	}
	//八文字以上か確認、あるならtrue
	checkEight = LengthEight(s)

	//すべての項目について確認し、安全性を標準出力
	if checkEight == true && capitalLetter == true && smallLetter == true {
		//secure = true
		fmt.Println("入力された文字列は安全です")
	} else {
		fmt.Println("以下の理由により入力された文字列は安全ではありません")
		if checkEight == false {
			fmt.Println("・8文字以上のパスワードを入力してください")
		}
		if capitalLetter == false {
			fmt.Println("・大文字のアルファベットが含まれていません")
		}
		if smallLetter == false {
			fmt.Println("・小文字のアルファベットが含まれていません")
		}
		if number == false {
			fmt.Println("・数字が含まれていません")
		}
	}

}
