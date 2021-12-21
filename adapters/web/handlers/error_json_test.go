package handlers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "Hello json"
	result := jsonError(msg)

	require.Equal(t, []byte(`{"message":"Hello json"}`), result)
}

