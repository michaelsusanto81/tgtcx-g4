package gqlserver

import "github.com/graphql-go/graphql"

var CouponType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Coupon",
		Description: "Detail of the coupon",
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

var ResultType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Result",
		Description: "Result of operation",
		Fields: graphql.Fields{
			"success": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)
