package expenses

import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	res := []Record{}
    for _, v := range in {
        if predicate(v) == true {
            res = append(res, v)
        }
    }
    return res
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
        return r.Day <= p.To && r.Day >= p.From
    }
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
        return r.Category == c
    }
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
    filter := ByDaysPeriod(p)
    total := 0.0
	for _, v := range in {
        if filter(v) == true {
        	total += v.Amount    
        }
    }
    return total
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	filter1 := ByCategory(c)
    filter2 := ByDaysPeriod(p)
    total := 0.0
    foundCategory := false
    for _, v := range in {
        if filter1(v) == true {
            foundCategory = true
            if filter2(v) == true {
            	total += v.Amount    
            }
        }
    }
    if !foundCategory {
        return total, errors.New("some error")
    }
    return total, nil
}
