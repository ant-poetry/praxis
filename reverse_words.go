package praxis

import "strings"

//给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
//示例 1:
//输入: "Let's take LeetCode contest"
//输出: "s'teL ekat edoCteeL tsetnoc"
//注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。

func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}
	builder := strings.Builder{}
	begin := 0
	for index, val := range s {
		if val == ' ' {
			for i := index - 1; i >= begin; i-- {
				builder.WriteByte(s[i])
			}
			builder.WriteByte(' ')
			begin = index + 1
		}
		if index+1 == len(s) {
			for i := index; i >= begin; i-- {
				builder.WriteByte(s[i])
			}
			break
		}
	}
	//splits := strings.Split(s," ")
	//for index, str := range splits {
	//	for i := len(str) - 1; i >= 0; i-- {
	//		builder.WriteByte(str[i])
	//	}
	//	if index < len(splits) - 1 {
	//		builder.WriteByte(' ')
	//	}
	//}
	return builder.String()
}
