package mocks

import (
	"integrationtests/pkg/storage"

	"github.com/stretchr/testify/mock"
)

type mockStorage struct {
	mock.Mock
	storage.Storage
}

func NewMockStorage() *mockStorage {
	return &mockStorage{}
}

func (ms *mockStorage) GetValue(key string) interface{} {
	args := ms.Called(key)
	return args.Get(0)
}

/*
func TestHunt_FastSharkShortDistance2_ReturnNotErr(t *testing.T) {
	// Arrange
	maxTimeToCatch := 8.5
	simulator := simulator.NewCatchSimulator(maxTimeToCatch)

	storageMock := newMockStorage()
	storageMock.On("GetValue", "white_shark_speed").Return(12.5)
	storageMock.On("GetValue", "white_shark_x").Return(1.1)
	storageMock.On("GetValue", "white_shark_y").Return(2.1)
	storageMock.On("GetValue", "tuna_speed").Return(5.5)

	whiteShark := shark.CreateWhiteShark(simulator, storageMock)
	tuna := prey.CreateTuna(storageMock)

	//Act
	err := whiteShark.Hunt(tuna)

	// Assert
	assert.NoError(t, err)
	assert.True(t, storageMock.AssertExpectations(t))
}
*/
