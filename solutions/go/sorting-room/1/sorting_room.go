package sorting

import (
    "strconv"
    "fmt"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %s", strconv.FormatFloat(f, 'f', 1, 32))
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %s", strconv.FormatFloat(float64(nb.Number()), 'f', 1, 32))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	fancyNum, ok := fnb.(FancyNumber)
    if !ok {
        return 0
    }
    i, _ := strconv.Atoi(fancyNum.Value())
    return i
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	fancyNum, ok := fnb.(FancyNumber)
    if !ok {
        return "This is a fancy box containing the number 0.0"
    }
    f, _ := strconv.ParseFloat(fancyNum.Value(), 64)
    return fmt.Sprintf("This is a fancy box containing the number %s", strconv.FormatFloat(f, 'f', 1, 32))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
	switch i.(type) {
    case int:
        i := i.(int)
        return DescribeNumber(float64(i))
    case float64:
        i := i.(float64)
        return DescribeNumber(i)
    case NumberBox:
        i := i.(NumberBox)
        return DescribeNumberBox(i)
    case FancyNumberBox:
        i := i.(FancyNumberBox)
        return DescribeFancyNumberBox(i)
    default:
        return "Return to sender"
    }
}
