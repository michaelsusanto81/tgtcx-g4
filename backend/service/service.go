package service

import (
	"errors"

	"github.com/tgtcx-g4/coupon/backend/database"
	"github.com/tgtcx-g4/coupon/backend/dictionary"
)

func UpdateCoupon(data dictionary.Coupon) (*dictionary.Coupon, error) {
	// get current db connection
	db := database.GetDB()

	// if status is live throw err
	query := `
		SELECT
			coupon_live
		FROM
			coupon
		WHERE
			coupon_id = $1
	`
	var is_live bool
	err_live := db.QueryRow(query, data.ID).Scan(&is_live)
	if err_live != nil || is_live {
		return nil, err_live
	}

	// query
	query = `
		UPDATE
			coupon
		SET
			coupon_name = $2
			coupon_desc = $3
			coupon_duration = $4
			coupon_live = $5
			coupon_start_date = $6
			coupon_end_date = $7
			coupon_min_transaction = $8
			coupon_discount = $9
			coupon_max_discount_price = $10
			coupon_category = $11
		WHERE
			coupon_id = $1
	`

	// exec
	result, err := db.Exec(query,
		data.ID,
		data.Name,
		data.Description,
		data.Duration,
		data.Live,
		data.StartDate,
		data.EndDate,
		data.MinTransaction,
		data.Discount,
		data.MaxDiscountPrice,
		data.Category,
	)
	if err != nil {
		return nil, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, errors.New("no row updated")
	}

	return &data, nil
}

func AddTargetUser(data dictionary.CouponTargetType) (*dictionary.CouponTargetType, error) {
	// get current db connection
	db := database.GetDB()

	// if already exists return err
	query := `
		SELECT coupontarget_id
		FROM coupun_target_type
		WHERE
			coupontarget_name = $2
			coupon_id = $1
	`
	var ctt dictionary.CouponTargetType
	err_exists := db.QueryRow(query, data.CouponID, data.Name).Scan(&ctt)
	if err_exists != nil || ctt.ID != 0 {
		return nil, err_exists
	}

	// query
	query = `
		INSERT INTO	coupon_target_type (coupontarget_name, coupon_id)
		VALUES ($2, $1)
	`

	// actual query process
	result, err := db.Exec(query,
		data.CouponID,
		data.Name,
	)
	if err != nil {
		return nil, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, errors.New("no row created")
	}

	return &data, nil
}

func SetTargetUser(data dictionary.CouponTargetType) (*dictionary.CouponTargetType, error) {
	// get current db connection
	db := database.GetDB()

	// query
	query := `
		UPDATE
			coupon_target_type
		SET
			coupontarget_name = $2
		WHERE
			coupon_id = $1
	`

	// exec
	result, err := db.Exec(query,
		data.CouponID,
		data.Name,
	)
	if err != nil {
		return nil, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, errors.New("no row updated")
	}

	return &data, nil
}
