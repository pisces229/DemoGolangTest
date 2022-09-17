package home

import (
	"testing"

	mock "demo.golang.test/mock/home"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_Home_GetAnimalName(test *testing.T) {
	controller := gomock.NewController(test)
	defer controller.Finish()
	// Arrange
	mockAnimal := mock.NewMockAnimal(controller)
	gomock.InOrder(
		mockAnimal.EXPECT().GetName().Times(2).Return("Golang"),
	)
	// Act
	home := NewHome(mockAnimal)
	result1 := home.GetAnimalName()
	result2 := home.GetAnimalName()
	// Assert
	assert.Equal(test, result1, "Golang")
	assert.Equal(test, result2, "Golang")
}

func Test_Home_GetAnimalSpeed(test *testing.T) {
	controller := gomock.NewController(test)
	defer controller.Finish()
	// Arrange
	mockAnimal := mock.NewMockAnimal(controller)
	gomock.InOrder(
		mockAnimal.EXPECT().GetName().Times(1).Return(""),
		mockAnimal.EXPECT().GetSpeed(gomock.Any()).Times(1).Return(10),
		mockAnimal.EXPECT().GetName().Times(1).Return(""),
		// mockAnimal.EXPECT().Get_Speed(1).Times(1).Return(20),
		mockAnimal.EXPECT().GetSpeed(gomock.Any()).Times(1).Return(20),
	)
	// Act
	home := NewHome(mockAnimal)
	result1 := home.GetAnimalSpeed(1)
	result2 := home.GetAnimalSpeed(2)
	// Assert
	assert.Equal(test, result1, 10)
	assert.Equal(test, result2, 20)
}
