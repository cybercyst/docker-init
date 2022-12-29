package types

type Target struct {
	TargetType TargetType
	Path       string
}

type TargetType uint8

const (
	Go TargetType = iota
	None
)
