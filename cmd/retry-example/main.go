// Backoffを使ったリトライの例
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/goaux/backoff"
)

func main() {
	// Backoffのインスタンスを作成
	exponentialBackOff := backoff.NewExponential(
		// 実行感覚を100msに設定
		backoff.WithInitialInterval(100*time.Millisecond),
		// 最大リトライ回数を5回に設定
		backoff.WithMaxRetries(5),
	)
	ctx := context.Background()

	// リトライのループ
	for i := range exponentialBackOff(ctx) {
		fmt.Printf("Attempt %d\n", i)
		// 何かの操作を実行する（例: ネットワークリクエスト）
		// 処理が成功したらbreak出抜ける
	}
}
