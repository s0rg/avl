package avl

import "cmp"

const (
	cBigger = 1
	cLesser = -1
)

type node[K cmp.Ordered, V any] struct {
	key    K
	val    V
	left   *node[K, V]
	right  *node[K, V]
	height int
}

func (n *node[K, V]) add(key K, val V) (rv *node[K, V]) {
	if n == nil {
		return &node[K, V]{
			key:    key,
			val:    val,
			height: 1,
		}
	}

	switch cmp.Compare(n.key, key) {
	case cBigger:
		n.left = n.left.add(key, val)
	case cLesser:
		n.right = n.right.add(key, val)
	default:
		n.val = val

		return
	}

	return n.rebalance()
}

func (n *node[K, V]) del(key K) (rv *node[K, V]) {
	if n == nil {
		return nil
	}

	switch cmp.Compare(n.key, key) {
	case cBigger:
		n.left = n.left.del(key)
	case cLesser:
		n.right = n.right.del(key)
	default:
		switch {
		case n.left != nil && n.right != nil:
			m := n.right.findSmallest()
			n.key, n.val = m.key, m.val
			n.right = n.right.del(m.key)
		case n.left != nil:
			n = n.left
		case n.right != nil:
			n = n.right
		default:
			return nil
		}
	}

	return n.rebalance()
}

func (n *node[K, V]) find(key K) (rv *node[K, V], found bool) {
	if n == nil {
		return nil, false
	}

	switch cmp.Compare(n.key, key) {
	case cBigger:
		return n.left.find(key)
	case cLesser:
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

func (n *node[K, V]) rebalance() (rv *node[K, V]) {
	n.updateHeight()

	switch b := n.left.getHeight() - n.right.getHeight(); b {
	case -2:
		if n.isLeftHeavy() {
			n.right = n.right.rotateRight()
		}

		return n.rotateLeft()
	case 2:
		if n.isRightHeavy() {
			n.left = n.left.rotateLeft()
		}

		return n.rotateRight()
	}

	return n
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
