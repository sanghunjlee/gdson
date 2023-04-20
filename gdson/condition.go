package gdson

type Hour struct {
	Start uint16
	End   uint16
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
	Day     []uint16
	Month   []uint16
	Weekday []Weekday
}

