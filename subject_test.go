package observer

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSubjectNew(t *testing.T) {
	subject := NewSubject(1)

	require.NotNil(t, subject)
}
