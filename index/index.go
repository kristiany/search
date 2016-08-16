package index

import (
	"github.com/timtadh/data-structures/trie"
	"strings"
)

type Index struct {
	content map[string]*trie.TST
}

type Result struct {
	Filename string
	Score    int
}

func New() *Index {
	return &Index{content: make(map[string]*trie.TST) }
}

func (i *Index) AddToIndex(name string, content string) {
	var root = trie.New()
	for _, word := range strings.Fields(content) {
		root.Put([]byte(word), 1.0)
	}
	i.content[name] = root
}

func (i *Index) Search(words []string) []Result {
	var result = make([]Result, 0)
	for filename, collection := range i.content {
		var score = 0
		for _, word := range words {
			kvi := collection.PrefixFind([]byte(word))
			key, _, _ := kvi()
			if key != nil { // We just care about any hit
				score += 1
			}
		}
		var normalized = score * 100 / len(words)
		result = append(result, Result { Filename: filename, Score: normalized})
	}
	return result
}

func (i *Index) String() string {
	var result string
	for key, value := range i.content {
		result += key + ":" + value.String() + "\n"
	}
	return result
}
