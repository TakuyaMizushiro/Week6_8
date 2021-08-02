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
	if checkCapitalLetter(s) != true {
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

func TestCount05(t *testing.T) {
	//小文字が含まれているかの確認（小文字あり）
	var s string
	s = "aIsy98!"
	if checkSmallLetter(s) != true {
		t.Error("TestCount05 is failed")
	}
}

func TestCount06(t *testing.T) {
	//小文字が含まれているかの確認（小文字なし）
	var s string
	s = "AHUFLIUGEF98!"
	if checkSmallLetter(s) == true {
		t.Error("TestCount06 is failed")
	}
}

func TestCount07(t *testing.T) {
	//数字が含まれているかの確認（数字あり）
	var s string
	s = "1isy98!"
	if checkNumber(s) != true {
		t.Error("TestCount07 is failed")
	}
}

func TestCount08(t *testing.T) {
	//数字が含まれているかの確認（数字なし）
	var s string
	s = "aisyJgkhdf"
	if checkNumber(s) == true {
		t.Error("TestCount08 is failed")
	}
}

func TestCount09(t *testing.T) {
	//記号が含まれているかの確認(記号なし)
	var s string
	s = "aisyJgkhdf"
	if checkSymbol(s) == true {
		t.Error("TestCount09 is failed")
	}
}

func TestCount10(t *testing.T) {
	//記号が含まれているかの確認（記号あり）
	var s string
	s = "%$ais#yJgkhdf"
	if checkSymbol(s) != true {
		t.Error("TestCount10 is failed")
	}
}
