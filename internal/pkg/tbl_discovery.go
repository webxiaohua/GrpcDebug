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

type TblDiscovery struct {
	ID          uint64    `json:"id"`
	DiscoveryID string    `json:"discovery_id"`
	Mtime       time.Time `json:"mtime"`
}

// 获取前100条服务发现信息
func DbDiscoveryList(db *sql.DB) (list []*TblDiscovery, err error) {
	list = make([]*TblDiscovery, 0)
	rows, rowsErr := db.Query("SELECT id,discovery_id FROM tbl_discovery order by discovery_id asc limit 100")
	if rowsErr != nil {
		err = rowsErr
		fmt.Println("db.Query() error:", err)
		return
	}
	for rows.Next() {
		tmp := &TblDiscovery{}
		err = rows.Scan(&tmp.ID, &tmp.DiscoveryID)
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
func DbDiscoveryInsert(db *sql.DB, discoveryId string) (id int64, err error) {
	result, dbErr := db.Exec("INSERT OR IGNORE INTO tbl_discovery (discovery_id) VALUES (?)", discoveryId)
	if dbErr != nil {
		err = dbErr
		return
	}
	id, err = result.LastInsertId()
	return
}
