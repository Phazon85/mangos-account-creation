package sqldb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSQLDBObject(t *testing.T) {
	t.Run("Postgres DB test", func(t *testing.T) {
		ps := NewSQLDBObject(postgresFile)
		assert.NotNil(t, ps)
	})
}
