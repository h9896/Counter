package database

import (
	"testing"
	"time"
)

func TestGetPermission(t *testing.T) {
	t.Log("Initialize Counter with interval 10")
	interval := 10
	c := NewCounter(interval)
	t.Log("Start test GetPermission Func with limit number 100")
	limitNum := 100
	aIP := "127.0.0.1"
	for i := 1; i <= limitNum; i++ {
		isTrue, err := c.GetPermission(aIP, limitNum)
		if err != nil {
			t.Error(err)
		}
		if !isTrue {
			t.Error("get wrong result when ", i)
		}
	}
	isTrue, err := c.GetPermission(aIP, limitNum)
	if err != nil {
		t.Error(err)
	}
	if isTrue {
		t.Error("get wrong result when out of limit number")
	}
	time.Sleep(time.Second * time.Duration(interval+1))
	isTrue, err = c.GetPermission(aIP, limitNum)
	if err != nil {
		t.Error(err)
	}
	if !isTrue {
		t.Error("get wrong result when clean limit number")
	}
}
func TestDeleteIPKey(t *testing.T) {
	t.Log("Initialize Counter with interval 2")
	interval := 2
	c := NewCounter(interval)
	t.Log("Start test GetPermission Func with limit number 3")
	limitNum := 3
	aIP := "127.0.0.1"
	_, err := c.GetPermission(aIP, limitNum)
	if err != nil {
		t.Error(err)
	}
	time.Sleep(time.Second * time.Duration((interval+1)*4))
	if _, ok := c.number[aIP]; ok {
		t.Error("Fail to delete IP key")
	}
}

func TestGetNumber(t *testing.T) {
	t.Log("Initialize Counter with interval 10")
	interval := 10
	c := NewCounter(interval)
	t.Log("Start test GetPermission Func with limit number 100")
	limitNum := 100
	aIP := "127.0.0.1"
	for i := 1; i <= limitNum; i++ {
		_, err := c.GetPermission(aIP, limitNum)
		if err != nil {
			t.Error(err)
		}
		val := c.GetNumber(aIP)
		if val != i {
			t.Error("GetNumber ", val, " and request times are not equal")
		}
	}
	nonIP := "127.0.0.2"
	val := c.GetNumber(nonIP)
	if val != 0 {
		t.Error("GetNumber ", val, " and request times are not equal")
	}
}

func TestGetAllNumber(t *testing.T) {
	t.Log("Initialize Counter with interval 10")
	interval := 10
	c := NewCounter(interval)
	t.Log("Start test GetPermission Func with limit number 100")
	limitNum := 100
	askA := 53
	askB := 23
	aIP := "127.0.0.1"
	bIP := "127.0.0.2"
	for i := 1; i < askA; i++ {
		_, err := c.GetPermission(aIP, limitNum)
		if err != nil {
			t.Error(err)
		}
		val := c.GetAllNumber()
		if val[aIP] != i {
			t.Error("GetAllNumber a ", val[aIP], " and request times are not equal")
		}
	}

	for i := 1; i < askB; i++ {
		_, err := c.GetPermission(bIP, limitNum)
		if err != nil {
			t.Error(err)
		}
		val := c.GetAllNumber()
		if val[bIP] != i {
			t.Error("GetAllNumber b ", val[bIP], " and request times are not equal")
		}
	}
	val := c.GetAllNumber()
	if val[aIP] != askA-1 || val[bIP] != askB-1 {
		t.Error("GetAllNumber ", val, " and request times are not equal")
	}
}
