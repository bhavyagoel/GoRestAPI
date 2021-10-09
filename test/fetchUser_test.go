package unittesting

import (
	"RestAPI/user"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func JSONEqual(a, b io.Reader) (bool, error) {
    var j, j2 interface{}
    d := json.NewDecoder(a)
    if err := d.Decode(&j); err != nil {
        return false, err
    }
    d = json.NewDecoder(b)
    if err := d.Decode(&j2); err != nil {
        return false, err
    }
    return reflect.DeepEqual(j2, j), nil
}

func JSONBytesEqual(a, b []byte) (bool, error) {
    var j, j2 interface{}
    if err := json.Unmarshal(a, &j); err != nil {
        return false, err
    }
    if err := json.Unmarshal(b, &j2); err != nil {
        return false, err
    }
    return reflect.DeepEqual(j2, j), nil
}

func TestFetchUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/user/getUserById", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("userID", "393d0be9-1091-4bd0-822b-915980d94b6d")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(user.GetUserById)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"handler Returned Wrong Status Code: got %v want %v", 
			status, http.StatusOK,
		)
	}	
	
	expected := []byte(`{
		"Email": "bgoel4132@gmail.com",
		"Password": "$2a$14$vJSwx9sOr2e9KMLiVji9tOfB5AbMjro69R/D1wN5Yqa/IpFMCr2Tq",
		"Posts": null,
		"_id": "61617e08b11882299ca8a3fa",
		"id": "393d0be9-1091-4bd0-822b-915980d94b6d",
		"name": "bhavya goel",
		"post": [
			"56f18233-07bc-4472-9f71-4192f3cc2a3f.jpg",
			"927d4eef-7a77-4865-9cb3-0907fa44a871.jpg"
		]
	}`)

	// eq, err := JSONBytesEqual(expected, rr.Body)
    // fmt.Println("a=b\t", eq, "with error", err)

	if rr.Body != expected {
		t.Errorf(
			"Handler Returned Unexpected Body: got %v want %v", 
			rr.Body.String(), expected,
		)
	}

}