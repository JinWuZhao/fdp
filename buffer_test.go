package fdp

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func Test_pktBuffer_Push(t *testing.T) {
	capacity := int32(500)
	buff := newPktBuffer(capacity)
	wg := new(sync.WaitGroup)
	threadCount := 10
	wg.Add(threadCount)
	for index := 0; index < threadCount; index++ {
		i := index
		go func() {
			for j := 0; j < 100; j++ {
				p := newDataPacket()
				p.ID = []byte(fmt.Sprintf("%d_%d", i, j))
				if err := buff.Push(p); err != nil {
					//t.Logf("failed to push into buff, %s", err.Error())
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	sortedBuff := make([]packet, 1000)
	for index := int32(0); index < 1000; index++ {
		p, err := buff.Get(int32(index))
		if err != nil {
			//t.Logf("failed to get from buff, %s", err.Error())
			continue
		}
		pkt := p.(*dataPacket)
		coms := strings.Split(string(pkt.ID), "_")
		id1, _ := strconv.Atoi(coms[0])
		id2, _ := strconv.Atoi(coms[1])
		sortedBuff[id1*100+id2] = pkt
	}
	for _, v := range sortedBuff {
		if v != nil {
			t.Log(v.String())
		} else {
			t.Log("empty")
		}
	}
}

func Test_pktBuffer_Pop(t *testing.T) {
	capacity := int32(500)
	buff := newPktBuffer(capacity)
	for index := int32(0); index < capacity; index++ {
		p := newDataPacket()
		p.ID = []byte(fmt.Sprintf("%d_%d", index/100, index%100))
		if err := buff.Push(p); err != nil {
			t.Logf("failed to push into buff, %s", err.Error())
			continue
		}

	}
	wg := new(sync.WaitGroup)
	threadCount := 10
	wg.Add(threadCount)
	for index := 0; index < threadCount; index++ {
		go func() {
			for j := 0; j < 100; j++ {
				pkt, err := buff.Pop()
				if err != nil {
					//t.Logf("failed to pop from buff, %s", err.Error())
					continue
				}
				t.Logf(pkt.String())
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func Test_pktBuffer_PushAndPop(t *testing.T) {
	capacity := int32(500)
	buff := newPktBuffer(capacity)
	wg := new(sync.WaitGroup)
	threadCount := 10
	wg.Add(threadCount * 100 * 2)
	for index := 0; index < threadCount; index++ {
		tIndex := index
		go func() {
			for i := 0; i < 100; i++ {
				ti := i
				go func() {
					defer wg.Done()
					p := newDataPacket()
					p.ID = []byte(fmt.Sprintf("%d_%d", tIndex, ti))
					if err := buff.Push(p); err != nil {
						t.Logf("failed to push to buff, %s", err.Error())
					}
				}()
				go func() {
					defer wg.Done()
					p, err := buff.Pop()
					if err != nil {
						t.Logf("failed to pop from buff, %s", err.Error())
						return
					}
					t.Logf(p.String())
				}()
			}
		}()
	}
	wg.Wait()
}
