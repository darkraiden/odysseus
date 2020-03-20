package odysseus_test

import (
	"errors"
	"github.com/darkraiden/odysseus/internal/mock/mockdns"
	mockip "github.com/darkraiden/odysseus/internal/mock/mockip"
	"github.com/darkraiden/odysseus/internal/odysseus"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewService(t *testing.T) {

}

func TestService_UpdateDNSWithLocalIP(t *testing.T) {
	t.Run("returns an error given an empty slice of records", func(t *testing.T) {
		s := odysseus.Service{}
		err := s.UpdateDNSWithLocalIP([]string{})

		assert.Error(t, err)
	})

	t.Run("returns an error given ipgetter fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ip := mockip.NewMockGetter(ctrl)
		dnsm := mockdns.NewMockManager(ctrl)
		someError := errors.New("some-error")
		someRecords := []string{"some-record"}

		ip.EXPECT().GetLocal().Return("", someError)

		s, err := odysseus.NewService(dnsm, ip)
		require.NoError(t, err)
		require.NotNil(t, s)

		err = s.UpdateDNSWithLocalIP(someRecords)

		assert.Error(t, err)
		assert.Equal(t, err, someError)
	})

}
