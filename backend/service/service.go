package service

import (
	"errors"
	"fmt"

	"github.com/tgtcx-g4/coupon/backend/database"
	"github.com/tgtcx-g4/coupon/backend/dictionary"
)

func UpdateCoupon(data dictionary.Coupon) (*dictionary.Coupon, error) {
	fmt.Printf("%v", data)
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
	if err_live != nil {
		return nil, errors.New("1")
	} else if is_live {
		return nil, errors.New("Coupon is live")
	}

	// query
	query = `
		UPDATE
			coupon
		SET
			coupon_name = $2,
			coupon_desc = $3,
			coupon_duration = $4,
			coupon_live = $5,
			coupon_start_date = $6,
			coupon_end_date = $7,
			coupon_min_transaction = $8,
			coupon_discount = $9,
			coupon_max_discount_price = $10,
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
		return nil, errors.New("2")
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return nil, errors.New("3")
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
		FROM coupon_target_type
		WHERE
			coupontarget_name = $2 AND
			coupon_id = $1
	`
	var ctt dictionary.CouponTargetType
	err_exists := db.QueryRow(query, data.CouponID, data.Name).Scan(&ctt)
	if err_exists == nil {
		return nil, err_exists
	} else if ctt.ID != 0 {
		return nil, errors.New("Coupon Target already exists")
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

func GetTargetUsers(coupon_id int) ([]string, error) {
	// get current database connection
	db := database.GetDB()

	// query
	query := `
		SELECT
			coupontarget_name
		FROM
			coupon_target_type
		WHERE
			coupon_id = $1
	`

	// actual query process
	rows, err := db.Query(query, coupon_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []string
	for rows.Next() {
		var data string
		rows.Scan(&data)
		if err != nil {
			return nil, err
		}

		result = append(result, data)
	}

	return result, nil
}

func GetUserCoupons(toppers_id int) ([]dictionary.CouponToppers, error) {
	db := database.GetDB()

	query := `
		SELECT
			coupontoppers_id,
			toppers_id,
			coupon_id
		FROM
			coupon_toppers
		WHERE
			toppers_id = $1
	`

	// actual query process
	rows, err := db.Query(query, toppers_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []dictionary.CouponToppers
	for rows.Next() {
		var data dictionary.CouponToppers
		rows.Scan(
			&data.ID,
			&data.ToppersID,
			&data.CouponID,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, data)
	}

	return result, nil
}

func GetCoupon(paramID int) (*dictionary.Coupon, error) {

	// you can connect and
	// get current database connection
	db := database.GetDB()

	// construct query
	query := `
	SELECT 
		coupon_id,
		coupon_name,
		coupon_desc,
		coupon_duration,
		coupon_live,
		coupon_start_date,
		coupon_end_date,
		coupon_min_transaction,
		coupon_discount,
		coupon_max_discount_price,
		coupon_category
	FROM coupon
	WHERE coupon_id = $1
	`
	// actual query process
	row := db.QueryRow(query, paramID)

	// read query result, and assign to variable(s)
	var result dictionary.Coupon
	err := row.Scan(
		&result.ID,
		&result.Name,
		&result.Description,
		&result.Duration,
		&result.Live,
		&result.StartDate,
		&result.EndDate,
		&result.MinTransaction,
		&result.Discount,
		&result.MaxDiscountPrice,
		&result.Category,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func GetCoupons() ([]dictionary.Coupon, error) {

	// you can connect and
	// get current database connection
	db := database.GetDB()

	// construct query
	query := `
	SELECT 
		coupon_id,
		coupon_name,
		coupon_desc,
		coupon_duration,
		coupon_live,
		coupon_start_date,
		coupon_end_date,
		coupon_min_transaction,
		coupon_discount,
		coupon_max_discount_price,
		coupon_category
	FROM coupon
	`
	// actual query process
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []dictionary.Coupon
	for rows.Next() {
		var data dictionary.Coupon
		rows.Scan(
			&data.ID,
			&data.Name,
			&data.Description,
			&data.Duration,
			&data.Live,
			&data.StartDate,
			&data.EndDate,
			&data.MinTransaction,
			&data.Discount,
			&data.MaxDiscountPrice,
			&data.Category,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, data)
	}

	return result, nil
}

func CreateCoupon(data dictionary.Coupon) (*dictionary.Coupon, error) {

	// connect and get current database connection
	db := database.GetDB()

	// construct query
	query := `
	INSERT INTO coupon (
		coupon_name,
		coupon_desc,
		coupon_duration,
		coupon_live,
		coupon_start_date,
		coupon_end_date,
		coupon_min_transaction,
		coupon_discount,
		coupon_max_discount_price,
		coupon_category
	) VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	// actual query process
	result, err := db.Exec(query,
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
			coupontarget_name = $2,
			coupon_id = $3
		WHERE
			coupontarget_id = $1
	`

	// exec
	result, err := db.Exec(query,
		data.ID,
		data.Name,
		data.CouponID,
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

func UpdateCouponDuration(data dictionary.Coupon) (*dictionary.Coupon, error) {

	// connect and get current database connection
	db := database.GetDB()

	// construct query
	query := `
	UPDATE 
		coupon
	SET 
		coupon_start_date = $2,
		coupon_end_date = $3,
		coupon_duration = $4
	WHERE
		coupon_id = $1
	`
	// actual query process
	result, err := db.Exec(query,
		data.ID,
		data.StartDate,
		data.EndDate,
		data.Duration,
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
