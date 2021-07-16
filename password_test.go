package main

import (
	//"fmt"
	//"reflect"
	"testing"
)

func TestStringReturn01(t *testing.T) {
	//出力確認
	main()
}

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
