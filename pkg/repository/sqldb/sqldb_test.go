package sqldb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSQLDBObject(t *testing.T) {
	t.Run("Postgres DB test", func(t *testing.T) {
		ps := New(postgresFile)
		assert.NotNil(t, ps)
	})

	t.Run("MySQL DB test", func(t *testing.T) {
		mysql := New(mysqlFile)
		assert.NotNil(t, mysql)
	})
}
