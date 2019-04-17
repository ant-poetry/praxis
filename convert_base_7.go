package praxis

//给定一个整数，将其转化为7进制，并以字符串形式输出。
//示例 1:
//输入: 100
//输出: "202"
//示例 2:
//输入: -7
//输出: "-10"
//注意: 输入范围是 [-1e7, 1e7] 。

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	buf := make([]byte, 24)
	negative := num < 0
	if negative {
		num = -num
	}
	i := len(buf)
	for num >= 7 {
		i--
		next := num / 7
		buf[i] = byte('0' + num - (next * 7))
		num = next
	}
	if num > 0 {
		i--
		buf[i] = byte('0' + num)
	}
	if negative {
		i--
		buf[i] = byte('-')
	}
	return string(buf[i:])
}
