package numbersapi_test

import (
	"fmt"
	"testing"

  gna "github.com/anders-14/gonumbersapi"
)

func TestMathFact(t *testing.T) {
	fmt.Println(gna.MathFact(39000))
}

func TestFunFact(t *testing.T) {
	fmt.Println(gna.FunFact(1))
}

func TestDateFact(t *testing.T) {
	fmt.Println(gna.DateFact("12/04"))
}

func TestYearFact(t *testing.T) {
	fmt.Println(gna.YearFact(2002))
}
