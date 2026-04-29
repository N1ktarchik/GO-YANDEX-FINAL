package repository

const schema string = `
CREATE TABLE IF NOT EXISTS scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT "",
    title VARCHAR NOT NULL DEFAULT "task_default_name",
    comment TEXT,
    repeat VARCHAR
);
CREATE INDEX IF NOT EXISTS idx_date ON scheduler(date);
`
