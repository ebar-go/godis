package types

type SdsFlag uint8

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

type SDS struct {
	len   uint
	alloc uint
	flags SdsFlag
	buf   []byte
}

func NewSDS(str string) *SDS {
	s := &SDS{}
	s.Update(str)
	return s
}

func (s *SDS) Len() uint { return s.len }

func (s *SDS) Update(str string) {
	avail := s.alloc - s.len
	sl := uint(len(str))
	if avail > sl-s.len {
		s.len = sl
		copy(s.buf, str)
		return
	}

	if sl < SdsMaxPreAlloc {
		s.alloc = sl * 2
	} else {
		s.alloc = sl + SdsMaxPreAlloc
	}
	s.len = sl
	s.buf = make([]byte, s.alloc)
	copy(s.buf, str)
}

func (s *SDS) String() string {
	return string(s.buf[:s.len])
}
