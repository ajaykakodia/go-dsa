package main

import (
	"errors"
	"fmt"
)

var mappingBraces = map[string]string{
	"}": "{",
	")": "(",
	"]": "[",
}

func Linter(str string) error {

	stack := StackStr{}
	for i := 0; i < len(str); i++ {
		char := string(str[i])
		if isOpeningBrace(char) {
			stack.Push(char)
			continue
		}

		if isClosingBrace(char) {
			poppedBrace, err := stack.Pop()
			if err != nil {
				return errors.New(fmt.Sprintf("no open brace for %s and stack is empty", char))
			}
			mappedOpeningBrace := mappingOpeningBeraces(char)
			if poppedBrace != mappedOpeningBrace {
				return errors.New(fmt.Sprintf("no open brace for %s, get %s", char, poppedBrace))
			}
		}
	}

	if stack.Size() > 0 {
		return errors.New("missing close braces")
	}

	return nil
}

func isOpeningBrace(str string) bool {
	if str == "{" || str == "(" || str == "[" {
		return true
	}
	return false
}

func isClosingBrace(str string) bool {
	if str == "}" || str == ")" || str == "]" {
		return true
	}
	return false
}

func mappingOpeningBeraces(brace string) string {
	return mappingBraces[brace]
}
