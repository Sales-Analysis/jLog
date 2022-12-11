package dotenv

import (
	"testing"
)

func TestIsNotContains(t *testing.T) {
	if ok := contains("test", []string{"test1", "test2"}); ok {
		t.Errorf("value 'test' is contains")
	}
}

func TestIsContains(t *testing.T) {
	if ok := contains("test", []string{"test1", "test"}); !ok {
		t.Errorf("value 'test' is not contains")
	}
}
