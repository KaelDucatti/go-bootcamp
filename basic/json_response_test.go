package basic_test

import (
	"bytes"
	"testing"

	. "github.com/KaelDucatti/go-bootcamp/basic"
	"github.com/stretchr/testify/require"
)

func TestJsonResponse(t *testing.T) {
	t.Run("should return HTTP Status Response", func(t *testing.T) {
		require := require.New(t)
		buffer := &bytes.Buffer{}
		err := JsonResponse(buffer)

		expected := "Hello, Go Standart Library\nHTTP Response Status: 404 Not Found\n"
		actual := buffer.String()

		require.NoError(err)
		require.Equal(expected, actual)
	})
}