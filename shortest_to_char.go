package praxis

//给定一个字符串 S 和一个字符 C。返回一个代表字符串 S 中每个字符到字符串 S 中的字符 C 的最短距离的数组。
//示例 1:
//输入: S = "loveleetcode", C = 'e'
//输出: [3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0]
//说明:
//字符串 S 的长度范围为 [1, 10000]。
//C 是一个单字符，且保证是字符串 S 里的字符。
//S 和 C 中的所有字母均为小写字母。

func shortestToChar(S string, C byte) []int {
	buf := []byte(S)
	if len(buf) == 0 {
		return []int{}
	}
	offsets := make([]int, len(buf))
	for key, value := range buf {
		if value == C {
			offsets[key] = 0
		} else {
			//找左边
			left, right := -1, -1
			if key > 0 {
				for index := key - 1; index >= 0; index-- {
					if buf[index] == C {
						left = key - index
						break
					}
				}
			}
			//找右边
			if key < len(buf) {
				for index := key + 1; index < len(buf); index++ {
					if buf[index] == C {
						right = index - key
						break
					}
				}
			}

			if left != -1 && right != -1 {
				if left <= right {
					offsets[key] = left
				} else {
					offsets[key] = right
				}
			} else {
				if left != -1 {
					offsets[key] = left
				} else {
					offsets[key] = right
				}
			}
		}
	}
	return offsets
}
