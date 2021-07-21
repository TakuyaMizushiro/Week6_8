package main

import (
	//"fmt"
	//"reflect"
	"testing"
)

/*func TestStringReturn01(t *testing.T) {
	//出力確認
	main()
}*/

func TestCount01(t *testing.T) {
	//８文字以上かどうかの確認(8文字以上)
	var s string
	s = "aisy983g24d9t8"
	if LengthEight(s) != true {
		t.Error("TestCount01 is failed")
	}
}

func TestCount02(t *testing.T) {
	//８文字以上かどうかの確認(8文字未満)
	var s string
	s = "aisy98!"
	if LengthEight(s) == true {
		t.Error("TestCount02 is failed")
	}
}

func TestCount03(t *testing.T) {
	//大文字が含まれているかの確認（大文字あり）
	var s string
	s = "Aisy98!"
	if checkCapitalLettergit(s) != true {
		t.Error("TestCount03 is failed")
	}
}

func TestCount04(t *testing.T) {
	//大文字が含まれているかの確認（大文字なし）
	var s string
	s = "aisy98!"
	if checkCapitalLetter(s) == true {
		t.Error("TestCount04 is failed")
	}
}
