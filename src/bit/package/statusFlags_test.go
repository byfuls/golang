package statusFlags

import (
	"fmt"
	"testing"
)

func ShowBitValue(s *Status) {
	fmt.Println("__________________________")
	fmt.Println("Eat    :    ", Eat)
	fmt.Println("Sleep  :    ", Sleep)
	fmt.Println("Sing   :    ", Sing)
	fmt.Println("Play   :    ", Play)
	fmt.Println("Call   :    ", Call)
	fmt.Println("Sms    :    ", Sms)
	fmt.Println("Shower :    ", Shower)
	fmt.Println("Work   :    ", Work)
	fmt.Println("Health :    ", Health)
	fmt.Println("__________________________")
}

func ShowBitStatus(s *Status) {
	fmt.Println("__________________________")
	fmt.Println("Eat    :    ", s.Has(Eat))
	fmt.Println("Sleep  :    ", s.Has(Sleep))
	fmt.Println("Sing   :    ", s.Has(Sing))
	fmt.Println("Play   :    ", s.Has(Play))
	fmt.Println("Call   :    ", s.Has(Call))
	fmt.Println("Sms    :    ", s.Has(Sms))
	fmt.Println("Shower :    ", s.Has(Shower))
	fmt.Println("Work   :    ", s.Has(Work))
	fmt.Println("Health :    ", s.Has(Health))
	fmt.Println("__________________________")
	fmt.Println("high byte value : ", s.GetValHigh())
	fmt.Println("low byte value  : ", s.GetValLow())
}

func TestMain(t *testing.T) {
	s := Status{}

	/* 초기 상태 확인 */
	ShowBitValue(&s)
	ShowBitStatus(&s)

	/* 비트 반전 후 상태 확인 */
	s.Toggle(Play)
	ShowBitStatus(&s)

	/* 비트 반전 후 상태 확인 */
	s.Toggle(Work)
	ShowBitStatus(&s)
}
