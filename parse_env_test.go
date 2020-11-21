package sec

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseString(t *testing.T) {
	type testStruct struct {
		StringData string
	}

	os.Setenv("STRINGDATA", "test")

	s := &testStruct{}

	err := Parse(s, nil)
	if err != nil {
		t.Log(err.Error())
	}

	require.Nil(t, err)
}

func TestParseBoolean(t *testing.T) {
	testCases := []struct {
		TestData  string
		ValidData bool
	}{
		{"1", true},
		{"t", true},
		{"T", true},
		{"TRUE", true},
		{"true", true},
		{"True", true},
		{"0", false},
		{"f", false},
		{"F", false},
		{"FALSE", false},
		{"false", false},
		{"False", false},
		{"not a boolean", false},
	}

	type testStruct struct {
		BoolData bool
	}

	for _, testCase := range testCases {
		t.Logf("Testing: %+v", testCase)
		os.Setenv("BOOLDATA", testCase.TestData)

		// If ErrorsAreCritical == false, then we should check only
		// equality of parsed data and valid data.
		s := &testStruct{}

		err := Parse(s, nil)
		if err != nil {
			t.Log(err.Error())
		}

		require.Nil(t, err)
		require.Equal(t, testCase.ValidData, s.BoolData)

		// If errors are critical - we should check if test data is within
		// int8 range and check for error if it isn't.
		s1 := &testStruct{}
		err1 := Parse(s1, &Options{ErrorsAreCritical: true})

		var checkNotBoolError bool

		_, err2 := strconv.ParseBool(testCase.TestData)
		if err2 != nil {
			checkNotBoolError = true
		}

		if checkNotBoolError {
			if err1 == nil {
				t.Log("No error returned!")
			}

			require.NotNil(t, err1)
			require.Equal(t, errNotBool, err1)
		}

		os.Unsetenv("BOOLDATA")
	}
}

func TestParseInt8(t *testing.T) {
	testCases := []struct {
		TestData  string
		ValidData int8
	}{
		{"0", 0},
		{"1", 1},
		{"2", 2},
		{"-128", -128},
		{"127", 127},
		{"-129", 0},
		{"128", 0},
		{"not an integer", 0},
	}

	type testStruct struct {
		IntData int8
	}

	for _, testCase := range testCases {
		t.Logf("Testing: %+v", testCase)
		os.Setenv("INTDATA", testCase.TestData)

		// If ErrorsAreCritical == false, then we should check only
		// equality of parsed data and valid data.
		s := &testStruct{}

		err := Parse(s, nil)
		if err != nil {
			t.Log(err.Error())
		}

		require.Nil(t, err)
		require.Equal(t, testCase.ValidData, s.IntData)

		// If errors are critical - we should check if test data is within
		// int8 range and check for error if it isn't.
		s1 := &testStruct{}
		err1 := Parse(s1, &Options{ErrorsAreCritical: true})

		var (
			checkNotIntError bool
			checkRangeError  bool
		)

		passedData, err2 := strconv.ParseInt(testCase.TestData, 10, 64)
		if err2 != nil {
			checkNotIntError = true
		}

		if passedData != int64(testCase.ValidData) {
			checkRangeError = true
		}

		if checkNotIntError || checkRangeError {
			if err1 == nil {
				t.Log("No error returned!")
			}

			require.NotNil(t, err1)

			if checkNotIntError {
				require.Equal(t, errNotInt, err1)
			}

			if checkRangeError {
				require.Equal(t, errNotInt8, err1)
			}
		}

		os.Unsetenv("INTDATA")
	}
}

