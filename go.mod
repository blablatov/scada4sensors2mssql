module github.com/blablatov/scada4sensors2mssql

go 1.16

require (
	github.com/blablatov/scada4sensors2mssql/mssqldsn v0.0.0-20220920112358-6f8e831869c7
	github.com/blablatov/scada4sensors2mssql/smssqlinsert v0.0.0-20220920112358-6f8e831869c7
)

replace github.com/blablatov/scada4sensors2mssql/mssqldsn => ./mssqldsn

replace github.com/blablatov/scada4sensors2mssql/smssqlinsert => ./smssqlinsert
