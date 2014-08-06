package binsearch

import (
"sort"
"bytes"
)

// ---------- Key_uint64 ----------

// Add this to any struct to make it binary searchable.
type Key_uint64 struct {
Key []uint64
keymax uint64
}

type sort_uint64 struct {
i int
k uint64
}
type sorter_uint64 []sort_uint64
func (a sorter_uint64) Len() int           { return len(a) }
func (a sorter_uint64) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sorter_uint64) Less(i, j int) bool { return a[i].k < a[j].k }

// Find returns the index based on the key.
func (f *Key_uint64) Find(thekey uint64) (uint64, bool) {
	var l uint64 = 0
	r := f.keymax
	tot := r
	mid := r/2
	for l<=r && mid>=0 && mid<=tot  {
		if (thekey < f.Key[mid]) {
			r = mid - 1;
		} else if (thekey > f.Key[mid]) {
			l = mid + 1;
		} else {
			return mid, true
		}
	mid = l+((r-l)/2)
	}
	return 0, false
}

// Add adds this index for later building
func (f *Key_uint64) AddKey(thekey uint64) {
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
	newkey := make([]uint64,l)
	for i:=0; i<l; i++ {
		imap[i]=temp[i].i
		newkey[i]=temp[i].k
	}
	f.Key = newkey
	f.keymax = uint64(l-1)
	return imap
}


// ---------- Key_uint32 ----------

// Add this to any struct to make it binary searchable.
type Key_uint32 struct {
Key []uint32
keymax uint64
keydistribution []uint32
}

type sort_uint32 struct {
i int
k uint32
}
type sorter_uint32 []sort_uint32
func (a sorter_uint32) Len() int           { return len(a) }
func (a sorter_uint32) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sorter_uint32) Less(i, j int) bool { return a[i].k < a[j].k }

// Find returns the index based on the key.
func (f *Key_uint32) Find(thekey uint32) (uint64, bool) {
	var l uint64
	r := f.keymax
	var mid uint64 = uint64(((float32(thekey - f.Key[l])/float32(f.Key[r] - f.Key[l]))*float32(r))+0.5)
	for l<=r && mid>=l && mid<=r {
		if (thekey < f.Key[mid]) {
			r = mid - 1;
		} else if (thekey > f.Key[mid]) {
			l = mid + 1;
		} else {
			return mid, true
		}
	mid = l + uint64(((float32(thekey - f.Key[l])/float32(f.Key[r] - f.Key[l]))*float32(r - l))+0.5) // +0.5 makes it round instead of floor
	}
	return 0, false
}

// Add adds this index for later building
func (f *Key_uint32) AddKey(thekey uint32) {
	f.Key = append(f.Key, thekey)
	return
}

// Build sorts the keys and returns an array telling you how to sort the values, you must do this yourself.
func (f *Key_uint32) Build() []int {
	l := len(f.Key)
	temp := make(sorter_uint32,l)
	var current uint
	for i,k := range f.Key {
		temp[current]=sort_uint32{i,k}
		current++
	}
	sort.Sort(temp)
	imap := make([]int,l)
	newkey := make([]uint32,l)
	for i:=0; i<l; i++ {
		imap[i]=temp[i].i
		newkey[i]=temp[i].k
	}
	f.Key = newkey
	f.keymax = uint64(l-1)
	return imap
}


// ---------- Key_uint16 ----------

// Add this to any struct to make it binary searchable.
type Key_uint16 struct {
Key []uint16
keymax uint64
}

type sort_uint16 struct {
i int
k uint16
}
type sorter_uint16 []sort_uint16
func (a sorter_uint16) Len() int           { return len(a) }
func (a sorter_uint16) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sorter_uint16) Less(i, j int) bool { return a[i].k < a[j].k }

// Find returns the index based on the key.
func (f *Key_uint16) Find(thekey uint16) (uint64, bool) {
	var min uint64
	max := f.keymax
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
func (f *Key_uint16) AddKey(thekey uint16) {
	f.Key = append(f.Key, thekey)
	return
}

// Build sorts the keys and returns an array telling you how to sort the values, you must do this yourself.
func (f *Key_uint16) Build() []int {
	l := len(f.Key)
	temp := make(sorter_uint16,l)
	var current uint
	for i,k := range f.Key {
		temp[current]=sort_uint16{i,k}
		current++
	}
	sort.Sort(temp)
	imap := make([]int,l)
	newkey := make([]uint16,l)
	for i:=0; i<l; i++ {
		imap[i]=temp[i].i
		newkey[i]=temp[i].k
	}
	f.Key = newkey
	f.keymax = uint64(l-1)
	return imap
}


// ---------- Key_uint8 ----------

// Add this to any struct to make it binary searchable.
type Key_uint8 struct {
Key []uint8
keymax uint64
}

type sort_uint8 struct {
i int
k uint8
}
type sorter_uint8 []sort_uint8
func (a sorter_uint8) Len() int           { return len(a) }
func (a sorter_uint8) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sorter_uint8) Less(i, j int) bool { return a[i].k < a[j].k }

