/* 创建sqlite库 */
sqlite3 grpc_debug.db

/* 创建表 */
CREATE TABLE IF NOT EXISTS tbl_discovery (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   discovery_id TEXT NOT NULL DEFAULT '',
   mtime TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
   UNIQUE (discovery_id)
);
CREATE INDEX ix_tbl_discovery_mtime ON tbl_discovery (mtime);

CREATE TABLE IF NOT EXISTS tbl_path (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  discovery_id TEXT NOT NULL DEFAULT '',
  path TEXT NOT NULL DEFAULT '',
  params TEXT NOT NULL DEFAULT '',
  mtime TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE (discovery_id, path)
);
CREATE INDEX ix_tbl_path_mtime ON tbl_path (mtime);