package utility

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestHkdf(t *testing.T) {
	emptyKey := make([]byte, 16)
	for i := 0; i < 6; i++ {
		id := uuid.New().String()
		key := HKDF16(id, "test", "test")
		assert.Equal(t, 16, len(key))
		assert.NotEqual(t, emptyKey, key)
		newKey := HKDF16(id, "test", "test")
		assert.Equal(t, key, newKey)

		fmt.Printf("id: %s, key: %x\n", id, key)
		fmt.Printf("new key: %x\n", newKey)
	}
}
