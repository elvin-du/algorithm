package sort

type Selection struct{}

func NewSelection() *Selection {
	return &Selection{}
}

func (s *Selection) ReadRawData() error {
	return nil
}

func (s *Selection) Sort() error {
	return nil
}
func (s *Selection) IsSorted() bool {
	return false
}
func (s *Selection) Show(limit int) {

}
