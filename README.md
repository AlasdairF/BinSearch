##BinSearch

BinSearch is a binary search package for Go. It is extremely efficient on memory, and so is excessively useful for key/value lookups where millions or billion of records need to be stored in memory and retrieved quickly. I wrote this package for the Forgotten Books search engine.

Compared to a map BinSearch uses an order of magnitude less memory. For []byte lookups BinSearch is around 5x slower than map[string]value, but it would be misleading to think that this means BinSearch will slow down your app, as I have yet to find an example where lookup speed is a bottleneck. On the contrary, the purpose of BinSearch is to allow everything to be realistically stored in-memory, and thereby speeding up the app, and the entire server, due to reduced IO. In the case of Forgotten Books this package allows an inverted index corresponding to 100,000,000,000 (one hundred billion) words to be kept in memory at all times, and retrieved in virtually no time at all.

##Installation

    go get github.com/AlasdairF/BinSearch

##Usage

First choose your key type, the options are:
`Key_uint64`, `Key_uint32`, `Key_uint16`, `Key_uint8`, `Key_string` & `Key_bytes`

Numeric keys are much faster than string or []byte. `Key_bytes` is the most highly optimized and is twice as fast as `Key_string` so the only reason to ever use `Key_string` over `Key_bytes` is if you intend to select a range of values between two records where the strings contain non-English UTF8 characters. If you are only searching for individual records, as in most cases, or are only working in English, then always use `Key_bytes`.

Just add your chosen key into your struct along with any values you want for it, as slices. If you can have multiple values for the same key, or even no values (if you just want to check for the existence of the key.)

	type MyStruct struct {
		binsearch.Key_bytes
		value []int
		secondvalue []uint8
	}
	
BinSearch package only handles the key, so whenever you add a key you must add the value yourself.

There are two ways to add a key:

`MyStruct.AddKeyUnsorted(key)` adds the key to the end, unsorted, which means you should not search the keys until you call `MyStruct.Build()`. This is more efficient if you intend to add all keys in one go and you already know they are unique.

`MyStruct.AddKeyAt(key,position)` adds the key in this exact position. This is useful for when you want to add keys as you go, in which case you should add the key at the index value returned by `Find(key)`. 

	if indx, exists := MyStruct.Find(key); exists {
		value := MyStruct.value[indx]
		fmt.Println(`Value is`,value)
	} else {
		// Add the key
		MyStruct.AddKeyAt(key,indx)
		// Add the value also in the same place using this code:
		MyStruct.value = append(MyStruct.value, 0) // Enlarge by 1
		copy(MyStruct.value[indx+1:], MyStruct,value[indx:]) // Make space at indx
		MyStruct.value[indx] = 123 // Add your value into the correct position
	}
	
If you used `AddKeyUnsorted` then before you use `Find()` you must run `MyStruct.Build()`, which will sort the keys so they can be searched. You must then sort the values. `MyStruct.Build()` returns a slice telling how each of the old indexes maps to the new (sorted) index. This is all unnecesary if you used `AddKeyAt` since that inserts the key into the correct sorted position as you go, `Build()` is used only if you used `AddKeyUnsorted`. Here is the code for how to use `Build()` for both keys and values:

	temp := make(int,len(MyStruct.value))
	newindexes := MyStruct.Build()
	for indx_new,indx_old := range newindexes {
		temp[indx_new] = MyStruct.value[indx_old]
	}
	MyStruct.value = temp
	
If you have two values for each key, then you can do both at the same time:

	temp := make(int,len(MyStruct.value))
	newindexes := MyStruct.Build()
	for indx_new,indx_old := range newindexes {
		temp[indx_new] = MyStruct.value[indx_old]
	}
	MyStruct.value = temp
	
Now you have all your keys added you can find the index of any key using:

	indx, exists := MyStruct.Find(key)
	
Assuming the key exists, your value is then located at:
	
	value = MyStruct.value[indx]
	
If `exists==false` then `indx` will be the position where the key *should be*, which is useful for inserting it with `AddKeyAt`.

For numeric keys only you can also use interpolation search:

	indx, exists := MyStruct.FindInterpolation(key)
	
However, I strongly recommend not to use interpolation search unless you are certain that your keys are evenly distributed. Interpolation search is only faster when keys are evenly distributed, but can be much, much slower if they are not. If you're unsure just use `Find()`. If you are using `Key_string` or `Key_bytes` you can only use `Find()` anyway.
