package main

import (
	_0064_minimum_path_sum "algorithm_go/src/dp/_0064.minimum-path-sum"
	"fmt"
)

func main() {
	//nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//max := _53_maximum_subarray.MaxSubArray(nums)
	//fmt.Println("max sub array result:", max)

	//m:=2
	//n:=3
	//res:=_0062_Unique_Paths.UniquePaths(m,n)
	//fmt.Println(res)

	//og:=[][]int{{0,0},{0,1}}
	//res:=_0063_unique_paths_ii.UniquePathsWithObstacles(og)
	//fmt.Println(res)
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	ret := _0064_minimum_path_sum.MinPathSum(grid)
	fmt.Println(ret)

}
