package main_test

import (
	mock_main "learnbytests/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestCharge(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockPaymentProcessor := mock_main.NewMockPaymentProcessor(mockCtrl)
	testPaymentProcessorClient := &main.PaymentProcessorClient{PaymentProcessor: mockPaymentProcessor}
	defer mockCtrl.Finish()
	mockPaymentProcessor.EXPECT().Charge(100.0, "test_token").Return(nil).Times(1)
	err := testPaymentProcessorClient.Charge(100.0, "test_token")
	if err != nil {
		t.Fail()
	}

	err = testPaymentProcessorClient.Charge(10.0, "test_token")
	if err.Error() != "Charge too low" {
		t.Errorf("Error returned was: %s", err.Error())
	}
}
