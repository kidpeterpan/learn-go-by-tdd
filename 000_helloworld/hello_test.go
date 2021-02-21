package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("Should say hello to deku", func(t *testing.T) {
		expected := "Hello, Deku"
		name := "Deku"

		actual := Hello("",name)

		assert.Equal(t,expected,actual,"want %q, got %q",expected,actual)
	})

	t.Run("Should say hello world when input with empty name", func(t *testing.T) {
		expected := "Hello, world"
		name := ""

		actual := Hello("",name)

		assert.Equal(t,expected,actual,"want %q, got %q",expected,actual)
	})

	t.Run("When talk with english language should say Hello, xxx", func(t *testing.T) {
		expected := "Hello, Deku"
		language := "english"
		name := "Deku"

		actual := Hello(language,name)

		assert.Equal(t,expected,actual,"want %q, got %q",expected,actual)
	})

	t.Run("When talk with thai language should say สวัสดี", func(t *testing.T) {
		expected := "สวัสดี, Deku"
		language := "thai"
		name := "Deku"

		actual := Hello(language,name)

		assert.Equal(t,expected,actual, "want %q, got %q",expected,actual)
	})

	t.Run("When not specific language should talk in english", func(t *testing.T) {
		expected := "Hello, Deku"
		language := ""
		name := "Deku"

		actual := Hello(language,name)

		assert.Equal(t,expected,actual, "want %q, got %q",expected,actual)
	})

}
