# 20180621

### Goの並列処理
Goで並列処理するにはGoRoutineを利用する。  
実行する関数の頭に`go`とつけるだけ。これで別スレッドで実行してくれる。

しかし、GoRoutineだけではその別スレッドを待たずにメインスレッドが終わってしまうので無意味。
これを待つようにするにはGoChannelかsync.WaitGroupを利用すると良さそう。

### Go Channel
厳密にはスレッドを待つためのものというより別スレッドとの通信用と言ったほうが正しいと思われる。  
とはいえ、メインスレッドは各GoRoutineから決められた数だけ（`<-c`の部分）送信を受けるまで、待機するので同期用にも使われる。

```go
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
```

なお、チャネルは`ch := make(chan int, 2)`のように第二引数を与えることでバッファをもたせることもできるが、バッファからあふれるとエラーが発生するので注意。
また、switch文のような形をとるselect文で複数あるチャネルのうち、どれかに送られてきたら所定の動作をして、終了という形にもできる。

### syncパッケージ
syncパッケージのwaitGroupを利用すると、チャネルを使わなくても待ってくれる。

```go
package main

import (
    "log"
    "sync"
    "time"
)

func main() {
    log.Print("started.")

    // 配列
    funcs := []func(){
        func() {
            // 1秒かかるコマンド
            log.Print("sleep1 started.")
            time.Sleep(1 * time.Second)
            log.Print("sleep1 finished.")
        },
        func() {
            // 2秒かかるコマンド
            log.Print("sleep2 started.")
            time.Sleep(2 * time.Second)
            log.Print("sleep2 finished.")
        },
        func() {
            // 3秒かかるコマンド
            log.Print("sleep3 started.")
            time.Sleep(3 * time.Second)
            log.Print("sleep3 finished.")
        },
    }

    var waitGroup sync.WaitGroup

    // 関数の数だけ並行化する
    for _, sleep := range funcs {
        waitGroup.Add(1) // 待つ数をインクリメント

        // Goルーチンに入る
        go func(function func()) {
            defer waitGroup.Done() // 待つ数をデクリメント
            function()
        }(sleep)

    }

    waitGroup.Wait() // 待つ数がゼロになるまで処理をブロックする

    log.Print("all finished.")
}
```

この辺はかなり奥が深そう。
https://hori-ryota.com/blog/golang-channel-pattern/

はかなり参考になる。

##### 参考
* https://go-tour-jp.appspot.com/concurrency/1
* https://qiita.com/suin/items/82ecb6f63ff4104d4f5d
* https://golang.org/pkg/sync/#WaitGroup