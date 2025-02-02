package main

import (
	"fmt"
	"github.com/edgardcham/go-pokedex/internal/pokecache"
	"testing"
	"time"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello world    ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("actual %v != expected %v", actual, c.expected)
			t.FailNow()
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("actual %v != expected %v", word, expectedWord)
				t.FailNow()
			}
		}
	}

}

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := pokecache.NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key %s", c.key)
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected value %s, got %s", c.val, val)
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	// waitTime must be longer than baseTime to allow reaping to occur
	const waitTime = baseTime + 5*time.Millisecond

	cache := pokecache.NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	// Initially, the entry should be present.
	if _, ok := cache.Get("https://example.com"); !ok {
		t.Errorf("expected to find key before reaping")
		return
	}

	// Wait for longer than the interval so that the entry should be reaped.
	time.Sleep(waitTime)

	// Now the entry should be gone.
	if _, ok := cache.Get("https://example.com"); ok {
		t.Errorf("expected key to have been removed by reapLoop")
		return
	}
}
