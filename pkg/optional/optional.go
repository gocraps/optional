package optional

type Optional struct {
	value interface{}
	Error error
}

type execute func(Optional)

func New() Optional {
	return Optional{}
}

func (optional *Optional) Set(value interface{}, err error) {
	optional.value = value
	optional.Error = err
}

func (optional Optional) Get() interface{} {
	if optional.Error != nil {
		return nil
	}
	return optional.value
}

func (optional Optional) GetOrDefault(defaultValue interface{}) interface{} {
	if optional.value == nil {
		return defaultValue
	}
	return optional.value
}

func (optional Optional) GetOrExecute(fn execute) interface{} {
	if optional.value == nil {
		fn(optional)
		return nil
	}
	return optional.value
}
