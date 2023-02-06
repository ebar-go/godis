package types

const (
	SdsMaxPreAlloc = 1 << 20 // 1M
	SdsMaxAlloc    = 1 << 29 // 512M
)

// SDS simple dynamic string, 字符串类型的底层存储结构，主要用于key和string类型的val
type SDS struct {
	len   uint64 // 记录字符串长度
	alloc uint64 // 记录bytes数组的长度
	//flags SdsFlag // 类型
	buf []byte // 保存实际数据
}

// NewSDS creates a new SDS object
func NewSDS(buf []byte) *SDS {
	s := &SDS{}
	s.Set(buf)
	return s
}

// Len returns the number of SDS bytes
func (s *SDS) Len() uint64 { return s.len }

// avail returns the available bytes of the SDS
func (s *SDS) avail() uint64 { return s.alloc - s.len }

// Update updates the SDS object value
func (s *SDS) Set(b []byte) {
	sl := uint64(len(b))
	if sl == 0 { // 代表传入的是空字符串
		s.len = 0
		return
	}

	if sl == SdsMaxAlloc { // 字符串最大长度为 512M
		return
	}

	// 检查剩余可用空间是否足够存储新增的字符串长度
	if s.avail() > sl-s.len {
		s.len = sl
		copy(s.buf, b)
		return
	}

	// 空间不足时，需要重新分配空间
	// 拼接后的字符串长度不超过1M，分配两倍的内存
	// 拼接够的字符串长度超过1M，多分配1M的内存
	if sl < SdsMaxPreAlloc {
		s.alloc = sl * 2
	} else {
		s.alloc = sl + SdsMaxPreAlloc
	}
	s.len = sl
	s.buf = make([]byte, s.alloc)
	copy(s.buf, b)
}

// String returns a string representation of SDS
func (s *SDS) String() string {
	return string(s.buf[:s.len])
}

func (s *SDS) Bytes() []byte {
	return s.buf[:s.len]
}
