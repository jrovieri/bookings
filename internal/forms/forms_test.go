package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postData := url.Values{}
	postData.Add("a", "a")
	postData.Add("b", "b")
	postData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/unknow", nil)
	r.PostForm = postData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows 'does not required fields' when it does")
	}
}

func TestForm_Has(t *testing.T) {
	postData := url.Values{}
	form := New(postData)

	hasField := form.Has("middle_name")
	if hasField {
		t.Error("form show has field when it does not")
	}

	postData = url.Values{}
	postData.Add("middle_name", "Smith")

	form = New(postData)
	hasField = form.Has("middle_name")
	if !hasField {
		t.Error("shows 'form does not have field' when it should")
	}
}

func TestForm_MinLength(t *testing.T) {
	postData := url.Values{}
	form := New(postData)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows minimum length for non-existent field")
	}

	errMsg := form.Errors.Get("x")
	if errMsg == "" {
		t.Error("should have an error but not get one")
	}

	postData = url.Values{}
	postData.Add("middle_name", "Smith")
	form = New(postData)

	form.MinLength("middle_name", 100)
	if form.Valid() {
		t.Error("shows minimum length of 100 met when data is shorter")
	}

	postData = url.Values{}
	postData.Add("middle_name", "Smith")
	form = New(postData)

	form.MinLength("middle_name", 3)
	if !form.Valid() {
		t.Error("shows minimum length of 3 is not met when it is")
	}

	errMsg = form.Errors.Get("middle_name")
	if errMsg != "" {
		t.Error("should not have an error but got one")
	}
}

func TestForm_Email(t *testing.T) {
	postData := url.Values{}
	form := New(postData)

	form.IsEmail("xxx_email")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	postData = url.Values{}
	postData.Add("email", "john.doe@gmail.com")
	form = New(postData)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got an invalid email when should not have")
	}

	postData = url.Values{}
	postData.Add("email", "john.doe#gmailcom")
	form = New(postData)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("got an valid for invalid email address")
	}
}
