package avl

import "cmp"

// AVL Tree.
type Tree[K cmp.Ordered, V any] struct {
	root *node[K, V]
}

// New creates empty AVL tree.
func New[K cmp.Ordered, V any]() (rv *Tree[K, V]) {
	return &Tree[K, V]{}
}

// Add places key and value into tree, if key exists - its value will be replaced.
func (t *Tree[K, V]) Add(key K, value V) {
	t.root = t.root.add(key, value)
}

// Del removes key from tree.
func (t *Tree[K, V]) Del(key K) {
	t.root = t.root.del(key)
}

// Get obtains value for specified key.
func (t *Tree[K, V]) Get(key K) (rv V, ok bool) {
	if node, ok := t.root.find(key); ok {
		return node.val, true
	}

	return
}

// Has checks for key existence in tree.
func (t *Tree[K, V]) Has(key K) (ok bool) {
	_, ok = t.root.find(key)

	return
}

// Iter implements range-iter interface.
func (t *Tree[K, V]) Iter(cb func(K, V) bool) {
	t.root.iterate(cb)
}
