package statusFlags

type Bits uint16

const (
	//AlterUser       Bits = 1 << iota /* 1 */
	//LurUpRequest                     /* 2 */
	//SmsSendRequest                   /* 4 */
	//CallSendRequest                  /* 8 */
	//CallDropRequest                  /* 16 */
	//PagingRequest                    /* 32 */
	//CallRecvRequest                  /* 64 */

	AlterUser       Bits = 1 << iota /* 1 */
	LurUpRequest                     /* 2 */
	SmsSend                          /* 4 */
	Call                             /* 8 */
	CallDrop                         /* 16 */
	PagingRequest                    /* 32 */
	CallRecvRequest                  /* 64 */
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

func (s *Status) SetVal(val Bits) {
	s.val = val
}

func (s *Status) GetVal() Bits {
	return s.val
}

func (s *Status) AllClear() {
	s.val = 0
}
