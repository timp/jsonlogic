package jsonlogic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDottedRef(t *testing.T) {
	tests := []struct {
		name   string
		data   interface{}
		ref    string
		expect interface{}
	}{
		{
			name:   "nil",
			data:   nil,
			ref:    "blah",
			expect: nil,
		},
		{
			name:   "non-map",
			data:   2.0,
			ref:    "blah",
			expect: nil,
		},
		{
			name:   "single-found",
			data:   map[string]interface{}{"one": 2.0},
			ref:    "one",
			expect: 2.0,
		},
		{
			name:   "single-failed",
			data:   map[string]interface{}{"one": 2.0},
			ref:    "two",
			expect: nil,
		},
		{
			name:   "deep-found",
			data:   map[string]interface{}{"one": map[string]interface{}{"two": 2.0}},
			ref:    "one.two",
			expect: 2.0,
		},
		{
			name:   "deep-miss",
			data:   map[string]interface{}{"one": map[string]interface{}{"two": 2.0}},
			ref:    "one.three",
			expect: nil,
		},
		{
			name:   "deep-non-trivial",
			data:   map[string]interface{}{"one": map[string]interface{}{"two": []interface{}{"hello", 2.0}}},
			ref:    "one.two",
			expect: []interface{}{"hello", 2.0},
		},
	}

	for _, st := range tests {
		t.Run(st.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				v := DottedRef(st.data, st.ref)
				assert.Equal(t, st.expect, v)
			})
		})
	}
}