package database

const DefaultFile  = "./storage/jirastats.db"
const DefaultDriver = "sqlite3"

// SQL statements
const SelectAllStatement = "SELECT * FROM %s"
const TruncateTable = "DELETE FROM %s"
