package csvparser

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestNewContextRecord(t *testing.T) {
	errObjConvertion := errors.New("bad type convertion for the recording")

	t.Run("Should be ok", func(t *testing.T) {
		var ctx *contextRecord
		var ok bool
		ret := NewContextRecord()
		if ctx, ok = ret.(*contextRecord); !ok {
			t.Error(errObjConvertion)
		}

		if assert.NotNil(t, ret) {
			assert.NotNil(t, ctx.params)
		}
	})
}

func TestContextRecord_Get(t *testing.T) {
	t.Run("should not find the value associated to key", func(t *testing.T) {
		ctx := &contextRecord{params: make(map[string]string)}

		v, ok := ctx.Get("key")
		if assert.False(t, ok) {
			assert.Equal(t, v, "")
		}
	})
	t.Run("should find the value associated to key", func(t *testing.T) {
		key := "key"
		value := "value"
		ctx := &contextRecord{params: make(map[string]string)}
		ctx.params[key] = value

		v, ok := ctx.Get(key)
		if assert.True(t, ok) {
			assert.Equal(t, v, value)
		}
	})
}

func TestContextRecord_Set(t *testing.T) {
	key := "key"
	value := "value"
	t.Run("should set the value for the key parameter", func(t *testing.T) {
		ctx := &contextRecord{params: make(map[string]string)}

		ctx.Set(key, value)
		if assert.NotEmpty(t, ctx.params) {
			assert.Equal(t, ctx.params[key], value)
		}
	})
	t.Run("should find the value associated to key", func(t *testing.T) {
		key := "key"
		value := "value"
		expectedValue := "value2"
		ctx := &contextRecord{params: make(map[string]string)}
		ctx.params[key] = value

		ctx.Set(key, expectedValue)
		if assert.NotEmpty(t, ctx.params) {
			assert.Equal(t, ctx.params[key], expectedValue)
		}
	})
}
