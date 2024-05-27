/**
 * @author:伯约
 * @date:2024/5/27
 * @note:
**/

package pkg

import (
	"database/sql"
	"fmt"
	"time"
)

type TblPath struct {
	ID          uint64    `json:"id"`
	DiscoveryID string    `json:"discovery_id"`
	Path        string    `json:"path"`
	Params      string    `json:"params"`
	Mtime       time.Time `json:"mtime"`
}

// 获取discoveryId对应的path列表
func DbPathListByDiscovery(db *sql.DB, discoveryId string) (list []*TblPath, err error) {
	list = make([]*TblPath, 0)
	rows, rowsErr := db.Query("SELECT id,discovery_id,path,params FROM tbl_path where discovery_id = ? order by mtime desc", discoveryId)
	if rowsErr != nil {
		err = rowsErr
		fmt.Println("db.Query() error:", err)
		return
	}
	for rows.Next() {
		tmp := &TblPath{}
		err = rows.Scan(&tmp.ID, &tmp.DiscoveryID, &tmp.Path, &tmp.Params)
		if err != nil {
			fmt.Println("rows.Scan() error:", err)
			return
		}
		list = append(list, tmp)
	}
	err = rows.Err()
	return
}

// 新增
func DbPathInsert(db *sql.DB, discoveryId, path, params string) (id int64, err error) {
	result, dbErr := db.Exec("INSERT INTO tbl_path (discovery_id,path,params) VALUES (?,?,?) ON DUPLICATE KEY UPDATE mtime=?,params=?", discoveryId, path, params, time.Now(), params)
	if dbErr != nil {
		err = dbErr
		return
	}
	id, err = result.LastInsertId()
	return
}

// 保留最近N条记录
func DbPathKeepLast(db *sql.DB, discoveryId string, reserveCnt int64) (err error) {
	_, err = db.Exec("DELETE FROM tbl_path where id  NOT IN (SELECT id from tbl_path WHERE discovery_id = ? order by mtime desc limit ?)", discoveryId, reserveCnt)
	return
}
