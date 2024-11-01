package log

type Config struct {
	segment struct {
		maxStoreBytes uint64
		MaxIndexBytes uint64
		InitialOffset uint64
	}
}
