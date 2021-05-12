package main

import "fmt"

func main() {
	messge := "uv vagreangvbany fcnpr fgngvba"
	//解密代码
	for i := 0; i < len(messge); i++ {
		c := messge[i]
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}

}
