package statusFlags

import (
	"fmt"
	"testing"
)

func ShowBitStatus(s *Status) {
	fmt.Println("__________________________")
	fmt.Println("AlterUser?       ", s.Has(AlterUser))
	fmt.Println("LurUpRequest?    ", s.Has(LurUpRequest))
	fmt.Println("SmsSend?         ", s.Has(SmsSend))
	fmt.Println("Call?            ", s.Has(Call))
	fmt.Println("CallDrop?        ", s.Has(CallDrop))
	fmt.Println("PagingRequest?   ", s.Has(PagingRequest))
	fmt.Println("CallRecvRequest? ", s.Has(CallRecvRequest))
	fmt.Println("Idle?            ", s.Has(Idle))
	fmt.Println("Active?          ", s.Has(Active))
	fmt.Println("Search?          ", s.Has(Search))
	fmt.Println("__________________________")
}

func TestMain(t *testing.T) {
	fmt.Println(AlterUser)
	fmt.Println(Idle)
	fmt.Println(Active)
	fmt.Println(Search)

	fmt.Println("_________________")
	s := Status{}
	s.Toggle(Idle)
	ShowBitStatus(&s)

	s.Toggle(LurUpRequest)
	ShowBitStatus(&s)

	s.Toggle(LurUpRequest)
	ShowBitStatus(&s)

	/*
		s := Status{}
		fmt.Println(s.GetVal())

		s.Toggle(AlterUser)
		fmt.Println("___________________")
		for i, flag := range []Bits{AlterUser, LurUpRequest, SmsSend, Call, CallDrop} {
			fmt.Println(i, s.Has(flag))
		}

		s.Toggle(SmsSend)
		fmt.Println("___________________")
		for i, flag := range []Bits{AlterUser, LurUpRequest, SmsSend, Call, CallDrop} {
			fmt.Println(i, s.Has(flag))
		}

		fmt.Println("___________________")
		fmt.Println("check: ", s.Has(AlterUser))
		fmt.Println("check: ", s.Has(SmsSend))
		fmt.Println("val: ", s.GetVal())

		s.Toggle(SmsSend)
		fmt.Println("___________________")
		for i, flag := range []Bits{AlterUser, LurUpRequest, SmsSend, Call, CallDrop} {
			fmt.Println(i, s.Has(flag))
		}

		fmt.Println("___________________")
		fmt.Println("check: ", s.Has(AlterUser))
		fmt.Println("check: ", s.Has(SmsSend))
		fmt.Println("val: ", s.GetVal())
	*/
}
