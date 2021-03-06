package statusFlags

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	s := Status{}
	fmt.Println(s.Val())

	s.Toggle(AlterUser)
	fmt.Println("___________________")
	for i, flag := range []Bits{AlterUser, LurUpRequest, SmsSendRequest, CallSendRequest, CallDropRequest} {
		fmt.Println(i, s.Has(flag))
	}

	s.Toggle(SmsSendRequest)
	fmt.Println("___________________")
	for i, flag := range []Bits{AlterUser, LurUpRequest, SmsSendRequest, CallSendRequest, CallDropRequest} {
		fmt.Println(i, s.Has(flag))
	}

	fmt.Println("___________________")
	fmt.Println("check: ", s.Has(AlterUser))
	fmt.Println("check: ", s.Has(SmsSendRequest))
	fmt.Println("val: ", s.Val())

	s.Toggle(SmsSendRequest)
	fmt.Println("___________________")
	for i, flag := range []Bits{AlterUser, LurUpRequest, SmsSendRequest, CallSendRequest, CallDropRequest} {
		fmt.Println(i, s.Has(flag))
	}

	fmt.Println("___________________")
	fmt.Println("check: ", s.Has(AlterUser))
	fmt.Println("check: ", s.Has(SmsSendRequest))
	fmt.Println("val: ", s.Val())
}
