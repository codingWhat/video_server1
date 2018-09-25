package main
/**
流控模块的具体实现如下：
 */
import "log"

// 流量控制
type ConnLimiter struct {
	concurrentConn int
	bucket chan int
}

func NewConnLimiter(cc int) *ConnLimiter{
	return &ConnLimiter{
		concurrentConn:cc,
		bucket:make(chan int,cc),
	}
}

// 判断是否已经辣到token
func (cl *ConnLimiter) GetConn() bool{
	if len(cl.bucket) >= cl.concurrentConn{
		log.Printf("Reached the rate limitation.")
		return false
	}

	cl.bucket <- 1
	return true
}

// 释放token
func (cl *ConnLimiter) ReleaseConn(){
	c:=<-cl.bucket
	log.Printf("New connction coming: %v",c)
}