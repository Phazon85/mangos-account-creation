package sqldb

//SQL contains the DB object from your SQL implementation of choice
type SQL struct {
	DB interface{}
}

//NewSQLDBObject returns a SQL struct containing the SQL implementation of your choice. Currenlty only support postgres
//TODO: Add MySQL and MicrosoftSQL support
func NewSQLDBObject(file string) *SQL {
	sql := NewSQLDBObject(file)
	return &SQL{
		DB: sql,
	}
}
