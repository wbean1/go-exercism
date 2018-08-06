package listops

type IntList []int
type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

func (a IntList) Foldl(f binFunc, i int) int {
	for _, x := range a {
		i = f(i, x)
	}
	return i
}

func (a IntList) Foldr(f binFunc, i int) int {
	for x := len(a) - 1; x >= 0; x-- {
		i = f(a[x], i)
	}
	return i
}

func (a IntList) Length() int {
	return 0
}

func (a IntList) Append(i IntList) IntList {
	return IntList{}
}
func (a IntList) Reverse() IntList {
	return IntList{}
}
func (a IntList) Concat(i []IntList) IntList {
	return IntList{}
}

func (a IntList) Filter(f predFunc) IntList {
	return IntList{}
}

func (a IntList) Map(f unaryFunc) IntList {
	return IntList{}
}
