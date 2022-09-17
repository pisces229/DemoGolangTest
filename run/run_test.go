package run

import (
	mock "demo.golang.test/mock/run"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Run_First(test *testing.T) {
	//test.Skip("Skip")
	controller := gomock.NewController(test)
	defer controller.Finish()
	// Arrange
	mockSecondRunner := mock.NewMockSecondRunner(controller)
	gomock.InOrder(
		mockSecondRunner.EXPECT().GetName().Times(2).Return("Second"),
		mockSecondRunner.EXPECT().SecondSpeedUp(gomock.Any()).Times(1).Return(20),
		mockSecondRunner.EXPECT().GetName().Times(2).Return("Second"),
	)
	// Act
	runner := &DefaultFirstRunner{Name: "First", Speed: 1, Runner: mockSecondRunner}
	result := runner.FirstSpeedUp(10)
	// Assert
	assert.Equal(test, result, 20)
}

func Test_Run_Second(test *testing.T) {
	//test.Skip("Skip")
	controller := gomock.NewController(test)
	defer controller.Finish()
	// Arrange
	mockThirdRunner := mock.NewMockThirdRunner(controller)
	gomock.InOrder(
		mockThirdRunner.EXPECT().GetName().Times(3).Return("Third"),
		mockThirdRunner.EXPECT().ThirdSpeedUp(gomock.Any()).Times(1).Return(30),
		mockThirdRunner.EXPECT().GetName().Times(3).Return("Third"),
	)
	// Act
	runner := &DefaultSecondRunner{Name: "Second", Speed: 2, Runner: mockThirdRunner}
	result := runner.SecondSpeedUp(10)
	// Assert
	assert.Equal(test, result, 60)
}
