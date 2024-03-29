package electricitymaps

import (
	"testing"

	"gotest.tools/assert"
)

func TestNew(t *testing.T) {
	want := ElectricityMaps{name: defaultName}
	got := New()
	assert.Equal(t, *got, want)
}

func TestName(t *testing.T) {
	want := defaultName
	got := New().Name()
	assert.Equal(t, got, want)
}
