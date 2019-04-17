package praxis

//给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
//示例 1 :
//输入:nums = [1,1,1], k = 2
//输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
//数组的长度为 [1, 20,000]。
//数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。

func subarraySum(nums []int, k int) int {
	count := 0
	for index, _ := range nums {
		sum := 0
		for nindex := index; nindex < len(nums); nindex++ {
			sum += nums[nindex]
			if sum == k {
				count++
			}
		}
	}
	//count := 0
	//sum := 0
	//pre := 0
	//for _, val := range nums {
	//	sum += val;
	//	if sum == k {
	//		count++
	//		pre++
	//	} else if sum > k {
	//		sum -= nums[pre]
	//		pre++
	//	}
	//}
	return count
}
