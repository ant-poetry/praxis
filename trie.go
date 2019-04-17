package praxis

//实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
//示例:
//Trie trie = new Trie();
//trie.insert("apple");
//trie.search("apple");   // 返回 true
//trie.search("app");     // 返回 false
//trie.startsWith("app"); // 返回 true
//trie.insert("app");
//trie.search("app");     // 返回 true
//说明:
//你可以假设所有的输入都是由小写字母 a-z 构成的。
//保证所有输入均为非空字符串。

const MAX = 26

type Trie struct {
	word bool
	next [MAX]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	n := this
	for _, val := range word {
		index := int(val - 'a')
		next := n.next[index]
		if next == nil {
			next = &Trie{word: false}
			n.next[index] = next
		}
		n = next
	}
	n.word = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	n := this
	for _, val := range word {
		index := int(val - 'a')
		if n.next[index] == nil {
			return false
		}
		n = n.next[index]
	}
	return n.word
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	n := this
	for _, val := range prefix {
		index := int(val - 'a')
		if n.next[index] == nil {
			return false
		}
		n = n.next[index]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
