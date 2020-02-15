package ameda

// OneUint try to return the first element, otherwise return zero value.
func OneUint(u []uint) uint {
	if len(u) > 0 {
		return u[0]
	}
	return 0
}

// UintsCopy creates a copy of the uint slice.
func UintsCopy(u []uint) []uint {
	b := make([]uint, len(u))
	copy(b, u)
	return b
}

// UintsToInterfaces converts uint slice to interface slice.
func UintsToInterfaces(u []uint) []interface{} {
	r := make([]interface{}, len(u))
	for k, v := range u {
		r[k] = UintToInterface(v)
	}
	return r
}

// UintsToStrings converts uint slice to string slice.
func UintsToStrings(u []uint) []string {
	r := make([]string, len(u))
	for k, v := range u {
		r[k] = UintToString(v)
	}
	return r
}

// UintsToBools converts uint slice to bool slice.
// NOTE:
//  0 is false, everything else is true
func UintsToBools(u []uint) []bool {
	r := make([]bool, len(u))
	for k, v := range u {
		r[k] = UintToBool(v)
	}
	return r
}

// UintsToFloat32s converts uint slice to float32 slice.
func UintsToFloat32s(u []uint) []float32 {
	r := make([]float32, len(u))
	for k, v := range u {
		r[k] = UintToFloat32(v)
	}
	return r
}

// UintsToFloat64s converts uint slice to float64 slice.
func UintsToFloat64s(u []uint) []float64 {
	r := make([]float64, len(u))
	for k, v := range u {
		r[k] = UintToFloat64(v)
	}
	return r
}

