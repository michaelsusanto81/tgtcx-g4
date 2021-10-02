package gqlserver

import "github.com/graphql-go/graphql"

var ToppersType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Toppers",
		Description: "Detail of the Toppers",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"type": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var CouponTargetTypeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "CouponTargetType",
		Description: "Detail of the Coupon Target Type",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"coupon_id": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var CouponType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Coupon",
		Description: "Detail of the Coupon",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"duration": &graphql.Field{
				Type: graphql.Int,
			},
			"live": &graphql.Field{
				Type: graphql.Boolean,
			},
			"start_date": &graphql.Field{
				Type: graphql.String,
			},
			"end_date": &graphql.Field{
				Type: graphql.String,
			},
			"min_transaction": &graphql.Field{
				Type: graphql.Int,
			},
			"discount": &graphql.Field{
				Type: graphql.Int,
			},
			"max_discount_price": &graphql.Field{
				Type: graphql.Int,
			},
			"category": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

<<<<<<< HEAD
var CouponToppersType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "CouponToppers",
		Description: "Detail of the Coupon Toppers",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"toppers_id": &graphql.Field{
				Type: graphql.Int,
			},
			"coupon_id": &graphql.Field{
				Type: graphql.Int,
=======
var ResultType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Result",
		Description: "Result of operation",
		Fields: graphql.Fields{
			"success": &graphql.Field{
				Type: graphql.Boolean,
>>>>>>> createcoupon-updateduration
			},
		},
	},
)
