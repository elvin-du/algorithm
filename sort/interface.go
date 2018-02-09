package sort

type Sorter interface {
	Sort() error
	IsSorted() bool
	Show(int)
}
