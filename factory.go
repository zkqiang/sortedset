package sortedset

func New[T comparable](sortFunc func(i, j T) bool) *SortedSet[T] {
	return &SortedSet[T]{
		elements: make(map[T]struct{}),
		order:    []T{},
		sortFunc: sortFunc,
	}
}

func NewInt() *SortedSet[int] {
	return &SortedSet[int]{
		elements: make(map[int]struct{}),
		order:    []int{},
		sortFunc: func(i, j int) bool {
			return i < j
		},
	}
}

func NewInt8() *SortedSet[int8] {
	return &SortedSet[int8]{
		elements: make(map[int8]struct{}),
		order:    []int8{},
		sortFunc: func(i, j int8) bool {
			return i < j
		},
	}
}

func NewInt16() *SortedSet[int16] {
	return &SortedSet[int16]{
		elements: make(map[int16]struct{}),
		order:    []int16{},
		sortFunc: func(i, j int16) bool {
			return i < j
		},
	}
}

func NewInt32() *SortedSet[int32] {
	return &SortedSet[int32]{
		elements: make(map[int32]struct{}),
		order:    []int32{},
		sortFunc: func(i, j int32) bool {
			return i < j
		},
	}
}

func NewInt64() *SortedSet[int64] {
	return &SortedSet[int64]{
		elements: make(map[int64]struct{}),
		order:    []int64{},
		sortFunc: func(i, j int64) bool {
			return i < j
		},
	}
}

func NewUint() *SortedSet[uint] {
	return &SortedSet[uint]{
		elements: make(map[uint]struct{}),
		order:    []uint{},
		sortFunc: func(i, j uint) bool {
			return i < j
		},
	}
}

func NewUint8() *SortedSet[uint8] {
	return &SortedSet[uint8]{
		elements: make(map[uint8]struct{}),
		order:    []uint8{},
		sortFunc: func(i, j uint8) bool {
			return i < j
		},
	}
}

func NewUint16() *SortedSet[uint16] {
	return &SortedSet[uint16]{
		elements: make(map[uint16]struct{}),
		order:    []uint16{},
		sortFunc: func(i, j uint16) bool {
			return i < j
		},
	}
}

func NewUint32() *SortedSet[uint32] {
	return &SortedSet[uint32]{
		elements: make(map[uint32]struct{}),
		order:    []uint32{},
		sortFunc: func(i, j uint32) bool {
			return i < j
		},
	}
}

func NewUint64() *SortedSet[uint64] {
	return &SortedSet[uint64]{
		elements: make(map[uint64]struct{}),
		order:    []uint64{},
		sortFunc: func(i, j uint64) bool {
			return i < j
		},
	}
}

func NewFloat32() *SortedSet[float32] {
	return &SortedSet[float32]{
		elements: make(map[float32]struct{}),
		order:    []float32{},
		sortFunc: func(i, j float32) bool {
			return i < j
		},
	}
}

func NewFloat64() *SortedSet[float64] {
	return &SortedSet[float64]{
		elements: make(map[float64]struct{}),
		order:    []float64{},
		sortFunc: func(i, j float64) bool {
			return i < j
		},
	}
}

func NewString() *SortedSet[string] {
	return &SortedSet[string]{
		elements: make(map[string]struct{}),
		order:    []string{},
		sortFunc: func(i, j string) bool {
			return i < j
		},
	}
}
