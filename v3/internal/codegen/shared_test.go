package codegen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnakeCaseToCamelCase(t *testing.T) {
	const input = "snake_case_name"
	const expected = "snakeCaseName"

	actual := snakeCaseToCamelCase(input)

	assert.Equal(t, expected, actual)
}

func TestSnakeCaseToCamelCaseLeadingUnderscore(t *testing.T) {
	const input = "_snake_case_name"
	const expected = "snakeCaseName"

	actual := snakeCaseToCamelCase(input)

	assert.Equal(t, expected, actual)
}

func TestSnakeCaseToCamelCaseTrailingUnderscore(t *testing.T) {
	const input = "snake_case_name_"
	const expected = "snakeCaseName"

	actual := snakeCaseToCamelCase(input)

	assert.Equal(t, expected, actual)
}

func TestSnakeCaseToCamelCaseTrailingUnderscoreSingleChar(t *testing.T) {
	const input = "snake_case_name_t"
	const expected = "snakeCaseNameT"

	actual := snakeCaseToCamelCase(input)

	assert.Equal(t, expected, actual)
}

func TestSnakeCaseToPascalCase(t *testing.T) {
	const input = "snake_case_name"
	const expected = "SnakeCaseName"

	actual := snakeCaseToPascalCase(input)

	assert.Equal(t, expected, actual)
}

func TestSnakeCaseToPascalCaseLeadingUnderscore(t *testing.T) {
	const input = "_snake_case_name"
	const expected = "SnakeCaseName"

	actual := snakeCaseToPascalCase(input)

	assert.Equal(t, expected, actual)
}

func TestSnakeCaseToPascalCaseTrailingUnderscore(t *testing.T) {
	const input = "snake_case_name_"
	const expected = "SnakeCaseName"

	actual := snakeCaseToPascalCase(input)

	assert.Equal(t, expected, actual)
}

func TestSnakeCaseToPascalCaseTrailingUnderscoreSingleChar(t *testing.T) {
	const input = "snake_case_name_t"
	const expected = "SnakeCaseNameT"

	actual := snakeCaseToPascalCase(input)

	assert.Equal(t, expected, actual)
}
