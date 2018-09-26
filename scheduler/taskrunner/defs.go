package taskrunner

// 定义control里面的三种类型的消息
const (
	READY_TO_DISPATCH = "d" // dispatcher
	READY_TO_EXECUTE  = "e" // executor
	CLOSE             = "c" // close

	VIDEO_PATH = "./videos/"
)

// controller chan
type controlChan chan string

// 需要下发的术语
type dataChan chan interface{}

// Dispatcher和Executor
type fn func(dc dataChan) error