func TestParseInt16(t *testing.T) {
	testCases := []struct {
		TestData  string
		ValidData int16
	}{
		{"0", 0},
		{"1", 1},
		{"2", 2},
		{"-32768", -32768},
		{"32767", 32767},
		{"-32770", 0},
		{"32770", 0},
		{"not an integer", 0},
	}

	type testStruct struct {
		IntData int16
	}

	for _, testCase := range testCases {
		t.Logf("Testing: %+v", testCase)
		os.Setenv("INTDATA", testCase.TestData)

		// If ErrorsAreCritical == false, then we should check only
		// equality of parsed data and valid data.
		s := &testStruct{}

		err := Parse(s, nil)
		if err != nil {
			t.Log(err.Error())
		}

		require.Nil(t, err)
		require.Equal(t, testCase.ValidData, s.IntData)

		// If errors are critical - we should check if test data is within
		// int8 range and check for error if it isn't.
		s1 := &testStruct{}
		err1 := Parse(s1, &Options{ErrorsAreCritical: true})

		var (
			checkNotIntError bool
			checkRangeError  bool
		)

		passedData, err2 := strconv.ParseInt(testCase.TestData, 10, 64)
		if err2 != nil {
			checkNotIntError = true
		}

		if passedData != int64(testCase.ValidData) {
			checkRangeError = true
		}

		if checkNotIntError || checkRangeError {
			if err1 == nil {
				t.Log("No error returned!")
			}

			require.NotNil(t, err1)

			if checkNotIntError {
				require.Equal(t, errNotInt, err1)
			}

			if checkRangeError {
				require.Equal(t, errNotInt16, err1)
			}
		}

		os.Unsetenv("INTDATA")
	}
}

func TestParseInt32(t *testing.T) {
	testCases := []struct {
		TestData  string
		ValidData int32
	}{
		{"0", 0},
		{"1", 1},
		{"2", 2},
		{"-2147483648", -2147483648},
		{"2147483647", 2147483647},
		{"-2147483650", 0},
		{"2147483650", 0},
		{"not an integer", 0},
	}

	type testStruct struct {
		IntData int32
	}

	for _, testCase := range testCases {
		t.Logf("Testing: %+v", testCase)
		os.Setenv("INTDATA", testCase.TestData)

		// If ErrorsAreCritical == false, then we should check only
		// equality of parsed data and valid data.
		s := &testStruct{}

		err := Parse(s, nil)
		if err != nil {
			t.Log(err.Error())
		}

		require.Nil(t, err)
		require.Equal(t, testCase.ValidData, s.IntData)

		// If errors are critical - we should check if test data is within
		// int8 range and check for error if it isn't.
		s1 := &testStruct{}
		err1 := Parse(s1, &Options{ErrorsAreCritical: true})

		var (
			checkNotIntError bool
			checkRangeError  bool
		)

		passedData, err2 := strconv.ParseInt(testCase.TestData, 10, 64)
		if err2 != nil {
			checkNotIntError = true
		}

		if passedData != int64(testCase.ValidData) {
			checkRangeError = true
		}

		if checkNotIntError || checkRangeError {
			if err1 == nil {
				t.Log("No error returned!")
			}

			require.NotNil(t, err1)

			if checkNotIntError {
				require.Equal(t, errNotInt, err1)
			}

			if checkRangeError {
				require.Equal(t, errNotInt32, err1)
			}
		}

		os.Unsetenv("INTDATA")
	}
}

func TestParseInt64(t *testing.T) {
	testCases := []struct {
		TestData  string
		ValidData int64
	}{
		{"0", 0},
		{"1", 1},
		{"2", 2},
		{"-9223372036854775808", -9223372036854775808},
		{"9223372036854775807", 9223372036854775807},
		{"-9223372036854775810", -9223372036854775808},
		{"9223372036854775810", 9223372036854775807},
		{"not an integer", 0},
	}

	type testStruct struct {
		IntData int64
	}

	for _, testCase := range testCases {
		t.Logf("Testing: %+v", testCase)
		os.Setenv("INTDATA", testCase.TestData)

		// If ErrorsAreCritical == false, then we should check only
		// equality of parsed data and valid data.
		s := &testStruct{}

		err := Parse(s, nil)
		if err != nil {
			t.Log(err.Error())
		}

		require.Nil(t, err)
		require.Equal(t, testCase.ValidData, s.IntData)

		// If errors are critical - we should check if test data is within
		// int8 range and check for error if it isn't.
		s1 := &testStruct{}
		err1 := Parse(s1, &Options{ErrorsAreCritical: true})

		var (
			checkNotIntError bool
			checkRangeError  bool
		)

		passedData, err2 := strconv.ParseInt(testCase.TestData, 10, 64)
		if err2 != nil {
			checkNotIntError = true
		}

		if passedData != testCase.ValidData {
			checkRangeError = true
		}

		if checkNotIntError || checkRangeError {
			if err1 == nil {
				t.Log("No error returned!")
			}

			require.NotNil(t, err1)

			if checkNotIntError {
				require.Equal(t, errNotInt, err1)
			}

			if checkRangeError {
				require.Equal(t, errNotInt64, err1)
			}
		}

		os.Unsetenv("INTDATA")
	}
}

