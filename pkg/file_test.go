package pkg_test

import (
	"bjss-go-training/pkg"
	"testing"
)

type testWriteFiles struct {
	input    []string
	fileName string
}

type testReadFiles struct {
	fileName string
}

func TestWriteStringsToFile(t *testing.T) {
	testCases := []testWriteFiles{
		{
			input:    []string{"Hello", "world", "from", "testcase1"},
			fileName: "file_test_case_1.txt",
		},
		{
			input:    []string{"Goodbye", "world", "from", "testcase2"},
			fileName: "file_test_case_2.txt",
		},
	}

	for _, tc := range testCases {
		err := pkg.WriteStringsToFile(tc.input, tc.fileName)
		if err != nil {
			t.Errorf("Expected no error but got %v", err)
		}
	}
}

func TestReadStringsFromFile(t *testing.T) {
	testCases := []testReadFiles{
		{
			fileName: "file_test_case_1.txt",
		},
		{
			fileName: "file_test_case_2.txt",
		},
	}

	for _, tc := range testCases {
		_, err := pkg.ReadStringsFromFile(tc.fileName)
		if err != nil {
			t.Errorf("Expected no error but got %v", err)
		}
	}
}

// Error scenarios
func TestWriteStringsToFileError(t *testing.T) {
	err := pkg.WriteStringsToFile(nil, "")
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

func TestReadStringsFromFileError(t *testing.T) {
	_, err := pkg.ReadStringsFromFile("file_test_case_3.txt")
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}
