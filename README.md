# pgpassthru

```zsh
curl -H 'Content-Type:application/json' --data '{"sql": "select 1, now(), '\''x'\'';", "dsn": "postgresql://eric@localhost:5432/test"}' localhost:8090

{"result":{"cols":[{"name":"?column?","datatype":"int32"},{"name":"now","datatype":"Time"},{"name":"?column?","datatype":"string"}],"rows":[["1","2020-11-01T19:49:52.17746-08:00","x"]]},"error":""}
```
