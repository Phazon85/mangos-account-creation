package sql

const (
	postgresFile = "postgres.yaml"
)

//SQL contains the DB object from your SQL implementation of choice
type SQL struct {
	DB interface{}
}

//NewSQLDBObject returns a SQL struct containing the SQL implementation of your choice. Currenlty only support postgres
//TODO: Add MySQL and MicrosoftSQL support
func NewSQLDBObject() *SQL {
	sql := NewSQLDBObject(postgresFile)
	return &SQL{
		DB: sql,
	}
}
