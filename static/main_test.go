package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// test verifie si la réponse contient la bonne URL longue
func TestRedirectToLongURL(t *testing.T) {
	urlShort := "http://localhost:8080/abc123"
	urlLong := "https://example.com"
	urlMap := map[string]string{"short_url": urlShort}

	reqBody, err := json.Marshal(urlMap)
	if err != nil {
		t.Fatalf("Erreur lors de la création du corps de la requête JSON: %v", err)
	}

	req := httptest.NewRequest("POST", "/get-long-url", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	// redirectToLongURL(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Erreur attendue: %v; Récu: %v", http.StatusOK, resp.Code)
	}

	var responseData map[string]string
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		t.Fatalf("Erreur lors de la décodage de la réponse JSON: %v", err)
	}

	if longURL, ok := responseData["long_url"]; !ok || longURL != urlLong {
		t.Errorf("URL longue incorrecte. Attendu: %s; Récu: %s", urlLong, longURL)
	}
}
