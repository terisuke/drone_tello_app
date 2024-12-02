package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/eiannone/keyboard"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
)

func main() {
	drone := tello.NewDriver("8888")
	var flightData *tello.FlightData
	var battery int8
	
	const speed = 25     // 移動速度 (cm/s)
	const distance = 100 // 移動距離 (cm)
	moveTime := time.Duration(distance/speed) * time.Second // 移動に必要な時間を計算

	// プログラム終了用のチャネル
	done := make(chan bool)
	interrupted := make(chan bool)

	// ESCキー監視の開始
	go func() {
		if err := keyboard.Open(); err != nil {
			panic(err)
		}
		defer keyboard.Close()

		fmt.Println("ESCキーで実行を終了できます")
		for {
			char, key, err := keyboard.GetKey()
			if err != nil {
				panic(err)
			}
			if key == keyboard.KeyEsc {
				fmt.Println("\n緊急停止します")
				drone.Land() // 安全のため着陸
				interrupted <- true
				return
			}
			_ = char // 他のキー入力は無視
		}
	}()

	// Ctrl+C などのシグナルハンドリング
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChan
		fmt.Println("\nシグナルを受信しました。安全に終了します")
		drone.Land()
		interrupted <- true
	}()

	work := func() {
		// フライトデータのモニタリング設定
		drone.On(tello.FlightDataEvent, func(data interface{}) {
			flightData = data.(*tello.FlightData)
			battery = flightData.BatteryPercentage
			fmt.Printf("Height: %d cm, Battery: %d%%\n", flightData.Height, battery)
		})

		// 離陸
		fmt.Println("離陸を開始します")
		drone.TakeOff()
		
		// 安定するまで待機
		gobot.After(3*time.Second, func() {
			fmt.Printf("%d cm 前進します\n", distance)
			drone.Forward(speed)
			
			// 指定距離を移動後、停止
			gobot.After(moveTime, func() {
				drone.Forward(0) // 前進停止
				fmt.Println("前進完了、一時停止")
				
				// 1秒待機後、後退開始
				gobot.After(1*time.Second, func() {
					fmt.Printf("%d cm 後退します\n", distance)
					drone.Backward(speed)
					
					// 指定距離を移動後、停止
					gobot.After(moveTime, func() {
						drone.Backward(0) // 後退停止
						fmt.Println("後退完了、着陸準備")
						
						// 1秒待機後、着陸
						gobot.After(1*time.Second, func() {
							fmt.Println("着陸します")
							drone.Land()
							fmt.Printf("フライト完了, バッテリー残量: %d%%\n", battery)
							// すべての動作が完了したら終了シグナルを送信
							done <- true
						})
					})
				})
			})
		})
	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		work,
	)

	// ロボットを非同期で開始
	go robot.Start()

	// 終了待機
	select {
	case <-done:
		fmt.Println("プログラムが正常に完了しました")
	case <-interrupted:
		fmt.Println("プログラムが中断されました")
	}

	// クリーンアップ処理
	fmt.Println("プログラムを終了します")
	os.Exit(0)
}