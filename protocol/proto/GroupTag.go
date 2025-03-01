// Code generated by gucooing DO NOT EDIT.

package proto

type GroupTag string

const (
    GroupTag_None = 0
    GroupTag_Group01 = 1
    GroupTag_Group02 = 2
    GroupTag_Group03 = 4
    GroupTag_Group04 = 8
    GroupTag_Group05 = 16
    GroupTag_Group06 = 32
    GroupTag_Group07 = 64
    GroupTag_Group08 = 128
    GroupTag_Group09 = 256
    GroupTag_Group10 = 512
    GroupTag_Group11 = 1024
    GroupTag_Group12 = 2048
    GroupTag_Group13 = 4096
    GroupTag_Group14 = 8192
    GroupTag_Group15 = 16384
    GroupTag_Group16 = 32768
)

var (
    GroupTag_name = map[int32]string{
      0 : "None",
      1 : "Group01",
      2 : "Group02",
      4 : "Group03",
      8 : "Group04",
      16 : "Group05",
      32 : "Group06",
      64 : "Group07",
      128 : "Group08",
      256 : "Group09",
      512 : "Group10",
      1024 : "Group11",
      2048 : "Group12",
      4096 : "Group13",
      8192 : "Group14",
      16384 : "Group15",
      32768 : "Group16",
   }
    GroupTag_value = map[string]int32{
      "None" : 0,
      "Group01" : 1,
      "Group02" : 2,
      "Group03" : 4,
      "Group04" : 8,
      "Group05" : 16,
      "Group06" : 32,
      "Group07" : 64,
      "Group08" : 128,
      "Group09" : 256,
      "Group10" : 512,
      "Group11" : 1024,
      "Group12" : 2048,
      "Group13" : 4096,
      "Group14" : 8192,
      "Group15" : 16384,
      "Group16" : 32768,
   }
)

func (x GroupTag) String() string {
  return string(x)
}

