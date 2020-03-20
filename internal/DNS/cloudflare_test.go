package DNS_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("InvalidParamError given an empty apiKey", func(t *testing.T) {
		a, err := New("", "", "")

		require.Nil(t, a)
		require.Error(t, err)

		var invalidError InvalidParamError
		assert.True(t, errors.As(err, &invalidError))
		assert.Equal(t, "apiKey", invalidError.Param)
	})

	t.Run("InvalidParamError given an empty email", func(t *testing.T) {
		a, err := New("some-valid-api-key", "", "")

		require.Nil(t, a)
		require.Error(t, err)

		var invalidError InvalidParamError
		assert.True(t, errors.As(err, &invalidError))
		assert.Equal(t, "email", invalidError.Param)
	})

	t.Run("InvalidParamError given an empty zoneName", func(t *testing.T) {
		a, err := New("some-valid-api-key", "some-valid-email", "")

		require.Nil(t, a)
		require.Error(t, err)

		var invalidError InvalidParamError
		assert.True(t, errors.As(err, &invalidError))
		assert.Equal(t, "zoneName", invalidError.Param)
	})
}
