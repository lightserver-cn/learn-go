package main

type Nums interface {
	~uintptr | ~int8 | ~uint8 | ~int16 | ~uint16 | ~int | ~uint | ~int32 | ~uint32 | ~int64 | ~uint64 | ~float32 | ~float64
}

// ComparedBase 可以比较的基础类型
type ComparedBase interface {
	~bool | Nums | ~string
}

// UniqueSliceIntersect 交集，切片去重（两个切片共同部分的值）
// UniqueSliceIntersect[T ComparedBase]
//
//	@Description: 从slice中移除某个slice中所有的值，并且保持slice中值都唯一
//				1、切片去重，slice2传空切片，slice2 := []T{}
//				2、切片去重，添加某个值，slice2 := []T{item}
//				3、切片去重，并添加slice2中的所有值，添加多个切片，可先append操作后再添加
//	@param slice
//	@param remove
//	@return []T
func UniqueSliceIntersect[T ComparedBase](slice1, slice2 []T) []T {
	// 创建了一个空的结果切片 result，长度为0，容量为原始切片 slice 的长度。
	// 这样，在追加元素时，切片不会频繁地重新分配内存。
	result := make([]T, 0, 0)

	// 创建了 setMap，用于快速查找需要添加的元素。
	setMap := make(map[T]bool)
	for _, value := range slice2 {
		setMap[value] = true
	}

	// 辅助 map，用于去重
	seen := make(map[T]bool)

	// 使用双指针的方法进行遍历
	for i := 0; i < len(slice1); i++ {
		// i 指针用于遍历原始切片 slice，并检查当前元素是否需要添加。
		// 如果需要添加，则将其追加到结果切片 result 中。
		if setMap[slice1[i]] && !seen[slice1[i]] {
			result = append(result, slice1[i])
			// 标记已经添加到结果中的元素
			seen[slice1[i]] = true
		}
	}

	return result
}

// UniqueSliceUnion 并集，切片去重（两个切片所有的值）
// UniqueSliceUnion[T ComparedBase]
//
//	@Description: 从slice中移除某个slice中所有的值，并且保持slice中值都唯一
//				1、切片去重，slice2传空切片，slice2 := []T{}
//				2、切片去重，添加某个值，slice2 := []T{item}
//				3、切片去重，并添加slice2中的所有值，添加多个切片，可先append操作后再添加
//	@param slice
//	@param remove
//	@return []T
func UniqueSliceUnion[T ComparedBase](slice1, slice2 []T) []T {
	set := make(map[T]bool, len(slice1)+len(slice2))
	union := make([]T, 0, len(slice1)+len(slice2))

	// 将第一个切片的元素添加到并集切片和 map 中
	for _, value := range slice1 {
		if !set[value] {
			set[value] = true
			union = append(union, value)
		}
	}

	// 遍历第二个切片，将不在 map 中的元素添加到并集切片中
	for _, value := range slice2 {
		if !set[value] {
			set[value] = true
			union = append(union, value)
		}
	}

	return union
}

// UniqueSliceComplement 差集，切片去重（移除另一个切片的所有值）
// UniqueSliceComplement[T ComparedBase]
//
//	@Description: 从slice中移除某个slice中所有的值，并且保持slice中值都唯一
//				1、切片去重，remove传空切片，remove := []T{}
//				2、切片去重，并移除某个值，remove := []T{item}
//				3、切片去重，并移除remove中的所有值，移除多个切片，可先append操作后再移除
//	@param slice
//	@param remove
//	@return []T
func UniqueSliceComplement[T ComparedBase](slice, remove []T) []T {
	// 创建了一个空的结果切片 result，长度为0，容量为原始切片 slice 的长度。
	// 这样，在追加元素时，切片不会频繁地重新分配内存。
	result := make([]T, 0, len(slice))

	// 创建了 removeMap，用于快速查找需要移除的元素。
	removeMap := make(map[T]bool)
	for _, value := range remove {
		removeMap[value] = true
	}

	// 辅助 map，用于去重
	seen := make(map[T]bool)

	// 使用双指针的方法进行遍历
	for i := 0; i < len(slice); i++ {
		// i 指针用于遍历原始切片 slice，并检查当前元素是否需要移除。
		// 如果不需要移除，则将其追加到结果切片 result 中。
		if !removeMap[slice[i]] && !seen[slice[i]] {
			result = append(result, slice[i])
			// 标记已经添加到结果中的元素
			seen[slice[i]] = true
		}
	}

	return result
}