// UintsToInts converts uint slice to int slice.
func UintsToInts(u []uint) ([]int, error) {
	var err error
	r := make([]int, len(u))
	for k, v := range u {
		r[k], err = UintToInt(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// UintsToInt8s converts uint slice to int8 slice.
func UintsToInt8s(u []uint) ([]int8, error) {
	var err error
	r := make([]int8, len(u))
	for k, v := range u {
		r[k], err = UintToInt8(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// UintsToInt16s converts uint slice to int16 slice.
func UintsToInt16s(u []uint) ([]int16, error) {
	var err error
	r := make([]int16, len(u))
	for k, v := range u {
		r[k], err = UintToInt16(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// UintsToInt32s converts uint slice to int32 slice.
func UintsToInt32s(u []uint) ([]int32, error) {
	var err error
	r := make([]int32, len(u))
	for k, v := range u {
		r[k], err = UintToInt32(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// UintsToInt64s converts uint slice to int64 slice.
func UintsToInt64s(u []uint) ([]int64, error) {
	var err error
	r := make([]int64, len(u))
	for k, v := range u {
		r[k], err = UintToInt64(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// UintsToUint8s converts uint slice to uint8 slice.
func UintsToUint8s(u []uint) ([]uint8, error) {
	var err error
	r := make([]uint8, len(u))
	for k, v := range u {
		r[k], err = UintToUint8(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// UintsToUint16s converts uint slice to uint16 slice.
func UintsToUint16s(u []uint) ([]uint16, error) {
	var err error
	r := make([]uint16, len(u))
	for k, v := range u {
		r[k], err = UintToUint16(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// UintsToUint32s converts uint slice to uint32 slice.
func UintsToUint32s(u []uint) ([]uint32, error) {
	var err error
	r := make([]uint32, len(u))
	for k, v := range u {
		r[k], err = UintToUint32(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// UintsToUint64s converts uint slice to uint64 slice.
func UintsToUint64s(u []uint) []uint64 {
	r := make([]uint64, len(u))
	for k, v := range u {
		r[k] = UintToUint64(v)
	}
	return r
}

// UintsConcat is used to merge two or more slices.
// This method does not change the existing slices, but instead returns a new slice.
func UintsConcat(u []uint, a ...[]uint) []uint {
	totalLen := len(u)
	for _, v := range a {
		totalLen += len(v)
	}
	ret := make([]uint, totalLen)
	n := copy(ret, u)
	dst := ret[n:]
	for _, v := range a {
		n := copy(dst, v)
		dst = dst[n:]
	}
	return ret
}

// UintsCopyWithin copies part of an slice to another location in the current slice.
// @target
//  Zero-based index at which to copy the sequence to. If negative, target will be counted from the end.
// @start
//  Zero-based index at which to start copying elements from. If negative, start will be counted from the end.
// @end
//  Zero-based index at which to end copying elements from. CopyWithin copies up to but not including end.
//  If negative, end will be counted from the end.
//  If end is omitted, CopyWithin will copy until the last index (default to len(s)).
func UintsCopyWithin(u []uint, target, start int, end ...int) {
	target = fixIndex(len(u), target, true)
	if target == len(u) {
		return
	}
	sub := UintsSlice(u, start, end...)
	for k, v := range sub {
		u[target+k] = v
	}
}

// UintsEvery tests whether all elements in the slice pass the test implemented by the provided function.
// NOTE:
//  Calling this method on an empty slice will return true for any condition!
func UintsEvery(u []uint, fn func(u []uint, k int, v uint) bool) bool {
	for k, v := range u {
		if !fn(u, k, v) {
			return false
		}
	}
	return true
}

// UintsFill changes all elements in the current slice to a value, from a start index to an end index.
// @value
//  Zero-based index at which to copy the sequence to. If negative, target will be counted from the end.
// @start
//  Zero-based index at which to start copying elements from. If negative, start will be counted from the end.
// @end
//  Zero-based index at which to end copying elements from. CopyWithin copies up to but not including end.
//  If negative, end will be counted from the end.
//  If end is omitted, CopyWithin will copy until the last index (default to len(s)).
func UintsFill(u []uint, value uint, start int, end ...int) {
	fixedStart, fixedEnd, ok := fixRange(len(u), start, end...)
	if !ok {
		return
	}
	for k := fixedStart; k < fixedEnd; k++ {
		u[k] = value
	}
}

// UintsFilter creates a new slice with all elements that pass the test implemented by the provided function.
func UintsFilter(u []uint, fn func(u []uint, k int, v uint) bool) []uint {
	ret := make([]uint, 0)
	for k, v := range u {
		if fn(u, k, v) {
			ret = append(ret, v)
		}
	}
	return ret
}

// UintsFind returns the key-value of the first element in the provided slice that satisfies the provided testing function.
// NOTE:
//  If not found, k = -1
func UintsFind(u []uint, fn func(u []uint, k int, v uint) bool) (k int, v uint) {
	for k, v := range u {
		if fn(u, k, v) {
			return k, v
		}
	}
	return -1, 0
}

// UintsIncludes determines whether an slice includes a certain value among its entries.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func UintsIncludes(u []uint, valueToFind uint, fromIndex ...int) bool {
	return UintsIndexOf(u, valueToFind, fromIndex...) > -1
}

// UintsIndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func UintsIndexOf(u []uint, searchElement uint, fromIndex ...int) int {
	idx := getFromIndex(len(u), fromIndex...)
	for k, v := range u[idx:] {
		if searchElement == v {
			return k + idx
		}
	}
	return -1
}

// UintsLastIndexOf returns the last index at which a given element can be found in the slice, or -1 if it is not present.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func UintsLastIndexOf(u []uint, searchElement uint, fromIndex ...int) int {
	idx := getFromIndex(len(u), fromIndex...)
	for k := len(u) - 1; k >= idx; k-- {
		if searchElement == u[k] {
			return k
		}
	}
	return -1
}

// UintsMap creates a new slice populated with the results of calling a provided function
// on every element in the calling slice.
func UintsMap(u []uint, fn func(u []uint, k int, v uint) uint) []uint {
	ret := make([]uint, len(u))
	for k, v := range u {
		ret[k] = fn(u, k, v)
	}
	return ret
}

// UintsPop removes the last element from an slice and returns that element.
// This method changes the length of the slice.
func UintsPop(u *[]uint) (uint, bool) {
	a := *u
	if len(a) == 0 {
		return 0, false
	}
	lastIndex := len(a) - 1
	last := a[lastIndex]
	a = a[:lastIndex]
	*u = a[:len(a):len(a)]
	return last, true
}

// UintsPush adds one or more elements to the end of an slice and returns the new length of the slice.
func UintsPush(u *[]uint, element ...uint) int {
	*u = append(*u, element...)
	return len(*u)
}

// UintsPushOnce adds one or more new elements that do not exist in the current slice at the end
// and returns the new length of the slice.
func UintsPushOnce(u *[]uint, element ...uint) int {
	a := *u
L:
	for _, v := range element {
		for _, vv := range a {
			if vv == v {
				continue L
			}
		}
		a = append(a, v)
	}
	*u = a
	return len(a)
}

// UintsReduce executes a reducer function (that you provide) on each element of the slice,
// resulting in a single output value.
// @accumulator
//  The accumulator accumulates callback's return values.
//  It is the accumulated value previously returned in the last invocation of the callback—or initialValue,
//  if it was supplied (see below).
// @initialValue
//  A value to use as the first argument to the first call of the callback.
//  If no initialValue is supplied, the first element in the slice will be used and skipped.
func UintsReduce(u []uint,
	fn func(u []uint, k int, v, accumulator uint) uint, initialValue ...uint,
) uint {
	if len(u) == 0 {
		return 0
	}
	start := 0
	acc := u[start]
	if len(initialValue) > 0 {
		acc = initialValue[0]
	} else {
		start += 1
	}
	for k := start; k < len(u); k++ {
		acc = fn(u, k, u[k], acc)
	}
	return acc
}

// UintsReduceRight applies a function against an accumulator and each value of the slice (from right-to-left)
// to reduce it to a single value.
// @accumulator
//  The accumulator accumulates callback's return values.
//  It is the accumulated value previously returned in the last invocation of the callback—or initialValue,
//  if it was supplied (see below).
// @initialValue
//  A value to use as the first argument to the first call of the callback.
//  If no initialValue is supplied, the first element in the slice will be used and skipped.
func UintsReduceRight(u []uint,
	fn func(u []uint, k int, v, accumulator uint) uint, initialValue ...uint,
) uint {
	if len(u) == 0 {
		return 0
	}
	end := len(u) - 1
	acc := u[end]
	if len(initialValue) > 0 {
		acc = initialValue[0]
	} else {
		end -= 1
	}
	for k := end; k >= 0; k-- {
		acc = fn(u, k, u[k], acc)
	}
	return acc
}

// UintsReverse reverses an slice in place.
func UintsReverse(u []uint) {
	first := 0
	last := len(u) - 1
	for first < last {
		u[first], u[last] = u[last], u[first]
		first++
		last--
	}
}

// UintsShift removes the first element from an slice and returns that removed element.
// This method changes the length of the slice.
func UintsShift(u *[]uint) (uint, bool) {
	a := *u
	if len(a) == 0 {
		return 0, false
	}
	first := a[0]
	a = a[1:]
	*u = a[:len(a):len(a)]
	return first, true
}

// UintsSlice returns a copy of a portion of an slice into a new slice object selected
// from begin to end (end not included) where begin and end represent the index of items in that slice.
// The original slice will not be modified.
func UintsSlice(u []uint, begin int, end ...int) []uint {
	fixedStart, fixedEnd, ok := fixRange(len(u), begin, end...)
	if !ok {
		return []uint{}
	}
	return UintsCopy(u[fixedStart:fixedEnd])
}

// UintsSome tests whether at least one element in the slice passes the test implemented by the provided function.
// NOTE:
//  Calling this method on an empty slice returns false for any condition!
func UintsSome(u []uint, fn func(u []uint, k int, v uint) bool) bool {
	for k, v := range u {
		if fn(u, k, v) {
			return true
		}
	}
	return false
}

// UintsSplice changes the contents of an slice by removing or replacing
// existing elements and/or adding new elements in place.
func UintsSplice(u *[]uint, start, deleteCount int, items ...uint) {
	a := *u
	if deleteCount < 0 {
		deleteCount = 0
	}
	start, end, _ := fixRange(len(a), start, start+1+deleteCount)
	deleteCount = end - start - 1
	for k := 0; k < len(items); k++ {
		if deleteCount > 0 {
			// replace
			a[start] = items[k]
			deleteCount--
			start++
		} else {
			// insert
			lastSlice := UintsCopy(a[start:])
			items = items[k:]
			a = append(a[:start], items...)
			a = append(a[:start+len(items)], lastSlice...)
			*u = a[:len(a):len(a)]
			return
		}
	}
	if deleteCount > 0 {
		a = append(a[:start], a[start+1+deleteCount:]...)
	}
	*u = a[:len(a):len(a)]
}

// UintsUnshift adds one or more elements to the beginning of an slice and returns the new length of the slice.
func UintsUnshift(u *[]uint, element ...uint) int {
	*u = append(element, *u...)
	return len(*u)
}

// UintsUnshiftOnce adds one or more new elements that do not exist in the current slice to the beginning
// and returns the new length of the slice.
func UintsUnshiftOnce(u *[]uint, element ...uint) int {
	a := *u
	if len(element) == 0 {
		return len(a)
	}
	m := make(map[uint]bool, len(element))
	r := make([]uint, 0, len(a)+len(element))
L:
	for _, v := range element {
		if m[v] {
			continue
		}
		m[v] = true
		for _, vv := range a {
			if vv == v {
				continue L
			}
		}
		r = append(r, v)
	}
	r = append(r, a...)
	*u = r[:len(r):len(r)]
	return len(r)
}

// UintsDistinct creates a new slice in place set that removes the same elements
// and returns the new length of the slice.
func UintsDistinct(u *[]uint) int {
	a := (*u)[:0]
	m := make(map[uint]bool, len(a))
	for _, v := range *u {
		if m[v] {
			continue
		}
		a = append(a, v)
		m[v] = true
	}
	n := len(m)
	*u = a[:n:n]
	return n
}

// UintsRemoveFirst removes the first matched elements from the slice,
// and returns the new length of the slice.
func UintsRemoveFirst(u *[]uint, element ...uint) int {
	a := *u
	m := make(map[uint]bool, len(element))
	for _, v := range element {
		if m[v] {
			continue
		}
		m[v] = true
		for kk, vv := range a {
			if vv == v {
				a = append(a[:kk], a[kk+1:]...)
				break
			}
		}
	}
	n := len(a)
	*u = a[:n:n]
	return n
}

// UintsRemoveEvery removes all the elements from the slice,
// and returns the new length of the slice.
func UintsRemoveEvery(u *[]uint, element ...uint) int {
	a := *u
	m := make(map[uint]bool, len(element))
	for _, v := range element {
		if m[v] {
			continue
		}
		m[v] = true
		for kk, vv := range a {
			if vv == v {
				a = append(a[:kk], a[kk+1:]...)
			}
		}
	}
	n := len(a)
	*u = a[:n:n]
	return n
}