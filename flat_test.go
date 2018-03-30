package flatjson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatten(t *testing.T) {

	var flagtests = []struct {
		jsonStr string
		eKey    string
		eValue  string
		hasErr  bool
	}{
		{`{"key": "hello"}`, "key", "hello", false},
		{`{"key": ""}`, "key", "", false},
		{`{"key": 0}`, "key", "0", false},
		{`{"key": 123}`, "key", "123", false},
		{`{"key": 123.0}`, "key", "123", false},
		{`{"key": 123.1}`, "key", "123.1", false},
		{`{"key": true}`, "key", "true", false},
		{`{"key": false}`, "key", "false", false},
		{`{"key": {"key": false}}`, "key.key", "false", false},
		{`{"key": {"key":  {"key": false}}}`, "key.key.key", "false", false},
		{`{"key": [{"key": false}, {"key": true}]}`, "key.#", "2", false},
		{`{"key": [{"key": false}, {"key": true}]}`, "key.0.key", "false", false},
		{`{"key": [{"key": false}, {"key": true}]}`, "key.1.key", "true", false},
		{`{"key" - "hello"}`, "", "", true},
	}
	for _, tt := range flagtests {
		m, err := Flatten(tt.jsonStr)
		if tt.hasErr {
			assert.Error(t, err)
			assert.Empty(t, m)
			continue
		}
		assert.NoError(t, err)
		assert.Equal(t, tt.eValue, m[tt.eKey], tt.eKey, m)
	}
}
