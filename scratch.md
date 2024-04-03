

## arrow extension
```
INSTALL arrow;
LOAD arrow;
create table data as select * from range(0,2000) tbl(col);
WITH data_union AS (
      SELECT * FROM data
      UNION ALL
      SELECT * FROM data
  )
  SELECT * FROM to_arrow_ipc((SELECT * FROM data_union ORDER BY col));
```

## iceberg

https://duckdb.org/data/iceberg_data.zip

```
INSTALL iceberg;
LOAD iceberg;
SELECT count(*)
FROM iceberg_scan('../data/iceberg/lineitem_iceberg', allow_moved_paths = true);
```

```
INSTALL iceberg;
LOAD iceberg;
SELECT *
FROM iceberg_metadata('../data/iceberg/lineitem_iceberg', allow_moved_paths = true);
```
```
INSTALL iceberg;
LOAD iceberg;

SELECT *
FROM iceberg_snapshots('../data/iceberg/lineitem_iceberg');
```
