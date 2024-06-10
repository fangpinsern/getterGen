package demo

type GetRestaurantInfoRequest struct {
	RestaurantID *uint64
}

type GetRestaurantInfoResponse struct {
	RestaurantInfo *RestaurantInfo `json:"restaurantInfo"`
}

type RestaurantInfo struct {
	ID             *uint64          `json:"id"`
	CreatedAt      *int64           `json:"createdAt"`
	Name           *string          `json:"name"`
	HeaderImage    *string          `json:"headerImage"`
	Country        *string          `json:"country"`
	Description    *string          `json:"description"`
	Address        *string          `json:"address"`
	GroupMin       *int32           `json:"groupMin"`
	GroupMax       *int32           `json:"groupMax"`
	Price          *int32           `json:"price"`
	OperatingHours []*OperatingHour `json:"operatingHours"`
	IsOpen         *bool            `json:"isOpen"`
	Tags           []*Tag           `json:"tags"`
	IsSaved        *bool            `json:"isSaved"`
}

type Tag struct {
	ID   *uint64 `json:"id"`
	Name *string `json:"name"`
	Type *int32  `json:"type"`
}

type OperatingHour struct {
	Day   *int32 `json:"day"`
	Start *int64 `json:"start"`
	End   *int64 `json:"end"`
}
