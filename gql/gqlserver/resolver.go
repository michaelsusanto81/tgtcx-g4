package gqlserver

import (
	"fmt"

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

func (r *Resolver) GetTargetUsers() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["coupon_id"].(int)
		return service.GetTargetUsers(id)
	}
}

func (r *Resolver) CreateCoupons() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		name, _ := p.Args["coupon_name"].(string)
		description, _ := p.Args["coupon_desc"].(string)
		duration, _ := p.Args["coupon_duration"].(int)
		live, _ := p.Args["coupon_live"].(bool)
		startDate, _ := p.Args["coupon_start_date"].(string)
		endDate, _ := p.Args["coupon_end_date"].(string)
		minTransaction, _ := p.Args["coupon_min_transaction"].(int)
		discount, _ := p.Args["coupon_discount"].(int)
		maxDiscount, _ := p.Args["coupon_max_discount_price"].(int)
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

func (r *Resolver) UpdateCoupon() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["coupon_id"].(int)
		name, _ := p.Args["coupon_name"].(string)
		desc, _ := p.Args["coupon_desc"].(string)
		duration, _ := p.Args["coupon_duration"].(int)
		live, _ := p.Args["coupon_live"].(bool)
		startDate, _ := p.Args["coupon_start_date"].(string)
		endDate, _ := p.Args["coupon_end_date"].(string)
		minTransaction, _ := p.Args["coupon_min_transaction"].(int)
		discount, _ := p.Args["coupon_discount"].(int)
		maxDiscountPrice, _ := p.Args["coupon_max_discount_price"].(int)
		category, _ := p.Args["coupon_category"].(string)
		fmt.Println(id)

		req := dictionary.Coupon{
			ID:               id,
			Name:             name,
			Description:      desc,
			Duration:         duration,
			Live:             live,
			StartDate:        startDate,
			EndDate:          endDate,
			MinTransaction:   minTransaction,
			Discount:         discount,
			MaxDiscountPrice: maxDiscountPrice,
			Category:         category,
		}

		_, err := service.UpdateCoupon(req)
		if err != nil {
			return false, err
		}

		return true, nil
	}
}

func (r *Resolver) UpdateCouponDuration() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["coupon_id"].(int)
		startDate, _ := p.Args["coupon_start_date"].(string)
		endDate, _ := p.Args["coupon_end_date"].(string)
		duration, _ := p.Args["coupon_duration"].(int)

		req := dictionary.Coupon{
			ID:        id,
			StartDate: startDate,
			EndDate:   endDate,
			Duration:  duration,
		}

		_, err := service.UpdateCouponDuration(req)
		if err != nil {
			return false, err
		}

		return true, nil
	}
}

func (r *Resolver) AddTargetUser() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		targetName, _ := p.Args["coupontarget_name"].(string)
		targetCouponID, _ := p.Args["coupon_id"].(int)

		req := dictionary.CouponTargetType{
			Name:     targetName,
			CouponID: targetCouponID,
		}

		_, err := service.AddTargetUser(req)
		if err != nil {
			return false, err
		}

		return true, nil
	}
}

func (r *Resolver) SetTargetUser() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["coupontarget_id"].(int)
		targetName, _ := p.Args["coupontarget_name"].(string)
		targetCouponID, _ := p.Args["coupon_id"].(int)

		req := dictionary.CouponTargetType{
			ID:       id,
			Name:     targetName,
			CouponID: targetCouponID,
		}

		_, err := service.SetTargetUser(req)
		if err != nil {
			return false, err
		}

		return true, nil
	}
}
