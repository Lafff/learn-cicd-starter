package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	emptyHeader := make(http.Header)
	noAuthHeader := make(http.Header)
	noAuthHeader.Add("SomeKey", "SomeValue")
	authHeader := make(http.Header)
	authHeader.Add("Authorization", "ApiKey 12354")
	noSplitAuthHeader := make(http.Header)
	noSplitAuthHeader.Add("Authorization", "ApiKey12354")
	emptyAuthHeader := make(http.Header)
	emptyAuthHeader.Add("Authorization", "")
	testCases := []struct {
		name        string
		header      http.Header
		expectedKey string
		expectedErr bool
	}{
		{
			name:        "emptyHeader",
			header:      emptyHeader,
			expectedKey: "",
			expectedErr: true,
		},
		{
			name:        "noAuthHeader",
			header:      noAuthHeader,
			expectedKey: "",
			expectedErr: true,
		},
		{
			name:        "authHeader",
			header:      authHeader,
			expectedKey: "12354",
			expectedErr: false,
		},
		{
			name:        "noSplitAuthHeader",
			header:      noSplitAuthHeader,
			expectedKey: "",
			expectedErr: true,
		},
		{
			name:        "emptyAuthHeader",
			header:      emptyAuthHeader,
			expectedKey: "",
			expectedErr: true,
		},
	}
	for _, testCase := range testCases {
		key, err := GetAPIKey(testCase.header)
		if key != testCase.expectedKey {
			t.Errorf("Key is not correct. Got key %s, but expected %s. Test case %s", key, testCase.expectedKey, testCase.name)
		}
		if testCase.expectedErr && err == nil {
			t.Errorf("Expected erorr, but there is no error. Test case %s", testCase.name)
		}
		if !testCase.expectedErr && err != nil {
			t.Errorf("Expected dont have error, but there is error. Test case %s", testCase.name)
		}
	}
}
