package sec

import (
	"os"
	"strconv"
	"testing"

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
	TestFloat64 float64
	TestUint64  uint64
	TestInt64   int64

	TestFloat32 float32
	TestUint32  uint32
	TestInt32   int32

	TestUint16 uint16
	TestInt16  int16

	TestUint8 uint8
	TestInt8  int8

	TestBool bool

	TestString string
}

type testStringType string

type testStruct1 struct {
	testDatas
	testStringType
	TestNestAnonymous struct {
		TestFloat64 float64
		TestUint64  uint64
		TestInt64   int64

		TestFloat32 float32
		TestUint32  uint32
		TestInt32   int32

		TestUint16 uint16
		TestInt16  int16

		TestUint8 uint8
		TestInt8  int8

		TestBool bool

		TestString string
	}
	TestNestAnonymousPointer *struct {
		TestFloat64 float64
		TestUint64  uint64
		TestInt64   int64

		TestFloat32 float32
		TestUint32  uint32
		TestInt32   int32

		TestUint16 uint16
		TestInt16  int16

		TestUint8 uint8
		TestInt8  int8

		TestBool bool

		TestString string
	}
	TestNestPointer          *testDatas
	TestNest                 testDatas
	TestNestInterfacePointer interface{}
	TestNestInterface        interface{}
	// testUnexported           string
	// testUnexportedNest       *testDatas
}

type testStructWithMap struct {
	MapConfig map[string]interface{}
}

func setenv(prefix string) {
	os.Setenv(prefix+"TESTSTRING", testString)
	os.Setenv(prefix+"TESTINT8", strconv.FormatInt(int64(testInt8), 10))
	os.Setenv(prefix+"TESTINT16", strconv.FormatInt(int64(testInt16), 10))
	os.Setenv(prefix+"TESTINT32", strconv.FormatInt(int64(testInt32), 10))
	os.Setenv(prefix+"TESTINT64", strconv.FormatInt(testInt64, 10))
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
	setenv("")
	setenv("TESTNEST_")
	setenv("TESTNESTANONYMOUS_")
	setenv("TESTNESTANONYMOUSPOINTER_")
	setenv("TESTNESTINTERFACE_")
	setenv("TESTNESTINTERFACEPOINTER_")
	setenv("TESTNESTPOINTER_")
	setenv("TESTUNEXPORTEDNEST_")
	setenv("MAPCONFIG_TESTSTRUCT_")
	setenv("MAPCONFIG_TESTSTRUCT_TESTNEST_")

	ts := &testStruct1{}
	err := Parse(ts, nil)
	t.Logf("Parsed data: %+v\n", ts)
	t.Logf("Parsed nested data: %+v\n", ts.TestNest)
	t.Logf("Parsed nested data as pointer: %+v\n", ts.TestNestPointer)
	t.Logf("Parsed nested interface data: %+v\n", ts.TestNestInterface)

	require.Nil(t, err)
	require.Equal(t, testBool, ts.TestBool)

	ts1 := &testStructWithMap{MapConfig: map[string]interface{}{
		"teststruct": &testStruct1{},
	}}

	err1 := Parse(ts1, nil)
	require.Nil(t, err1)

	t.Logf("Parsed struct with map data: %+v\n", ts1.MapConfig["teststruct"])

	unsetenv("")
	unsetenv("TESTNEST_")
	unsetenv("TESTNESTANONYMOUS_")
	unsetenv("TESTNESTANONYMOUSPOINTER_")
	unsetenv("TESTNESTINTERFACE_")
	unsetenv("TESTNESTINTERFACEPOINTER_")
	unsetenv("TESTNESTPOINTER_")
	unsetenv("TESTUNEXPORTEDNEST_")
	unsetenv("MAPCONFIG_TESTSTRUCT_")
	unsetenv("MAPCONFIG_TESTSTRUCT_TESTNEST_")
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
