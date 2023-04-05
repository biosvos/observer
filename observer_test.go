package observer

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var _ Observer[int] = &intObserver{}

type intObserver struct {
	value int
}

func newIntObserver() *intObserver {
	return &intObserver{}
}

func (i *intObserver) Notify(subject *Subject[int]) {
	i.value = subject.value
}

func TestObserverNew(t *testing.T) {
	observer := newIntObserver()

	require.NotNil(t, observer)
}

func TestObserverUpdate(t *testing.T) {
	t.Run("", func(t *testing.T) {
		subject := NewSubject(0)
		observer := newIntObserver()
		subject.RegisterObserver(observer)

		subject.Set(1)

		require.Equal(t, 1, observer.value)
	})

	t.Run("", func(t *testing.T) {
		subject := NewSubject(0)
		observer := newIntObserver()
		subject.RegisterObserver(observer)

		subject.Set(2)

		require.Equal(t, 2, observer.value)
	})

	t.Run("multiple", func(t *testing.T) {
		subject := NewSubject(0)
		observer1 := newIntObserver()
		observer2 := newIntObserver()
		subject.RegisterObserver(observer1)
		subject.RegisterObserver(observer2)

		subject.Set(1)

		require.Equal(t, 1, observer1.value)
		require.Equal(t, 1, observer2.value)
	})
}
