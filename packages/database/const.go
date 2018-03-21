package database

const DefaultStoragePath = "./storage"
const DefaultStorageFile = DefaultStoragePath + "/jirastats.db"
const DefaultDriver = "sqlite3"

// SQL statements
const SelectAllStatement = "SELECT * FROM %s"
const SelectCountStatement = "SELECT COUNT(*) FROM %s"
const TruncateTable = "DELETE FROM %s"
