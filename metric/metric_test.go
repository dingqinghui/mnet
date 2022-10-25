/**
 * @Author: dingQingHui
 * @Description:
 * @File: metric_test
 * @Version: 1.0.0
 * @Date: 2022/10/10 9:59
 */

package metric

import (
	metrics "github.com/rcrowley/go-metrics"
	"log"
	"os"
	"testing"
	"time"
)

func TestNewPropsWithFunc(t *testing.T) {
	m := metrics.NewMeter()
	metrics.Register("quux", m)
	m.Mark(1)
	log := log.New(os.Stdout, "metrics: ", log.Lmicroseconds)
	go metrics.Log(metrics.DefaultRegistry,
		1*time.Second,
		log,
	)
	var j int64
	j = 1
	for true {
		time.Sleep(time.Second * 1)
		j++
		m.Mark(10)
	}
}
