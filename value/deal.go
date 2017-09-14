package value

import (
	"math/big"
	"math/rand"
)

type shuffleMap map[Int]Int

// Swap the number in i-th and subIdx-th position, and return the i-th position value.
func (s shuffleMap) Swap(i, subIdx Int) Int {
	// number in s[subIdx] will never be selected, so it does not need to store in map
	if i == subIdx {
		if v, ok := s[subIdx]; ok {
			delete(s, subIdx)
			return v
		}
		return subIdx
	}

	v := i
	if sv, ok := s[i]; ok {
		v = sv
	}
	sub := subIdx
	if sv, ok := s[subIdx]; ok {
		sub = sv
		delete(s, subIdx)
	}
	s[i] = sub
	return v
}

// deal is a variant of Fisher–Yates shuffle, which it does not prepare a slice
// of all available items, but use a map to store the swapped numbers.
//
// deal will assign distinct random numbers between [origin, origin+count) into b
func deal(b []Value, origin, count Int, rn *rand.Rand) {
	// m stores the swapped numbers
	m := make(shuffleMap, len(b))
	for i := range b {
		r := Int(rn.Int63n(int64(count - Int(i))))
		last := count - Int(i) - 1

		v := m.Swap(r, last)
		b[i] = v + origin
	}
}

type bigIntShuffleMap map[string]*big.Int

// Swap the number in i-th and subIdx-th position, and return the i-th position value.
func (s bigIntShuffleMap) Swap(i, subIdx *big.Int) *big.Int {
	skey := string(subIdx.Bytes())
	// number in s[subIdx] will never be selected, so it does not need to store in map
	if i.Cmp(subIdx) == 0 {
		if v, ok := s[skey]; ok {
			delete(s, skey)
			return v
		}
		return subIdx
	}

	ikey := string(i.Bytes())
	v := i
	if sv, ok := s[ikey]; ok {
		v = sv
	}
	sub := subIdx
	if sv, ok := s[skey]; ok {
		sub = sv
		delete(s, skey)
	}
	s[ikey] = sub
	return v
}

// bigIntDeal is a variant of Fisher–Yates shuffle, which it does not prepare a slice
// of all available items, but use a map to store the swapped numbers.
//
// bigIntDeal will assign distinct random numbers between [origin, origin+count) into b
func bigIntDeal(b []Value, origin, count BigInt, rn *rand.Rand) {
	// m stores the swapped numbers
	m := make(bigIntShuffleMap, len(b))
	r := new(big.Int)

	rmax := new(big.Int)
	last := new(big.Int)
	for i := range b {
		rmax.SetInt64(int64(i))
		r.Rand(rn, rmax.Sub(count.Int, rmax))
		last.Sub(rmax, bigOne.Int)

		v := m.Swap(r, last)
		b[i] = BigInt{new(big.Int).Add(v, origin.Int)}
	}
}
