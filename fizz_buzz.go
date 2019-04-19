package praxis

import "fmt"

//写一个程序，输出从 1 到 n 数字的字符串表示。
//
//1. 如果 n 是3的倍数，输出“Fizz”；
//
//2. 如果 n 是5的倍数，输出“Buzz”；
//
//3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。
//
//示例：
//
//n = 15,
//
//返回:
//[
//"1",
//"2",
//"Fizz",
//"4",
//"Buzz",
//"Fizz",
//"7",
//"8",
//"Fizz",
//"Buzz",
//"11",
//"Fizz",
//"13",
//"14",
//"FizzBuzz"
//]

func fizzBuzz(n int) []string {
	strs := make([]string, 0, 32)
	for i := 1; i <= n; i++ {
		f := i%3 == 0
		b := i%5 == 0
		if f && b {
			strs = append(strs, "FizzBuzz")
		} else if f {
			strs = append(strs, "Fizz")
		} else if b {
			strs = append(strs, "Buzz")
		} else {
			strs = append(strs, fmt.Sprint(i))
		}
	}
	return strs
}
