package leftright

import (
	"testing"
)

func TestIndexOfLastChar(t *testing.T) {
	testSentence := "Test sentence passing by, dont mind me. Hey! mind your own business"

	BasicTestIndexOfLastChar(t, []byte("hhh"), testSentence, 4, "")
	BasicTestIndexOfLastChar(t, []byte("111"), testSentence, 1, "")
	BasicTestIndexOfLastChar(t, []byte("hh1"), testSentence, 2, "")
	BasicTestIndexOfLastChar(t, []byte("hh11"), testSentence, 1, "")

	BasicTestIndexOfLastChar(t, []byte("hhsh"), testSentence, -1, "contains unknown command : s")

	BasicTestIndexOfLastChar(t, []byte("w"), testSentence, 6, "")
	BasicTestIndexOfLastChar(t, []byte("wwwwwwwwwwwwwwwwwwww"), "ats as", 5, "")
	BasicTestIndexOfLastChar(t, []byte("wwwwwwwwwwwwwwwwwwww"), testSentence, 60, "")
	BasicTestIndexOfLastChar(t, []byte("wwww"), testSentence, 27, "")
	BasicTestIndexOfLastChar(t, []byte("hhhhw"), testSentence, 6, "")
	BasicTestIndexOfLastChar(t, []byte("whww1"), testSentence, 22, "")

	BasicTestIndexOfLastChar(t, []byte("b"), testSentence, 1, "")
	BasicTestIndexOfLastChar(t, []byte("bbb"), testSentence, 1, "")
	BasicTestIndexOfLastChar(t, []byte("wb"), testSentence, 1, "")
	BasicTestIndexOfLastChar(t, []byte("wwb"), testSentence, 6, "")
	BasicTestIndexOfLastChar(t, []byte("wwwb"), testSentence, 15, "")
	BasicTestIndexOfLastChar(t, []byte("wwwbb"), testSentence, 6, "")

	BasicTestIndexOfLastChar(t, []byte("hhhbh"), testSentence, 2, "")
	BasicTestIndexOfLastChar(t, []byte("hhhhwwwbb"), testSentence, 6, "")
	BasicTestIndexOfLastChar(t, []byte("wwwbb11"), testSentence, 4, "")

	BasicTestIndexOfLastCharString(t, "hhh "+testSentence, 4, "")
	BasicTestIndexOfLastCharString(t, "111 "+testSentence, 1, "")
	BasicTestIndexOfLastCharString(t, "hhs1 "+testSentence, -1, "contains unknown command : s")
	BasicTestIndexOfLastCharString(t, "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"+testSentence, -1, "too much command")
	BasicTestIndexOfLastCharString(t, "hh11 "+testSentence+testSentence+testSentence+testSentence+testSentence+testSentence+testSentence+testSentence+testSentence+testSentence+testSentence+testSentence, -1, "too much word")
}

func BasicTestIndexOfLastChar(t *testing.T, commands []byte, sentence string, expected int, expectedError string) {
	if output, err := IndexOfLastChar(commands, sentence); output != expected || err != nil {
		if expectedError == "" || err.Error() != expectedError {
			t.Errorf("Test failed with sentence : \n\t\t%s\ncommands : \n\t\t%s\n Expected %v but got %v\nand expected error :\n\t\t%s \nerror: \n\t\t%v", sentence, commands, expected, output, expectedError, err)
		}
	}
}

func BasicTestIndexOfLastCharString(t *testing.T, s string, expectedOutput int, expectedError string) {
	if output, err := IndexOfLastCharString(s); output != expectedOutput || err != nil {
		if expectedError == "" || err.Error() != expectedError {
			t.Errorf("Test failed with string : \n\t\t%s\nExpected Output : \n\t\t%v \nExpected Error : \n\t\t%s\nbut got \n\t\t%v \nAnd Error \n\t\t%v", s, expectedOutput, expectedError, output, err)
		}
	}
}
