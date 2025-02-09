// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	apimodels "github.com/keptn/go-utils/pkg/api/models"
	"sync"
)

// ISequenceTaskStartedHookMock is a mock implementation of controller.ISequenceTaskStartedHook.
//
// 	func TestSomethingThatUsesISequenceTaskStartedHook(t *testing.T) {
//
// 		// make and configure a mocked controller.ISequenceTaskStartedHook
// 		mockedISequenceTaskStartedHook := &ISequenceTaskStartedHookMock{
// 			OnSequenceTaskStartedFunc: func(keptnContextExtendedCE apimodels.KeptnContextExtendedCE)  {
// 				panic("mock out the OnSequenceTaskStarted method")
// 			},
// 		}
//
// 		// use mockedISequenceTaskStartedHook in code that requires controller.ISequenceTaskStartedHook
// 		// and then make assertions.
//
// 	}
type ISequenceTaskStartedHookMock struct {
	// OnSequenceTaskStartedFunc mocks the OnSequenceTaskStarted method.
	OnSequenceTaskStartedFunc func(keptnContextExtendedCE apimodels.KeptnContextExtendedCE)

	// calls tracks calls to the methods.
	calls struct {
		// OnSequenceTaskStarted holds details about calls to the OnSequenceTaskStarted method.
		OnSequenceTaskStarted []struct {
			// KeptnContextExtendedCE is the keptnContextExtendedCE argument value.
			KeptnContextExtendedCE apimodels.KeptnContextExtendedCE
		}
	}
	lockOnSequenceTaskStarted sync.RWMutex
}

// OnSequenceTaskStarted calls OnSequenceTaskStartedFunc.
func (mock *ISequenceTaskStartedHookMock) OnSequenceTaskStarted(keptnContextExtendedCE apimodels.KeptnContextExtendedCE) {
	if mock.OnSequenceTaskStartedFunc == nil {
		panic("ISequenceTaskStartedHookMock.OnSequenceTaskStartedFunc: method is nil but ISequenceTaskStartedHook.OnSequenceTaskStarted was just called")
	}
	callInfo := struct {
		KeptnContextExtendedCE apimodels.KeptnContextExtendedCE
	}{
		KeptnContextExtendedCE: keptnContextExtendedCE,
	}
	mock.lockOnSequenceTaskStarted.Lock()
	mock.calls.OnSequenceTaskStarted = append(mock.calls.OnSequenceTaskStarted, callInfo)
	mock.lockOnSequenceTaskStarted.Unlock()
	mock.OnSequenceTaskStartedFunc(keptnContextExtendedCE)
}

// OnSequenceTaskStartedCalls gets all the calls that were made to OnSequenceTaskStarted.
// Check the length with:
//     len(mockedISequenceTaskStartedHook.OnSequenceTaskStartedCalls())
func (mock *ISequenceTaskStartedHookMock) OnSequenceTaskStartedCalls() []struct {
	KeptnContextExtendedCE apimodels.KeptnContextExtendedCE
} {
	var calls []struct {
		KeptnContextExtendedCE apimodels.KeptnContextExtendedCE
	}
	mock.lockOnSequenceTaskStarted.RLock()
	calls = mock.calls.OnSequenceTaskStarted
	mock.lockOnSequenceTaskStarted.RUnlock()
	return calls
}
