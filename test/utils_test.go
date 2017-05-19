package test

import (
    "testing"
)

func TestAdd(t *testing.T) {

	if Add(1, 1) != 2 {
		t.Error("test add fail")
	}else{
		t.Log("test add succ")
	}
}

func TestMax(t *testing.T){
	if Max(1, 2) == 1 {
		t.Error("test max fail")
	}else{
		t.Log("test max succ")
	}
}

