package trace

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

var goroutineSpace = []byte("goroutine ")

func curGoroutineID() uint64 {
	b := make([]byte, 64)

	/*Stack 将调用其的go进程,调用栈踪迹格式化后写入到buf中并返回写入的字节数。
	若all为true，函数会在写入当前go程的踪迹信息后，将其它所有go程的调用栈踪迹都格式化写入到buf中。*/
	b = b[:runtime.Stack(b, false)]

	b = bytes.TrimPrefix(b, goroutineSpace) //返回去除s可能的前缀prefix的子切片。（共用底层数组）
	i := bytes.IndexByte(b, ' ')            //字符c在s中第一次出现的位置，不存在则返回-1。
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}
	b = b[:i]
	n, err := strconv.ParseUint(string(b), 10, 64) //ParseUint类似ParseInt但不接受正负号，用于无符号整型。
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}
	return n
}

func printTrace(id uint64, name, arrow string, indent int) {
	indents := ""
	for i := 0; i < indent; i++ {
		indents += "    "
	}
	fmt.Printf("g[%05d]:%s%s%s\n", id, indents, arrow, name)
}

var mu sync.Mutex //Mutex是一个互斥锁，可以创建为其他结构体的字段；零值为解锁状态。
// Mutex类型的锁和线程无关，可以由不同的线程加锁和解锁。

var m = make(map[uint64]int)

func Trace() func() {

	/*Caller报告当前go程调用栈所执行的函数的文件和行号信息。实参skip为上溯的栈帧数，
	0表示Caller的调用者（Caller所在的调用栈）。（由于历史原因，skip的意思在Caller和Callers中并不相同。）
	函数的返回值为调用栈标识符、文件名、该调用在文件中的行号。如果无法获得信息，ok会被设为false。*/
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}

	fn := runtime.FuncForPC(pc)
	name := fn.Name()
	gid := curGoroutineID()

	mu.Lock() //Lock方法锁住m，如果m已经加锁，则阻塞直到m解锁。
	indents := m[gid]
	m[gid] = indents + 1
	mu.Unlock() //Unlock方法解锁m，如果m未加锁会导致运行时错误。锁和线程无关，可以由不同的线程加锁和解锁。
	printTrace(gid, name, "->", indents+1)
	return func() {
		mu.Lock()
		indents := m[gid]
		m[gid] = indents - 1
		mu.Unlock()
		printTrace(gid, name, "<-", indents)
	}
}
