package service_test

import (
	service "Spam-Masker/basic/services"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/stretchr/testify/mock"
)

type MockProducer struct {
	mock.Mock
}

func (m *MockProducer) Produce() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

type MockPresenter struct {
	mock.Mock
}

func (m *MockPresenter) Present(byteText []string) error {
	args := m.Called(byteText)
	return args.Error(0)
}

func Test_Run_Service(t *testing.T) {
	mockProducer := new(MockProducer)
	mockPresenter := new(MockPresenter)

	testData := []string{
		"test http://example.com",
		"secure https://test.com",
		"no url here",
	}

	expectedResults := []string{
		"test http://***********",
		"secure https://test.com",
		"no url here",
	}
	mockProducer.On("Produce").Return(testData, nil)
	mockPresenter.On("Present", expectedResults).Return(nil)

	serv := service.NewService(mockProducer, mockPresenter)
	err := serv.Run()

	assert.NoError(t, err)
	mockProducer.AssertExpectations(t)
	mockPresenter.AssertExpectations(t)
}

func TestMask(t *testing.T) {
	tests := []struct {
		name string
		input string
		output string
	}{
		{
			name: "Текст без URL",
			input: "The usual test",
			output: "The usual test",
		},
		{
			name: "Текст с URL",
			input: "Check out my wedsite http://vk.com friend",
			output: "Check out my wedsite http://****** friend",
		},
		{
			name: "Текст с 2 URL",
			input: "Check out my wedsites http://vk.com http://instagram.com Do you say?",
			output: "Check out my wedsites http://****** http://************* Do you say?",
		},
	}
	for _, tt := range tests{
		serv := service.NewService(nil, nil)
		assert.Equal(t, tt.output, serv.Mask(tt.input), tt.name)
	}
}