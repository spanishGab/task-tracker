package ports

type IFileHandler interface {
	Open() error
	Close() error
	Write(data []byte) (int, error)
	Read() ([]byte, error)
}
