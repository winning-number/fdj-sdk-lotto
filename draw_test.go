package lotto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func messyDraws() []Draw {
	return []Draw{
		dataSuperLottoV0()[0],
		dataClassicLottoV1()[1],
		dataClassicLottoV1()[0],
		dataSuperLottoV0()[1],
		dataClassicLottoV1()[1],
	}
}

func TestOrderDraws(t *testing.T) {
	t.Run("Should be ok to order draws by ASC order", func(t *testing.T) {
		expectedDraw := []Draw{
			dataClassicLottoV1()[0],
			dataClassicLottoV1()[1],
			dataClassicLottoV1()[1],
			dataSuperLottoV0()[0],
			dataSuperLottoV0()[1],
		}

		draws := messyDraws()
		OrderDraws(&draws, OrderASC)
		assert.Equal(t, expectedDraw, draws)

	})
	t.Run("Should be ok to order draws by DESC order", func(t *testing.T) {
		expectedDraw := []Draw{
			dataSuperLottoV0()[1],
			dataSuperLottoV0()[0],
			dataClassicLottoV1()[1],
			dataClassicLottoV1()[1],
			dataClassicLottoV1()[0],
		}

		draws := messyDraws()
		OrderDraws(&draws, OrderDESC)
		assert.Equal(t, expectedDraw, draws)
	})
}
