package storage

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestGetValue(t *testing.T){


	t.Run("ok", func(t *testing.T){
		
		stor := NewStorage()
		resul := 1.1
		result := stor.GetValue("white_shark_speed")

		assert.Nil(t, resul, result)
		
		//stubbie := stubRepository{file: "hola"}
		//stubbietwo:= NewStorage()
	
		//newer := stubbie.GetValue(stubbie.file)
		//stubbietwo.GetValue("hola.json")
		//assert.Error(t, ErrInternal)
		//assert.NotEqual(t, stubbietwo, stubbie)
		/*mockservice := NewStorage()
		
		newer := mockservice.GetValue("hola.json")
		assert.Error*/
	})

	t.Run("GetValue error", func(t *testing.T){
		
	})
}