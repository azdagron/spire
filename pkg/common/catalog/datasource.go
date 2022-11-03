package catalog

type dataSource interface {
	Data() (string, error)
	Updated() (<-chan struct{}, bool)
	Close() error
}

type fixedData string

func (d fixedData) Data() (string, error) {
	return string(d), nil
}

func (d fixedData) Updated() (<-chan struct{}, bool) {
	return nil, false
}

func (d fixedData) Close() error {
	return nil
}

type fileData string

func (d fixedData) Data() (string, error) {
	return string(d), nil
}

func (d fixedData) Updated() (<-chan struct{}, bool) {
	return nil, false
}

func (d fixedData) Close() error {
	return nil
}
