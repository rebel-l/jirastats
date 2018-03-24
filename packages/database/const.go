package database

const DefaultStoragePath = "./storage"
const DefaultStorageFile = DefaultStoragePath + "/jirastats.db"
const DefaultDriver = "sqlite3"

// SQL statements
const SelectStatement = "SELECT %s FROM %s %s %s %s"
const SelectAllStatement = "SELECT * FROM %s"
const TruncateTable = "DELETE FROM %s"
const exprCount = "COUNT(*)"

// Formats
const dateTimeFormat = "2006-01-02T15:04:05Z"
