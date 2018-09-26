package dbops

import "log"

// api->videoid->mysql
// dispatcher->mysql-diveoid->datachannel
// executor->datachannel->videoid->delete videos

// 读
func ReadVideoDeletionRecord(count int) ([]string, error) {
	stmtOut, err := dbConn.Prepare("select video_id from video_del_rec limit ?")
	defer stmtOut.Close()

	var ids []string
	if err != nil {
		return ids, err
	}

	rows, err := stmtOut.Query(count)
	if err != nil {
		log.Printf("Query VideoDeletionRecord Error: %v", err)
		return ids, err
	}

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}
	return ids, nil
}

// 删除
func DelVideoDeletionRecord(vid string) error {
	stmtDel, err := dbConn.Prepare("delete from video_del_rec where vid = ?")
	defer stmtDel.Close()
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		log.Printf("Delting VideoDelitionRecord Error: %v", err)
		return err
	}

	return nil
}
