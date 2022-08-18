package main

import (
	"fmt"
	"github.com/bigwhite/tcp-server-demo1/frame"
	"github.com/bigwhite/tcp-server-demo1/packet"
	"github.com/lucasepe/codename"
	"net"
	"sync"
	"time"
)

func main() {
	/*WaitGroup用于等待一组线程的结束。父线程调用Add方法来设定应等待的线程的数量。
	每个被等待的线程在结束时应调用Done方法。同时，主线程里可以调用Wait方法阻塞至所有线程结束。*/
	var wg sync.WaitGroup
	var num int = 5

	/*Add方法向内部计数加上delta，delta可以是负数；如果内部计数器变为0，
	Wait方法阻塞等待的所有线程都会释放，如果计数器小于0，方法panic。
	注意Add加上正数的调用应在Wait之前，否则Wait可能只会等待很少的线程。
	一般来说本方法应在创建新的线程或者其他应等待的事件之前调用。*/
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func(i int) {
			defer wg.Done() //Done方法减少WaitGroup计数器的值，应在线程的最后执行。
			startClient(i)
		}(i + 1)
	}
	wg.Wait() //Wait方法阻塞直到WaitGroup计数器减为0。

}

func startClient(i int) {
	quit := make(chan struct{})
	done := make(chan struct{})
	conn, err := net.Dial("tcp", ":8888") //Dial函数和服务端建立连接：
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}

	defer conn.Close() // Close方法关闭conn连接,并会导致任何阻塞中的Read或Write方法不再阻塞并返回错误
	fmt.Printf("[client %d]: dial ok", i)

	// 生成payload
	rng, err := codename.DefaultRNG()
	if err != nil {
		panic(err)
	}

	frameCodec := frame.NewMyFrameCodec()
	var counter int

	go func() {
		// handle ack
		for {
			select {
			case <-quit:
				done <- struct{}{}
				return
			default:
			}
			// SetReadDeadline设定该连接的读操作deadline，参数t为零值表示不设置期限
			conn.SetReadDeadline(time.Now().Add(time.Second * 5))
			ackFramePayLoad, err := frameCodec.Decode(conn)
			if err != nil {
				if e, ok := err.(net.Error); ok {
					if e.Timeout() { // 错误是否为超时？
						continue
					}
				}
				panic(err)
			}

			p, err := packet.Decode(ackFramePayLoad)
			submitAck, ok := p.(*packet.SubmitAck)
			if !ok {
				panic("not submitack")
			}
			fmt.Printf("[client %d]: the result of submit ack[%s] is %d\n", i, submitAck.ID, submitAck.Result)
		}
	}()

	for {
		// send submit
		counter++
		id := fmt.Sprintf("%08d", counter) // 8 byte string
		payload := codename.Generate(rng, 4)
		s := &packet.Submit{
			ID:      id,
			Payload: []byte(payload),
		}

		framePayload, err := packet.Encode(s)
		if err != nil {
			panic(err)
		}

		fmt.Printf("[client %d]: send submit id = %s, payload=%s, frame length = %d\n",
			i, s.ID, s.Payload, len(framePayload)+4)

		err = frameCodec.Encode(conn, framePayload)
		if err != nil {
			panic(err)
		}

		time.Sleep(1 * time.Second) //Sleep阻塞当前go进程至少d代表的时间段。d<=0时，Sleep会立刻返回。
		if counter >= 10 {
			quit <- struct{}{}
			<-done
			fmt.Printf("[client %d]: exit ok\n", i)
			return
		}
	}
}
