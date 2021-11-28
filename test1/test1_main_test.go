package main

import "testing"

func TestGet_hello_正常系(t *testing.T) {
	expect := "Hello"
	actual := get_hello()
	if expect != actual {
		t.Errorf("expect:%v, actual:%v\n", expect, actual)
	}
}

func TestGet_world_正常系(t *testing.T) {
	expect := "World"
	actual := get_world()
	if expect != actual {
		t.Errorf("expect:%v, actual:%v\n", expect, actual)
	}
}
