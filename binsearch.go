package binsearch

import (
"sort"
"fmt"
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
type sorter_uint64 []sort_uint64
func (a sorter_uint64) Len() int           { return len(a) }
func (a sorter_uint64) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sorter_uint64) Less(i, j int) bool { return a[i].k < a[j].k }

// Find returns the index based on the key.
func (f *Key_uint64) Find(thekey uint64) (uint64, bool) {
	var min,at uint64
	var current uint64
	max := uint64(len(f.Key))
	if max>0 {
		max--
	} else {
		return 0, false
	}
	for min<=max {
		at = min+((max-min)/2)
		if current=f.Key[at]; thekey<current {
			if at==0 {
				return 0, false
			}
			max = at-1
		} else {
		if thekey>current {
			min = at+1
			} else {
				return at, true // found
			}
		}
	}
	return min, false // doesn't exist
}

// Find returns the index based on the key, using Interpolation.
func (f *Key_uint64) FindInterpolation(thekey uint64) (uint64, bool) {
	var l uint64
	r := uint64(len(f.Key))
	if r>0 {
		r--
	} else {
		return 0, false
	}
	var mid uint64 = uint64(((float32(thekey - f.Key[l])/float32(f.Key[r] - f.Key[l]))*float32(r))+0.5) // +0.5 makes it round instead of floor
	for l<=r && mid>=l && mid<=r {
		if (thekey < f.Key[mid]) {
			r = mid - 1;
		} else if (thekey > f.Key[mid]) {
			l = mid + 1;
		} else {
			return mid, true
		}
	mid = l + uint64(((float32(thekey - f.Key[l])/float32(f.Key[r] - f.Key[l]))*float32(r - l))+0.5)
	}
	return 0, false
}

// AddKeyUnsorted adds this key to the end of the index for later building with Build.
func (f *Key_uint64) AddKeyUnsorted(thekey uint64) {
	f.Key = append(f.Key, thekey)
	return
}

// AddKeyAt adds this key to the index in this exact position, so it does not require later rebuilding.
func (f *Key_uint64) AddKeyAt(thekey uint64, i uint64) {
	f.Key = append(f.Key, 0)
	copy(f.Key[i+1:], f.Key[i:])
	f.Key[i] = thekey
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
	return imap
}


// ---------- Key_uint32 ----------

// Add this to any struct to make it binary searchable.
type Key_uint32 struct {
Key []uint32
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
	var min,at uint64
	var current uint32
	max := uint64(len(f.Key))
	if max>0 {
		max--
	} else {
		return 0, false
	}
	for min<=max {
		at = (max+min)/2
		if current=f.Key[at]; thekey<current {
			if at==0 {
				return 0, false
			}
			max = at-1
		} else {
		if thekey>current {
			min = at+1
			} else {
				return at, true // found
			}
		}
	}
	return min, false // doesn't exist
}

// Find returns the index based on the key, using Interpolation.
func (f *Key_uint32) FindInterpolation(thekey uint32) (uint64, bool) {
	var l uint64
	r := uint64(len(f.Key))
	if r>0 {
		r--
	} else {
		return 0, false
	}
	var mid uint64 = uint64(((float32(thekey - f.Key[l])/float32(f.Key[r] - f.Key[l]))*float32(r))+0.5) // +0.5 makes it round instead of floor
	for l<=r && mid>=l && mid<=r {
		if (thekey < f.Key[mid]) {
			r = mid - 1;
		} else if (thekey > f.Key[mid]) {
			l = mid + 1;
		} else {
			return mid, true
		}
	mid = l + uint64(((float32(thekey - f.Key[l])/float32(f.Key[r] - f.Key[l]))*float32(r - l))+0.5)
	}
	return 0, false
}

// AddKeyUnsorted adds this key to the end of the index for later building with Build.
func (f *Key_uint32) AddKeyUnsorted(thekey uint32) {
	f.Key = append(f.Key, thekey)
	return
}

// AddKeyAt adds this key to the index in this exact position, so it does not require later rebuilding.
func (f *Key_uint32) AddKeyAt(thekey uint32, i uint64) {
	f.Key = append(f.Key, 0)
	copy(f.Key[i+1:], f.Key[i:])
	f.Key[i] = thekey
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
	return imap
}


