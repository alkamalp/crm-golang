package customers

import (
	"errors"
	"testing"
	"time"

	"github.com/alkamalp/crm-golang/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCustomerRepo struct {
	mock.Mock
}

func (m *MockCustomerRepo) CreateCustomer(customer *entity.Customer) (*entity.Customer, error) {
	args := m.Called(customer)
	result := args.Get(0)
	err := args.Error(1)
	if result == nil {
		return nil, err
	}
	return result.(*entity.Customer), err
}

func TestCreateCustomer(t *testing.T) {
	// Create a new instance of the mock repository
	mockRepo := new(MockCustomerRepo)

	// Create an instance of the use case with the mock repository
	useCase := useCaseCustomer{
		customerRepo: mockRepo,
	}

	// Create test input data
	customer := CustomerParam{
		First_name: "John",
		Last_name:  "Doe",
		Email:      "john.doe@example.com",
		Avatar:     "avatar.jpg",
	}

	// Set up expectations for the mock repository
	// hashedPassword := "hashed_password"
	mockRepo.On("CreateCustomer", mock.AnythingOfType("*entity.Customer")).Return(&entity.Customer{}, nil)
	// middleware.On("HashPassword", customerParam.Password).Return(hashedPassword, nil)

	// Call the function being tested
	createdCustomer, err := useCase.CreateCustomer(customer)

	// Assert that the expected repository method was called with the correct input
	mockRepo.AssertCalled(t, "CreateCustomer", mock.AnythingOfType("*entity.Customer"))

	// Assert the expected output
	assert.NotNil(t, createdCustomer)
	assert.NoError(t, err)
}

func TestCreateCustomer_Error(t *testing.T) {
	// Create a new instance of the mock repository
	mockRepo := new(MockCustomerRepo)

	// Create an instance of the use case with the mock repository
	useCase := useCaseCustomer{
		customerRepo: mockRepo,
	}

	// Create test input data
	customer := CustomerParam{
		First_name: "John",
		Last_name:  "Doe",
		Email:      "john.doe@example.com",
		Avatar:     "avatar.jpg",
	}

	// Create a new instance of the expected customer returned by the repository
	expectedCustomer := &entity.Customer{
		First_name: customer.First_name,
		Last_name:  customer.Last_name,
		Email:      customer.Email,
		Avatar:     customer.Avatar,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	expectedError := errors.New("create customer failed")

	// Set up expectations for the mock repository
	mockRepo.On("CreateCustomer", expectedCustomer).Return(nil, expectedError)

	// Call the function being tested
	result, err := useCase.CreateCustomer(customer)

	// Assert that the expected repository method was called with the correct input
	mockRepo.AssertCalled(t, "CreateCustomer", expectedCustomer)

	// Assert the expected error
	assert.Error(t, err)
	assert.EqualError(t, err, expectedError.Error())
	assert.Equal(t, entity.Customer{}, result)
}

func (m *MockCustomerRepo) GetCustomerById(id uint) (entity.Customer, error) {
	args := m.Called(id)
	result := args.Get(0)
	err := args.Error(1)
	if result == nil {
		return entity.Customer{}, err
	}
	return result.(entity.Customer), err
}

func TestGetCustomerById(t *testing.T) {
	// Create a new instance of the mock repository
	mockRepo := new(MockCustomerRepo)

	// Create an instance of the use case with the mock repository
	useCase := useCaseCustomer{
		customerRepo: mockRepo,
	}

	// Create test input data
	customerID := uint(1)
	customer := entity.Customer{
		First_name: "John",
		Last_name:  "Doe",
		Email:      "johndoe@gmail.com",
		Avatar:     "avatar.jpg",
	}

	// Set up expectations for the mock repository
	mockRepo.On("GetCustomerById", customerID).Return(customer, nil)

	// Call the function being tested
	result, err := useCase.GetCustomerById(customerID)

	// Assert that the expected repository method was called with the correct input
	mockRepo.AssertCalled(t, "GetCustomerById", customerID)

	// Assert the expected output
	assert.Equal(t, customer, result)
	assert.NoError(t, err)
}

func TestGetCustomerById_Error(t *testing.T) {
	// Create a new instance of the mock repository
	mockRepo := new(MockCustomerRepo)

	// Create an instance of the use case with the mock repository
	useCase := useCaseCustomer{
		customerRepo: mockRepo,
	}

	// Create test input data
	customerID := uint(1)
	expectedError := errors.New("failed to get customer")

	// Set up expectations for the mock repository
	mockRepo.On("GetCustomerById", customerID).Return(entity.Customer{}, expectedError)

	// Call the function being tested
	result, err := useCase.GetCustomerById(customerID)

	// Assert that the expected repository method was called with the correct input
	mockRepo.AssertCalled(t, "GetCustomerById", customerID)

	// Assert the expected error
	assert.Error(t, err)
	assert.EqualError(t, err, expectedError.Error())
	assert.Equal(t, entity.Customer{}, result)
}
