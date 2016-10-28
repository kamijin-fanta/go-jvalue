package jvalue

import "testing"

var testJson = `
[{
    "arr": [1,2,3],
    "obj": {
        "a":1,
        "b":2.34,
        "c":"foo"
    }
}]`

func TestDecode(t *testing.T) {
	_, err := DecodeJSONString(testJson)
	if err != nil {
		t.Error("Decode Error", err)
	}
}

func TestAccess(t *testing.T) {
	jvalue, _ := DecodeJSONString(testJson)

	{
		obj, err := jvalue.Index(0).Key("obj").Key("a").ToInt()
		if err != nil {
			t.Error("Int cast error", obj)
		} else {
			if *obj != 1 {
				t.Error("must root[0].obj.a is 1", *obj)
			}
		}
	}
	{
		obj, err := jvalue.Index(0).Key("obj").Key("b").ToInt()
		if err != nil {
			t.Error("Float to Int cast error", obj)
		} else {
			if *obj != 2 {
				t.Error("must root[0].obj.b is 2", *obj)
			}
		}
	}
	{
		obj, err := jvalue.Index(0).Key("obj").Key("c").ToString()
		if err != nil {
			t.Error("Float to Int cast error", obj)
		} else {
			if *obj != "foo" {
				t.Error("must root[0].obj.c is 'foo'", *obj)
			}
		}
	}
}

func TestSpeed(t *testing.T) {
	jvalue, _ := DecodeJSONString(testJson)

	for i := 0; i < 10000000; i++ {
		if i%1000000 == 0 {
			t.Log("loop ...", i)
		}
		obj, _ := jvalue.Index(0).Key("obj").Key("a").ToInt()
		if *obj != 1 {
			t.Error("parse error")
			break
		}
	}
}
