package issue

// Recorder godoc
type Recorder interface {
	Record(err error) error
}
