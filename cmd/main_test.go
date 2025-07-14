package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mocks for entity package
var (
	mockIsValidCEP func(string) bool
	mockGetCep     func(string) (string, bool, error)
	mockGetWeather func(string) (interface{}, error)
)

func newRequestWithPathValue(method, target, cep string) *http.Request {
	req := httptest.NewRequest(method, target, nil)
	// Simulate r.PathValue("cep")
	req.SetPathValue("cep", cep)
	return req
}

func TestHandle_InvalidCEP(t *testing.T) {
	mockIsValidCEP = func(cep string) bool { return false }

	rr := httptest.NewRecorder()
	req := newRequestWithPathValue("GET", "/12345678", "1234567")

	Handle(rr, req)

	if rr.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected status %d, got %d", http.StatusUnprocessableEntity, rr.Code)
	}
	if body := rr.Body.String(); body != "invalid zipcode\n" {
		t.Errorf("expected body %q, got %q", "invalid zipcode\n", body)
	}
}

func TestHandle_GetCepError(t *testing.T) {
	mockIsValidCEP = func(cep string) bool { return true }
	mockGetCep = func(cep string) (string, bool, error) {
		return "", false, errors.New("not found")
	}

	rr := httptest.NewRecorder()
	req := newRequestWithPathValue("GET", "/12345678", "12345678")

	Handle(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, rr.Code)
	}
	if body := rr.Body.String(); body != "can not find zipcode\n" {
		t.Errorf("expected body %q, got %q", "can not find zipcode\n", body)
	}
}

func TestHandle_GetCepErroTrue(t *testing.T) {
	mockIsValidCEP = func(cep string) bool { return true }
	mockGetCep = func(cep string) (string, bool, error) {
		return "", true, nil
	}

	rr := httptest.NewRecorder()
	req := newRequestWithPathValue("GET", "/12345678", "12345678")

	Handle(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, rr.Code)
	}
	if body := rr.Body.String(); body != "can not find zipcode\n" {
		t.Errorf("expected body %q, got %q", "can not find zipcode\n", body)
	}
}

func TestHandle_GetWeatherError(t *testing.T) {
	mockIsValidCEP = func(cep string) bool { return true }
	mockGetCep = func(cep string) (string, bool, error) {
		return "location", false, nil
	}
	mockGetWeather = func(location string) (interface{}, error) {
		return nil, errors.New("weather error")
	}

	rr := httptest.NewRecorder()
	req := newRequestWithPathValue("GET", "/12345678", "12345678")

	Handle(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, rr.Code)
	}
	if body := rr.Body.String(); body != "can not find zipcode\n" {
		t.Errorf("expected body %q, got %q", "can not find zipcode\n", body)
	}
}

func TestHandle_Success(t *testing.T) {
	mockIsValidCEP = func(cep string) bool { return true }
	mockGetCep = func(cep string) (string, bool, error) {
		return "location", false, nil
	}
	expectedWeather := map[string]interface{}{"temp": 25.0}
	mockGetWeather = func(location string) (interface{}, error) {
		return expectedWeather, nil
	}

	rr := httptest.NewRecorder()
	req := newRequestWithPathValue("GET", "/12345678", "12345678")

	Handle(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, rr.Code)
	}
	if ct := rr.Header().Get("Content-Type"); ct != "text/plain; charset=utf-8" {
		t.Errorf("expected Content-Type text/plain; charset=utf-8, got %s", ct)
	}
}
