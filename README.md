# go-jvalue

```go
var testJson = `
[{
    "arr": [1,2,3],
    "obj": {
        "a":1,
        "b":2.34,
        "c":"foo"
    }
}]`

value, _ := jvalue.DecodeJSONString(testJson)
obj, _ := value.Index(0).Key("obj").Key("a").ToInt()
fmt.Println("root[0].obj.a", *obj)
```

## cast

- ToString
- ToInt
- ToBool

# api

- IsArray() bool
- HasIndex(index int) bool
- Index(index int) JValue
- IsMap() bool
- HasKey(key string) bool
- Key(index int) JValue
