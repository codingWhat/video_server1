package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

/**
1. user->api service->delete video
2. api service->scheduler->write video deletion record
3. timer
4. timer->runner->read wvdr->exec->delete video from folder
*/

// 由于这个api比较简单，我们把鉴权session,错误码等都去掉了
func AddVideoDeletionRecord(vid string) error {
	stmtIns, err := dbConn.Prepare("insert into video_del_rec (video_id) values (?)")
	defer stmtIns.Close()
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(vid)
	if err != nil {
		log.Printf("AddVideoDeletionRecord error: %v", err)
		return err
	}

	return nil
}
