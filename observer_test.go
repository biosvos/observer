package observer

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestObserverNew(t *testing.T) {
	observer := NewObserver()

	require.NotNil(t, observer)
}

func TestObserverUpdate(t *testing.T) {
	t.Run("", func(t *testing.T) {
		subject := NewSubject()
		observer := NewObserver()
		subject.RegisterObserver(observer)

		subject.Set(1)

		require.Equal(t, 1, observer.value)
	})

	t.Run("", func(t *testing.T) {
		subject := NewSubject()
		observer := NewObserver()
		subject.RegisterObserver(observer)

		subject.Set(2)

		require.Equal(t, 2, observer.value)
	})

	t.Run("multiple", func(t *testing.T) {
		subject := NewSubject()
		observer1 := NewObserver()
		observer2 := NewObserver()
		subject.RegisterObserver(observer1)
		subject.RegisterObserver(observer2)

		subject.Set(1)

		require.Equal(t, 1, observer1.value)
		require.Equal(t, 1, observer2.value)
	})
}
