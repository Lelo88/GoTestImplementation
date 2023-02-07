package mocks

import (
	"integrationtests/pkg/storage"
)

type storageMock struct {
	Data map[string]interface{}
}

func NewStorageMock(data map[string]interface{}) storage.Storage {
	return &storageMock{
		Data: data,
	}
}

func (s *storageMock) GetValue(key string) interface{} {
	return s.Data[key]
}
