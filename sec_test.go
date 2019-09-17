package sec

import (
	// stdlib
	"os"
	"strconv"
	"testing"

	// other
	"github.com/stretchr/testify/require"
)

var (
	testString  = "Test string"
	testInt8    = int8(8)
	testInt16   = int16(16)
	testInt32   = int32(32)
	testInt64   = int64(64)
	testUint8   = uint8(8)
	testUint16  = uint16(16)
	testUint32  = uint32(32)
	testUint64  = uint64(64)
	testFloat32 = float32(32.00)
	testFloat64 = float64(64.00)
	testBool    = true
)

type testDatas struct {
	TestString  string
	TestInt8    int8
	TestInt16   int16
	TestInt32   int32
	TestInt64   int64
	TestUint8   uint8
	TestUint16  uint16
	TestUint32  uint32
	TestUint64  uint64
	TestFloat32 float32
	TestFloat64 float64
	TestBool    bool
}

type testStringType string

type testStruct1 struct {
	testDatas
	testStringType
	TestNestAnonymous struct {
		TestString  string
		TestInt8    int8
		TestInt16   int16
		TestInt32   int32
		TestInt64   int64
		TestUint8   uint8
		TestUint16  uint16
		TestUint32  uint32
		TestUint64  uint64
		TestFloat32 float32
		TestFloat64 float64
		TestBool    bool
	}
	TestNestAnonymousPointer *struct {
		TestString  string
		TestInt8    int8
		TestInt16   int16
		TestInt32   int32
		TestInt64   int64
		TestUint8   uint8
		TestUint16  uint16
		TestUint32  uint32
		TestUint64  uint64
		TestFloat32 float32
		TestFloat64 float64
		TestBool    bool
	}
	TestNestPointer          *testDatas
	TestNest                 testDatas
	TestNestInterfacePointer interface{}
	TestNestInterface        interface{}
	testUnexported           string
	testUnexportedNest       *testDatas
}

func setenv(prefix string) {
	os.Setenv(prefix+"TESTSTRING", testString)
	os.Setenv(prefix+"TESTINT8", strconv.FormatInt(int64(testInt8), 10))
	os.Setenv(prefix+"TESTINT16", strconv.FormatInt(int64(testInt16), 10))
	os.Setenv(prefix+"TESTINT32", strconv.FormatInt(int64(testInt32), 10))
	os.Setenv(prefix+"TESTINT64", strconv.FormatInt(int64(testInt64), 10))
	os.Setenv(prefix+"TESTUINT8", strconv.FormatInt(int64(testUint8), 10))
	os.Setenv(prefix+"TESTUINT16", strconv.FormatInt(int64(testUint16), 10))
	os.Setenv(prefix+"TESTUINT32", strconv.FormatInt(int64(testUint32), 10))
	os.Setenv(prefix+"TESTUINT64", strconv.FormatInt(int64(testUint64), 10))
	os.Setenv(prefix+"TESTFLOAT32", strconv.FormatFloat(float64(testFloat32), 'f', 2, 32))
	os.Setenv(prefix+"TESTFLOAT64", strconv.FormatFloat(testFloat64, 'f', 2, 64))
	os.Setenv(prefix+"TESTBOOL", "true")

	os.Setenv(debugFlagEnvName, "true")
}

func unsetenv(prefix string) {
	os.Unsetenv(prefix + "TESTSTRING")
	os.Unsetenv(prefix + "TESTINT8")
	os.Unsetenv(prefix + "TESTINT16")
	os.Unsetenv(prefix + "TESTINT32")
	os.Unsetenv(prefix + "TESTINT64")
	os.Unsetenv(prefix + "TESTUINT8")
	os.Unsetenv(prefix + "TESTUINT16")
	os.Unsetenv(prefix + "TESTUINT32")
	os.Unsetenv(prefix + "TESTUINT64")
	os.Unsetenv(prefix + "TESTFLOAT32")
	os.Unsetenv(prefix + "TESTFLOAT64")
	os.Unsetenv(prefix + "TESTBOOL")

	os.Unsetenv(debugFlagEnvName)
}

func TestParseValidData(t *testing.T) {
	setenv("TESTSTRUCT1_")
	setenv("TESTSTRUCT1_TESTNEST_")
	setenv("TESTSTRUCT1_TESTNESTANONYMOUS_")
	setenv("TESTSTRUCT1_TESTNESTANONYMOUSPOINTER_")
	setenv("TESTSTRUCT1_TESTNESTINTERFACE_")
	setenv("TESTSTRUCT1_TESTNESTINTERFACEPOINTER_")
	setenv("TESTSTRUCT1_TESTNESTPOINTER_")
	setenv("TESTSTRUCT1_TESTUNEXPORTEDNEST_")

	ts := &testStruct1{}
	err := Parse(ts, nil)
	t.Logf("Parsed data: %+v\n", ts)
	t.Logf("Parsed nested data: %+v\n", ts.TestNest)
	t.Logf("Parsed nested interface data: %+v\n", ts.TestNestInterface)

	require.Nil(t, err)
	require.Equal(t, testBool, ts.TestBool)

	unsetenv("TESTSTRUCT1_")
	unsetenv("TESTSTRUCT1_TESTNEST_")
	unsetenv("TESTSTRUCT1_TESTNESTANONYMOUS_")
	unsetenv("TESTSTRUCT1_TESTNESTANONYMOUSPOINTER_")
	unsetenv("TESTSTRUCT1_TESTNESTINTERFACE_")
	unsetenv("TESTSTRUCT1_TESTNESTINTERFACEPOINTER_")
	unsetenv("TESTSTRUCT1_TESTNESTPOINTER_")
	unsetenv("TESTSTRUCT1_TESTUNEXPORTEDNEST_")
}

func TestParseNotPointerToStructurePassed(t *testing.T) {
	setenv("")

	var data string
	err := Parse(&data, nil)

	require.NotNil(t, err)
	require.Equal(t, errNotStructure, err)

	unsetenv("")
}

func TestParseNotPointerPassed(t *testing.T) {
	setenv("")

	c := testStruct1{}
	err := Parse(c, nil)

	require.NotNil(t, err)
	require.Equal(t, errNotPTR, err)

	unsetenv("")
}

func TestParseNotStructurePassed(t *testing.T) {
	d := "invalid data"
	err := Parse(d, nil)
	t.Log(err.Error())

	require.NotNil(t, err)
	require.Equal(t, errNotPTR, err)
}

func TestInvalidDebugFlagValue(t *testing.T) {
	_ = os.Setenv(debugFlagEnvName, "INVALID")
	c := &testStruct1{}
	err := Parse(c, nil)

	require.Nil(t, err)
	require.False(t, debug)

	os.Unsetenv(debugFlagEnvName)
}

func TestInvalidDebugFlagValueWithErrorsAreCritical(t *testing.T) {
	_ = os.Setenv(debugFlagEnvName, "INVALID")
	c := &testStruct1{}
	err := Parse(c, &Options{ErrorsAreCritical: true})

	if err != nil {
		t.Log(err.Error())
	}

	require.NotNil(t, err)
	require.False(t, debug)

	os.Unsetenv(debugFlagEnvName)
}
