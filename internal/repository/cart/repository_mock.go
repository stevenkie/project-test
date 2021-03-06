// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package cart

import (
	cartModel "github.com/stevenkie/project-test/internal/model/cart"
	"sync"
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of Repository.
//
// 	func TestSomethingThatUsesRepository(t *testing.T) {
//
// 		// make and configure a mocked Repository
// 		mockedRepository := &RepositoryMock{
// 			EmptyCartFunc: func(userID string) error {
// 				panic("mock out the EmptyCart method")
// 			},
// 			GetCartFunc: func(userID string) (cartModel.Cart, error) {
// 				panic("mock out the GetCart method")
// 			},
// 			SaveCartFunc: func(userID string, cart cartModel.Cart) error {
// 				panic("mock out the SaveCart method")
// 			},
// 		}
//
// 		// use mockedRepository in code that requires Repository
// 		// and then make assertions.
//
// 	}
type RepositoryMock struct {
	// EmptyCartFunc mocks the EmptyCart method.
	EmptyCartFunc func(userID string) error

	// GetCartFunc mocks the GetCart method.
	GetCartFunc func(userID string) (cartModel.Cart, error)

	// SaveCartFunc mocks the SaveCart method.
	SaveCartFunc func(userID string, cart cartModel.Cart) error

	// calls tracks calls to the methods.
	calls struct {
		// EmptyCart holds details about calls to the EmptyCart method.
		EmptyCart []struct {
			// UserID is the userID argument value.
			UserID string
		}
		// GetCart holds details about calls to the GetCart method.
		GetCart []struct {
			// UserID is the userID argument value.
			UserID string
		}
		// SaveCart holds details about calls to the SaveCart method.
		SaveCart []struct {
			// UserID is the userID argument value.
			UserID string
			// Cart is the cart argument value.
			Cart cartModel.Cart
		}
	}
	lockEmptyCart sync.RWMutex
	lockGetCart   sync.RWMutex
	lockSaveCart  sync.RWMutex
}

// EmptyCart calls EmptyCartFunc.
func (mock *RepositoryMock) EmptyCart(userID string) error {
	if mock.EmptyCartFunc == nil {
		panic("RepositoryMock.EmptyCartFunc: method is nil but Repository.EmptyCart was just called")
	}
	callInfo := struct {
		UserID string
	}{
		UserID: userID,
	}
	mock.lockEmptyCart.Lock()
	mock.calls.EmptyCart = append(mock.calls.EmptyCart, callInfo)
	mock.lockEmptyCart.Unlock()
	return mock.EmptyCartFunc(userID)
}

// EmptyCartCalls gets all the calls that were made to EmptyCart.
// Check the length with:
//     len(mockedRepository.EmptyCartCalls())
func (mock *RepositoryMock) EmptyCartCalls() []struct {
	UserID string
} {
	var calls []struct {
		UserID string
	}
	mock.lockEmptyCart.RLock()
	calls = mock.calls.EmptyCart
	mock.lockEmptyCart.RUnlock()
	return calls
}

// GetCart calls GetCartFunc.
func (mock *RepositoryMock) GetCart(userID string) (cartModel.Cart, error) {
	if mock.GetCartFunc == nil {
		panic("RepositoryMock.GetCartFunc: method is nil but Repository.GetCart was just called")
	}
	callInfo := struct {
		UserID string
	}{
		UserID: userID,
	}
	mock.lockGetCart.Lock()
	mock.calls.GetCart = append(mock.calls.GetCart, callInfo)
	mock.lockGetCart.Unlock()
	return mock.GetCartFunc(userID)
}

// GetCartCalls gets all the calls that were made to GetCart.
// Check the length with:
//     len(mockedRepository.GetCartCalls())
func (mock *RepositoryMock) GetCartCalls() []struct {
	UserID string
} {
	var calls []struct {
		UserID string
	}
	mock.lockGetCart.RLock()
	calls = mock.calls.GetCart
	mock.lockGetCart.RUnlock()
	return calls
}

// SaveCart calls SaveCartFunc.
func (mock *RepositoryMock) SaveCart(userID string, cart cartModel.Cart) error {
	if mock.SaveCartFunc == nil {
		panic("RepositoryMock.SaveCartFunc: method is nil but Repository.SaveCart was just called")
	}
	callInfo := struct {
		UserID string
		Cart   cartModel.Cart
	}{
		UserID: userID,
		Cart:   cart,
	}
	mock.lockSaveCart.Lock()
	mock.calls.SaveCart = append(mock.calls.SaveCart, callInfo)
	mock.lockSaveCart.Unlock()
	return mock.SaveCartFunc(userID, cart)
}

// SaveCartCalls gets all the calls that were made to SaveCart.
// Check the length with:
//     len(mockedRepository.SaveCartCalls())
func (mock *RepositoryMock) SaveCartCalls() []struct {
	UserID string
	Cart   cartModel.Cart
} {
	var calls []struct {
		UserID string
		Cart   cartModel.Cart
	}
	mock.lockSaveCart.RLock()
	calls = mock.calls.SaveCart
	mock.lockSaveCart.RUnlock()
	return calls
}