// Find returns the index based on the key.
func (f *Key_uint8) Find(thekey uint8) (uint64, bool) {
	var min uint64
	max := f.keymax
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
func (f *Key_uint8) AddKey(thekey uint8) {
	f.Key = append(f.Key, thekey)
	return
}

// Build sorts the keys and returns an array telling you how to sort the values, you must do this yourself.
func (f *Key_uint8) Build() []int {
	l := len(f.Key)
	temp := make(sorter_uint8,l)
	var current uint
	for i,k := range f.Key {
		temp[current]=sort_uint8{i,k}
		current++
	}
	sort.Sort(temp)
	imap := make([]int,l)
	newkey := make([]uint8,l)
	for i:=0; i<l; i++ {
		imap[i]=temp[i].i
		newkey[i]=temp[i].k
	}
	f.Key = newkey
	f.keymax = uint64(l-1)
	return imap
}


// ---------- Key_string ----------

// Add this to any struct to make it binary searchable.
type Key_string struct {
Key []string
keymax uint64
}

type sort_string struct {
i int
k string
}
type sorter_string []sort_string
func (a sorter_string) Len() int           { return len(a) }
func (a sorter_string) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sorter_string) Less(i, j int) bool { return a[i].k < a[j].k }

// Find returns the index based on the key.
func (f *Key_string) Find(thekey string) (uint64, bool) {
	var min uint64
	max := f.keymax
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
		at = min+((max-min)/2) // so as not to go out of bounds
	}
}

// Add adds this index for later building
func (f *Key_string) AddKey(thekey string) {
	f.Key = append(f.Key, thekey)
	return
}

// Build sorts the keys and returns an array telling you how to sort the values, you must do this yourself.
func (f *Key_string) Build() []int {
	l := len(f.Key)
	temp := make(sorter_string,l)
	var current uint
	for i,k := range f.Key {
		temp[current]=sort_string{i,k}
		current++
	}
	sort.Sort(temp)
	imap := make([]int,l)
	newkey := make([]string,l)
	for i:=0; i<l; i++ {
		imap[i]=temp[i].i
		newkey[i]=temp[i].k
	}
	f.Key = newkey
	f.keymax = uint64(l-1)
	return imap
}


// ---------- Key_byte ----------

// Add this to any struct to make it binary searchable.
type Key_bytes struct {
Key [][]byte
keymax uint64
}

type sort_bytes struct {
i int
k []byte
}
type sorter_bytes []sort_bytes
func (a sorter_bytes) Len() int           { return len(a) }
func (a sorter_bytes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sorter_bytes) Less(i, j int) bool { return bytes.Compare(a[i].k,a[j].k)==-1 }

// Find returns the index based on the key.
func (f *Key_bytes) Find(thekey []byte) (uint64, bool) {
	var min uint64
	max := f.keymax
	at := max/2
	for {
		what := bytes.Compare(thekey,f.Key[at])
		switch(what) {
			case -1: max = at-1
			case 1: min = at+1
			default: return at, true
		}
		if min>max {
			return min, false // doesn't exist
		}
		at = min+((max-min)/2) // so as not to go out of bounds
	}
}

// Add adds this index for later building
func (f *Key_bytes) AddKey(thekey []byte) {
	f.Key = append(f.Key, thekey)
	return
}

// Build sorts the keys and returns an array telling you how to sort the values, you must do this yourself.
func (f *Key_bytes) Build() []int {
	l := len(f.Key)
	temp := make(sorter_bytes,l)
	var current uint
	for i,k := range f.Key {
		temp[current]=sort_bytes{i,k}
		current++
	}
	sort.Sort(temp)
	imap := make([]int,l)
	newkey := make([][]byte,l)
	for i:=0; i<l; i++ {
		imap[i]=temp[i].i
		newkey[i]=temp[i].k
	}
	f.Key = newkey
	f.keymax = uint64(l-1)
	return imap
}