package main

import (
	"strings"
)

type Trie struct {
	root *node
}

type node struct {
	children map[rune]*node
}

func InitTrie() *Trie {
	return &Trie{&node{children: make(map[rune]*node)}}
}

func (t *Trie) Insert(word string) {
	word = strings.ToLower(word)
	currentNode := t.root

	for _, ch := range word {
		if _, ok := currentNode.children[ch]; ok {
			currentNode = currentNode.children[ch]
			continue
		}
		newNode := node{children: make(map[rune]*node)}
		currentNode.children[ch] = &newNode
		currentNode = &newNode
	}
	currentNode.children['*'] = &node{}
}

func (t *Trie) Search(word string) *node {
	word = strings.ToLower(word)
	currentNode := t.root
	for _, ch := range word {
		if _, ok := currentNode.children[ch]; ok {
			currentNode = currentNode.children[ch]
			continue
		}
		return nil
	}
	return currentNode
}

func (t *Trie) AutoCorrection(word string) *[]string {
	word = strings.ToLower(word)
	currentNode := t.root
	wordFound := ""
	for _, ch := range word {
		if _, ok := currentNode.children[ch]; ok {
			currentNode = currentNode.children[ch]
			wordFound = wordFound + string(ch)
			continue
		}
		return getAllWords(currentNode, wordFound, &[]string{})
	}
	return getAllWords(currentNode, word, &[]string{})
}

func (t *Trie) AutoComplete(word string) *[]string {
	word = strings.ToLower(word)
	currentNode := t.root
	for _, ch := range word {
		if _, ok := currentNode.children[ch]; ok {
			currentNode = currentNode.children[ch]
			continue
		}
		return nil
	}
	return getAllWords(currentNode, word, &[]string{})
}

func getAllWords(n *node, word string, words *[]string) *[]string {

	for ch, val := range n.children {

		if ch == '*' {
			*words = append(*words, word)
		}
		getAllWords(val, word+string(ch), words)
	}
	return words

}
