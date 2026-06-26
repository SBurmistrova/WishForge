package service_test

import (
	"WishForge/internal/service"
	"testing"
)

func TestCheckTitle(t *testing.T) {
	tests := []struct {
		name     string
		title    string
		expected string
		err      error
	}{
		{"TestTitleWithoutSpace", "ValidTitle", "ValidTitle", nil},
		{"TestTitleWithSpase", "   ValidTitleWithSpace   ", "ValidTitleWithSpace", nil},
		{"TestSpace", "", "", service.ErrorTitleEmpty},
		{"TestEmptyTitle", "    ", "", service.ErrorTitleEmpty},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			err := service.CheckTitle(&test.title)

			if test.title != test.expected {
				t.Errorf("Expected title: %s, got: %s", test.expected, test.title)
			}

			if err != nil || test.err != nil {
				if err != nil && test.err == nil {
					t.Errorf("Error not expected, got: %v", err)
				}
				if err == nil && test.err != nil {
					t.Errorf("Expected err: %v", test.err)
				}
				if err != test.err {
					t.Errorf("Expected err: %v, got: %v", test.err, err)
				}
			}
		})
	}
}
