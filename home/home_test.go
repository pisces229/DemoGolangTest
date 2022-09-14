package home

import (
	"testing"

	"demo.golang.test/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test__Home__Get_AnimalName(test *testing.T) {
	cntroller := gomock.NewController(test)
	defer cntroller.Finish()
	// Arrange
	mockAniaml := mock.NewMockAnimal(cntroller)
	gomock.InOrder(
		mockAniaml.EXPECT().Get_Name().Times(2).Return("Golang"),
	)
	// Act
	home := NewHome(mockAniaml)
	reuslt1 := home.Get_AnimalName()
	reuslt2 := home.Get_AnimalName()
	// Assert
	assert.Equal(test, reuslt1, "Golang")
	assert.Equal(test, reuslt2, "Golang")
}

func Test__Home__Get_AnimalSpeed(test *testing.T) {
	cntroller := gomock.NewController(test)
	defer cntroller.Finish()
	// Arrange
	mockAniaml := mock.NewMockAnimal(cntroller)
	gomock.InOrder(
		mockAniaml.EXPECT().Get_Name().Times(1).Return(""),
		mockAniaml.EXPECT().Get_Speed(gomock.Any()).Times(1).Return(10),
		mockAniaml.EXPECT().Get_Name().Times(1).Return(""),
		// mockAniaml.EXPECT().Get_Speed(1).Times(1).Return(20),
		mockAniaml.EXPECT().Get_Speed(gomock.Any()).Times(1).Return(20),
	)
	// Act
	home := NewHome(mockAniaml)
	reuslt1 := home.Get_AnimalSpeed(1)
	reuslt2 := home.Get_AnimalSpeed(2)
	// Assert
	assert.Equal(test, reuslt1, 10)
	assert.Equal(test, reuslt2, 20)
}
