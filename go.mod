module github.com/blablatov/scada4sensors2mssql

go 1.16

require (
	github.com/blablatov/scada4sensors2mssql/mssqldsn v0.0.0-00010101000000-000000000000
	github.com/blablatov/scada4sensors2mssql/smssqlinsert v0.0.0-00010101000000-000000000000
	github.com/denisenkom/go-mssqldb v0.12.2
)

replace github.com/blablatov/scada4sensors2mssql/mssqldsn => ./mssqldsn

replace github.com/blablatov/scada4sensors2mssql/smssqlinsert => ./smssqlinsert
