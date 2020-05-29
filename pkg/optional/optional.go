package optional

import "errors"

type Optional struct {
	value interface{}
}

type ExecuteIfPresentFunc func(interface{}) error

type FilterFunc func(interface{}) bool

type MapFunc func(interface{}) interface{}

type GetFunc func() interface{}

type ErrorFunc func() error

func Of(value interface{}) (*Optional, error) {
	if value == nil {
		return nil, errors.New("'value' cannot be nil")
	}
	return &Optional{
		value: value,
	}, nil
}

func OfNillable(value interface{}) *Optional {
	return &Optional{
		value: value,
	}
}

func Empty() *Optional {
	return &Optional{
		value: nil,
	}
}

func (optional Optional) Get() (interface{}, error) {
	if optional.value == nil {
		return nil, errors.New("no value present")
	}
	return optional.value, nil
}

func (optional Optional) IsPresent() bool {
	return optional.value != nil
}

func (optional Optional) IfPresent(fn ExecuteIfPresentFunc) error {
	if optional.value == nil {
		return nil
	}
	if fn == nil {
		return errors.New("'fn' cannot be nil")
	}
	return fn(optional.value)
}

func (optional Optional) Filter(fn FilterFunc) (*Optional, error) {
	if fn == nil {
		return nil, errors.New("'fn' cannot be nil")
	}
	if optional.value == nil {
		return Empty(), nil
	}
	if fn(optional.value) {
		return Of(optional.value)
	} else {
		return Empty(), nil
	}
}

func (optional Optional) Map(fn MapFunc) (*Optional, error) {
	if fn == nil {
		return nil, errors.New("'fn' cannot be nil")
	}
	if optional.value == nil {
		return Empty(), nil
	}
	newVal := fn(optional.value)
	if newVal == nil {
		return Empty(), nil
	} else {
		return Of(newVal)
	}
}

func (optional Optional) OrElse(elseVal interface{}) interface{} {
	if optional.value == nil {
		return elseVal
	}
	return optional.value
}

func (optional Optional) OrElseGet(fn GetFunc) (interface{}, error) {
	if optional.value != nil {
		return optional.value, nil
	}
	if fn == nil {
		return nil, errors.New("'fn' cannot be nil")
	}
	return fn(), nil
}

func (optional Optional) OrElseThrow(fn ErrorFunc) (interface{}, error) {
	if optional.value != nil {
		return optional.value, nil
	}
	if fn == nil {
		return nil, errors.New("'fn' cannot be nil")
	}
	return nil, fn()
}
