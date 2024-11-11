package random

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandStringByte(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{
			name: "size = 1",
			size: 1,
		},
		{
			name: "size = 2",
			size: 2,
		},
		{
			name: "size = 5",
			size: 5,
		},
		{
			name: "size = 30",
			size: 30,
		},
		{
			name: "size = 999",
			size: 999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str1 := RandStringByte(tt.size)
			str2 := RandStringByte(tt.size)

			assert.Len(t, str1, tt.size)
			assert.Len(t, str2, tt.size)

			assert.NotEqual(t, str1, str2)
			fmt.Println(str1, str2)
		})
	}
}