// ---------- Key_uint16 ----------

// Add this to any struct to make it binary searchable.
type Key_uint16 struct {
Key []uint16
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
	var min,at uint64
	var current uint16
	max := uint64(len(f.Key))
	if max>0 {
		max--
	} else {
		return 0, false
	}
	for min<=max {
		at = (max+min)/2
		if current=f.Key[at]; thekey<current {
			if at==0 {
				return 0, false
			}
			max = at-1
		} else {
		if thekey>current {
			min = at+1
			} else {
				return at, true // found
			}
		}
	}
	return min, false // doesn't exist
}

// Find returns the index based on the key, using Interpolation.
func (f *Key_uint16) FindInterpolation(thekey uint16) (uint64, bool) {
	var l uint64
	r := uint64(len(f.Key))
	if r>0 {
		r--
	} else {
		return 0, false
	}
	var mid uint64 = uint64(((float32(thekey - f.Key[l])/float32(f.Key[r] - f.Key[l]))*float32(r))+0.5) // +0.5 makes it round instead of floor
	for l<=r && mid>=l && mid<=r {
		if (thekey < f.Key[mid]) {
			r = mid - 1;
		} else if (thekey > f.Key[mid]) {
			l = mid + 1;
		} else {
			return mid, true
		}
	mid = l + uint64(((float32(thekey - f.Key[l])/float32(f.Key[r] - f.Key[l]))*float32(r - l))+0.5)
	}
	return 0, false
}

// AddKeyUnsorted adds this key to the end of the index for later building with Build.
func (f *Key_uint16) AddKeyUnsorted(thekey uint16) {
	f.Key = append(f.Key, thekey)
	return
}

// AddKeyAt adds this key to the index in this exact position, so it does not require later rebuilding.
func (f *Key_uint16) AddKeyAt(thekey uint16, i uint64) {
	f.Key = append(f.Key, 0)
	copy(f.Key[i+1:], f.Key[i:])
	f.Key[i] = thekey
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
	return imap
}


// ---------- Key_uint8 ----------

// Add this to any struct to make it binary searchable.
type Key_uint8 struct {
Key []uint8
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
	var min,at uint64
	var current uint8
	max := uint64(len(f.Key))
	if max>0 {
		max--
	} else {
		return 0, false
	}
	for min<=max {
		at = (max+min)/2
		if current=f.Key[at]; thekey<current {
			if at==0 {
				return 0, false
			}
			max = at-1
		} else {
		if thekey>current {
			min = at+1
			} else {
				return at, true // found
			}
		}
	}
	return min, false // doesn't exist
}

// Find returns the index based on the key, using Interpolation.
func (f *Key_uint8) FindInterpolation(thekey uint8) (uint64, bool) {
	var l uint64
	r := uint64(len(f.Key))
	if r>0 {
		r--
	} else {
		return 0, false
	}
	var mid uint64 = uint64(((float32(thekey - f.Key[l])/float32(f.Key[r] - f.Key[l]))*float32(r))+0.5) // +0.5 makes it round instead of floor
	for l<=r && mid>=l && mid<=r {
		if (thekey < f.Key[mid]) {
			r = mid - 1;
		} else if (thekey > f.Key[mid]) {
			l = mid + 1;
		} else {
			return mid, true
		}
	mid = l + uint64(((float32(thekey - f.Key[l])/float32(f.Key[r] - f.Key[l]))*float32(r - l))+0.5)
	}
	return 0, false
}

// AddKeyUnsorted adds this key to the end of the index for later building with Build.
func (f *Key_uint8) AddKeyUnsorted(thekey uint8) {
	f.Key = append(f.Key, thekey)
	return
}

// AddKeyAt adds this key to the index in this exact position, so it does not require later rebuilding.
func (f *Key_uint8) AddKeyAt(thekey uint8, i uint64) {
	f.Key = append(f.Key, 0)
	copy(f.Key[i+1:], f.Key[i:])
	f.Key[i] = thekey
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
	return imap
}


// ---------- Key_string ----------

