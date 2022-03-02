/**
 * @date: 2022/2/20
 * @desc: ...
 */

package utils

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

//
// BuildOrderId
// @Description: 获取订单号,固定24位长度订单号，毫秒+进程id+序号,同一毫秒内只要不超过一万次并发，则订单号不会重复
// @param t:
// @return string:
//
func BuildOrderId(t time.Time) string {
	s := t.Format("20060102150405")
	m := t.UnixNano()/1e6 - t.UnixNano()/1e9*1e3
	ms := sup(m, 3)
	p := os.Getpid() % 1000
	ps := sup(int64(p), 3)
	var num int64
	i := atomic.AddInt64(&num, 1)
	r := i % 10000
	rs := sup(r, 4)
	n := fmt.Sprintf("%s%s%s%s", s, ms, ps, rs)
	return n
}

//对长度不足n的数字前面补0
func sup(i int64, n int) string {
	m := fmt.Sprintf("%d", i)
	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}
