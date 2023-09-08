package handler

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestErrorHandler(t *testing.T) {
	msg := "Not Found"
	result := jsonError(msg)

	require.Equal(t, []byte(`{"message":"Not Found"}`), result)
}
