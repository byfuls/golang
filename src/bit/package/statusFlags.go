package statusFlags

type Bits uint16

const (
	AlterUser       Bits = 1 << iota /* 1 */
	LurUpRequest                     /* 2 */
	SmsSendRequest                   /* 4 */
	CallSendRequest                  /* 8 */
	CallDropRequest                  /* 16 */
)

type Status struct {
	val Bits
}

func (s *Status) Set(flag Bits) {
	s.val = s.val | flag
}

func (s *Status) Clear(flag Bits) {
	s.val = s.val &^ flag
}

func (s *Status) Toggle(flag Bits) {
	s.val = s.val ^ flag
}

func (s *Status) Has(flag Bits) bool {
	return s.val&flag != 0
}

func (s *Status) Val() Bits {
	return s.val
}
