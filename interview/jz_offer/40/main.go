//输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

package main

import (
	"fmt"
	"sort"
	
)

// sort算法底层是快排 quickSort()

func getLeastNumbers(arr []int, k int) []int {
	// sort.Sort(arr,len(arr))

	// cp(arr[:k],)


}

func PrintSlice(data []int){
	for idx,d := range data{
		fmt.Printf("data[%v]=%v\n",idx,d)
	}
}

func main() {
	vec := []int{4,5,1,6,2,7,3,8}
	PrintSlice(vec)
	min_vec := getLeastNumbers(vec,4)
	PrintSlice(min_vec)
} 