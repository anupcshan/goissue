Reproduce Go 1.20 compiler crash
--------------------------------

```
$ go build
./dag/dag.go:12:2: internal compiler error: '(*DAG[go.shape.int]).RemoveNodeFromParent.func1': value dag..dict (nil) incorrectly live at entry

Please file a bug report including a short program that triggers the error.
https://go.dev/issue/new
```
