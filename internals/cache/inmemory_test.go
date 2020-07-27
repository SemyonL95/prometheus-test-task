package cache

import (
	"sync"
	"testing"
)

var ipAdresses = []string{
	"0.0.0.0",
	"1.1.1.1",
	"2.2.2.2",
}

func Test_SetError(t *testing.T) {
	c := New()
	err := c.Set(ipAdresses[0])
	if err != nil {
		t.Log("Error insert first value")
		t.Fail()
	}
	err = c.Set(ipAdresses[0])
	if err == nil {
		t.Log("Expect error, instead of nil")
		t.Fail()
	}

	if len(c.data) != 1 {
		t.Log("Wrong len of data")
		t.Fail()
	}
}

func Test_Set(t *testing.T) {
	c := New()

	for _, val := range ipAdresses {
		err := c.Set(val)
		if err != nil {
			t.Logf("Error insert value: %s", val)
			t.Fail()
		}
	}

	if len(c.data) != len(ipAdresses) {
		t.Log("Wrong len of data")
		t.Fail()
	}
}

func Test_SetConcurently(t *testing.T) {
	c := New()

	wg := sync.WaitGroup{}
	for _, val := range ipAdresses {
		wg.Add(1)
		go func(val string) {
			err := c.Set(val)
			if err != nil {
				t.Logf("Error insert value: %s", val)
				t.Fail()
			}
			wg.Done()
		}(val)
	}

	wg.Wait()
	if len(c.data) != len(ipAdresses) {
		t.Log("Wrong len of data")
		t.Fail()
	}
}