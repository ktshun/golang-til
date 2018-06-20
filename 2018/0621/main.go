package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var fp *os.File
	var err error

	if len(os.Args) < 2 {
		fp = os.Stdin
	} else {
		fp, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		// 関数内でdeferを書いておくと、この関数が終了するときは必ず呼ばれる
		defer fp.Close()
	}

	queue := make(chan string, 10)

	go readFile(fp, queue)

	for i := range queue {
		fmt.Println(i)
	}
}

/**
 * ファイルを1行ずつ読んでチャネルに詰める
 */
func readFile(fp *os.File, queue chan string) {

	defer close(queue)

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		queue <- scanner.Text()
	}
}
