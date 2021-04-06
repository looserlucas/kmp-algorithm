package main

import "fmt"

func main() {
	t := "asbasdasadasdawesafbasdasbasd"
	p := "aabaabaaa"

	pre := make([]int, len(p))

	// preprocess
	pre[0] = 0
	i := 1
	j := 0
	for i < len(p) {
		if p[i] == p[j] {
			j++
			pre[i] = j
			i++
		} else {
			if j != 0 {
				j = pre[j-1]
			} else {
				pre[i] = 0
				i++
			}
		}
	}

	i = 0
	j = 0
	for i < len(t) {
		if p[j] == t[i] {
			i++
			j++
		}

		if j == len(p) {
			fmt.Println(i - j)
			j = pre[j-1]
		} else if i < len(t) && p[j] != t[i] {
			if j != 0 {
				j = pre[j-1]
			} else {
				i++
			}
		}
	}

	fmt.Println(p)
	fmt.Println(pre)
}

//a s b s a s d

//0 0 0 0 1 2
