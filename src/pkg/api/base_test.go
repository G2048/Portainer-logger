package api

import (
	"fmt"
	"testing"
)

const URL = `http://httpbin.org/`

func TestHttpApi(t *testing.T) {
	client := NewHttpApi(URL)
	res, err := client.Get("/get", nil)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if res.Status() != 200 {
		t.Fail()
	}

	fmt.Printf("%#+v\n", res.Body())

	// 200
	res, err = client.Get("/status/200", nil)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if res.Status() != 200 {
		t.Fail()
	}
	fmt.Printf("%#+v\n", res.Body())

	// 500
	res, err = client.Get("/status/500", nil)
	t.Log(err)
	if err == nil {
		t.Errorf("Error by request to /status/500: %s", err.Error())
	}
	if res.Status() != 500 {
		t.Error(err)
	}

	//400
	res, err = client.Get("/status/400", nil)
	if err == nil {
		t.Error(err)
	}
	if res.Status() != 400 {
		t.Error(err)
	}
}
