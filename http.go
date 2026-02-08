package gosugar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GET requests

// GetBody sends an HTTP GET request to the given URL
// and returns the response body as a string.
//
// It returns an error if the request fails,
// the response status code is not 200 OK,
// or the response body cannot be read.
func GetBody(url string) (string, error) {
	// Send an HTTP GET request to the given URL
	resp, err := http.Get(url)

	// Return an error if the request fails
	if err != nil {
		return "", err
	}

	// Ensure the response body is closed
	// when the function exits
	defer resp.Body.Close()

	// Treat non-200 OK status codes as errors
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(
			"status code: %d",
			resp.StatusCode,
		)
	}

	// Read the entire response body
	body, err := io.ReadAll(resp.Body)

	// Return an error if reading the body fails
	if err != nil {
		return "", err
	}

	// Convert body from bytes to string and return it
	return string(body), nil
}

// MustGetBody sends an HTTP GET request to the given URL
// and returns the response body as a string.
//
// It panics if the request fails, the status code is not 200 OK,
// or the response body cannot be read.
// Use this function when a failure should stop the program.
func MustGetBody(url string) string {
	body, err := GetBody(url)

	if err != nil {
		panic(err)
	}

	// Return body
	return body
}

// GetJSON sends an HTTP GET request to the given URL
// and decodes the JSON response into the specified type T.
//
// The type T is a generic type parameter, allowing the caller
// to define the expected response structure at compile time.
//
// This function returns an error if:
// - the HTTP request fails
// - the response status code is not successful (handled in GetBody)
// - the response body cannot be read
// - the JSON cannot be unmarshaled into type T
func GetJSON[T any](url string) (T, error) {
	// Declare a zero-value variable of type T.
	// This will be returned in case of an error.
	var result T

	// Perform an HTTP GET request and retrieve the response body as a string
	body, err := GetBody(url)
	if err != nil {
		// Return the zero-value result along with the error
		return result, err
	}

	// Convert the response body (string) into a byte slice
	// and unmarshal the JSON data into the result variable
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		// Return the zero-value result if JSON decoding fails
		return result, err
	}

	// Return the successfully decoded result
	// and a nil error to indicate success
	return result, nil
}

// GetHeader sends an HTTP GET request to the given URL
// and returns the response headers.
//
// It returns an error if the request fails
// or the response status code is not 200 OK.
func GetHeader(url string) (http.Header, error) {
	// Send an HTTP GET request to the given URL
	resp, err := http.Get(url)

	// Return an error if the request fails
	if err != nil {
		return nil, err
	}

	// Ensure the response body is closed
	// even though the body is not read
	defer resp.Body.Close()

	// Treat non-200 OK status codes as errors
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"status code: %d",
			resp.StatusCode,
		)
	}

	// Return the response headers
	// http.Header is a map[string][]string
	// and is case-insensitive by design
	return resp.Header, nil
}

// MustGetHeader sends an HTTP GET request to the given URL
// and returns the response headers.
//
// It panics if the request fails or the response
// status code is not 200 OK.
// Use this function when a failure should stop the program.
func MustGetHeader(url string) http.Header {
	headers, err := GetHeader(url)

	// Panic if an error occurs
	if err != nil {
		panic(err)
	}

	return headers
}

// POST requests

// PostBody sends an HTTP POST request to the given URL
// with the provided body and content type,
// and returns the response body as a string.
func PostBody(url string, body io.Reader, contentType string) (string, error) {
	resp, err := http.Post(url, contentType, body)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(
			"status code: %d",
			resp.StatusCode,
		)
	}

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(respBody), nil

}

// MustPostBody sends an HTTP POST request and panics on error.
func MustPostBody(url string, body io.Reader, contentType string) string {
	result, err := PostBody(url, body, contentType)

	if err != nil {
		panic(err)
	}

	return result
}

