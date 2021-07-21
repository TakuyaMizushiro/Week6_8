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

func main() {
	var s string

	var checkEight bool = false
	var largeCap bool = false

	fmt.Print("Input your password: ")
	fmt.Scan(&s)

	//文字列をスライスに一文字ずつ代入
	arr := strings.Split(s, "")

	//スライスの要素をひとつずつ取り出して、最後の要素までfor文を回す
	//largeCap：大文字が入っているか判断
	for _, str := range arr {
		if largeCap == false && checkCapitalLetter(str) {
			largeCap = true
		}
	}
	//八文字以上か確認、あるならtrue
	checkEight = LengthEight(s)

	//すべての項目について確認し、安全性を標準出力
	if checkEight == true && largeCap == true {
		//secure = true
		fmt.Println("入力された文字列は安全です")
	} else {
		fmt.Println("以下の理由により入力された文字列は安全ではありません")
		if checkEight == false {
			fmt.Println("・8文字以上のパスワードを入力してください")
		}
		if largeCap == false {
			fmt.Println("・大文字のアルファベットが含まれていません")
		}
	}

}
