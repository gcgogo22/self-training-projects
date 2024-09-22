package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
)

/*
SNI server name indication example,

// Set up the TLS configuration with SNI support

	config := &tls.Config{
		// GetCertificate dynamically selects the certificate based on the SNI
		// With callback function
		GetCertificate: func(hello *tls.ClientHelloInfo) (*tls.Cerficate, error) {
			// Look up the certificate based on the requested hostname
			// Default to the first certificate if no match
			if cert, ok := certMap[hello.ServerName]; ok {
				return &cert, nil
			}
			// Return the first certificate as fallback
			for _, cert := range certMap {
				return &cert, nil
			}
			return nil, fmt.Errorf("no certificates found")
		}

}
*/
type Certificates struct {
	CertFile string
	KeyFile  string
}

func main() {
	httpsServer := &http.Server{
		Addr: ":8080",
	}
	var certs []Certificates
	certs = append(certs, Certificates{
		CertFile: "../etc/yourSite.pem", // Your site certificate key
		KeyFile:  "../ect/yourSite.key", // Your site private key
	})

	config := &tls.Config{}

	config.Certificates = make([]tls.Certificate, len(certs))
	for i, v := range certs {
		config.Certificates[i], _ = tls.LoadX509KeyPair(v.CertFile, v.KeyFile)
	}

	conn, _ := net.Listen("tcp", ":8080")
	// During the handshake, the client sends the requested hostname, the server selectes the appropriate certificates from the list based on the requested hostname.
	tlsListener := tls.NewListener(conn, config)
	httpsServer.Serve(tlsListener)
	fmt.Println("Listening on port 8080...")
}
