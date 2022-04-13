package logger

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/stank0s/zap-logger/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Test_NewLogger(t *testing.T) {
	// when: creating new logger object
	l := NewLogger()

	// then: object has corect type
	assert.IsType(t, &Logger{}, l)
	assert.NotEmpty(t, l.cfg)
}

func Test_Logger_Happy(t *testing.T) {
	// given: mock controler
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// and: necessary mocks
	mockConfig := mocks.NewMockConfig(ctrl)
	mockConfig.EXPECT().CreateLoggerConfig("info", true).Return(zap.Config{
		Encoding: "json",
		Level:    zap.NewAtomicLevelAt(zapcore.InfoLevel),
	}, nil)

	// and: test subject
	l := Logger{
		cfg: mockConfig,
	}

	// when: calling function
	log, err := l.Logger("info", true)

	// then: no error returned
	assert.NoError(t, err)
	assert.IsType(t, &zap.Logger{}, log)
}

func Test_Logger_Config_Error(t *testing.T) {
	// given: mock controler
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// and: necessary mocks
	mockConfig := mocks.NewMockConfig(ctrl)
	mockConfig.EXPECT().CreateLoggerConfig("info", false).Return(zap.Config{}, errors.New("test error"))

	// and: test subject
	l := Logger{
		cfg: mockConfig,
	}

	// when: calling function
	_, err := l.Logger("info", false)

	// then: error returned
	assert.Error(t, err)
}

func Test_Logger_Builder_Error(t *testing.T) {
	// given: mock controler
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// and: necessary mocks
	mockConfig := mocks.NewMockConfig(ctrl)
	mockConfig.EXPECT().CreateLoggerConfig("info", true).Return(zap.Config{
		Encoding: "json",
		Level:    zap.NewAtomicLevelAt(zapcore.InfoLevel),
	}, nil)

	// and: mock build function
	tmp := build
	build = func(c zap.Config) (*zap.Logger, error) {
		return nil, errors.New("test error")
	}

	// and: test subject
	l := Logger{
		cfg: mockConfig,
	}

	// when: calling function
	_, err := l.Logger("info", true)

	// then: error returned
	assert.Error(t, err)

	build = tmp
}

func Test_SugaredLogger_Happy(t *testing.T) {
	// given: mock controler
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// and: necessary mocks
	mockConfig := mocks.NewMockConfig(ctrl)
	mockConfig.EXPECT().CreateLoggerConfig("info", false).Return(zap.Config{
		Encoding: "json",
		Level:    zap.NewAtomicLevelAt(zapcore.InfoLevel),
	}, nil)

	// and: test subject
	l := Logger{
		cfg: mockConfig,
	}

	// when: calling function
	log, err := l.SugaredLogger("info", false)

	// then: no error returned
	assert.NoError(t, err)
	assert.IsType(t, &zap.SugaredLogger{}, log)
}

func Test_SugaredLogger_Unhappy(t *testing.T) {
	// given: mock controler
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// and: necessary mocks
	mockConfig := mocks.NewMockConfig(ctrl)
	mockConfig.EXPECT().CreateLoggerConfig("info", true).Return(zap.Config{}, errors.New("test error"))

	// and: test subject
	l := Logger{
		cfg: mockConfig,
	}

	// when: calling function
	_, err := l.SugaredLogger("info", true)

	// then: error returned
	assert.Error(t, err)
}
