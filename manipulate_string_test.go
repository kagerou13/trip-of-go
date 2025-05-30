package tripofgo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
Fail() --> still execute after run t.Fail()
FailNow() --> directly stop after run t.FailNow()
Fail()  vs FailNow()
still      stop

Fatal() == Fail()  --> Fatal display detail from failed test reason
Error() == FailNow() --> Error display detail from failed test reason
Fatal() vs Error()

assert  vs require
assert  --> display spesific detail from failed test
require --> display spesific detail from failed test

*/

func TestTable(t *testing.T) {
	tests := []struct {
		name      string
		request   string
		expected  string
		request2  string
		expected2 string
	}{
		{
			name:     "Reverse String",
			request:  "kagerou",
			expected: "uoregak",
		},
		{
			name:     "valhal",
			request:  "valha",
			expected: "ahlav",
		},
		{
			name:      "Palindrome test",
			request2:  "lol",
			expected2: "lol",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ReverseString(test.request)
			require.Equal(t, test.expected, result, "Testing Done with table test")
		})
		t.Run(test.name, func(t *testing.T) {
			result2 := Palindrome(test.request2)
			assert.Equal(t, test.expected2, result2, "Testing Palindrome done with table test")
		})
	}
}

func TestReverseString(t *testing.T) {
	result := ReverseString("kagerou") // input and output string
	if result != "uoregak" {
		t.Fail()
	}

	fmt.Println("Testing Reverse String done")
}

func TestCountWords(t *testing.T) {
	result := CountWords("hello world") // input string and output int
	if result != 2 {
		t.FailNow()
	}

	fmt.Println("Testing Count Words done")
}

func TestPalindrome(t *testing.T) {
	result := Palindrome("aka") // input and output string
	if result != "Palindrome" {
		t.Error("String is not palindrome")
	}

	fmt.Println("Testing Palindrome done")
}

func TestChangeWithOtherChar(t *testing.T) {
	result := ChangeWithOtherChar("k", "m", "kagerou")
	if result != "magerou" {
		t.Fatal("Error occured while changing")
	}

	fmt.Println("Testing ChangeWithOtherChar done")
}

func TestChangeWithOthers(t *testing.T) {
	result := ChangeWithOthers("ka", "mi", "kagerou")
	assert.Equal(t, "migerou", result, "Result must be migerou")
	fmt.Println("Testing done")
}

func TestCountsFrecWords(t *testing.T) {
	result := CountsFrecWords("I am a beginner Golang developer")
	require.Equal(t, map[string]int{"Golang": 1, "I": 1, "a": 1, "am": 1, "beginner": 1, "developer": 1}, result, "Result is on map data")
	fmt.Println("Testing done")
}

// func TestMain for before and after test

func TestMain(m *testing.M) {
	fmt.Println("BEFORE TEST")

	m.Run()

	fmt.Println("AFTER TEST")
}

func TestSubTest(t *testing.T) {
	t.Run("Reverse String", func(t *testing.T) {
		result := ReverseString("kagerou")
		if result != "uoregak" {
			t.Fail()
		}

		fmt.Println("Testing Reverse String done")
	})
	t.Run("Palindrome", func(t *testing.T) {
		result := Palindrome("aka")
		if result != "Palindrome" {
			t.Error("String is not palindrome")
		}

		fmt.Println("Testing Palindrome done")
	})
}

func TestSimpleRegex(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "First Testing for Simple Regex",
			request:  "I'm Golang Developer!",
			expected: "Im Golang Developer",
		},
		{
			name:     "Second Testing Simple Regex",
			request:  "Hello Guys! What's up?",
			expected: "Hello Guys Whats up",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SimpleRegex(test.request)
			require.Equal(t, test.expected, result, "Testing Simple Regex Done")
		})
	}
}

func TestPairSumSorted(t *testing.T) {
	testsPair := []struct {
		name     string
		req1     []int
		req2     int
		expected [][]int
	}{
		{
			name:     "First Testing for Pair Sum Sorted",
			req1:     []int{-5, -2, 3, 4, 6},
			req2:     7,
			expected: [][]int{{2, 3}},
		},
		{
			name:     "Second Testing for Pair Sum Sorted",
			req1:     []int{1, 1, 1},
			req2:     2,
			expected: [][]int{[]int{0, 1}, []int{0, 2}, []int{1, 0}, []int{1, 2}, []int{2, 0}, []int{2, 1}},
		},
	}

	for _, test := range testsPair {
		t.Run(test.name, func(t *testing.T) {
			result := PairSumSorted(test.req1, test.req2)
			require.Equal(t, test.expected, result, "Testing Done")
		})
	}
}
