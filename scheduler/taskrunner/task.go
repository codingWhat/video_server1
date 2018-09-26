package taskrunner

import (
	"github.com/pkg/errors"
	"log"
	"os"
	"sync"
	"video_server/scheduler/dbops"
)

// 删除文件的函数
func deleteVideo(vid string) error {
	err := os.Remove(VIDEO_PATH + vid)
	// 文件不存在也会删除成功，所以需要判断下是否存在
	if err != nil && !os.IsNotExist(err) {
		log.Printf("Deleting video error: %v", err)
		return err
	}
	return nil
}

func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3) // 每次只读3条
	if err != nil {
		log.Printf("Video clear dispatcher error: %v", err)
		return err
	}

	if len(res) == 0 { // 如果没有任何东西取出来
		return errors.New("All task finished")
	}

	for _, id := range res {
		dc <- id
	}

	return nil
}

func VideoClearExecutor(dc dataChan) error {
	// 用一个map把所有的goroutine执行的错误结果带出去
	errMap := &sync.Map{} // 线程安全的map,多个goroutine往里面写都是没有问题的
	var err error

forloop:
	for {
		select {
		case vid := <-dc:
			go func(id interface{}) {
				if err := deleteVideo(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
				if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
			}(vid)
		default:
			break forloop
		}
	}

	errMap.Range(func(key, value interface{}) bool {
		err = value.(error)
		// 只要有一个err存在则停止range过程，返回
		if err != nil {
			return false
		}
		return true
	})
	return err
}
