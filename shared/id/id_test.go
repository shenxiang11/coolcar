package id

import (
	"fmt"
	"testing"
)

type MyString string

func haha(a MyString) {

}

func TestConvert(t *testing.T) {
	var str any = "123456"

	sid, ok := str.(string)
	if !ok {
		t.Fatalf("string 类型转换失败")
	}

	//haha(sid)

	aid, ok := str.(MyString)
	if !ok {
		t.Fatalf("AccountID 类型转换失败")
	}

	fmt.Println(sid, aid)
}
