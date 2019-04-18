package praxis

//给定一个整数 n, 返回从 1 到 n 的字典顺序。
//例如，
//给定 n =1 3，返回 [1,10,11,12,13,2,3,4,5,6,7,8,9] 。
//请尽可能的优化算法的时间复杂度和空间复杂度。 输入的数据 n 小于等于 5,000,000。

type Array struct {
	Arr    []int
	Offset int
}

func dfs(t int, n int, m int, arr *Array) {
	if m > n {
		return
	}
	index := 0
	if t == 0 {
		index = 1
	} else {
		arr.Arr[arr.Offset] = m
		arr.Offset++
	}
	for ; index < 10; index++ {
		dfs(t+1, n, m*10+index, arr)
	}
}

func lexicalOrder(n int) []int {
	//arr := make([]int, n)
	arr := &Array{make([]int, n), 0}
	dfs(0, n, 0, arr)
	return arr.Arr
}
