package optional

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_GivenNoObjectPresent_WhenNewExecuted_ThenEmptyObjectCreated(t *testing.T) {
	//act
	op := New()

	//assert
	assert.NotNil(t, op)
	assert.Empty(t, op)
}

func TestUnit_GivenSomeValue_WhenSetExecuted_ThenValueAndErrorNotNil(t *testing.T) {
	//arrange
	op := New()
	value := "Hello GoCrapes"
	err := errors.New("Do not scare")

	//act
	op.Set(value, err)

	//assert
	assert.NotNil(t, op.value)
	assert.NotNil(t, op.Error)
}

func TestUnit_GivenTypeInt_WhenValueSet_ThenGetReturnsIntValue(t *testing.T) {
	//arrange
	randomInt := rand.Intn(100)

	//act
	op := Optional{randomInt, nil}

	//assert
	assert.Equal(t, randomInt, op.Get())
}

func TestUnit_GivenTypeString_WhenValueSet_ThenGetReturnsStringValue(t *testing.T) {
	//arrange
	aStrignValue := "This is a string"

	//act
	op := Optional{aStrignValue, nil}

	//assert
	assert.Equal(t, aStrignValue, op.Get())
}

func TestUnit_GivenNoValue_WhenNoDeafult_ThenGetReturnsNil(t *testing.T) {
	//act
	op := New()

	//assert
	assert.Nil(t, op.Get())
}

func TestUnit_GivenNoValuePresent_WhenDefaultValuePassed_ThenGetOrDefaultReturnsDefault(t *testing.T) {
	//arrange
	deafultValue := "GoCraps"
	//act
	op := Optional{}

	//assert
	assert.Equal(t, deafultValue, op.GetOrDefault(deafultValue))
}

func TestUnit_GivenValuePresent_WhenDefaultValuePassed_ThenGetOrDefaultReturnsValue(t *testing.T) {
	//arrange
	deafultValue := "GoCraps"
	value := "Optional"
	//act
	op := Optional{value, nil}

	//assert
	assert.Equal(t, value, op.GetOrDefault(deafultValue))
}

func TestUnit_GivenNoValuePresent_WhenFunctionPassed_ThenGetOrExecuteRunsFunction(t *testing.T) {
	//arrange
	var someString string
	expected := "Hello GoCraps"
	op := New()

	//act
	op.GetOrExecute(func(op Optional) {
		someString = expected
	})

	//assert
	assert.Equal(t, expected, someString)
}

func TestUnit_GivenValuePresent_WhenFunctionPassed_ThenGetOrExecuteReturnsValue(t *testing.T) {
	//arrange
	someString := "Initial"
	expected := "Hello GoCraps"
	op := Optional{expected, nil}

	//act
	op.GetOrExecute(func(op Optional) {
		someString = expected
	})

	//assert
	assert.NotEqual(t, expected, someString)
}
