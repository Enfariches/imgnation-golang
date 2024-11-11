package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandStringByte(t *testing.T) {
	tests := []struct {
		Name string
		Size int
	}{
		{
			Name: "Size = 1",
			Size: 1,
		},
		{
			Name: "Size = 2",
			Size: 2,
		},
		{
			Name: "Size = 3",
			Size: 3,
		},
		{
			Name: "Size = 5321",
			Size: 5321,
		},
		{
			Name: "Size = 99999",
			Size: 99999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			str1 := RandStringByte(tt.Size)
			str2 := RandStringByte(tt.Size)

			assert.Len(t, str1, tt.Size)
			assert.Len(t, str2, tt.Size)

			assert.NotEqual(t, str1, str2)
		})
	}
}
