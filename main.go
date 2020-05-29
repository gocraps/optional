package main

import (
	"errors"
	"fmt"

	"github.com/gocraps/optional.git/pkg/optional"
)

func main() {
	intValue := getInt().Get()
	fmt.Println(intValue)

	stringValue := getString().Get()
	fmt.Println(stringValue)

	defaultValue := getStringDefaultValue().GetOrDefault("Hello World")
	fmt.Println(defaultValue)

	getOrExecuteOnError().GetOrExecute(func(op optional.Optional) {
		fmt.Println(op.Error.Error())
	})
}

func getInt() optional.Optional {
	op := optional.New()
	op.Set(100, nil)
	return op
}

func getString() optional.Optional {
	op := optional.New()
	op.Set("Hello Optional", nil)
	return op
}

func getStringDefaultValue() optional.Optional {
	op := optional.New()
	return op
}

func getOrExecuteOnError() optional.Optional {
	op := optional.New()
	op.Error = errors.New("Error in doing something good")
	return op
}
