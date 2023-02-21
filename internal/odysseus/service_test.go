package odysseus_test

import (
	"errors"
	"testing"

	"github.com/darkraiden/odysseus/internal/mock/mockdns"
	mockip "github.com/darkraiden/odysseus/internal/mock/mockip"
	"github.com/darkraiden/odysseus/internal/odysseus"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("returns an error when DNS.Manager is nil", func(t *testing.T) {
		mockIP := mockip.NewMockGetter(ctrl)
		s, err := odysseus.NewService(nil, mockIP)
		require.Nil(t, s)
		require.Error(t, err)
	})

	t.Run("returns an error when ipaddress.Getter is nil", func(t *testing.T) {
		mockDNS := mockdns.NewMockManager(ctrl)
		s, err := odysseus.NewService(mockDNS, nil)
		require.Nil(t, s)
		require.Error(t, err)
	})

	t.Run("returns an error when both ipaddress.Getter and DNS.Manager are nil", func(t *testing.T) {
		s, err := odysseus.NewService(nil, nil)
		require.Nil(t, s)
		require.Error(t, err)
	})

	t.Run("test content of *Service", func(t *testing.T) {
		mockDNS := mockdns.NewMockManager(ctrl)
		mockIP := mockip.NewMockGetter(ctrl)

		s, err := odysseus.NewService(mockDNS, mockIP)
		require.Nil(t, err)
		require.NotNil(t, s)
	})
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
	t.Run("returns an error given a failure to get DNS Records", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ip := mockip.NewMockGetter(ctrl)
		dnsm := mockdns.NewMockManager(ctrl)
		someErr := errors.New("some-error")
		someRecords := []string{"some-record"}

		s, err := odysseus.NewService(dnsm, ip)
		require.NoError(t, err)

		gomock.InOrder(
			ip.EXPECT().GetLocal().Return("192.168.0.1", nil),
			dnsm.EXPECT().GetDNSRecords(someRecords).Return(nil, someErr),
		)

		err = s.UpdateDNSWithLocalIP(someRecords)

		assert.Error(t, err)

	})

}
