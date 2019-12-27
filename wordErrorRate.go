/*
Name: Ciaran Cooney
Date: 20/12/2019
Description: Implementation of Levenstein Distance as described below
 */
package main

import (
	"fmt"
)

func zeroMat(x int, y int) [][]int {
	/*
	generates an x-by-y dimensional matrix of zeros
	input x (int): length of x-axis
	      y (int): length of y-axis
	output d [x][y] matrix of zeros
	 */
	d := make([][]int, x)
	for i := range d {
		d[i] = make([]int, y)
	}
	return d
}

func minOfArray(arr []int) int {
	min := arr[0]

	for _, num := range arr {
		if min > num {
			min = num
		}
	}
	return min
}

func maxOfArray(arr []int) int {
	max := arr[0]

	for _, num := range arr {
		if max < num {
			max = num
		}
	}
	return max
}

func maxOfInts(x, y int) int {
	if x < y {
		return y
	}
	return x
}

type updates struct {
	a int
	b int
	c int
}

var substitutionCost int

func WordErrorRateMain(str1 string, str2 string) (int, float64, float64) {
	/*
	Simple implementation of the Levenstein Distance and word error rate (WER)
	Distance computed with iterative matrix
	WER computed using formula from [2]
		[1] Levenshtein, Vladimir I. (February 1966). "Binary codes capable of correcting
		    deletions, insertions, and reversals". Soviet Physics Doklady. 10 (8): 707–710.
		[2] Soukoreff, R. W., & MacKenzie, I. S. (2001). Measuring errors in text entry tasks:
		    An application of the Levenshtein string distance statistic. Extended Abstracts of the
		    ACM Conference on Human Factors in Computing Systems - CHI 2001, pp. 319-320
			(WER: https://www.yorku.ca/mack/CHI01a.html)
		Python implementation of the levenshtein distance computation:
		input: str1: string
			   str2: string
	    output: Levenstein Distance (int)
				Word Error Rate (float64)
				Word Accuracy (float64)
		"""
	 */
	if str1 == str2 {
		return 0, 0 , 0
	}
	if len(str1) == 0 {
		return len(str2), 100, 0
	}
	if len(str2) == 0 {
		return len(str1), 100, 0
	}

	d := zeroMat(len(str1)+1, len(str2)+1)

	for i := range d[:][0]{
		d[i][0] = i
	}

	for j := range d[0][:] {
		d[0][j] = j
	}

	for j := 1; j<=len(str2); j++{

		for i := 1; i<=len(str1); i++ {

			if str1[i-1] == str2[j-1] {
				substitutionCost = 0
			} else {
				substitutionCost = 1
			}
			aU := d[i-1][j] + 1
			bU := d[i][j-1] + 1
			cU := d[i-1][j-1] + substitutionCost
			update := []int{aU,bU,cU}
			updateMin := minOfArray(update)
			d[i][j] = updateMin

		}
	}

	WER := float64(d[len(str1)][len(str2)]) / float64(maxOfInts(len(str1), len(str2))) * 100
	WACC := 100 - WER
	fmt.Println(WACC)
	return d[len(str1)][len(str2)], WER, WACC
}

func main () {
	distance, WER, WAcc := WordErrorRateMain("hello", "world")
	fmt.Println("Levenstein Distance: ", distance) // 4
	fmt.Println("Word Error Rate: ", WER)          // 80
	fmt.Println("Word Accuracy: ", WAcc)           // 20
}