package site

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const baseLayout = "../../../ui/html/site/base.layout.html"

func TestNewHandlers(t *testing.T) {
	h := NewMainHandlers(slog.Default())
	assert.NotNil(t, h)
	assert.NotNil(t, h.log)
}

func TestAbout(t *testing.T) {
	handler := &MainHandlers{log: slog.Default()}

	ts := httptest.NewServer(http.HandlerFunc(handler.MenuForCatering))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(t, err)
	defer resp.Body.Close()

}

func TestMenuForCatering(t *testing.T) {
	handler := &MainHandlers{log: slog.Default()}

	ts := httptest.NewServer(http.HandlerFunc(handler.MenuForCatering))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(t, err)
	defer resp.Body.Close()
}

func TestContacts(t *testing.T) {
	hadler := &MainHandlers{log: slog.Default()}

	ts := httptest.NewServer(http.HandlerFunc(hadler.Contacts))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(t, err)
	defer resp.Body.Close()
}

func TestIndex(t *testing.T) {
	hadler := &MainHandlers{log: slog.Default()}

	ts := httptest.NewServer(http.HandlerFunc(hadler.Index))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(t, err)
	defer resp.Body.Close()
}

func TestMainMenu(t *testing.T) {
	hadler := &MainHandlers{log: slog.Default()}

	ts := httptest.NewServer(http.HandlerFunc(hadler.MainMenu))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(t, err)
	defer resp.Body.Close()
}

func TestPrivacyPolicy(t *testing.T) {
	hadler := &MainHandlers{log: slog.Default()}

	ts := httptest.NewServer(http.HandlerFunc(hadler.PrivacyPolicy))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(t, err)
	defer resp.Body.Close()
}

func TestUserAgreement(t *testing.T) {
	hadler := &MainHandlers{log: slog.Default()}

	ts := httptest.NewServer(http.HandlerFunc(hadler.UserAgreement))

	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(t, err)
	defer resp.Body.Close()
}

func TestNewsAndEvents(t *testing.T) {
	hadler := &MainHandlers{log: slog.Default()}

	ts := httptest.NewServer(http.HandlerFunc(hadler.NewsAndEvents))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(t, err)
	defer resp.Body.Close()
}
