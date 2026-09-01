package roman

import (
	"errors"
	"fmt"
	"strings"
)

type (
	Number int
	Roman  string
)

var ErrInvalidNumber = errors.New("invalid number")

func (r Roman) String() string {
	return fmt.Sprintf("%s Roman Value", string(r))
}




func ConvertToRoman(number Number) (Roman, error) {
	type dict struct {
		Value  Number
		Symbol Roman
	}
	dictSlice := []dict{
		{
			Value:  10,
			Symbol: "X",
		},
				{
			Value:  9,
			Symbol: "IX",
		},
		{
			Value:  5,
			Symbol: "V",
		},
				{
			Value:  4,
			Symbol: "IV",
		},
		{
			Value:  1,
			Symbol: "I",
		},
	}

	var sb strings.Builder
	if number == 0 {
		return "", ErrInvalidNumber
	}
	for number != 0 {
		var max Number = 0
		var best string = ""
		for _,val:=range dictSlice{
			if val.Value >= max && val.Value<=number {
				max = val.Value
				best = string(val.Symbol)
			} 
		}
		sb.WriteString(best)
		number = number - max
	}
	return Roman(sb.String()),nil
}
