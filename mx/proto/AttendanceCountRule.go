// Code generated by gucooing DO NOT EDIT.

package proto

type AttendanceCountRule int32

const (
    AttendanceCountRule_Accumulation = 0
    AttendanceCountRule_Date = 1
)

var (
    AttendanceCountRule_name = map[int32]string{
      0 : "Accumulation",
      1 : "Date",
   }
    AttendanceCountRule_value = map[string]int32{
      "Accumulation" : 0,
      "Date" : 1,
   }
)

func (x AttendanceCountRule) String() string {
  return AttendanceCountRule_name[int32(x)]
}

