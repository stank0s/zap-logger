package builder

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/stank0s/zap-logger/builder/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func Test_NewBuilder(t *testing.T) {
	// given: mock controler
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// and: config mock
	configMock := mocks.NewMockConfig(ctrl)

	// when: creating new builder object
	b := NewBuilder(configMock)

	// then: object has correct type
	assert.IsType(t, &Builder{}, b)
}

func Test_Build_Happy(t *testing.T) {
	// given: mock controler
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// and: config mock
	configMock := mocks.NewMockConfig(ctrl)
	configMock.EXPECT().Build(gomock.Any()).Return(&zap.Logger{}, nil)

	// and: test subject
	b := NewBuilder(configMock)

	// when: calling function
	log, err := b.Build()

	//then: no error returned
	assert.NoError(t, err)
	assert.IsType(t, &zap.Logger{}, log)
}

func Test_Build_Unhappy(t *testing.T) {
	// given: mock controler
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// and: config mock
	configMock := mocks.NewMockConfig(ctrl)
	configMock.EXPECT().Build(gomock.Any()).Return(&zap.Logger{}, errors.New("test error"))

	// and: test subject
	b := NewBuilder(configMock)

	// when: calling function
	_, err := b.Build()

	//then: error returned
	assert.Error(t, err)
}
