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

TestMain -- Before and After Test
use parameter m *testing.M -> m.Run()

TestSubTest --
use parameter t *testing.T -> t.Run()
can testing many test function

Table Test
use slice struct
how to declare
var := []struct {
	field
}{
	{
		field : value,
	}
}

after that, iterate table test
for _, v := range var {
	t.Run(name string, func (t *testing.T) bool {
		result := ()
		assert or require
	})
}


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
	if result != "aka" {
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
		{
			name:     "Third Testing for Simple Regex",
			request:  "I like Go! why? because it's simple",
			expected: "I like Go why because its simple",
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
			expected: [][]int{{0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}},
		},
	}

	for _, test := range testsPair {
		t.Run(test.name, func(t *testing.T) {
			result := PairSumSorted(test.req1, test.req2)
			require.Equal(t, test.expected, result, "Testing Done")
		})
	}
}

func TestShiftZeroToEnd(t *testing.T) {
	tests := []struct {
		name     string
		request  []int
		expected []int
	}{
		{
			name:     "First Testing for shift zero",
			request:  []int{0, 1, 0, 3, 2},
			expected: []int{1, 3, 2, 0, 0},
		},
		{
			name:     "Second Testing for shift zero",
			request:  []int{0, 0, 4, 6, 0, 1, 0, 3, 2},
			expected: []int{4, 6, 1, 3, 2, 0, 0, 0, 0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ShiftZeroToEnd(test.request)
			require.Equal(t, test.expected, result, "Testing Done")
		})
	}
}

func TestCheckBiggerNumber(t *testing.T) {
	tests := []struct {
		name     string
		request  []int
		expected int
	}{
		{
			name:     "first testing",
			request:  []int{4, 2, 5},
			expected: 5,
		},
		{
			name:     "second testing",
			request:  []int{1, 5, 2, 6},
			expected: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CheckBiggerNumber(test.request...)
			assert.Equal(t, test.expected, result, "Testing Done")
		})
	}
}
