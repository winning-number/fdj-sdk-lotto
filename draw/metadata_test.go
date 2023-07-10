package draw

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetadata_setID(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		var m Metadata

		m.FDJID = "123456"
		m.TirageOrder = 1
		m.setID("suffix")
		assert.Equal(t, "1234561suffix", m.ID)
	})
}
