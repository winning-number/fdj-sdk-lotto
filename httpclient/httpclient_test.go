package httpclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		client := New()
		assert.NotNil(t, client)
	})
}
