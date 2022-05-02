package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	arr []int
}

func (st *Stack) Push(val int) {
	st.arr = append(st.arr, val)
}

func (st *Stack) Pop() int {
	arrlen := len(st.arr)
	if arrlen == 0 {
		fmt.Println("no data exists in stack")
		return 0
	}
	le := st.arr[arrlen-1 : arrlen]
	st.arr = st.arr[:arrlen-1]
	return le[0]
}

func (st *Stack) Seek() int {
	arrlen := len(st.arr)
	if arrlen == 0 {
		fmt.Println("no data exists in stack")
		return 0
	}
	le := st.arr[arrlen-1 : arrlen]
	return le[0]
}

func (st *Stack) Size() int {
	return len(st.arr)
}

type StackStr struct {
	arr []string
}

func (st *StackStr) Push(val string) {
	st.arr = append(st.arr, val)
}

func (st *StackStr) Pop() (string, error) {
	arrlen := len(st.arr)
	if arrlen == 0 {
		fmt.Println("no data exists in stack")
		return "", errors.New("no data exists")
	}
	le := st.arr[arrlen-1 : arrlen]
	st.arr = st.arr[:arrlen-1]
	return le[0], nil
}

func (st *StackStr) Seek() string {
	arrlen := len(st.arr)
	if arrlen == 0 {
		fmt.Println("no data exists in stack")
		return ""
	}
	le := st.arr[arrlen-1 : arrlen]
	return le[0]
}

func (st *StackStr) Size() int {
	return len(st.arr)
}
