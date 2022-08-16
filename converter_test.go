package lotto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDateFormat(t *testing.T) {
	t.Run("Should return an error because the format is invalid", func(t *testing.T) {
		expectedErr := "parsing time \"20210213\" as \"2006-01-02\": cannot parse \"0213\" as \"-\""
		date, err := DateFormat("-", "20210213", false)

		if assert.Error(t, err) {
			assert.EqualError(t, err, expectedErr)
			assert.Empty(t, date)
		}
	})
	t.Run("Should be ok without separator", func(t *testing.T) {
		expectedDate := "2021-02-13 00:00:00 +0000 UTC"
		date, err := DateFormat("", "20210213", false)

		if assert.NoError(t, err) {
			assert.Equal(t, expectedDate, date.String())
		}
	})
	t.Run("Should be with a '/' separator", func(t *testing.T) {
		expectedDate := "2021-02-13 00:00:00 +0000 UTC"
		date, err := DateFormat("/", "13/02/2021", true)

		if assert.NoError(t, err) {
			assert.Equal(t, expectedDate, date.String())
		}
	})
}
