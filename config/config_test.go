package config

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/stank0s/zap-logger/config/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func Test_NewConfig(t *testing.T) {
	// when: creating new config object
	c := NewConfig(zap.AtomicLevel{})

	// then: object has correct type
	assert.IsType(t, &Config{}, c)
}

func Test_CreateLoggerConfig_Happy(t *testing.T) {
	// given: mock controler
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// and: necessary mocks
	mockLevel := mocks.NewMockLevel(ctrl)
	mockLevel.EXPECT().UnmarshalText([]byte("info")).Return(nil)

	// and: test subject
	c := Config{
		lvl: mockLevel,
	}

	// when: calling function
	cfg, err := c.CreateLoggerConfig("info")

	// then: no error returned
	assert.NoError(t, err)
	assert.IsType(t, zap.Config{}, cfg)
}

func Test_CreateLoggerConfig_Unhappy(t *testing.T) {
	// given: mock controler
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// and: necessary mocks
	mockLevel := mocks.NewMockLevel(ctrl)
	mockLevel.EXPECT().UnmarshalText([]byte("nonsense")).Return(errors.New("test error"))

	// and: test subject
	c := Config{
		lvl: mockLevel,
	}

	// when: calling function
	_, err := c.CreateLoggerConfig("nonsense")

	// then: error returned
	assert.Error(t, err)
}
