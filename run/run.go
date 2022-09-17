package run

import "fmt"

func Run() {
	thirdRunner := &DefaultThirdRunner{Name: "Third", Speed: 3}
	secondRunner := &DefaultSecondRunner{Name: "Second", Speed: 2, Runner: thirdRunner}
	firstRunner := &DefaultFirstRunner{Name: "First", Speed: 1, Runner: secondRunner}
	fmt.Println("Speed:", firstRunner.FirstSpeedUp(10))
}
