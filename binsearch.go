package binsearch

import (
"sort"
)

// ---------- Key_uint64 ----------

// Add this to any struct to make it binary searchable.
type Key_uint64 struct {
Key []uint64
}

type sort_uint64 struct {
i int
k uint64
}
type sorter_uint64 []keyVal_uint
func (a *sorter_uint64) Len() int           { return len(a) }
func (a *sorter_uint64) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a *sorter_uint64) Less(i, j int) bool { return a[i].k < a[j].k }

// Find returns the index based on the key.
func (f *Key_uint64) Find(thekey uint64) (int, bool) {
	min := 0
	max := len(f.Key)-1
	at := max/2
	for {
		current := f.Key[at]
		if thekey<current {
			max = at-1
		} else {
		if thekey>current {
			min = at+1
			} else {
				return at, true // found
			}
		}
		if min>max {
			return min, false // doesn't exist
		}
		at = (max+min)/2
	}
}

// Add adds this index for later building
func (f *Key_uint64) Add(thekey uint64) {
	f.Key = append(f.Key, thekey)
	return
}

// Build sorts the keys and returns an array telling you how to sort the values, you must do this yourself.
func (f *Key_uint64) Build() []int {
	l := len(f.Key)
	temp := make(sorter_uint64,l)
	var current uint
	for i,k := range f.Key {
		temp[current]=sort_uint64{i,k}
		current++
	}
	sort.Sort(temp)
	imap := make([]int,l)
	newkey := make([]int64,l)
	for i:=0; i<l; i++ {
		imap[i]=temp[i].i
		newkey[i]=temp[i].k
	}
	f.Key = newkey
	return imap
}