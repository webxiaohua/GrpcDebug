/**
 * @author:伯约
 * @date:2024/5/27
 * @note:
**/

package dao

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var (
	DB *sql.DB
)

func init() {
	// 连接数据库
	var dbErr error
	DB, dbErr = sql.Open("sqlite3", "./sql/grpc_debug.db")
	if dbErr != nil {
		panic(dbErr)
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	// 启动时自动检查表是否创建
	createTblDiscoverySQL := `CREATE TABLE IF NOT EXISTS tbl_discovery (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		discovery_id TEXT NOT NULL DEFAULT '',
		mtime TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
		UNIQUE (discovery_id)
	);`
	createTblPathSQL := `CREATE TABLE IF NOT EXISTS tbl_path (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  discovery_id TEXT NOT NULL DEFAULT '',
  path TEXT NOT NULL DEFAULT '',
  params TEXT NOT NULL DEFAULT '',
  mtime TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE (discovery_id, path)
);`
	_, err := DB.Exec(createTblDiscoverySQL)
	if err != nil {
		log.Fatalf("创建 tbl_discovery 表失败: %s\n", err)
	}
	_, err = DB.Exec(createTblPathSQL)
	if err != nil {
		log.Fatalf("创建 tbl_path 表失败: %s\n", err)
	}
}
