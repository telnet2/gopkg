package panicutil

import (
	"fmt"
	"runtime/debug"
	"sync"
)

var (
	l            sync.Mutex
	nopPanicHook = func(stackTrace string, err error) {}

	// globalPanicHook is the function to be called when a panic occurs.
	// It is set to nopPanicHook if no hook is set or nil hook is set via InstallGlobalPanicHook.
	globalPanicHook func(stackTrace string, err error) = nopPanicHook
)

func InstallGlobalPanicHook(hook func(stackTrace string, err error)) {
	l.Lock()
	defer l.Unlock()
	if hook == nil {
		globalPanicHook = nopPanicHook
	} else {
		globalPanicHook = hook
	}
}

// RecoverFromPanic is a helper function to recover from panics.
// It helps handling the panic consistently with the global recovery hook.
func RecoverFromPanic() (stackTrace string, err error) {
	if panicMsg := recover(); panicMsg != nil {
		stackTrace = string(debug.Stack())
		err = fmt.Errorf("panic occurred: %v", panicMsg)
		globalPanicHook(stackTrace, err)
		return stackTrace, err
	}
	return "", nil
}
