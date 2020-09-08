package merchant

import (
	"errors"

	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/domain/merchant"
	merch "github.com/williamchang80/sea-apd/domain/merchant"
)

type MockRepository struct {
	ctrl *gomock.Controller
}

func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{
		ctrl: ctrl,
	}
	return mock
}

func (m MockRepository) UpdateMerchantBalance(amount int, merchantId string) error {
	if len(merchantId) == 0 || amount == 0 {
		return errors.New("Id and amount cannot be empty")
	}
	return nil
}

func (m MockRepository) GetMerchantBalance(merchantId string) (int, error) {
	if len(merchantId) == 0 {
		return 0, errors.New("Id cannot be empty")
	}
	return 100, nil
}

func (m MockRepository) RegisterMerchant(merchant merchant.Merchant) error {
	var mh = merch.Merchant{}
	if merchant == mh {
		return errors.New("Cannot Register Merchant")
	}
	return nil
}

func (m MockRepository) GetMerchants() ([]merchant.Merchant, error) {
	m.ctrl.T.Helper()
	return []merchant.Merchant{}, nil
}

func (m MockRepository) GetMerchantById(merchantId string) (*merchant.Merchant, error) {
	if merchantId != "" {
		return &merchant.Merchant{}, nil
	}
	return nil, errors.New("Cannot Get Merchant By Id")
}

func (m MockRepository) GetMerchantsByUser(userId string) ([]merchant.Merchant, error) {
	if userId != "" {
		return []merchant.Merchant{}, nil
	}
	return nil, errors.New("Cannot Get Merchants By User")
}
