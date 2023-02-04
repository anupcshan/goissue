package main

func main() {
	_ = DAG[int]{}
}

type DAG[T comparable] struct {
	inverse map[T]T
}

func (g *DAG[T]) RemoveNodeFromParent(node T) {
	defer delete(g.inverse, node)
	// To fix, remove "defer" from the above line, like below.
	// delete(g.inverse, node)
}
