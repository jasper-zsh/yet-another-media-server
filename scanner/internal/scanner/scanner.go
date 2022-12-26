package scanner

type Scanner interface {
	Scan(libraryID int32, options map[string]string) error
}
