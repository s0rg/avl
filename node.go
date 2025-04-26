package avl

import "cmp"

const (
	cmpBigger = 1
	cmpLesser = -1
)

type node[K cmp.Ordered, V any] struct {
	key    K
	val    V
	left   *node[K, V]
	right  *node[K, V]
	height int
}

func (n *node[K, V]) add(key K, val V) (rv *node[K, V], changed bool) {
	if n == nil {
		return &node[K, V]{
			key:    key,
			val:    val,
			height: 1,
		}, true
	}

	switch cmp.Compare(n.key, key) {
	case cmpBigger:
		n.left, changed = n.left.add(key, val)
	case cmpLesser:
		n.right, changed = n.right.add(key, val)
	default:
		n.val = val

		return n, false
	}

	tmp := changed

	rv, changed = n.rebalance()

	return rv, changed || tmp
}

func (n *node[K, V]) del(key K) (rv *node[K, V], changed bool) {
	if n == nil {
		return nil, false
	}

	switch cmp.Compare(n.key, key) {
	case cmpBigger:
		n.left, changed = n.left.del(key)
	case cmpLesser:
		n.right, changed = n.right.del(key)
	default:
		switch {
		case n.left != nil && n.right != nil:
			m := n.right.findSmallest()
			n.key, n.val = m.key, m.val
			n.right, changed = n.right.del(m.key)
		case n.left != nil:
			n = n.left
		case n.right != nil:
			n = n.right
		default:
			return nil, false
		}
	}

	tmp := changed
	rv, changed = n.rebalance()

	return rv, changed || tmp
}

func (n *node[K, V]) find(key K) (rv *node[K, V], found bool) {
	if n == nil {
		return nil, false
	}

	switch cmp.Compare(n.key, key) {
	case cmpBigger:
		return n.left.find(key)
	case cmpLesser:
		return n.right.find(key)
	}

	return n, true
}

func (n *node[K, V]) iterate(cb func(K, V) bool) {
	if n == nil {
		return
	}

	if n.left != nil {
		n.left.iterate(cb)
	}

	if !cb(n.key, n.val) {
		return
	}

	if n.right != nil {
		n.right.iterate(cb)
	}
}

func (n *node[K, V]) rotateLeft() (rv *node[K, V]) {
	rv = n.right
	n.right = rv.left
	rv.left = n

	n.updateHeight()
	rv.updateHeight()

	return rv
}

func (n *node[K, V]) rotateRight() (rv *node[K, V]) {
	rv = n.left
	n.left = rv.right
	rv.right = n

	n.updateHeight()
	rv.updateHeight()

	return rv
}

func (n *node[K, V]) getHeight() (rv int) {
	if n == nil {
		return
	}

	return n.height
}

func (n *node[K, V]) rebalance() (rv *node[K, V], changed bool) {
	n.updateHeight()

	const (
		heavyLeft  = 2
		heavyRight = -2
	)

	switch n.left.getHeight() - n.right.getHeight() {
	case heavyLeft:
		if n.isRightHeavy() {
			n.left = n.left.rotateLeft()
		}

		return n.rotateRight(), true
	case heavyRight:
		if n.isLeftHeavy() {
			n.right = n.right.rotateRight()
		}

		return n.rotateLeft(), true
	}

	return n, false
}

func (n *node[K, V]) findSmallest() (rv *node[K, V]) {
	if n.left != nil {
		return n.left.findSmallest()
	}

	return n
}

func (n *node[K, V]) updateHeight() {
	n.height = 1 + max(n.left.getHeight(), n.right.getHeight())
}

func (n *node[K, V]) isLeftHeavy() (yes bool) {
	t := n.right

	return t.left.getHeight() > t.right.getHeight()
}

func (n *node[K, V]) isRightHeavy() (yes bool) {
	t := n.left

	return t.right.getHeight() > t.left.getHeight()
}
