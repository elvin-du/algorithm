package sort

type Sorter interface {
	ReadRawData() error
	Sort() error
	IsSorted() bool
	Show(int)
}
