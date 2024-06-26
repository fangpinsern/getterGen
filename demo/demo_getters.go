package demo

// Code generated by getterGen. DO NOT EDIT

func (i *GetRestaurantInfoRequest) GetRestaurantID() uint64 {
	if i != nil && i.RestaurantID != nil {
		return *i.RestaurantID
	}
	return 0
}
func (i *GetRestaurantInfoResponse) GetRestaurantInfo() *RestaurantInfo {
	if i != nil && i.RestaurantInfo != nil {
		return i.RestaurantInfo
	}
	return nil
}
func (i *RestaurantInfo) GetID() uint64 {
	if i != nil && i.ID != nil {
		return *i.ID
	}
	return 0
}
func (i *RestaurantInfo) GetCreatedAt() int64 {
	if i != nil && i.CreatedAt != nil {
		return *i.CreatedAt
	}
	return 0
}
func (i *RestaurantInfo) GetName() string {
	if i != nil && i.Name != nil {
		return *i.Name
	}
	return ""
}
func (i *RestaurantInfo) GetHeaderImage() string {
	if i != nil && i.HeaderImage != nil {
		return *i.HeaderImage
	}
	return ""
}
func (i *RestaurantInfo) GetCountry() string {
	if i != nil && i.Country != nil {
		return *i.Country
	}
	return ""
}
func (i *RestaurantInfo) GetDescription() string {
	if i != nil && i.Description != nil {
		return *i.Description
	}
	return ""
}
func (i *RestaurantInfo) GetAddress() string {
	if i != nil && i.Address != nil {
		return *i.Address
	}
	return ""
}
func (i *RestaurantInfo) GetGroupMin() int32 {
	if i != nil && i.GroupMin != nil {
		return *i.GroupMin
	}
	return 0
}
func (i *RestaurantInfo) GetGroupMax() int32 {
	if i != nil && i.GroupMax != nil {
		return *i.GroupMax
	}
	return 0
}
func (i *RestaurantInfo) GetPrice() int32 {
	if i != nil && i.Price != nil {
		return *i.Price
	}
	return 0
}
func (i *RestaurantInfo) GetOperatingHours() []*OperatingHour {
	if i != nil {
		return i.OperatingHours
	}
	return nil
}
func (i *RestaurantInfo) GetIsOpen() bool {
	if i != nil && i.IsOpen != nil {
		return *i.IsOpen
	}
	return false
}
func (i *RestaurantInfo) GetTags() []*Tag {
	if i != nil {
		return i.Tags
	}
	return nil
}
func (i *RestaurantInfo) GetIsSaved() bool {
	if i != nil && i.IsSaved != nil {
		return *i.IsSaved
	}
	return false
}
func (i *Tag) GetID() uint64 {
	if i != nil && i.ID != nil {
		return *i.ID
	}
	return 0
}
func (i *Tag) GetName() string {
	if i != nil && i.Name != nil {
		return *i.Name
	}
	return ""
}
func (i *Tag) GetType() int32 {
	if i != nil && i.Type != nil {
		return *i.Type
	}
	return 0
}
func (i *OperatingHour) GetDay() int32 {
	if i != nil && i.Day != nil {
		return *i.Day
	}
	return 0
}
func (i *OperatingHour) GetStart() int64 {
	if i != nil && i.Start != nil {
		return *i.Start
	}
	return 0
}
func (i *OperatingHour) GetEnd() int64 {
	if i != nil && i.End != nil {
		return *i.End
	}
	return 0
}
func (i *Complex) GetNumberI64() int64 {
	if i != nil {
		return i.NumberI64
	}
	return 0
}
func (i *Complex) GetPtrNumberI64() int64 {
	if i != nil && i.PtrNumberI64 != nil {
		return *i.PtrNumberI64
	}
	return 0
}
func (i *Complex) GetArrNumberI64() []int64 {
	if i != nil {
		return i.ArrNumberI64
	}
	return nil
}
func (i *Complex) GetArrPtrNumberI64() []*int64 {
	if i != nil {
		return i.ArrPtrNumberI64
	}
	return nil
}
func (i *Complex) GetNumberI32() int32 {
	if i != nil {
		return i.NumberI32
	}
	return 0
}
func (i *Complex) GetPtrNumberI32() int32 {
	if i != nil && i.PtrNumberI32 != nil {
		return *i.PtrNumberI32
	}
	return 0
}
func (i *Complex) GetNumberI16() int16 {
	if i != nil {
		return i.NumberI16
	}
	return 0
}
func (i *Complex) GetPtrNumberI16() int16 {
	if i != nil && i.PtrNumberI16 != nil {
		return *i.PtrNumberI16
	}
	return 0
}
func (i *Complex) GetStr() string {
	if i != nil {
		return i.Str
	}
	return ""
}
func (i *Complex) GetPtrStr() string {
	if i != nil && i.PtrStr != nil {
		return *i.PtrStr
	}
	return ""
}
func (i *Complex) GetArrStr() []string {
	if i != nil {
		return i.ArrStr
	}
	return nil
}
func (i *Complex) GetArrStruct() []*Complex {
	if i != nil {
		return i.ArrStruct
	}
	return nil
}
func (i *Complex) GetBaseStruct() SecondStruct {
	if i != nil {
		return i.BaseStruct
	}
	return SecondStruct{}
}
func (i *Complex) GetPtrStruct() *SecondStruct {
	if i != nil && i.PtrStruct != nil {
		return i.PtrStruct
	}
	return nil
}
func (i *SecondStruct) GetArrStr() []string {
	if i != nil {
		return i.ArrStr
	}
	return nil
}
func (i *SecondStruct) GetArrStruct() []*SecondStruct {
	if i != nil {
		return i.ArrStruct
	}
	return nil
}
