package recover

import (
	"fmt"
	"runtime/debug"
	"sync"
)

var (
	l                  sync.Mutex
	globalRecoveryHook func(stackTrace string, err error) = func(stackTrace string, err error) {}
)

func InstallGlobalRecoveryHook(hook func(stackTrace string, err error)) {
	l.Lock()
	defer l.Unlock()
	globalRecoveryHook = hook
}

// Recovery is a helper function to recover from panics.
// It helps handling the panic consistently with the global recovery hook.
func Recovery() (stackTrace string, err error) {
	if panicMsg := recover(); panicMsg != nil {
		stackTrace = string(debug.Stack())
		err = fmt.Errorf("panic occurred: %v", panicMsg)
		globalRecoveryHook(stackTrace, err)
		return stackTrace, err
	}
	return "", nil
}
