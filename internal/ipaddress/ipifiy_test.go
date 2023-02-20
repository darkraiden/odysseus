package ipaddress_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/darkraiden/odysseus/internal/ipaddress"
	"github.com/darkraiden/odysseus/internal/mock/mockip"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIpify_NewService(t *testing.T) {
	t.Run("TestNewService_WithErrors", func(t *testing.T) {
		_, err := ipaddress.NewService(nil)
		require.Error(t, err)
	})

	t.Run("TestNewService_WithoutErrors", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		doer := mockip.NewMockDoer(ctrl)

		i, err := ipaddress.NewService(doer)
		require.NoError(t, err)
		require.NotNil(t, i)
	})
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
		if err != nil {
			t.Fatal(err)
		}

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
		require.NoError(t, err)
		require.NotNil(t, s)

		i, err := s.GetLocal()
		require.NoError(t, err)
		require.NotEmpty(t, i)
		assert.Equal(t, validIP, i)
	})
}
