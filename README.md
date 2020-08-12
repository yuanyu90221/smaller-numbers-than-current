# smaller-numbers-than-current

## 題目解讀：

### 題目來源:
[how-many-numbers-are-smaller-than-the-current-number](https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/)

### 原文:
Given the array nums, for each nums[i] find out how many numbers in the array are smaller than it. That is, for each nums[i] you have to count the number of valid j's such that j != i and nums[j] < nums[i].

Return the answer in an array.

Constraints:

2 <= nums.length <= 500
0 <= nums[i] <= 100
### 解讀：
給定一個正整數的陣列 nums, 對每個陣列元素 nums[i]

回傳一個陣列正整數陣列 results, 每個陣列元素 results[i] 代表在nums陣列中所有小於nums[i]的元素個數


## 初步解法:
### 初步觀察:

目前自己的想法:

對每個專注到都是比自己小的元素

可以考慮用map紀錄所有值

然後每次新增loop每個map裡面的結果

參考別人的想法：

把每個元素值得次數記錄下來

然後把每個小於自己值得次數加起來


### 初步設計:

原本自己的設計:

given a integer array nums

step 0: create a integer array result with length of nums

step 1: create an map valueMap = make(map[int]int)

step 2: create an map resultMap = make(map[int]int)

step 3: let idx = 0 

step 4: check idx >= length of nums go to step 6

step 5: set valueMap[nums[idx]] += 1, set resultMap[nums[idx]] = 0 go to step 4

step 6: loop valueMap for each k,v in loop resultMap for k1, v1 if k1 < k resultMap[k] += valueMap[k]

step 7: loop nums for each idx , result[idx] = valueMap[result[idx]]

step 8: return result

參考別人的設計：

given a integer array nums

step 0: create a integer array result with length of nums

step 1: create a integer array valueArray with length of 101 

step 2: create a integer array accuArray with length of 101

step 3: loop nums for idx = 0 to idx < nums.length set accuArray[nums[idx]]+= 1

step 4: loop accuArray for index = 1 to index < numaccuArrays.length set accuArray[idx] = accuArray[idx-1], valueArray[idx] = accuArray[idx-1]

step 5: loop result for idx = 0 to idx < result.length set result[idx] = valueArray[nums[idx]]

step 6: return result 

## 遇到的困難
### 題目上理解的問題
因為英文不是筆者母語

所以在題意解讀上 容易被英文用詞解讀給搞模糊

### pseudo code撰寫

一開始不習慣把pseudo code寫下來

因此 不太容易把自己的code做解析

### golang table driven test不熟
對於table driven test還不太熟析

所以對於寫test還是耗費不少時間
## 我的github source code
原本自己想法的解法
```golang
package smaller_numbers

func smallerNumbersThanCurrent(nums []int) []int {
	result := make([]int, len(nums))
	valueMap := make(map[int]int)
	resultMap := make(map[int]int)
	for _, val := range nums {
	 	valueMap[val] += 1
	 	resultMap[val] = 0
    }
    for k, _ := range valueMap {
	    for k1, _ := range valueMap {
            if k1 < k {
                resultMap[k] += valueMap[k1]
            }
 	    }
	}
    for idx, _ := range nums {
        result[idx] = resultMap[nums[idx]]
    }
    return result
}
```
參考其他人的寫法
```golang
package smaller_numbers

func smallerNumbersThanCurrent(nums []int) []int {
	result := make([]int, len(nums))
	//array to keep value
	accuArray := make([]int, 101)
	valueArray := make([]int, 101)
	for _, val := range nums {
		accuArray[val] += 1
	}
	for idx, _ := range valueArray {
		if idx > 0 {
			accuArray[idx] += accuArray[idx-1]
			valueArray[idx] = accuArray[idx-1]
		}
	}
	for idx, _ := range result {
		result[idx] = valueArray[nums[idx]]
	}
	return result
}
```

## 測資的撰寫

```golang
package smaller_numbers

import (
	"reflect"
	"testing"
)

func Test_smallerNumbersThanCurrent(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				nums: []int{8, 1, 2, 2, 3},
			},
			want: []int{4, 0, 1, 1, 3},
		},
		{
			name: "Example2",
			args: args{
				nums: []int{6, 5, 4, 8},
			},
			want: []int{2, 1, 0, 3},
		},
		{
			name: "Example3",
			args: args{
				nums: []int{7, 7, 7, 7},
			},
			want: []int{0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallerNumbersThanCurrent(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallerNumbersThanCurrent() = %v, want %v", got, tt.want)
			}
		})
	}
}


```

## my self record

[golang ironman 30 day 9th day](https://hackmd.io/yZ6r1OgoQHGGsrk8wFf8JQ?view)

## 參考文章

[golang test](https://ithelp.ithome.com.tw/articles/10204692)