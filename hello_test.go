package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}
	t.Run("hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("'Hello World' when no name is given", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("hello to people in spanish", func(t *testing.T) {
		got := Hello("Chris", "Spanish")
		want := "Hola, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("'Hola World' when no name is given and lang is spanish", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("hello to people in french", func(t *testing.T) {
		got := Hello("Chris", "French")
		want := "Bonjour, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("'Hola World' when no name is given and lang is french", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, World"
		assertCorrectMessage(t, got, want)
	})
}