// PostJSON sends an HTTP POST request with a JSON payload
// and decodes the JSON response into type T.
// and decodes the JSON response into type T.
func PostJSON[T any](url string, payload any) (T, error) {
	var result T

	// Encode payload to JSON
	data, err := json.Marshal(payload)
	if err != nil {
		return result, err
	}

	respBody, err := PostBody(
		url,
		bytes.NewReader(data),
		"application/json",
	)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal([]byte(respBody), &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// PostHeader sends an HTTP POST request
// and returns the response headers.
func PostHeader(url string, body io.Reader, contentType string) (http.Header, error) {
	resp, err := http.Post(url, contentType, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"status code: %d",
			resp.StatusCode,
		)
	}

	return resp.Header, nil
}

// MustPostHeader sends an HTTP POST request
// and panics on error.
func MustPostHeader(url string, body io.Reader, contentType string) http.Header {
	headers, err := PostHeader(url, body, contentType)
	if err != nil {
		panic(err)
	}
	return headers
}

// PUT requests

// PutBody sends an HTTP PUT request to the given URL
// with the provided body and content type,
// and returns the response body as a string.
func PutBody(url string, body io.Reader, contentType string) (string, error) {
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(
			"status code: %d",
			resp.StatusCode,
		)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}

// MustPutBody sends an HTTP PUT request and panics on error.
func MustPutBody(url string, body io.Reader, contentType string) string {
	result, err := PutBody(url, body, contentType)

	if err != nil {
		panic(err)
	}

	return result
}

// PutJSON sends an HTTP PUT request with a JSON payload
// and decodes the JSON response into type T.
func PutJSON[T any](url string, payload any) (T, error) {
	var result T

	// Encode payload to JSON
	data, err := json.Marshal(payload)
	if err != nil {
		return result, err
	}

	respBody, err := PutBody(
		url,
		bytes.NewReader(data),
		"application/json",
	)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal([]byte(respBody), &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// PutHeader sends an HTTP PUT request
// and returns the response headers.
func PutHeader(url string, body io.Reader, contentType string) (http.Header, error) {
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"status code: %d",
			resp.StatusCode,
		)
	}

	return resp.Header, nil
}

// MustPutHeader sends an HTTP PUT request
// and panics on error.
func MustPutHeader(url string, body io.Reader, contentType string) http.Header {
	headers, err := PutHeader(url, body, contentType)
	if err != nil {
		panic(err)
	}
	return headers
}

// DELETE requests

// DeleteBody sends an HTTP DELETE request to the given URL
// with the provided body and content type,
// and returns the response body as a string.
func DeleteBody(url string, body io.Reader, contentType string) (string, error) {
	req, err := http.NewRequest("DELETE", url, body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(
			"status code: %d",
			resp.StatusCode,
		)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}

// MustDeleteBody sends an HTTP DELETE request and panics on error.
func MustDeleteBody(url string, body io.Reader, contentType string) string {
	result, err := DeleteBody(url, body, contentType)

	if err != nil {
		panic(err)
	}

	return result
}

// DeleteJSON sends an HTTP DELETE request with a JSON payload
// and decodes the JSON response into type T.
func DeleteJSON[T any](url string, payload any) (T, error) {
	var result T

	// Encode payload to JSON
	data, err := json.Marshal(payload)
	if err != nil {
		return result, err
	}

	respBody, err := DeleteBody(
		url,
		bytes.NewReader(data),
		"application/json",
	)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal([]byte(respBody), &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// DeleteHeader sends an HTTP DELETE request
// and returns the response headers.
func DeleteHeader(url string, body io.Reader, contentType string) (http.Header, error) {
	req, err := http.NewRequest("DELETE", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"status code: %d",
			resp.StatusCode,
		)
	}

	return resp.Header, nil
}

// MustDeleteHeader sends an HTTP DELETE request
// and panics on error.
func MustDeleteHeader(url string, body io.Reader, contentType string) http.Header {
	headers, err := DeleteHeader(url, body, contentType)
	if err != nil {
		panic(err)
	}
	return headers
}
