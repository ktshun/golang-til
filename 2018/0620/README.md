# 20180620

### ファイル入出力
bufioを利用することでファイルの入出力が可能。  
1行ずつ読むには
ReadLineとReadStringの二種類存在しているが、ReadStringは改行コードが\nのみの対応なのでReadLine推奨。

bufio.Scannerを利用するのが最も推奨される模様。

参考: https://qiita.com/ikawaha/items/28186d965780fab5533d

### panicとdefer
まず、panicは端的に言うとRuntimeエラーのようなもの。Exceptionよりさらにひどく、プログラムを停止させるもの。  
**基本的に使わない**と書かれていることが多い模様。

deferはpanicに対処できるものの一つ。  

```go
defer fp.Close()
```

のように書いておくと、記載した関数が終了するときにdeferで定義した文が呼ばれる。これはpanicで終了してしまったときも例外ではない。上記のようにファイルクローズに使うと便利かもしれない。
これを遅延指定関数と呼ぶ。


さらにdeferで定義した遅延指定関数内でのみrecoverが利用できる。  
recoverはpanicに渡された引数を取得することができる。つまり

```go
func main() {
	// 遅延指定
    defer func() {
    	// panicの引数がerrに入る
        if err := recover(); err != nil {
            log.Println("work failed:", err)
        }
    }()
    ...
    panic("Panic!!!")
}
```

のようにすると`work failed:Panic!!!`のように表示される事となる