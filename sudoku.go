package main

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

// 0:未入力
// 1-9:入力済み
type Board [9][9]int

func pretty(b Board) string {
	var buf bytes.Buffer
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			buf.WriteString("+---+---+---+\n")
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				buf.WriteString("|")
			}
			buf.WriteString(strconv.Itoa(b[i][j]))

		}
		buf.WriteString("|\n")
	}
	buf.WriteString("+---+---+---+\n")
	return buf.String()
}

func duplicated(c [10]int) bool {
	for k, v := range c {
		if k == 0 {
			continue
		}
		if v >= 2 {

			return true
		}
	}
	return false
}

func verify(b Board) bool {
	//行チェック
	for i := 0; i < 9; i++ {
		//出現回数
		var c [10]int
		for j := 0; j < 9; j++ {
			c[b[i][j]]++
		}
		if duplicated(c) {
			fmt.Println("duplicated gyo check")
			return false
		}
	}
	//列チェック
	for i := 0; i < 9; i++ {
		//出現回数
		var d [10]int
		for j := 0; j < 9; j++ {
			d[b[j][i]]++
		}
		if duplicated(d) {
			fmt.Println("retus check")
			return false
		}
	}

	// 3x3チェック
	for i := 0; i < 9; i += 3 {
		//fmt.Println(i)
		for j := 0; j < 9; j += 3 {
			//fmt.Println(j)
			var e [10]int
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					//fmt.Println(b[row][col])
					e[b[row][col]]++
				}
			}
			if duplicated(e) {
				fmt.Println("duplicated  3x3 check")
				return false
			}
			fmt.Println("-------")
		}
	}

	return true
}

func solved(b Board) bool {
	if !verify(b) {
		return false
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func backtrack(b *Board) bool {

	if solved(*b) {
		return true
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			//0なら探索を開始する
			if b[i][j] == 0 {
				for c := 9; c >= 1; c-- {
					b[i][j] = c
					fmt.Printf("%+v\n", pretty(*b))
					time.Sleep(time.Second * 1)
					//もし数字がルールに適合するなら
					if verify(*b) {
						fmt.Println("verify true")
						//更に深く探索する
						if backtrack(b) {
							return true
						}
					}
				}
				return false
			}
		}
	}
	return false
}

func main() {
	b := Board{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	fmt.Printf("%+v\n", pretty(b))
}
