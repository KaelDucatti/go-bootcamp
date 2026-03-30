package basic_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	. "github.com/KaelDucatti/go-bootcamp/basic"
)

func TestJsonResponse(t *testing.T) {
	t.Run("Success Cases", func(t *testing.T) {
		t.Run("should return HTTP Status Response", func(t *testing.T) {
			require := require.New(t)
			buffer := &bytes.Buffer{}
			err := Json

		})
	})
}
