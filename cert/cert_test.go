package cert

import (
	"crypto/x509"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/telnet2/gopkg/ioutil"
)

func TestGenerateCert(t *testing.T) {
	rootPEM := filepath.Join(os.TempDir(), "root.pem")
	err := GenerateCertFile("test-ag.tiktok-us.org", "", time.Hour*24*365, "", rootPEM, "", "", "", "")
	require.NoError(t, err)
	rootPEMCert := ioutil.MustReadFile(rootPEM)

	pool := x509.NewCertPool()
	ok := pool.AppendCertsFromPEM([]byte(rootPEMCert))
	require.True(t, ok)
}
