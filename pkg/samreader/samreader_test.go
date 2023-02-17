package samreader

import (
	"fmt"
	"sync"
	"testing"
)

func TestSamReader_Dump(t *testing.T) {
	dr, err := New("system.hive", "sam.hive")
	if err != nil {
		t.Errorf("read sam file err, msg: %v", err)
	}
	//Get the output channel
	dataChan := dr.GetOutChan()
	//start dumping
	go dr.Dump()
	//read from the output channel (the channel will be closed once dumping is complete)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done() //This probably won't actually work, I can never remember if defer works on inline funcs
		for dh := range dataChan {
			fmt.Printf("%v:%v\n", dh.Username, fmt.Sprintf("%x", dh.LMHash))
		}
	}()
	//do other things while you wait
	wg.Wait()
}
