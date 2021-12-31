package lib

import (
	"os"
	"testing"
)

func TestFormatMessage(t *testing.T) {

	testCases := []struct {
		env     string
		message string
		want    string
	}{
		{
			"",
			"ABCD",
			"[Env: UNKNOWN] Message: ABCD",
		},
		{
			"FOO",
			"Message",
			"[Env: FOO] Message: Message",
		},
		{
			"ENV",
			"FooBar",
			"[Env: ENV] Message: FooBar",
		},
	}

	for _, tc := range testCases {
		os.Setenv("DEPLOY_ENV", tc.env)
		if got := FormatMessage(tc.message); got != tc.want {
			t.Errorf("got '%s', want '%s'", got, tc.want)
		}
	}
}
