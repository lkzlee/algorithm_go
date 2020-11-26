package main

import (
	"algorithm_go/src/back_tracking"
	"flag"
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
	//grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	//ret := _0064_minimum_path_sum.MinPathSum(grid)
	//fmt.Println(ret)
	var filePath string
	var targetScore float64
	var diff float64
	flag.StringVar(&filePath, "f", "bin\\file.xlsx", "文件路径为空")
	flag.Float64Var(&targetScore, "target", 3900, "额度为空，默认:3900")
	flag.Float64Var(&diff, "diff", 0.5, "计算差值精度，默认:0.5")
	// 解析命令行参数写入注册的flag里
	flag.Parse()
	if filePath == "" {
		fmt.Println("文件路径不能为空filepath")
		return
	}
	fmt.Printf("文件路径file:%v\n报销额度Money:%v\n报销额度允许的精度差diff:%v\n", filePath, targetScore, diff)
	tickets := back_tracking.ParseXlsx(filePath)
	tmpList := make([]back_tracking.Ticket, 0, len(tickets))
	back_tracking.IfCalcResult = false
	back_tracking.CalcTicket(tickets, tmpList, 0, targetScore, diff)
	if back_tracking.IfCalcResult == false {
		fmt.Println("*********************")
		fmt.Println("计算不出来，o(╥﹏╥)o")
		fmt.Println("*********************")
	}
}
