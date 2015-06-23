package contacts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"code.google.com/p/go.net/context"
	webctx "github.com/goji/context"
	"github.com/zenazn/goji/web"

	"github.com/reinbach/zenpager/auth"
	"github.com/reinbach/zenpager/database"
)

func SetupWebContext() web.C {
	var ctx = context.Background()
	c := web.C{}
	db := database.Connect()
	ctx = database.NewContext(ctx, db)
	webctx.Set(&c, ctx)
	return c
}

// Routes
func TestRoutes(t *testing.T) {
	Routes()
}

// List
func TestList(t *testing.T) {
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	List(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}
}

// Item
func TestItem(t *testing.T) {
	db := database.Connect()
	u := auth.User{
		Email: "contact-api-10@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	cu := Contact{
		Name: "Joe",
		User: u,
	}
	cu.Create(db)

	r, err := http.NewRequest("POST", fmt.Sprintf("/%d", cu.ID), nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	Item(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", cu.ID)}
	Item(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}

	// not found
	r, err = http.NewRequest("POST", "/321", nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": "321"}
	Item(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}
}

// Add
func TestAdd(t *testing.T) {
	db := database.Connect()
	u := auth.User{
		Email: "contact-api-1@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	type Data struct {
		Name  string
		Email string
	}
	d := Data{Name: "test", Email: u.Email}
	j, _ := json.Marshal(d)
	b := bytes.NewBuffer(j)

	r, err := http.NewRequest("POST", "/", b)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	Add(c, w, r)
	if w.Code != http.StatusAccepted {
		t.Errorf("%v expected, got %v instead", http.StatusAccepted, w.Code)
	}

	cu := Contact{
		User: u,
	}
	cu.GetByUser(db)
	if cu.Name != d.Name {
		t.Errorf("Expected data to be saved")
	}

	// attempt to add again
	w = httptest.NewRecorder()
	Add(c, w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("%v expected, got %v instead", http.StatusBadRequest, w.Code)
	}
}

// Add invalid data
func TestAddInvalidData(t *testing.T) {
	db := database.Connect()
	u := auth.User{
		Email: "contact-api-2@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	type Data struct {
		ID int64
	}
	d := Data{ID: 321}
	j, _ := json.Marshal(d)
	b := bytes.NewBuffer(j)

	r, err := http.NewRequest("POST", "/", b)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	Add(c, w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("%v expected, got %v instead", http.StatusBadRequest, w.Code)
	}
}

// Add that creates user
func TestAddCreateUser(t *testing.T) {
	db := database.Connect()
	u := auth.User{
		Email: "contact-api-3@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	type Data struct {
		Name  string
		Email string
	}
	d := Data{Name: "test", Email: "contact-api-4@example.com"}
	j, _ := json.Marshal(d)
	b := bytes.NewBuffer(j)

	r, err := http.NewRequest("POST", "/", b)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	Add(c, w, r)
	if w.Code != http.StatusAccepted {
		t.Errorf("%v expected, got %v instead", http.StatusAccepted, w.Code)
	}
}

// Update
func TestUpdate(t *testing.T) {
	db := database.Connect()
	u := auth.User{
		Email: "contact-api-5@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	cu := Contact{
		Name: "Joe",
		User: u,
	}
	cu.Create(db)

	cu.Name = "Jane"
	j, _ := json.Marshal(cu)
	b := bytes.NewBuffer(j)

	r, err := http.NewRequest("PUT", fmt.Sprintf("/%d", cu.ID), b)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	Update(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", cu.ID)}
	Update(c, w, r)
	if w.Code != http.StatusAccepted {
		t.Errorf("%v expected, got %v instead", http.StatusAccepted, w.Code)
	}

	cu.Get(db)
	if cu.Name != "Jane" {
		t.Errorf("Expected data to be updated")
	}
}

// Delete
func TestDelete(t *testing.T) {
	db := database.Connect()
	u := auth.User{
		Email: "contact-api-6@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	cu := Contact{
		Name: "Joe",
		User: u,
	}
	cu.Create(db)

	r, err := http.NewRequest("DELETE", fmt.Sprintf("/%d", cu.ID), nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	Delete(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", cu.ID)}
	Delete(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}
}
