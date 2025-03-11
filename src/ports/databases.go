package ports

type IFileHandler interface {
	Write([]byte) (int, error)
	Read() ([]byte, error)
}
