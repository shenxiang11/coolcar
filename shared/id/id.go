package id

type AccountID string

func (a AccountID) String() string {
	return string(a)
}
