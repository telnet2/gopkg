package cert

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const Organization = "GatewayAgentCommon"

// generateKey generates a key and store it as a file if keyFile is non-empty.
func generateKey(keyFile string) (*ecdsa.PrivateKey, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	if keyFile != "" {
		err = keyToFile(keyFile, key)
		if err != nil {
			return nil, err
		}
	}
	return key, nil
}

func GenerateCertFile(host, validFrom string, validFor time.Duration,
	outRootKeyFile, outRootPEMFile, outLeafKeyFile, outLeafPEMFile, outClientKeyFile, outClientPEMFile string) error {
	if len(host) == 0 {
		return errors.New("empty host not allowed")
	}
	var err error
	var notBefore time.Time
	if len(validFrom) == 0 {
		notBefore = time.Now()
	} else {
		notBefore, err = time.Parse("Jan 2 15:04:05 2006", validFrom)
		if err != nil {
			return fmt.Errorf("failed to parse creation date: %w", err)
		}
	}
	notAfter := notBefore.Add(validFor)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return fmt.Errorf("failed to generate serial number: %w", err)
	}

	rootKey, err := generateKey(outRootKeyFile)
	if err != nil {
		return err
	}

	rootTemplate := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{Organization},
			CommonName:   "Root CA",
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &rootTemplate, &rootTemplate, &rootKey.PublicKey, rootKey)
	if err != nil {
		return err
	}
	if outRootPEMFile != "" {
		err = certToFile(outRootPEMFile, derBytes)
		if err != nil {
			return err
		}
	}

	leafKey, err := generateKey(outLeafKeyFile)
	if err != nil {
		return err
	}

	serialNumber, err = rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return fmt.Errorf("failed to generate serial number: %w", err)
	}
	leafTemplate := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{Organization},
			CommonName:   "TestCert1",
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  false,
	}
	hosts := strings.Split(host, ",")
	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			leafTemplate.IPAddresses = append(leafTemplate.IPAddresses, ip)
		} else {
			leafTemplate.DNSNames = append(leafTemplate.DNSNames, h)
		}
	}

	derBytes, err = x509.CreateCertificate(rand.Reader, &leafTemplate, &rootTemplate, &leafKey.PublicKey, rootKey)
	if err != nil {
		return err
	}
	if outLeafPEMFile != "" {
		err = certToFile(outLeafPEMFile, derBytes)
		if err != nil {
			return err
		}
	}

	clientKey, err := generateKey(outClientKeyFile)
	if err != nil {
		return err
	}

	clientTemplate := x509.Certificate{
		SerialNumber: new(big.Int).SetInt64(4),
		Subject: pkix.Name{
			Organization: []string{Organization},
			CommonName:   "client_auth_test_cert",
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
		IsCA:                  false,
	}

	derBytes, err = x509.CreateCertificate(rand.Reader, &clientTemplate, &rootTemplate, &clientKey.PublicKey, rootKey)
	if err != nil {
		return err
	}
	if outClientPEMFile != "" {
		err = certToFile(outClientPEMFile, derBytes)
		if err != nil {
			return err
		}
	}

	return nil
}

// keyToFile writes a PEM serialization of a key.
func keyToFile(filename string, key *ecdsa.PrivateKey) error {
	file, err := os.Create(filepath.Clean(filename))
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()
	b, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return err
	}
	if err := pem.Encode(file, &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}); err != nil {
		return err
	}
	return nil
}

// certToFile creates a certficate PEM file from bytes.
func certToFile(filename string, derBytes []byte) error {
	certOut, err := os.Create(filepath.Clean(filename))
	if err != nil {
		return err
	}
	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		return err
	}
	if err := certOut.Close(); err != nil {
		return err
	}
	return nil
}
