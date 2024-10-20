package main

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestSplit(t *testing.T) {
	type test struct {
		input         http.Header
		wantString    string
		wantErrString string
	}

	req1, _ := http.NewRequest("GET", "", nil)
	req1.Header.Set("Authorization", "ApiKey HAHAHAHAHA")
	req2, _ := http.NewRequest("GET", "", nil)
	req3, _ := http.NewRequest("GET", "", nil)
	req3.Header.Set("Authorization", "HAHAHAHAH")
	tests := []test{

		{input: req1.Header, wantString: "HAHAHAHAHA", wantErrString: ""},
		{input: req2.Header, wantString: "", wantErrString: "no authorization header included"},
		{input: req3.Header, wantString: "", wantErrString: "malformed authorization header"},
	}
	for _, tc := range tests {
		gotString, gotErr := auth.GetAPIKey(tc.input)
		gotErrString := ""
		if gotErr != nil {
			gotErrString = gotErr.Error()
		}

		if !reflect.DeepEqual(tc.wantString, gotString) {
			t.Fatalf("expected: %v, got: %v", tc.wantString, gotString)
		}
		if !reflect.DeepEqual(tc.wantErrString, gotErrString) {
			t.Fatalf("expected: %v, got: %v", tc.wantErrString, gotErrString)
		}
	}

}
