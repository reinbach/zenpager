package database

import (
	"testing"

	"code.google.com/p/go.net/context"
	webctx "github.com/goji/context"
	"github.com/zenazn/goji/web"
)

func TestContext(t *testing.T) {
	var ctx = context.Background()
	c := web.C{}
	db := Connect()
	ctx = NewContext(ctx, db)
	webctx.Set(&c, ctx)

	db2 := FromContext(c)

	if db != db2 {
		t.Errorf("Expected db connection from context")
	}
}
