package extra

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/manilion/json-iterator-go"
)

func Test_private_fields(t *testing.T) {
	type TestObject struct {
		field1 string `json:"field_1"`
	}
	SupportPrivateFields()
	should := require.New(t)
	obj := TestObject{}
	should.Nil(jsoniter.UnmarshalFromString(`{"field_1":"Hello"}`, &obj))
	should.Equal("Hello", obj.field1)
}
