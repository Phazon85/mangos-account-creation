package sqldb

import "testing"

const (
	postgresFile = "tests/postgres.yaml"
)

func TestNewConfig(t *testing.T) {
	t.Run("Postgres config", func(t *testing.T) {
		driver, cfg := newConfig(postgresFile)
		want := "host=localhost port=5432 user=postgres password=changeme dbname=test sslmode=disable"
		if cfg != want {
			t.Errorf("Error loading config, got: %s, want %s", cfg, want)
		}
	})

	t.Run("Nil file name", func(t *testing.T) {
		driver, cfg := newConfig("")
		if cfg != "" {
			t.Errorf("Wanted to get an empty string, but didn't")
		}
	})
}
