package infoblox

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestARecordService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/wapi/v%s/record:a", WAPIVersion), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name":"host01"}`)
	})

	a, _, err := client.A.Get("host01")
	if err != nil {
		t.Errorf("DNS.Get returned error: %v", err)
	}

	expected := &A{Name: String("host01")}
	if !reflect.DeepEqual(a, expected) {
		t.Errorf("DNS.Get returned %#v, expected %#v", a, expected)
	}
}

func TestARecordService_Create(t *testing.T) {
	setup()
	defer teardown()

	input := &A{Name: String("host01")}

	mux.HandleFunc(fmt.Sprintf("/wapi/v%s/record:a", WAPIVersion), func(w http.ResponseWriter, r *http.Request) {
		v := new(A)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, expected %+v", v, input)
		}

		fmt.Fprint(w, `{"name":"host01"}`)
	})

	a, _, err := client.A.Create(input)
	if err != nil {
		t.Errorf("DNS.Create returned error: %v", err)
	}

	expected := &A{Name: String("host01")}
	if !reflect.DeepEqual(a, expected) {
		t.Errorf("DNS.Create returned %+v, expected %+v", a, expected)
	}
}
