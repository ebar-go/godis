package types

type SdsFlag uint8 // 定义类型

const (
	SdsHDR5 SdsFlag = iota
	SdsHDR8
	SdsHDR16
	SdsHDR32
	SdsHDR64
)

const (
	SdsMaxPreAlloc = 1 << 20
)

// SDS simple dynamic string, 字符串类型的底层存储结构，主要用于key和string类型的val
type SDS struct {
	len   uint    // 记录字符串长度
	alloc uint    // 记录bytes数组的长度
	flags SdsFlag // 类型
	buf   []byte  // 保存实际数据
}

// NewSDS creates a new SDS object
func NewSDS(str string) *SDS {
	s := &SDS{}
	s.Update(str)
	return s
}

// Len returns the number of SDS bytes
func (s *SDS) Len() uint { return s.len }

// avail returns the available bytes of the SDS
func (s *SDS) avail() uint { return s.alloc - s.len }

// Update updates the SDS object value
func (s *SDS) Update(str string) {
	sl := uint(len(str))
	if sl == 0 { // 代表传入的是空字符串
		s.len = 0
		return
	}

	// 检查剩余可用空间是否足够存储新增的字符串长度
	if s.avail() > sl-s.len {
		s.len = sl
		copy(s.buf, str)
		return
	}

	// 空间不足时，需要重新分配空间
	if sl < SdsMaxPreAlloc {
		s.alloc = sl * 2
	} else {
		s.alloc = sl + SdsMaxPreAlloc
	}
	s.len = sl
	s.buf = make([]byte, s.alloc)
	copy(s.buf, str)
}

// String returns a string representation of SDS
func (s *SDS) String() string {
	return string(s.buf[:s.len])
}
