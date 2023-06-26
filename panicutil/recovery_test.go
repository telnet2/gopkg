package panicutil

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGlobalPanicHook(t *testing.T) {
	var (
		gst  string
		gerr error
	)

	wg := sync.WaitGroup{}
	wg.Add(2)
	InstallGlobalPanicHook(func(stackTrace string, err error) {
		gst = stackTrace
		gerr = err
		wg.Done()
	})
	go func() {
		// It MUST defer RecoverFromPanic directly. RecoverFromPanic must not be called from another defer function.
		defer RecoverFromPanic(nil)
		panic("hello")
	}()

	go func() {
		defer func() {
			_, err := RecoverFromPanic(nil) // this won't work
			require.Nil(t, err)
			_ = recover()
			// the global panic hook was not called because `recover()` within the RecoverFromPanic returned nil
			wg.Done()
		}()
		panic("incorrect usage")
	}()

	wg.Wait()
	require.NotEmpty(t, gst)
	require.Error(t, gerr)
	require.Contains(t, gerr.Error(), "hello")
}
