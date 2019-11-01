package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newChunkedCounter(test *testing.T) {
	got := newChunkedCounter(23)
	assert.Equal(test, chunkedCounter{step: 23}, got)
}

func Test_chunkedCounter_isOver(test *testing.T) {
	type fields struct {
		step    uint64
		current uint64
		final   uint64
	}

	for _, data := range []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "isn't over",
			fields: fields{
				step:    23,
				current: 42,
				final:   65,
			},
			want: false,
		},
		{
			name: "is over",
			fields: fields{
				step:    23,
				current: 65,
				final:   65,
			},
			want: true,
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			counter := chunkedCounter{
				step:    data.fields.step,
				current: data.fields.current,
				final:   data.fields.final,
			}
			got := counter.isOver()

			assert.Equal(test, data.want, got)
		})
	}
}

func Test_chunkedCounter_increase(test *testing.T) {
	counter := chunkedCounter{current: 23}
	previous := counter.increase()

	assert.Equal(test, chunkedCounter{current: 24}, counter)
	assert.Equal(test, uint64(23), previous)
}

func Test_chunkedCounter_reset(test *testing.T) {
	counter := chunkedCounter{step: 23, current: 42, final: 65}
	counter.reset(100)

	assert.Equal(test, chunkedCounter{step: 23, current: 100, final: 123}, counter)
}
