// Code generated by gucooing DO NOT EDIT.

package proto

type PurchaseStatusCode int32

const (
    PurchaseStatusCode_None = 0
    PurchaseStatusCode_Start = 1
    PurchaseStatusCode_PublishSuccess = 2
    PurchaseStatusCode_End = 3
    PurchaseStatusCode_Error = 4
    PurchaseStatusCode_DuplicateOrder = 5
    PurchaseStatusCode_Refund = 6
)

var (
    PurchaseStatusCode_name = map[int32]string{
      0 : "None",
      1 : "Start",
      2 : "PublishSuccess",
      3 : "End",
      4 : "Error",
      5 : "DuplicateOrder",
      6 : "Refund",
   }
    PurchaseStatusCode_value = map[string]int32{
      "None" : 0,
      "Start" : 1,
      "PublishSuccess" : 2,
      "End" : 3,
      "Error" : 4,
      "DuplicateOrder" : 5,
      "Refund" : 6,
   }
)

func (x PurchaseStatusCode) String() string {
  return PurchaseStatusCode_name[int32(x)]
}

