SET sql_safe_updates = FALSE;

USE defaultdb
DROP DATABASE IF EXISTS turbotodo CASCADE;
CREATE DATABASE IF NOT EXISTS turbotodo;

USE turbotodo;

CREATE TABLE todos (
    id STRING(36) PRIMARY KEY,
    datecreated TIMESTAMP,
    title STRING,
    content STRING,
    tags STRING[],
    status STRING,
    version INT8
)

