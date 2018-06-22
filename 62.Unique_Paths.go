package main

import (
	"fmt"
)

/*
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?

+------+-------+-------+--------+-------+--------+---------
|start |       |       |        |       |        |        |
+---------------------------------------------------------+
|      |       |       |        |       |        |        |
+---------------------------------------------------------+
|      |       |       |        |       |        | finish |
+------+-------+-------+--------+-------+--------+--------+
Above is a 3 x 7 grid. How many possible unique paths are there?
Note: m and n will be at most 100.
*/


type Matrix [][]int

func constructMatrix(rows, cols int) Matrix {
    m := make([][]int, rows)
    for i := 0; i < rows; i++ {
        m[i] = make([]int, cols)
    }
    return m
}

//DP is a better way, and recursion
func uniquePaths(m int, n int) int {
    x :=  constructMatrix(m+1, n+1)
    for i:=1; i<m+1; i++{
        x[i][1] =  1
    }
    for i:=1; i<n+1; i++{
        x[1][i] =  1
    }
    for i:=1; i<m+1; i++{
        for j:=1; j<n+1; j++{
        }
    }


    for i:=2; i<m+1; i++{
        for j:=2; j<n+1; j++{
            x[i][j] = x[i-1][j] + x[i][j-1]
        }
    }
    return x[m][n]
}

func main() {
    fmt.Println(uniquePaths(3, 7))
}