func TestParseUint8(t *testing.T) {
	testCases := []struct {
		TestData  string
		ValidData uint8
	}{
		{"0", 0},
		{"1", 1},
		{"2", 2},
		{"255", 255},
		{"256", 0},
		{"-1", 0},
		{"not an integer", 0},
	}

	type testStruct struct {
		UintData uint8
	}

	for _, testCase := range testCases {
		t.Logf("Testing: %+v", testCase)
		os.Setenv("UINTDATA", testCase.TestData)

		// If ErrorsAreCritical == false, then we should check only
		// equality of parsed data and valid data.
		s := &testStruct{}

		err := Parse(s, nil)
		if err != nil {
			t.Log(err.Error())
		}

		require.Nil(t, err)
		require.Equal(t, testCase.ValidData, s.UintData)

		// If errors are critical - we should check if test data is within
		// int8 range and check for error if it isn't.
		s1 := &testStruct{}
		err1 := Parse(s1, &Options{ErrorsAreCritical: true})

		var (
			checkNotIntError bool
			checkRangeError  bool
		)

		passedData, err2 := strconv.ParseUint(testCase.TestData, 10, 64)
		if err2 != nil {
			checkNotIntError = true
		}

		if passedData != uint64(testCase.ValidData) {
			checkRangeError = true
		}

		if checkNotIntError || checkRangeError {
			if err1 == nil {
				t.Log("No error returned!")
			}

			require.NotNil(t, err1)

			if checkNotIntError {
				require.Equal(t, errNotUint, err1)
			}

			if checkRangeError {
				require.Equal(t, errNotUint8, err1)
			}
		}

		os.Unsetenv("UINTDATA")
	}
}

func TestParseUint16(t *testing.T) {
	testCases := []struct {
		TestData  string
		ValidData uint16
	}{
		{"0", 0},
		{"1", 1},
		{"2", 2},
		{"65535", 65535},
		{"65536", 0},
		{"-1", 0},
		{"not an integer", 0},
	}

	type testStruct struct {
		UintData uint16
	}

	for _, testCase := range testCases {
		t.Logf("Testing: %+v", testCase)
		os.Setenv("UINTDATA", testCase.TestData)

		// If ErrorsAreCritical == false, then we should check only
		// equality of parsed data and valid data.
		s := &testStruct{}

		err := Parse(s, nil)
		if err != nil {
			t.Log(err.Error())
		}

		require.Nil(t, err)
		require.Equal(t, testCase.ValidData, s.UintData)

		// If errors are critical - we should check if test data is within
		// int8 range and check for error if it isn't.
		s1 := &testStruct{}
		err1 := Parse(s1, &Options{ErrorsAreCritical: true})

		var (
			checkNotIntError bool
			checkRangeError  bool
		)

		passedData, err2 := strconv.ParseUint(testCase.TestData, 10, 64)
		if err2 != nil {
			checkNotIntError = true
		}

		if passedData != uint64(testCase.ValidData) {
			checkRangeError = true
		}

		if checkNotIntError || checkRangeError {
			if err1 == nil {
				t.Log("No error returned!")
			}

			require.NotNil(t, err1)

			if checkNotIntError {
				require.Equal(t, errNotUint, err1)
			}

			if checkRangeError {
				require.Equal(t, errNotUint16, err1)
			}
		}

		os.Unsetenv("UINTDATA")
	}
}

