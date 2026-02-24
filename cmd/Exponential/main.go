// 指数バックオフを使ったリトライの例
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/goaux/backoff"
)

func main() {
	unit := 50 * time.Millisecond
	// Backoffのインスタンスを作成
	exponentialBackOff := backoff.NewExponential(
		// 実行感覚を100msに設定
		backoff.WithInitialInterval(2*unit),
		// リトライ間隔の倍率を2.0に設定
		backoff.WithMultiplier(2.0),
		// 最大リトライ回数を7回に設定
		backoff.WithMaxRetries(7),
	)
	ctx := context.Background()
	start := time.Now()
	// リトライのループ
	for i := range exponentialBackOff(ctx) {
		fmt.Printf("Attempt %d, Elapsed Time: %s\n", i, (time.Since(start) / unit * unit).String())
		// 何かの操作を実行する（例: ネットワークリクエスト）
		// 処理が成功したらbreak出抜ける
	}
}
