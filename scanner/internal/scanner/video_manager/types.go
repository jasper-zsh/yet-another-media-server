package video_manager

type Video struct {
	ID               uint
	SourceIdentifier string
	Title            string
	Cover            []byte
	VideoFiles       []byte
}
