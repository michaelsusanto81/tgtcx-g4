package gqlserver

import (
	"github.com/graphql-go/graphql"
	"github.com/tgtcx-g4/coupon/backend/dictionary"
	"github.com/tgtcx-g4/coupon/backend/service"
)

type Resolver struct {
}

func NewResolver() *Resolver {
	return &Resolver{}
}

func (r *Resolver) GetCoupons() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		return service.GetCoupons()
	}
}

func (r *Resolver) CreateCoupons() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		name, _ := p.Args["coupon_name"].(string)
		description, _ := p.Args["coupon_desc"].(string)
		duration, _ := p.Args["coupon_duration"].(uint64)
		live, _ := p.Args["coupon_live"].(bool)
		startDate, _ := p.Args["coupon_start_date"].(string)
		endDate, _ := p.Args["coupon_end_date"].(string)
		minTransaction, _ := p.Args["coupon_min_transaction"].(uint64)
		discount, _ := p.Args["coupon_discount"].(uint8)
		maxDiscount, _ := p.Args["coupon_max_discount_price"].(uint64)
		category, _ := p.Args["coupon_category"].(string)

		req := dictionary.Coupon{
			Name:             name,
			Description:      description,
			Duration:         duration,
			Live:             live,
			StartDate:        startDate,
			EndDate:          endDate,
			MinTransaction:   minTransaction,
			Discount:         discount,
			MaxDiscountPrice: maxDiscount,
			Category:         category,
		}

		_, err := service.CreateCoupon(req)
		if err != nil {
			return false, err
		}

		return true, nil
	}
}

func (r *Resolver) UpdateCouponDuration() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["coupon_id"].(uint64)
		startDate, _ := p.Args["coupon_start_date"].(string)
		endDate, _ := p.Args["coupon_end_date"].(string)

		req := dictionary.Coupon{
			ID:        id,
			StartDate: startDate,
			EndDate:   endDate,
		}

		_, err := service.UpdateCouponDuration(req)
		if err != nil {
			return false, err
		}

		return true, nil
	}
}
