package handlers

import (
	"testing"

	"github.com/odpf/columbus/record"
)

func TestTypeHandler(t *testing.T) {
	t.Run("validateType", func(t *testing.T) {
		t.Run("should fail", func(t *testing.T) {
			types := []record.Type{
				{
					Name: "",
				},
				{
					Name:           "four",
					Classification: "",
				},
			}

			handler := new(TypeHandler)
			for _, recordType := range types {
				err := handler.validateType(recordType)
				if err == nil {
					t.Errorf("expected type %#v to fail validation", recordType)
				}
			}
		})
	})
}
