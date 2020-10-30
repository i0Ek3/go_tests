package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		/*
			// DO NOT USE ANONYMOUS NESTED
			// STRUCT IN THE COMPLICATED CODE
			{
				"use pointer",
				&Person{
					"i0Ek3"
					Profile{30, "Tokoy"}
				},
				[]string{"i0Ek3", "Tokoy"},
			},
			{
				"nested anonymous struct",
				Person{
					"Satoshi",
					Profile{30, "Tokoy"},
				},
				[]string{"Satoshi", "Tokoy"},
			},
			{
				"slices test",
				[]Profile{
					{33, "Paris"},
					{32, "Berlin"},
				},
				[]string{"Paris", "Berlin"},
			},
			{
				"arrays test",
				[2]Profile{
					{33, "Paris"},
					{32, "Berlin"},
				},
				[]string{"Paris", "Berlin"},
			},*/
		{
			"maps test",
			map[int]string{
				1: "tokoy",
				2: "paris",
			},
			[]string{"tokoy", "paris"},
		},
	}

	var got []string
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			assertReflect(t, got, test.ExpectedCalls)
		})

		t.Run("maps test", func(t *testing.T) {
			m := map[int]string{
				1: "tokoy",
				2: "paris",
			}

			walk(m, func(input string) {
				got = append(got, input)
			})

			assertContains(t, got, "tokoy")
			assertContains(t, got, "paris")
		})
	}
}

func assertReflect(t *testing.T, got, expected []string) {
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v but we expected %v", got, expected)
	}
}

func assertContains(t *testing.T, haystack []string, needle string) {
	flag := false
	for _, v := range haystack {
		if v == needle {
			flag = true
		}
	}
	if !flag {
		t.Errorf("got %+v but we expected %s", haystack, needle)
	}
}
