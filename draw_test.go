package lotto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func messyDraws() []Draw {
	return []Draw{
		TestDataSuperLottoV0[0],
		TestDataClassicLottoV1[1],
		TestDataClassicLottoV1[0],
		TestDataSuperLottoV0[1],
		TestDataClassicLottoV1[1],
	}
}

func TestOrderDraws(t *testing.T) {
	t.Run("Should be ok to order draws by ASC order", func(t *testing.T) {
		expectedDraw := []Draw{
			TestDataClassicLottoV1[0],
			TestDataClassicLottoV1[1],
			TestDataClassicLottoV1[1],
			TestDataSuperLottoV0[0],
			TestDataSuperLottoV0[1],
		}

		draws := messyDraws()
		OrderDraws(&draws, OrderASC)
		assert.Equal(t, expectedDraw, draws)

	})
	t.Run("Should be ok to order draws by DESC order", func(t *testing.T) {
		expectedDraw := []Draw{
			TestDataSuperLottoV0[1],
			TestDataSuperLottoV0[0],
			TestDataClassicLottoV1[1],
			TestDataClassicLottoV1[1],
			TestDataClassicLottoV1[0],
		}

		draws := messyDraws()
		OrderDraws(&draws, OrderDESC)
		assert.Equal(t, expectedDraw, draws)
	})
}
