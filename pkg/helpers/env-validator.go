package helpers

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type EnvValidator struct {
	variableName string
	value        string
	isRequired   bool
}

type OptionsValidator struct {
	Value        string
	VariableName string
}

func NewValidator(options OptionsValidator) *EnvValidator {
	return &EnvValidator{
		value:        options.Value,
		variableName: options.VariableName,
	}
}

func (v *EnvValidator) Required() *EnvValidator {
	v.isRequired = true
	return v
}

func (v *EnvValidator) Text() string {
	if v.isRequired && len(strings.TrimSpace(v.value)) == 0 {
		panic(fmt.Sprintf("Environment variable '%s' is required but not set.", v.variableName))
	}

	return v.value
}

func (v *EnvValidator) Number() int {

	if v.isRequired && v.value == "" {
		panic(fmt.Sprintf("Environment variable '%s' is required but not set.", v.variableName))
	}

	if v.value == "" {
		return 0
	}

	num, err := strconv.Atoi(v.value)

	if err != nil {
		panic(fmt.Sprintf("Environment variable '%s' must be a valid number. Got: '%s'", v.variableName, v.value))
	}

	return num
}

func (v *EnvValidator) Boolean() bool {
	if v.isRequired && v.value == "" {
		panic(fmt.Sprintf("Environment variable '%s' is required but not set.", v.variableName))
	}

	if v.value == "" {
		return false
	}

	boolValue, err := strconv.ParseBool(v.value)

	if err != nil {
		panic(fmt.Sprintf("Environment variable '%s' must be a valid boolean (e.g., true, false, 1, 0). Got: '%s'", v.variableName, v.value))
	}

	return boolValue
}

func (v *EnvValidator) Email() string {
	if v.isRequired && strings.TrimSpace(v.value) == "" {
		panic(fmt.Sprintf("Environment variable '%s' is required but not set.", v.variableName))
	}

	if v.value == "" {
		return ""
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if !emailRegex.MatchString(v.value) {
		panic(fmt.Sprintf("Environment variable '%s' must be a valid email address. Got: '%s'", v.variableName, v.value))
	}

	return v.value
}

func (v *EnvValidator) Default(data any) {
	if v.value != "" {
		return
	}

	switch reflect.TypeOf(data).Kind() {
	case reflect.String:
		v.value = data.(string)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.value = fmt.Sprintf("%d", data)

	case reflect.Bool:
		v.value = fmt.Sprintf("%v", data)

	default:
		panic(fmt.Sprintf("Unsupported default data type: %s", reflect.TypeOf(data)))
	}

}