// Add this to any struct to make it binary searchable.
type Key_string struct {
Key []string
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
	var min,at uint64
	max := uint64(len(f.Key))
	if max>0 {
		max--
	} else {
		return 0, false
	}
	for min<=max {
		at = min+((max-min)/2)
		if thekey<f.Key[at] {
			if at==0 {
				return 0, false
			}
			max = at-1
		} else {
		if thekey>f.Key[at] {
			min = at+1
			} else {
				return at, true // found
			}
		}
	}
	return min, false // doesn't exist
}

// AddKeyUnsorted adds this key to the end of the index for later building with Build.
func (f *Key_string) AddKeyUnsorted(thekey string) {
	f.Key = append(f.Key, thekey)
	return
}

// AddKeyAt adds this key to the index in this exact position, so it does not require later rebuilding.
func (f *Key_string) AddKeyAt(thekey string, i uint64) {
	f.Key = append(f.Key, ``)
	copy(f.Key[i+1:], f.Key[i:])
	f.Key[i] = thekey
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
	return imap
}


// ---------- Key_byte ----------

// Add this to any struct to make it binary searchable.
type Key_bytes struct {
Key [][]byte
Keyindex []uint64
}

type sort_bytes struct {
i int
k []byte
}
type sorter_bytes []sort_bytes
func (a sorter_bytes) Len() int           { return len(a) }
func (a sorter_bytes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sorter_bytes) Less(i, j int) bool {
	aa:=a[i].k
	bb:=a[j].k
	if len(aa)<len(bb) {
		return true
	} else {
		if len(aa)>len(bb) {
			return false
		} else {
			for i:=0; i<len(aa); i++ {
				if aa[i]<bb[i] {
					return true
				} else {
					if aa[i]>bb[i] {
					return false
					}
				}
			}
			return false
		}
	}
}

// Find returns the index based on the key.
func (f *Key_bytes) Find(thekey []byte) (uint64, bool) {
	var at uint64
	keylen := len(thekey)
	min := f.Keyindex[keylen]
	max := f.Keyindex[keylen+1]
	fmt.Println(`min`,min,`max`,max)
	if max>0 {
		max--
	} else {
		return 0, false
	}
	Outer:
	for min<=max {
		at = min+((max-min)/2)
		fmt.Println(`min`,min,`max`,max,`at`,at)
		for i:=0; i<keylen; i++ {
			if thekey[i]<f.Key[at][i] {
				if at==0 {
					return 0, false
				}
				max = at-1
				continue Outer
			} else {
				if thekey[i]>f.Key[at][i] {
				min = at+1
				}
				continue Outer
			}
		}
		return at, true
	}
	return min, false // doesn't exist
}

// AddKeyUnsorted adds this key to the end of the index for later building with Build.
func (f *Key_bytes) AddKeyUnsorted(thekey []byte) {
	f.Key = append(f.Key, thekey)
	return
}

// AddKeyAt adds this key to the index in this exact position, so it does not require later rebuilding.
func (f *Key_bytes) AddKeyAt(thekey []byte, i uint64) {
	temp := make([]byte,0)
	f.Key = append(f.Key, temp)
	copy(f.Key[i+1:], f.Key[i:])
	f.Key[i] = thekey
	// Now modify the keyindex
	l := len(thekey)
	if l+1<len(f.Keyindex) { // first key of this length
		newar := make([]uint64,l+2)
		copy(newar,f.Keyindex)
		newar[l] = i
		newar[l+1] = uint64(len(f.Key))
	} else { // already have keys of this length
		for r:=l+1; r<l+2; r++ {
			f.Keyindex[r]++
		}
	}
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
	keyindex := make([]uint64,100)
	var max uint64
	for i:=0; i<l; i++ {
		imap[i]=temp[i].i
		newkey[i]=temp[i].k
		l2 := uint64(len(temp[i].k))
		if l2>max {
			max = l2
			if l2>uint64(len(keyindex)-2) {
				temp := make([]uint64,l*2)
				copy(temp,keyindex)
				keyindex = temp
			}
		}
		keyindex[l2]++
	}
	f.Key = newkey
	var at uint64
	newar := make([]uint64,max+2)
	for i:=uint64(0); i<max+2; i++ {
		newar[i]=at
		at+=keyindex[i]
	}
	f.Keyindex = newar
	return imap
}