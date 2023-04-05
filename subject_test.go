package observer

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSubjectNew(t *testing.T) {
	subject := NewSubject()

	require.NotNil(t, subject)
}
