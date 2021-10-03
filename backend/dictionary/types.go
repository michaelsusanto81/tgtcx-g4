package dictionary

type Toppers struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type CouponTargetType struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	CouponID int    `json:"coupon_id"`
}

type Coupon struct {
	ID                 int      `json:"id"`
	Name               string   `json:"name"`
	Description        string   `json:"description"`
	Duration           int      `json:"duration"`
	Live               bool     `json:"live"`
	StartDate          string   `json:"start_date"`
	EndDate            string   `json:"end_date"`
	MinTransaction     int      `json:"min_transaction"`
	Discount           int      `json:"discount"`
	MaxDiscountPrice   int      `json:"max_discount_price"`
	Category           string   `json:"category"`
	TargetToppersTypes []string `json:"target_toppers_types"`
}

type CouponToppers struct {
	ID        int `json:"id"`
	ToppersID int `json:"toppers_id"`
	CouponID  int `json:"coupon_id"`
}
