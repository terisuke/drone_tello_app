package main

// import (
// 	"fmt"
// 	"os/exec"
// 	"time"

// 	"gobot.io/x/gobot"
// 	"gobot.io/x/gobot/platforms/dji/tello"
// 	"gobot.io/x/gobot/platforms/keyboard"
// )

// // 機体の移動と旋回に使用するパラメータ
// type KeyboardParams struct {
// 	move     int
// 	rotation int
// }

// func NewKeyboardParams(move int, rotation int) *KeyboardParams {
// 	return &KeyboardParams{
// 		move:     move,
// 		rotation: rotation,
// 	}
// }

// func main() {
// 	drone := tello.NewDriver("8888")
// 	keys := keyboard.NewDriver()
// 	kp := NewKeyboardParams(3, 20)
// 	var height int16
// 	var battery int8

// 	mplayer := exec.Command("mplayer", "-fps", "60", "-")
// 	mplayerIn, _ := mplayer.StdinPipe()
// 	if err := mplayer.Start(); err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	work := func() {
// 		drone.On(tello.FlightDataEvent, func(data interface{}) {
// 			flightData := data.(*tello.FlightData)
// 			height = flightData.Height
// 			battery = flightData.BatteryPercentage
// 		})

// 		keys.On(keyboard.Key, func(data interface{}) {
// 			key := data.(keyboard.KeyEvent)

// 			switch key.Key {
// 			// フライトデータ
// 			case keyboard.H:
// 				fmt.Println("Height:", height)
// 			case keyboard.B:
// 				fmt.Println("Battery:", battery)

// 			// 離陸・着陸
// 			case keyboard.Spacebar:
// 				fmt.Println("TakeOff")
// 				drone.TakeOff()
// 			case keyboard.Escape:
// 				fmt.Println("Land")
// 				drone.Land()

// 			// 上昇・下降
// 			case keyboard.ArrowUp:
// 				fmt.Println("Up")
// 				drone.Up(kp.move)
// 			case keyboard.ArrowDown:
// 				fmt.Println("Down")
// 				drone.Down(kp.move)

// 			// 進行
// 			case keyboard.W:
// 				fmt.Println("Forward")
// 				drone.Forward(kp.move)
// 			case keyboard.A:
// 				fmt.Println("Left")
// 				drone.Left(kp.move)
// 			case keyboard.S:
// 				fmt.Println("Backward")
// 				drone.Backward(kp.move)
// 			case keyboard.D:
// 				fmt.Println("Right")
// 				drone.Right(kp.move)
// 			case keyboard.F:
// 				fmt.Println("Hover")
// 				drone.Hover()

// 			// 旋回
// 			case keyboard.Q:
// 				fmt.Println("Clockwise")
// 				drone.Clockwise(kp.rotation)
// 			case keyboard.E:
// 				fmt.Println("CounterClockwise")
// 				drone.CounterClockwise(kp.rotation)
// 			case keyboard.R:
// 				fmt.Println("CeaseRotation")
// 				drone.CeaseRotation()

// 			// 特殊
// 			case keyboard.I:
// 				fmt.Println("FrontFlip")
// 				drone.FrontFlip()
// 			case keyboard.J:
// 				fmt.Println("LeftFlip")
// 				drone.LeftFlip()
// 			case keyboard.K:
// 				fmt.Println("BackFlip")
// 				drone.BackFlip()
// 			case keyboard.L:
// 				fmt.Println("RightFlip")
// 				drone.RightFlip()
// 			case keyboard.U:
// 				fmt.Println("Bounce")
// 				drone.Bounce()

// 			// 進行値セット (1〜100)
// 			case keyboard.One:
// 				kp.move = 1
// 				fmt.Println("move:", kp.move)
// 			case keyboard.Two:
// 				kp.move = 2
// 				fmt.Println("move:", kp.move)
// 			case keyboard.Three:
// 				kp.move = 3
// 				fmt.Println("move:", kp.move)
// 			case keyboard.Four:
// 				kp.move = 4
// 				fmt.Println("move:", kp.move)
// 			case keyboard.Five:
// 				kp.move = 5
// 				fmt.Println("move:", kp.move)
// 			case keyboard.Six:
// 				kp.move = 6
// 				fmt.Println("move:", kp.move)
// 			case keyboard.Seven:
// 				kp.move = 7
// 				fmt.Println("move:", kp.move)
// 			case keyboard.Eight:
// 				kp.move = 8
// 				fmt.Println("move:", kp.move)
// 			case keyboard.Nine:
// 				kp.move = 9
// 				fmt.Println("move:", kp.move)
// 			case keyboard.Zero:
// 				kp.move *= 10
// 				if kp.move > 100 {
// 					kp.move = 100
// 				}
// 				fmt.Println("move:", kp.move)

// 			// 旋回値セット (10〜100)
// 			case keyboard.C:
// 				kp.rotation += 10
// 				if kp.rotation > 100 {
// 					kp.rotation = 10
// 				}
// 				fmt.Println("rotation:", kp.rotation)
// 			case keyboard.X:
// 				kp.rotation -= 10
// 				if kp.rotation < 10 {
// 					kp.rotation = 100
// 				}
// 				fmt.Println("rotation:", kp.rotation)

// 			// 設定パラメータ確認
// 			case keyboard.P:
// 				fmt.Println("move:", kp.move, "rotation:", kp.rotation)
// 			}
// 		})

// 		drone.On(tello.ConnectedEvent, func(data interface{}) {
// 			fmt.Println("Connected")
// 			drone.StartVideo()
// 			drone.SetVideoEncoderRate(tello.VideoBitRateAuto)
// 			gobot.Every(100*time.Millisecond, func() {
// 				drone.StartVideo()
// 			})
// 		})

// 		drone.On(tello.VideoFrameEvent, func(data interface{}) {
// 			pkt := data.([]byte)
// 			if _, err := mplayerIn.Write(pkt); err != nil {
// 				fmt.Println(err)
// 			}
// 		})
// 	}

// 	robot := gobot.NewRobot("tello",
// 		[]gobot.Connection{},
// 		[]gobot.Device{drone, keys},
// 		work,
// 	)

// 	robot.Start()
// }
