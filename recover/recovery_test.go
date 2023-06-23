package recover

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGlobalRecoveryHook(t *testing.T) {
	var (
		gst  string
		gerr error
	)

	wg := sync.WaitGroup{}
	wg.Add(1)
	InstallGlobalRecoveryHook(func(stackTrace string, err error) {
		gst = stackTrace
		gerr = err
		wg.Done()
	})
	go func() {
		defer Recovery()
		panic("hello")
	}()

	wg.Wait()
	require.NotEmpty(t, gst)
	require.Error(t, gerr)
	require.Contains(t, gerr.Error(), "hello")
	fmt.Println(gst)
}
