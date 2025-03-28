package template_method

type QuotesInterface interface {
	Open() string
	Close() string
}

type Quotes struct {
	QuotesInterface
}

func (q *Quotes) Quotes(str string) string {
	return q.Open() + str + q.Close()
}

func NewQuotes(qt QuotesInterface) *Quotes {
	return &Quotes{qt}
}

type FrenchQuotes struct {
}

func (q *FrenchQuotes) Open() string {
	return "<"
}

func (q *FrenchQuotes) Close() string {
	return ">"
}

type GermanQuotes struct {
}

func (q *GermanQuotes) Open() string {
	return "„"
}

func (q *GermanQuotes) Close() string {
	return "“"
}
