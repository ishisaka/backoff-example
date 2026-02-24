// Backoffでの処理中にキャンセル付きコンテキストでのキャンセル処理の例。
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/goaux/backoff"
)

func main() {
	ctx := context.TODO()
	unit := 50 * time.Millisecond

	// [ExponentialBackOff](https://pkg.go.dev/github.com/cenkalti/backoff/v4#ExponentialBackOff)
	exponentialBackOff := backoff.NewExponential(
		// 100ミリ秒ごとに実行するように設定
		backoff.WithInitialInterval(2*unit),
		// リトライの間隔をランダム化しないように設定
		backoff.WithRandomizationFactor(0),
		// 最大リトライ回数を7回に設定
		backoff.WithMaxRetries(7),
	)
	start := time.Now()
	// キャンセル付きコンテキストを作成して、150ミリ秒後にキャンセルするように設定
	ctx, cancel := context.WithTimeout(ctx, 3*unit)
	defer cancel()
	for i := range exponentialBackOff(ctx) {
		fmt.Println("exponential#false", i, (time.Since(start) / unit * unit).String())
		if false {
			break
		}
	}
	fmt.Println("cancel", (time.Since(start) / unit * unit).String())
}
