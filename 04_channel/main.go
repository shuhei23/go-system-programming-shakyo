package main

import (
	//"context"
	"fmt"
	//"math"
	"time"
)

/* 4.1 */
// func sub() {
// 	fmt.Println("sub() is running")
// 	time.Sleep(time.Second)
// 	fmt.Println("sub() is finished")
// }

/* 4.2.3*/
// func primeNumber() chan int {
// 	result := make(chan int)
// 	go func() {
// 		result <- 2
// 		for i := 3; i < 100000; i += 2 {
// 			l := int(math.Sqrt(float64(i)))
// 			found := false // 約数見つけた
// 			for j := 3; j < l + 1; j += 2 {
// 				if i%j == 0{
// 					found = true;
// 				}
// 			}
// 			if !found {
// 				result <- i
// 			}
// 		}
// 		close(result)
// 	}()
// 	return result
// }

func main() {
	/* 4.1 */
	// fmt.Println("start sub()")
	// // goroutineを作って関数を実行
	// //go sub()
	// // インラインで無名関数を作ってその場でgoroutineで実行
	// go func(){
	// 	fmt.Println("sub() is running")
	// 	time.Sleep(time.Second)
	// 	fmt.Println("sub() is finished")
	// }()
	// time.Sleep(2 * time.Second)

	/* 4.2 */
	// fmt.Println("Start sub()")
	// // 終了を受け取るためのチャネル
	// done := make(chan bool)
	// go func() {
	// 	fmt.Println("sub() is finished")
	// 	// 終了を通知
	// 	done <- true
	// }()

	// <-done
	// fmt.Println("All tasks are finished")

	/* 4.2.3 */
	// pn := primeNumber()
	// // ここがポイント
	// for n := range pn{
	// 	fmt.Println(n)
	// }

	/* 4.2.5 */
	// fmt.Println("start sub()")
	// // 終了を受け取るための終了関数付きコンテキスト
	// ctx, ctxCancel := context.WithCancel(context.Background())

	// go func() {
	// 	fmt.Println("sub() is finished")
	// 	// 終了を通知
	// 	ctxCancel()
	// }()
	// // 終了を待つ
	// <-ctx.Done() // cancel()されるまでブロック
	// fmt.Println("all tasks are finished")

	/* Q4.1 */
	fmt.Println("----start----")
	now := time.Now()
	// <- time.Tick(3 * time.Second)
	<-time.After(3 * time.Second)
	fmt.Printf("精度悪いんですか？\n")
	fmt.Printf("経過: %vs\n", time.Since(now).Seconds())
}
