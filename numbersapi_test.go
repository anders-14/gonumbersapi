package numbersapi_test

import (
	"fmt"
	"testing"

	gnu "github.com/anders-14/gonumbersapi"
)

func TestMathFact(t *testing.T) {
	fmt.Println(gnu.MathFact(39000))
}

func TestFunFact(t *testing.T) {
	fmt.Println(gnu.FunFact(1))
}

func TestDateFact(t *testing.T) {
	fmt.Println(gnu.DateFact("12/04"))
}

func TestYearFact(t *testing.T) {
	fmt.Println(gnu.YearFact(2002))
}
