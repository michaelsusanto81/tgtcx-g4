package dictionary

type Toppers struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type CouponTargetType struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	CouponID uint64 `json:"coupon_id"`
}

type Coupon struct {
	ID                 uint64   `json:"id"`
	Name               string   `json:"name"`
	Description        string   `json:"description"`
	Duration           uint64   `json:"duration"`
	Live               bool     `json:"live"`
	StartDate          string   `json:"start_date"`
	EndDate            string   `json:"end_date"`
	MinTransaction     uint64   `json:"min_transaction"`
	Discount           uint8    `json:"discount"`
	MaxDiscountPrice   uint64   `json:"max_discount_price"`
	Category           string   `json:"category"`
	TargetToppersTypes []string `json:"target_toppers_types"`
}

type CouponToppers struct {
	ID        uint64 `json:"id"`
	ToppersID uint64 `json:"toppers_id"`
	CouponID  uint64 `json:"coupon_id"`
}
