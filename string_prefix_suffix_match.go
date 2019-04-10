package praxis

import "strings"

//给定多个 words，words[i] 的权重为 i 。
//设计一个类 WordFilter 实现函数WordFilter.f(String prefix, String suffix)。这个函数将返回具有前缀 prefix 和后缀suffix 的词的最大权重。如果没有这样的词，返回 -1。
//例子:
//输入:
//WordFilter(["apple"])
//WordFilter.f("a", "e") // 返回 0
//WordFilter.f("b", "") // 返回 -1
//注意:
//words的长度在[1, 15000]之间。
//对于每个测试用例，最多会有words.length次对WordFilter.f的调用。
//words[i]的长度在[1, 10]之间。
//prefix, suffix的长度在[0, 10]之前。
//words[i]和prefix, suffix只包含小写字母。

type WordFilter struct {
	words []string
}

func Constructor(words []string) WordFilter {
	return WordFilter{words}
}

func (this *WordFilter) F(prefix string, suffix string) int {
	for index := len(this.words) - 1; index >= 0; index-- {
		val := this.words[index]
		if strings.HasPrefix(val, prefix) && strings.HasSuffix(val, suffix) {
			return index
		}
	}
	return -1
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */
