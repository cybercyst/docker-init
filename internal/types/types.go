package types

type Target struct {
	TargetType TargetType
	Path       string
}

type TargetType uint8

func (t TargetType) ToString() string {
	return []string{"Go"}[t]
}

const (
	Go TargetType = iota
	None
)
