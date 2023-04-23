package gdson

type Hour struct {
	Start int
	End   int
}

type Weekday int16

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Null Weekday = -1
)

type Condition struct {
	Hour    []Hour
	Day     []int
	Month   []int
	Weekday []Weekday
}
