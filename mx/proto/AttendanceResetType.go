// Code generated by gucooing DO NOT EDIT.

package proto

type AttendanceResetType int32

const (
    AttendanceResetType_User = 0
    AttendanceResetType_Server = 1
)

var (
    AttendanceResetType_name = map[int32]string{
      0 : "User",
      1 : "Server",
   }
    AttendanceResetType_value = map[string]int32{
      "User" : 0,
      "Server" : 1,
   }
)

func (x AttendanceResetType) String() string {
  return AttendanceResetType_name[int32(x)]
}

