package types

// Money for dirams, sents, kopeyks 
type Money int64

// Category describes payment category
type Category string

//Status info about payment status
type Status string

//Statuses for payments
const (
	StatusOk         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

// Payment describes structure for payments
type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}
