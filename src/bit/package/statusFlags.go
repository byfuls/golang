package statusFlags

type UByte uint8
type UShort uint16

const (
	/*_______________LOBYTE_________________*/
	AlterUser       UShort = 1 << iota /* 1 */
	LurUpRequest                       /* 2 */
	SmsSend                            /* 4 */
	Call                               /* 8 */
	CallDrop                           /* 16 */
	PagingRequest                      /* 32 */
	CallRecvRequest                    /* 64 */
	/*_______________HIBYTE_________________*/
	_
	Idle   /* 128 */
	Active /* 256 */
	Search /* 512 */
)

type Status struct {
	val UShort
	hi  UByte
	lo  UByte
}

func (s *Status) Set(flag UShort) {
	s.val = s.val | flag
}

func (s *Status) Clear(flag UShort) {
	s.val = s.val &^ flag
}

func (s *Status) Toggle(flag UShort) {
	s.val = s.val ^ flag
}

func (s *Status) Has(flag UShort) bool {
	return s.val&flag != 0
}

func (s *Status) SetVal(val UShort) {
	s.val = val
}

func (s *Status) GetVal() UShort {
	return s.val
}

func (s *Status) AllClear() {
	s.val = 0
}
