package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/test-url", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("Got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("Form shows valid when required fields are missing")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	form = New(postedData)

	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("Form shows 'does not have required fields' when it does.")
	}

}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	if form.Has("a") {
		t.Error("Form shows 'has field' when field is missing")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")

	form = New(postedData)

	if !form.Has("a") {
		t.Error("Form shows 'does not have field' when field exists")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	if form.MinLength("a", 3) {
		t.Error("Form shows 'min length valid' when field does not exist")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")

	form = New(postedData)

	if form.MinLength("a", 3) {
		t.Error("Form shows 'min length valid' when field is too short")
	}

	postedData = url.Values{}
	postedData.Add("a", "abc")

	form = New(postedData)

	if !form.MinLength("a", 3) {
		t.Error("Form shows 'min length Invalid' when field is long enough")
	}

}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.IsEmail("a")
	if form.Valid() {
		t.Error("Form shows 'email valid' when field does not exist")
	}

	isError := form.Errors.Get("a")
	if isError == "" {
		t.Error("Should be an error but did not get one")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")

	form = New(postedData)

	form.IsEmail("a")
	if form.Valid() {
		t.Error("Form shows 'email valid' when field is not an email")
	}

	postedData = url.Values{}
	postedData.Add("a", "test@email.com")

	form = New(postedData)

	form.IsEmail("a")
	if !form.Valid() {
		t.Error("Form shows 'Invalid email' when field is an email")
	}

	isError = form.Errors.Get("a")
	if isError != "" {
		t.Error("Should not have an error, but got one")
	}

}
