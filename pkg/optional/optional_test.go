package optional

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnit_Of_WHEN_NonNilValue_THEN_ReturnsOptional(t *testing.T) {
	//act
	op, err := Of(10)

	//assert
	assert.NoError(t, err)
	assert.Equal(t, 10, op.value)
}

func TestUnit_Of_WHEN_NilValue_THEN_ReturnsError(t *testing.T) {
	//act
	op, err := Of(nil)

	//assert
	assert.Error(t, err)
	assert.Nil(t, op)
}

func TestUnit_OfNillable_WHEN_NonNilValue_THEN_ReturnsOptional(t *testing.T) {
	//act
	op := OfNillable(10)

	//assert
	assert.NotNil(t, op)
	assert.Equal(t, 10, op.value)
}

func TestUnit_OfNillable_WHEN_NilValue_THEN_ReturnsOptional(t *testing.T) {
	//act
	op := OfNillable(nil)

	//assert
	assert.NotNil(t, op)
	assert.Nil(t, op.value)
}

func TestUnit_Empty_THEN_ReturnsOptional(t *testing.T) {
	//act
	op := Empty()

	//assert
	assert.NotNil(t, op)
	assert.Nil(t, op.value)
}

func TestUnit_Get_WHEN_NonNilValue_THEN_ReturnsValue(t *testing.T) {
	//arrange
	op := Optional{
		value: 10,
	}

	//act
	val, err := op.Get()

	//assert
	assert.NoError(t, err)
	assert.Equal(t, 10, val)
}

func TestUnit_Get_WHEN_NilValue_THEN_ReturnsError(t *testing.T) {
	//arrange
	op := Optional{
		value: nil,
	}

	//act
	val, err := op.Get()

	//assert
	assert.Error(t, err)
	assert.Nil(t, val)
}

func TestUnit_IsPresent_WHEN_NilValue_THEN_ReturnsFalse(t *testing.T) {
	//arrange
	op := Optional{
		value: nil,
	}

	//act
	val := op.IsPresent()

	//assert
	assert.False(t, val)
}

func TestUnit_IsPresent_WHEN_NonNilValue_THEN_ReturnsTrue(t *testing.T) {
	//arrange
	op := Optional{
		value: "abcd",
	}

	//act
	val := op.IsPresent()

	//assert
	assert.True(t, val)
}

func TestUnit_IfPresent_WHEN_NonNilValue_THEN_ExecuteFunctionAndReturnNoError(t *testing.T) {
	//arrange
	op := Optional{
		value: "abcd",
	}

	//act
	err := op.IfPresent(func(val interface{}) error {
		return nil
	})

	//assert
	assert.NoError(t, err)
}

func TestUnit_IfPresent_WHEN_NilValue_THEN_ExecuteFunctionAndReturnNoError(t *testing.T) {
	//arrange
	op := Optional{
		value: nil,
	}

	//act
	err := op.IfPresent(func(val interface{}) error {
		return nil
	})

	//assert
	assert.NoError(t, err)
}

func TestUnit_IfPresent_WHEN_NonNilValueAndNilFunction_THEN_ReturnsError(t *testing.T) {
	//arrange
	op := Optional{
		value: "abcd",
	}

	//act
	err := op.IfPresent(nil)

	//assert
	assert.Error(t, err)
}
func TestUnit_IfPresent_WHEN_NilValueAndNilFUnction_THEN_ReturnNoError(t *testing.T) {
	//arrange
	op := Optional{
		value: nil,
	}

	//act
	err := op.IfPresent(nil)

	//assert
	assert.NoError(t, err)
}

func TestUnit_Filter_WHEN_NilFunction_THEN_ReturnError(t *testing.T) {
	//arrange
	op := Optional{
		value: "abcd",
	}

	//act
	val, err := op.Filter(nil)

	//assert
	assert.Error(t, err)
	assert.Nil(t, val)
}

func TestUnit_Filter_WHEN_NilValueAndNonNilFunction_THEN_ReturnEmpty(t *testing.T) {
	//arrange
	op := Optional{
		value: nil,
	}

	//act
	val, err := op.Filter(func(obj interface{}) bool {
		return true
	})

	//assert
	assert.NoError(t, err)
	assert.Equal(t, &Optional{
		value: nil,
	}, val)
}

func TestUnit_Filter_WHEN_NonNilValueAndNonNilFunctionAndFunctionReturnsFalse_THEN_ReturnEmpty(t *testing.T) {
	//arrange
	op := Optional{
		value: "abcd",
	}

	//act
	val, err := op.Filter(func(obj interface{}) bool {
		return false
	})

	//assert
	assert.NoError(t, err)
	assert.Equal(t, &Optional{
		value: nil,
	}, val)
}

