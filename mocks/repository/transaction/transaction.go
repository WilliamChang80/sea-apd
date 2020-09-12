package transaction

import (
	"errors"
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/domain"
	"github.com/williamchang80/sea-apd/domain/transaction"
)

var (
	emptyTransaction = transaction.Transaction{
		Base:       domain.Base{},
		Status:     "",
		BankNumber: "",
		BankName:   "",
		Amount:     0,
		CustomerId: "",
		MerchantId: "",
	}
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

func (m MockRepository) CreateTransaction(transaction transaction.Transaction) error {
	if reflect.DeepEqual(transaction, emptyTransaction) {
		return errors.New("Transaction cannot be empty")
	}
	return nil
}

func (m MockRepository) GetTransactionById(id string) (*transaction.Transaction, error) {
	if len(id) == 0 {
		return nil, errors.New("Id cannot be empty")
	}
	return &emptyTransaction, nil
}

func (m MockRepository) UpdateTransactionStatus(status string, id string) (*transaction.Transaction, error) {
	if len(status) == 0 || len(id) == 0 {
		return nil, errors.New("Cannot Update with empty object")
	}
	return &emptyTransaction, nil
}

func (m MockRepository) GetTransactionByRequiredStatus(requiredStatus []string, userId string) ([]transaction.Transaction, error) {
	if len(userId) == 0 || len(requiredStatus) == 0 {
		return nil, errors.New("Cannot Get Required status with empty user id")
	}
	return []transaction.Transaction{}, nil
}

func (m MockRepository) GetMerchantRequestItem(merchantId string) ([]transaction.Transaction, error) {
	panic("implement me")
}

func (m MockRepository) UpdateTransaction(transaction transaction.Transaction) error {
	panic("implement me")
}

func (m MockRepository) AddCartItem(cart transaction.ProductTransaction) error {
	panic("implement me")
}

func (m MockRepository) RemoveCartItem(cart transaction.ProductTransaction) error {
	panic("implement me")
}

func (m MockRepository) UpdateCartItem(cart transaction.ProductTransaction) error {
	panic("implement me")
}

func (m MockRepository) GetCartItems(id string) ([]transaction.ProductTransaction, error) {
	panic("implement me")
}
