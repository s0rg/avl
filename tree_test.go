package avl_test

import (
	"math/rand"
	"testing"

	"github.com/s0rg/avl"
)

const (
	opAdd = iota
	opDel
	opGet
)

func TestTreeMany(t *testing.T) {
	t.Parallel()

	tree := avl.New[int, int]()
	m := make(map[int]int)

	const (
		maxOps = 3
		maxKey = 1_000
		nOps   = 10_000
	)

	for range nOps {
		k := rand.Intn(maxKey)

		switch rand.Intn(maxOps) {
		case opAdd:
			v := rand.Int()
			tree.Add(k, v)
			m[k] = v
		case opDel:
			tree.Del(k)
			delete(m, k)
		case opGet:
			var tv int

			if val, ok := tree.Get(k); ok {
				tv = val
			} else {
				continue
			}

			mv := m[k]
			if tv != mv {
				t.Errorf("[-] key: %d want: %d got: %d", k, mv, tv)
			}
		}
	}
}

func TestTreeCRUD(t *testing.T) {
	t.Parallel()

	tree := avl.New[int, string]()

	for k := range tree.Iter {
		_ = k
	}

	tree.Add(4, "four")
	tree.Add(2, "two")
	tree.Add(5, "five")
	tree.Add(1, "one")
	tree.Add(3, "three")

	if tree.Has(6) {
		t.Fail()
	}

	if !tree.Has(2) {
		t.Fail()
	}

	tree.Del(2)
	tree.Del(6)

	if tree.Has(2) {
		t.Fail()
	}

	tree.Add(6, "six")

	if !tree.Has(6) {
		t.Fail()
	}

	for range tree.Iter {
	}

	for k := range tree.Iter {
		if k == 5 {
			break
		}
	}
}
