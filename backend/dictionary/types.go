package dictionary

type User struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Coupon struct {
	ID               uint64   `json:"id"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Duration         uint64   `json:"duration"`
	Live             bool     `json:"live"`
	StartDate        string   `json:"start_date"`
	EndDate          string   `json:"end_date"`
	MinTransaction   uint64   `json:"min_transaction"`
	Discount         uint8    `json:"discount"`
	MaxDiscountPrice uint64   `json:"max_discount_price"`
	Category         string   `json:"category"`
	TargetUserTypes  []string `json:"target_user_types"`
}

type CouponUser struct {
	ID       uint64 `json:"id"`
	UserID   uint64 `json:"user_id"`
	CouponID uint64 `json:"coupon_id"`
}