func TestParseUint32(t *testing.T) {
	testCases := []struct {
		TestData  string
		ValidData uint32
	}{
		{"0", 0},
		{"1", 1},
		{"2", 2},
		{"4294967295", 4294967295},
		{"4294967296", 0},
		{"-1", 0},
		{"not an integer", 0},
	}

	type testStruct struct {
		UintData uint32
	}

	for _, testCase := range testCases {
		t.Logf("Testing: %+v", testCase)
		os.Setenv("UINTDATA", testCase.TestData)

		// If ErrorsAreCritical == false, then we should check only
		// equality of parsed data and valid data.
		s := &testStruct{}

		err := Parse(s, nil)
		if err != nil {
			t.Log(err.Error())
		}

		require.Nil(t, err)
		require.Equal(t, testCase.ValidData, s.UintData)

		// If errors are critical - we should check if test data is within
		// int8 range and check for error if it isn't.
		s1 := &testStruct{}
		err1 := Parse(s1, &Options{ErrorsAreCritical: true})

		var (
			checkNotIntError bool
			checkRangeError  bool
		)

		passedData, err2 := strconv.ParseUint(testCase.TestData, 10, 64)
		if err2 != nil {
			checkNotIntError = true
		}

		if passedData != uint64(testCase.ValidData) {
			checkRangeError = true
		}

		if checkNotIntError || checkRangeError {
			if err1 == nil {
				t.Log("No error returned!")
			}

			require.NotNil(t, err1)

			if checkNotIntError {
				require.Equal(t, errNotUint, err1)
			}

			if checkRangeError {
				require.Equal(t, errNotUint32, err1)
			}
		}

		os.Unsetenv("UINTDATA")
	}
}

func TestParseUint64(t *testing.T) {
	testCases := []struct {
		TestData  string
		ValidData uint64
	}{
		{"0", 0},
		{"1", 1},
		{"2", 2},
		{"18446744073709551615", 18446744073709551615},
		{"18446744073709551616", 18446744073709551615},
		{"-1", 0},
		{"not an integer", 0},
	}

	type testStruct struct {
		UintData uint64
	}

	for _, testCase := range testCases {
		t.Logf("Testing: %+v", testCase)
		os.Setenv("UINTDATA", testCase.TestData)

		// If ErrorsAreCritical == false, then we should check only
		// equality of parsed data and valid data.
		s := &testStruct{}

		err := Parse(s, nil)
		if err != nil {
			t.Log(err.Error())
		}

		require.Nil(t, err)
		require.Equal(t, testCase.ValidData, s.UintData)

		// If errors are critical - we should check if test data is within
		// int8 range and check for error if it isn't.
		s1 := &testStruct{}
		err1 := Parse(s1, &Options{ErrorsAreCritical: true})

		var (
			checkNotIntError bool
			checkRangeError  bool
		)

		passedData, err2 := strconv.ParseUint(testCase.TestData, 10, 64)
		if err2 != nil {
			checkNotIntError = true
		}

		if passedData != testCase.ValidData {
			checkRangeError = true
		}

		if checkNotIntError || checkRangeError {
			if err1 == nil {
				t.Log("No error returned!")
			}

			require.NotNil(t, err1)

			if checkNotIntError {
				require.Equal(t, errNotUint, err1)
			}

			if checkRangeError {
				require.Equal(t, errNotUint64, err1)
			}
		}

		os.Unsetenv("UINTDATA")
	}
}

// Next tests should be improved.
func TestParseFloat32(t *testing.T) {
	testCases := []struct {
		TestData  string
		ValidData float32
	}{
		{"0.00", 0.00},
		{"1.00", 1.00},
		{"2.00", 2.00},
		{"-1", -1},
		{"not a float", 0.00},
	}

	type testStruct struct {
		FloatData float32
	}

	for _, testCase := range testCases {
		t.Logf("Testing: %+v", testCase)
		os.Setenv("FLOATDATA", testCase.TestData)

		// If ErrorsAreCritical == false, then we should check only
		// equality of parsed data and valid data.
		s := &testStruct{}

		err := Parse(s, nil)
		if err != nil {
			t.Log(err.Error())
		}

		require.Nil(t, err)
		require.Equal(t, testCase.ValidData, s.FloatData)

		// If errors are critical - we should check if test data is within
		// int8 range and check for error if it isn't.
		s1 := &testStruct{}
		err1 := Parse(s1, &Options{ErrorsAreCritical: true})

		var (
			checkNotIntError bool
			checkRangeError  bool
		)

		passedData, err2 := strconv.ParseFloat(testCase.TestData, 64)
		if err2 != nil {
			checkNotIntError = true
		}

		if passedData != float64(testCase.ValidData) {
			checkRangeError = true
		}

		if checkNotIntError || checkRangeError {
			if err1 == nil {
				t.Log("No error returned!")
			}

			require.NotNil(t, err1)

			if checkNotIntError {
				require.Equal(t, errNotFloat, err1)
			}

			if checkRangeError {
				require.Equal(t, errNotFloat32, err1)
			}
		}

		os.Unsetenv("FLOATDATA")
	}
}

