package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <url>\n", os.Args[0])
		return
	}

	url := os.Args[1]

	// Create a new TLS configuration that uses SSL 3.0 or TLS 1.0
	tlsConfig := &tls.Config{
		MinVersion:               tls.VersionSSL30,
		MaxVersion:               tls.VersionTLS10,
		PreferServerCipherSuites: true,
	}

	// Create an HTTP transport with the custom TLS configuration
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	// Create an HTTP client with the custom transport
	client := &http.Client{
		Transport: transport,
	}

	// Make a request using the custom client and the provided URL
	response, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer response.Body.Close()

	// Print the response status and body
	fmt.Printf("Status: %s\n", response.Status)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}
	fmt.Printf("Body: %s\n", body)
}
