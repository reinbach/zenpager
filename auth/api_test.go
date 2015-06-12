package auth

import (
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"code.google.com/p/go.net/context"
	webctx "github.com/goji/context"
	"github.com/zenazn/goji/web"

	"github.com/reinbach/zenpager/database"
)

func SetupWebContext() web.C {
	var db *sql.DB
	var ctx = context.Background()
	c := web.C{}
	ctx = database.NewContext(ctx, db)
	webctx.Set(&c, ctx)
	return c
}

// login no payload
func TestLoginNoPayload(t *testing.T) {
	body := url.Values{}
	body.Set("email", "test@example.com")
	body.Set("password", "123")
	b := bytes.NewBufferString(body.Encode())

	r, err := http.NewRequest("POST", "/login", b)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	Login(c, w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("400 expected, got %v instead", w.Code)
	}
}
