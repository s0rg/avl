package avl

import "cmp"

// AVL Tree.
type Tree[K cmp.Ordered, V any] struct {
	root  *node[K, V]
	count int
}

// New creates empty AVL tree.
func New[K cmp.Ordered, V any]() (rv *Tree[K, V]) {
	return &Tree[K, V]{}
}

// Add places key and value into tree, if key exists - its value will be replaced.
func (t *Tree[K, V]) Add(key K, value V) {
	var ok bool

	if t.root, ok = t.root.add(key, value); ok {
		t.count++
	}
}

// Del removes key from tree.
func (t *Tree[K, V]) Del(key K) {
	var ok bool

	if t.root, ok = t.root.del(key); ok {
		t.count--
	}
}

// Get obtains value for specified key.
func (t *Tree[K, V]) Get(key K) (rv V, ok bool) {
	if node, found := t.root.find(key); found {
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

// Clear drops tree contents, by resetting root.
func (t *Tree[K, V]) Clear() {
	t.root, t.count = nil, 0
}

// Len returns number of elements in tree.
func (t *Tree[K, V]) Len() (rv int) {
	return t.count
}
