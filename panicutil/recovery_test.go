package panicutil

import (
	"fmt"
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
	wg.Add(1)
	InstallGlobalPanicHook(func(stackTrace string, err error) {
		gst = stackTrace
		gerr = err
		wg.Done()
	})
	go func() {
		defer RecoverFromPanic()
		panic("hello")
	}()

	wg.Wait()
	require.NotEmpty(t, gst)
	require.Error(t, gerr)
	require.Contains(t, gerr.Error(), "hello")
	fmt.Println(gst)
}
