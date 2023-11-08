[![Go](https://github.com/blablatov/scada4sensors2mssql/actions/workflows/scada-sensors2mssql-action.yml/badge.svg)](https://github.com/blablatov/scada4sensors2mssql/actions/workflows/scada-sensors2mssql-action.yml)
## scada
### RU

Демо пакеты для для записи данных в таблицу MSSQL за транзакцию.
Модуль `sensors2mssql` посредством модуля `mssqldsn` формирует DSN для подключения к БД, далее вызывается модуль `smssqlinsert` для записи данных. 


***Схема обмена данными (scheme exchange of data):***
			

```mermaid
graph TB

  SubGraph1 --> SubGraph1Flow
  subgraph "DSN"
  SubGraph1Flow(Create DSN)
  end
  
  SubGraph3 --> SubGraph2Flow
  subgraph "DBMS MSSQL"
  SubGraph2Flow(Tables of data SCADA in MSSQL)
  end

  subgraph "Module MSSQL"
  Node1[Module write to MSSQL `sensors2mssql`] --> SubGraph1[Request to create DSN `mssqldsn`]
  SubGraph1Flow -- Response with DSN data --> Node1 
  Node1 --> SubGraph3[Insert-method of goroutine `smssqlinsert`]
end
```			

Для проверки, запустить модуль [main4sensors](https://github.com/blablatov/scada4modbus2sensors.git), из строки браузера создать запрос:

	https://localhost:8443


### EN

Demo packages for writing data to MSSQL table per transaction.
The `sensors2mssql` module, using the `mssqldsn` module, forms a DSN for connecting to the database, next calls the `smssqlinsert` module to write data.

To check, run the [main4sensors](https://github.com/blablatov/scada4modbus2sensors.git) module, create a request from the browser line:

	https://localhost:8443

