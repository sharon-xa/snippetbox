package main

import (
	"testing"

	"github.com/sharon-xa/snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	assert.Equal(t, code, 200)
	assert.Equal(t, body, "OK")
}
