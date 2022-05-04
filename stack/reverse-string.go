package main

func reverseString(str string) string {

	queue := StackStr{}
	for i := 0; i < len(str); i++ {
		queue.Push(string(str[i]))
	}
	var str1 string
	for queue.Size() > 0 {
		v, _ := queue.Pop()
		str1 = str1 + v
	}

	return str1
}

func reverseString2(str string) string {

	var str1 []byte
	for i := len(str) - 1; i >= 0; i-- {
		str1 = append(str1, str[i])
	}
	return string(str1)
}
