package value

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
)

type customInt struct {
	Int int
}

func (v *customInt) IntValue() int {
	return v.Int
}

func (v *customInt) ExprValue() any {
	return v.Int
}

type customTypedInt struct {
	Int int
}

func (v *customTypedInt) IntValue() int {
	return v.Int
}

type customUntypedInt struct {
	Int int
}

func (v *customUntypedInt) ExprValue() any {
	return v.Int
}

type customString struct {
	String string
}

func (v *customString) StringValue() string {
	return v.String
}

func (v *customString) ExprValue() any {
	return v.String
}

type customTypedString struct {
	String string
}

func (v *customTypedString) StringValue() string {
	return v.String
}

type customUntypedString struct {
	String string
}

func (v *customUntypedString) ExprValue() any {
	return v.String
}

type customTypedArray struct {
	Array []any
}

func (v *customTypedArray) ArrayValue() []any {
	return v.Array
}

type customTypedMap struct {
	Map map[string]any
}

func (v *customTypedMap) MapValue() map[string]any {
	return v.Map
}

func Test_valueAddInt(t *testing.T) {
	env := make(map[string]any)
	env["ValueOne"] = &customInt{1}
	env["ValueTwo"] = &customInt{2}

	program, err := expr.Compile("ValueOne + ValueTwo", expr.Env(env), Patcher)
	require.NoError(t, err)

	out, err := vm.Run(program, env)

	require.NoError(t, err)
	require.Equal(t, 3, out.(int))
}

func Test_valueUntypedAddInt(t *testing.T) {
	env := make(map[string]any)
	env["ValueOne"] = &customUntypedInt{1}
	env["ValueTwo"] = &customUntypedInt{2}

	program, err := expr.Compile("ValueOne + ValueTwo", expr.Env(env), Patcher)
	require.NoError(t, err)

	out, err := vm.Run(program, env)

	require.NoError(t, err)
	require.Equal(t, 3, out.(int))
}

func Test_valueTypedAddInt(t *testing.T) {
	env := make(map[string]any)
	env["ValueOne"] = &customTypedInt{1}
	env["ValueTwo"] = &customTypedInt{2}

	program, err := expr.Compile("ValueOne + ValueTwo", expr.Env(env), Patcher)
	require.NoError(t, err)

	out, err := vm.Run(program, env)

	require.NoError(t, err)
	require.Equal(t, 3, out.(int))
}

func Test_valueTypedAddMismatch(t *testing.T) {
	env := make(map[string]any)
	env["ValueOne"] = &customTypedInt{1}
	env["ValueTwo"] = &customTypedString{"test"}

	_, err := expr.Compile("ValueOne + ValueTwo", expr.Env(env), Patcher)
	require.Error(t, err)
}

func Test_valueUntypedAddMismatch(t *testing.T) {
	env := make(map[string]any)
	env["ValueOne"] = &customUntypedInt{1}
	env["ValueTwo"] = &customUntypedString{"test"}

	program, err := expr.Compile("ValueOne + ValueTwo", expr.Env(env), Patcher)
	require.NoError(t, err)

	_, err = vm.Run(program, env)

	require.Error(t, err)
}

func Test_valueTypedArray(t *testing.T) {
	env := make(map[string]any)
	env["ValueOne"] = &customTypedArray{[]any{1, 2}}

	program, err := expr.Compile("ValueOne[0] + ValueOne[1]", expr.Env(env), Patcher)
	require.NoError(t, err)

	out, err := vm.Run(program, env)

	require.NoError(t, err)
	require.Equal(t, 3, out.(int))
}

func Test_valueTypedMap(t *testing.T) {
	env := make(map[string]any)
	env["ValueOne"] = &customTypedMap{map[string]any{"one": 1, "two": 2}}

	program, err := expr.Compile("ValueOne.one + ValueOne.two", expr.Env(env), Patcher)
	require.NoError(t, err)

	out, err := vm.Run(program, env)

	require.NoError(t, err)
	require.Equal(t, 3, out.(int))
}