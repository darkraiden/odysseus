package ipaddress_test

import (
	"bytes"
	"errors"
	"github.com/darkraiden/odysseus/internal/ipaddress"
	"github.com/darkraiden/odysseus/internal/mock/mockip"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestNewService(t *testing.T) {
	//TODO: Davide to test
}

func TestIpify_GetLocal(t *testing.T) {
	const getIPURL = "https://api.ipify.org?format=text"
	t.Run("returns an error given the get request fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		doer := mockip.NewMockDoer(ctrl)
		someErr := errors.New("some-error")

		req, err := http.NewRequest("GET", getIPURL, nil)
		require.NoError(t, err)
		require.NotNil(t, req)

		doer.EXPECT().Do(req).Return(nil, someErr)

		s, err := ipaddress.NewService(doer)
		i, err := s.GetLocal()

		assert.Empty(t, i)
		assert.Error(t, err)
	})

	t.Run("returns an ip address given a successful call", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		doer := mockip.NewMockDoer(ctrl)

		req, err := http.NewRequest("GET", getIPURL, nil)
		require.NoError(t, err)
		require.NotNil(t, req)

		validIP := "192.168.0.1"

		validRes := &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(validIP)),
		}

		doer.EXPECT().Do(req).Return(validRes, nil)

		s, err := ipaddress.NewService(doer)
		i, err := s.GetLocal()

		require.NoError(t, err)
		require.NotEmpty(t, i)

		assert.Equal(t, validIP, i)
	})
}
