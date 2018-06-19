package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	var fp *os.File
	var err error

	if len(os.Args) < 2 {
		fp = os.Stdin
	}else{
		fp,err = os.Open(os.Args[1])
		if err != nil {
			// panicってなんぞ。exception的な?
			panic(err)
		}
		// 関数内でdeferを書いておくと、この関数が終了するときは必ず呼ばれる
		defer fp.Close()
	}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
