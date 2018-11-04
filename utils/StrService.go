package utils

type IString interface {
	SubString(index int, end int) string
	IndexOf(index string) int
	SliceIndex(target string, index string)
}

type Stringer struct {
	Str string
}

func (str Stringer) SubString(index int, end int) string {

	if index > end {
		panic(nil)

	}
	str1 := str.Str[index:end]
	return str1

}

//切割数据我的数据
func (str Stringer) IndexOf(index string) int {

	for head, str1 := range str.Str {

		if string(str1) == index {
			return head
		}
	}

	return -1
}

//切割数据
func (str Stringer) SliceIndex( index string) {





}
