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

func (es *ErrorStack) Push(err error) error {
	es.errors = append(es.errors, err)
	return nil
}

func (es *ErrorStack) Size() int {
	return len(es.errors)
}
