package shark

import (
	"errors"
	"integrationtests/mocks"
	"integrationtests/prey"
	"integrationtests/simulator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHunt_FastSharkShortDistance_ReturnNotErr(t *testing.T) {
	// Arrange
	config := map[string]interface{}{
		"white_shark_speed": 12.5,
		"white_shark_x":     1.1,
		"white_shark_y":     2.1,
		"tuna_speed":        5.5,
	}
	maxTimeToCatch := 8.5
	simulator := simulator.NewCatchSimulator(maxTimeToCatch)
	storageMock := mocks.NewStorageMock(config)
	whiteShark := CreateWhiteShark(simulator, storageMock)
	tuna := prey.CreateTuna(storageMock)

	//Act
	err := whiteShark.Hunt(tuna)

	// Assert
	assert.NoError(t, err)
}

// Con testify
func TestHunt_FastSharkShortDistance2_ReturnNotErr(t *testing.T) {
	// Arrange
	maxTimeToCatch := 8.5
	simulator := simulator.NewCatchSimulator(maxTimeToCatch)

	storageMock := mocks.NewMockStorage()
	storageMock.On("GetValue", "white_shark_speed").Return(12.5)
	storageMock.On("GetValue", "white_shark_x").Return(1.1)
	storageMock.On("GetValue", "white_shark_y").Return(2.1)
	storageMock.On("GetValue", "tuna_speed").Return(5.5)

	whiteShark := CreateWhiteShark(simulator, storageMock)
	tuna := prey.CreateTuna(storageMock)

	//Act
	err := whiteShark.Hunt(tuna)

	// Assert
	assert.NoError(t, err)
	assert.True(t, storageMock.AssertExpectations(t))
}

// Sin testify como el primero 
func TestHunt_FastSharkShortDistance_ReturnErr(t *testing.T) {
	// Arrange
	config := map[string]interface{}{
		"white_shark_speed": 12.5,
		"white_shark_x":     1.1,
		"white_shark_y":     2.1,
		"tuna_speed":        25.5,
	}
	maxTimeToCatch := 8.5
	simulator := simulator.NewCatchSimulator(maxTimeToCatch)
	storageMock := mocks.NewStorageMock(config)
	whiteShark := CreateWhiteShark(simulator, storageMock)
	tuna := prey.CreateTuna(storageMock)

	//Act
	err := whiteShark.Hunt(tuna)

	// Assert
	assert.Error(t, err)
	assert.EqualError(t, errors.New("could not hunt the prey"), err.Error())
}
