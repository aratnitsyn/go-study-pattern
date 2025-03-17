package iterator

type Iterator interface {
	Index() int
	Value() interface{}
	Has() bool
	Next()
	Prev()
	Reset()
	End()
}

type Aggregate interface {
	Iterator() Iterator
}

type BookIterator struct {
	shelf    *BookShelf
	index    int
	internal int
}

func (i *BookIterator) Index() int {
	return i.index
}

func (i *BookIterator) Value() interface{} {
	return i.shelf.Books[i.index]
}

func (i *BookIterator) Has() bool {
	if i.internal < 0 || i.internal >= len(i.shelf.Books) {
		return false
	}
	return true
}

func (i *BookIterator) Next() {
	i.internal++
	if i.Has() {
		i.index++
	}
}

func (i *BookIterator) Prev() {
	i.internal--
	if i.Has() {
		i.index--
	}
}

func (i *BookIterator) Reset() {
	i.index = 0
	i.internal = 0
}

func (i *BookIterator) End() {
	i.index = len(i.shelf.Books) - 1
	i.internal = i.index
}

type BookShelf struct {
	Books []*Book
}

func (b *BookShelf) Iterator() Iterator {
	return &BookIterator{shelf: b}
}

func (b *BookShelf) Add(book *Book) {
	b.Books = append(b.Books, book)
}

type Book struct {
	Name string
}