func TestUnit_Filter_WHEN_NonNilValueAndNonNilFunctionAndFunctionReturnsTrue_THEN_ReturnsOptional(t *testing.T) {
	//arrange
	op := Optional{
		value: "abcd",
	}

	//act
	val, err := op.Filter(func(obj interface{}) bool {
		return true
	})

	//assert
	assert.NoError(t, err)
	assert.Equal(t, &Optional{
		value: "abcd",
	}, val)
}

func TestUnit_Map_WHEN_NilFunction_THEN_ReturnError(t *testing.T) {
	//arrange
	op := Optional{
		value: "abcd",
	}

	//act
	val, err := op.Map(nil)

	//assert
	assert.Error(t, err)
	assert.Nil(t, val)
}

func TestUnit_Map_WHEN_NilValueAndNonNilFunction_THEN_ReturnEmpty(t *testing.T) {
	//arrange
	op := Optional{
		value: nil,
	}

	//act
	val, err := op.Map(func(obj interface{}) interface{} {
		return true
	})

	//assert
	assert.NoError(t, err)
	assert.Equal(t, &Optional{
		value: nil,
	}, val)
}

func TestUnit_Map_WHEN_NonNilValueAndNonNilFunctionAndFunctionReturnsNil_THEN_ReturnEmpty(t *testing.T) {
	//arrange
	op := Optional{
		value: "abcd",
	}

	//act
	val, err := op.Map(func(obj interface{}) interface{} {
		return nil
	})

	//assert
	assert.NoError(t, err)
	assert.Equal(t, &Optional{
		value: nil,
	}, val)
}

func TestUnit_Map_WHEN_NonNilValueAndNonNilFunctionAndFunctionReturnsValue_THEN_ReturnsOptional(t *testing.T) {
	//arrange
	op := Optional{
		value: "abcd",
	}

	//act
	val, err := op.Map(func(obj interface{}) interface{} {
		return 10
	})

	//assert
	assert.NoError(t, err)
	assert.Equal(t, &Optional{
		value: 10,
	}, val)
}

func TestUnit_OrElse_WHEN_NonNilValue_THEN_ReturnsOrignialValue(t *testing.T) {
	//arrange
	op := Optional{
		value: "abcd",
	}

	//act
	val := op.OrElse(10)

	//assert
	assert.Equal(t, "abcd", val)
}

func TestUnit_OrElse_WHEN_NilValue_THEN_ReturnsElseValue(t *testing.T) {
	//arrange
	op := Optional{
		value: nil,
	}

	//act
	val := op.OrElse(10)

	//assert
	assert.Equal(t, 10, val)
}

func TestUnit_OrElseGet_WHEN_NonNilValue_THEN_ReturnsOriginalValue(t *testing.T) {
	//arrange
	op := Optional{
		value: "abcd",
	}

	//act
	val, err := op.OrElseGet(func() interface{} {
		return 10
	})

	//assert
	assert.NoError(t, err)
	assert.Equal(t, "abcd", val)
}

func TestUnit_OrElseGet_WHEN_NilValueAndNilFunction_THEN_ReturnsError(t *testing.T) {
	//arrange
	op := Optional{
		value: nil,
	}

	//act
	val, err := op.OrElseGet(nil)

	//assert
	assert.Error(t, err)
	assert.Nil(t, val)
}

func TestUnit_OrElseGet_WHEN_NilValue_THEN_ReturnsValueReturnedByFunction(t *testing.T) {
	//arrange
	op := Optional{
		value: nil,
	}

	//act
	val, err := op.OrElseGet(func() interface{} {
		return 10
	})

	//assert
	assert.NoError(t, err)
	assert.Equal(t, 10, val)
}

func TestUnit_OrElseThrow_WHEN_NonNilValue_THEN_ReturnsOriginalValue(t *testing.T) {
	//arrange
	op := Optional{
		value: "abcd",
	}

	//act
	val, err := op.OrElseThrow(func() error {
		return errors.New("New Error")
	})

	//assert
	assert.NoError(t, err)
	assert.Equal(t, "abcd", val)
}

func TestUnit_OrElseThrow_WHEN_NilValueAndNilFunction_THEN_ReturnsError(t *testing.T) {
	//arrange
	op := Optional{
		value: nil,
	}

	//act
	val, err := op.OrElseThrow(nil)

	//assert
	assert.Error(t, err)
	assert.Equal(t, "'fn' cannot be nil", err.Error())
	assert.Nil(t, val)
}

func TestUnit_OrElseGet_WHEN_NilValueAndNonNillErrorFunction_THEN_ReturnsErrorReturnedByFunction(t *testing.T) {
	//arrange
	op := Optional{
		value: nil,
	}

	//act
	val, err := op.OrElseThrow(func() error {
		return errors.New("New Error")
	})

	//assert
	assert.Error(t, err)
	assert.Equal(t, "New Error", err.Error())
	assert.Nil(t, val)
}
