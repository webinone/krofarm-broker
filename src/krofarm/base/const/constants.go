package constants

// 센서 ValueType Code 선언
//-----------------------------------------------------
type TypeValueType struct {
	DOUBLE int32
	FLOAT  int32
	INT64  int32
	INT32  int32
	BOOL   int32
	STRING int32
	BYTES  int32
}

var ValueType = &TypeValueType {
	DOUBLE : 0,
	FLOAT: 1,
	INT64: 2,
	INT32: 3,
	BOOL:4,
	STRING:5,
	BYTES:6,
}

//-----------------------------------------------------

// 작동모드 코드
//-----------------------------------------------------
type TypeFnctngModeCd struct {
	MANNUAL string
	AUTOMATIC string
}

var FnctngModeCd = &TypeFnctngModeCd{
	MANNUAL:"10180001",
	AUTOMATIC : "10180002",
}
//------------------------------------------------------

// 센서 속성코드
//------------------------------------------------------
type TypeAttrbCd struct {
	TMPERATURE int32
	HUMIDITY int32
	CO2 int32
	RADIATION int32
	OUT_TEMPERATURE int32
	IN_TEMPERATURE int32
	OUT_HUMIDITY int32
	IN_HUMIDITY int32
	OUT_RADIATION int32
	IN_RADIATION int32
	WIND_DIRECTION int32
	WIND_SPEED int32
	RAIN_LEVEL int32
	RAIN_QTY int32
	HEAT_PIPE int32
	CONTROL_ON_OFF int32
	CONTROL_RANGE int32
}

var AttrbCd = &TypeAttrbCd{
	TMPERATURE: 10010001,
	HUMIDITY : 10010002,
	CO2 : 10010003,
	RADIATION : 10010004,
	OUT_TEMPERATURE : 10010017,
	IN_TEMPERATURE : 10010018,
	OUT_HUMIDITY : 10010019,
	IN_HUMIDITY : 10010020,
	OUT_RADIATION : 10010021,
	IN_RADIATION : 10010022,
	WIND_DIRECTION : 10010007,
	WIND_SPEED : 10010008,
	RAIN_LEVEL : 10010009,
	RAIN_QTY : 10010010,
	HEAT_PIPE : 10010028,
	CONTROL_ON_OFF : 10020001,
	CONTROL_RANGE : 10020002,
}
//------------------------------------------------------

// (센서, 제어기) 상태 코드
type TypeAttrbStatCd struct {
	CONTROL_READY int32
	CONTROL_ING int32
	CONTROL_COMPLETE int32
	SENSOR_NORMAL int32
	SENSOR_ALERT int32
	MODE_ERROR int32
	RESET_ING int32
	RESET_COMPLETE int32
}

var AttrbStatCd = &TypeAttrbStatCd{
	CONTROL_READY : 10100001,
	CONTROL_ING : 10100002,
	CONTROL_COMPLETE : 10100003,
	SENSOR_NORMAL : 10100004,
	SENSOR_ALERT : 10100005,
	MODE_ERROR : 10100006,
	RESET_ING : 10100007,
	RESET_COMPLETE : 10100008,
}

//------------------------------------------------------
// 통신 상태 코드
type TypeCommStatCd struct {
	NORMAL int32
	ERROR int32
}

var CommStatCd = &TypeCommStatCd{
	NORMAL : 10060001,
	ERROR : 10060002,
}
//------------------------------------------------------

// 작동 상태 코드
type TypeFnctngStatCd struct {
	NORMAL int32
	ELECT_WARINING int32
	ALL_ERROR int32
	PARTIAL_ERROR int32
	BLACK_OUT int32
}

var FnctngStatCd = &TypeFnctngStatCd{
	NORMAL : 10080001,
	ELECT_WARINING : 10080002,
	ALL_ERROR : 10080003,
	PARTIAL_ERROR : 10080005,
	BLACK_OUT : 10080006,
}
//------------------------------------------------------
