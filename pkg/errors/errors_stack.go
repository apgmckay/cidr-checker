package errors

/* TODO:
- ErrorStack should be a singleton?
- Write Pop function
*/
type ErrorStack struct {
	errors []error
}

func New() *ErrorStack {
	es := &ErrorStack{}
	return es
}

func (es *ErrorStack) Pop() error {
	if es.Size() == 0 {
		return nil
	}
	result := es.errors[es.Size()-1]

	es.errors = es.errors[:len(es.errors)-1]

	return result
}

func (es *ErrorStack) Push(err error) error {
	es.errors = append(es.errors, err)
	return nil
}

func (es *ErrorStack) Size() int {
	return len(es.errors)
}