func TestParseFloat64(t *testing.T) {
	testCases := []struct {
		TestData  string
		ValidData float64
	}{
		{"0.00", 0.00},
		{"1.00", 1.00},
		{"2.00", 2.00},
		{"-1", -1},
		{"not a float", 0.00},
	}

	type testStruct struct {
		FloatData float64
	}

	for _, testCase := range testCases {
		t.Logf("Testing: %+v", testCase)
		os.Setenv("FLOATDATA", testCase.TestData)

		// If ErrorsAreCritical == false, then we should check only
		// equality of parsed data and valid data.
		s := &testStruct{}

		err := Parse(s, nil)
		if err != nil {
			t.Log(err.Error())
		}

		require.Nil(t, err)
		require.Equal(t, testCase.ValidData, s.FloatData)

		// If errors are critical - we should check if test data is within
		// int8 range and check for error if it isn't.
		s1 := &testStruct{}
		err1 := Parse(s1, &Options{ErrorsAreCritical: true})

		var (
			checkNotIntError bool
			checkRangeError  bool
		)

		passedData, err2 := strconv.ParseFloat(testCase.TestData, 64)
		if err2 != nil {
			checkNotIntError = true
		}

		if passedData != testCase.ValidData {
			checkRangeError = true
		}

		if checkNotIntError || checkRangeError {
			if err1 == nil {
				t.Log("No error returned!")
			}

			require.NotNil(t, err1)

			if checkNotIntError {
				require.Equal(t, errNotFloat, err1)
			}

			if checkRangeError {
				require.Equal(t, errNotFloat64, err1)
			}
		}

		os.Unsetenv("FLOATDATA")
	}
}

func TestParseStructWithInterfaceFields(t *testing.T) {
	type testStruct struct {
		Data interface{}
	}

	os.Setenv(debugFlagEnvName, "true")

	testCase := &testStruct{}
	testCase.Data = 0

	os.Setenv("DATA", "64")

	err := Parse(testCase, nil)
	require.Nil(t, err)
	require.NotEqual(t, 64, testCase.Data)

	testCase1 := &testStruct{}
	d := 0
	shouldBeRaw := 64
	shouldBe := &shouldBeRaw
	testCase1.Data = &d

	err1 := Parse(testCase1, nil)
	require.Nil(t, err1)
	require.Equal(t, (*shouldBe), (*testCase1.Data.(*int)))

	os.Unsetenv("DATA")
	os.Unsetenv(debugFlagEnvName)
}

func TestParseStructWitStructAsInterface(t *testing.T) {
	type testStruct struct {
		Data interface{}
		Int  int
	}

	type testUnderlyingStruct struct {
		Data string
	}

	os.Setenv(debugFlagEnvName, "true")
	os.Setenv("INT", "64")
	os.Setenv("DATA_DATA", "Test data")

	testCase := &testStruct{}
	testUnderlyingCase := &testUnderlyingStruct{}
	testCase.Data = testUnderlyingCase
	err := Parse(testCase, nil)

	require.Nil(t, err)
	require.Equal(t, testCase.Int, 64)
	require.Equal(t, testCase.Data.(*testUnderlyingStruct).Data, "Test data")

	os.Unsetenv("INT")
	os.Unsetenv("DATA_DATA")
	os.Unsetenv(debugFlagEnvName)
}
