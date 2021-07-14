package mypackage

import "fmt"

func UpperCase(x string) string { //大文字があるか判定

}

func LowerCase(x string) string { //小文字があるか判定

}

func LengthEight(x string) string { //8文字以上か判定

}

func Number(x string) string { //数字が入っているか判定

}

func Symbol(x string) string { //記号が入っているか判定

}

func main() {
	var s string

	fmt.Print("Input your password: ")
	fmt.Scanf("%d", &s)

	UpperCase(s)
	LowerCase(s)
	LengthEight(s)
	Number(s)
	Symbol(s)
}
