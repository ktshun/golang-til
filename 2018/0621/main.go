package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
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
	waitGroup := new(sync.WaitGroup)

	go readFile(fp, queue)

	for i := range queue {
		waitGroup.Add(1)
		go waitAndPrintStr(i, waitGroup)
	}

	fmt.Println("読み込みは終わったよ")

	waitGroup.Wait()

	fmt.Println("表示まで終わったよ")
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

/**
 * 100ms待って標準出力に出力
 */
func waitAndPrintStr(s string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	time.Sleep(time.Millisecond * 1000)
	fmt.Println(s)
}
