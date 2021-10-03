package gqlserver

import (
	"github.com/graphql-go/graphql"
)

type SchemaWrapper struct {
	couponResolver *Resolver
	Schema         graphql.Schema
}

func NewSchemaWrapper() *SchemaWrapper {
	return &SchemaWrapper{}
}

func (s *SchemaWrapper) WithCouponResolver(pr *Resolver) *SchemaWrapper {
	s.couponResolver = pr

	return s
}

func (s *SchemaWrapper) Init() error {
	// add gql schema as needed
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:        "CouponGetter",
			Description: "All query related to getting coupon data",
			Fields: graphql.Fields{
				"Coupons": &graphql.Field{
					Type:        graphql.NewList(CouponType),
					Description: "Get all coupons",
					Resolve:     s.couponResolver.GetCoupons(),
				},
				"CouponTargetUsers": &graphql.Field{
					Type:        graphql.NewList(graphql.String),
					Description: "Get target users from specific coupon",
					Args: graphql.FieldConfigArgument{
						"coupon_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: s.couponResolver.GetTargetUsers(),
				},
			},
		}),
		// uncomment this and add objects for mutation
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:        "CouponSetter",
			Description: "All query related to modify coupon data",
			Fields: graphql.Fields{
				"CreateCoupon": &graphql.Field{
					Type:        graphql.Boolean,
					Description: "Create coupon",
					Args: graphql.FieldConfigArgument{
						"coupon_name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"coupon_description": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"coupon_duration": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"coupon_live": &graphql.ArgumentConfig{
							Type: graphql.Boolean,
						},
						"coupon_start_date": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"coupon_end_date": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"coupon_min_transaction": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"coupon_discount": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"coupon_max_discount_price": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"coupon_category": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: s.couponResolver.CreateCoupons(),
				},
				"UpdateCoupon": &graphql.Field{
					Type:        graphql.Boolean,
					Description: "Update coupon by ID",
					Args: graphql.FieldConfigArgument{
						"coupon_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
						"coupon_name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"coupon_description": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"coupon_duration": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"coupon_live": &graphql.ArgumentConfig{
							Type: graphql.Boolean,
						},
						"coupon_start_date": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"coupon_end_date": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"coupon_min_transaction": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"coupon_discount": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"coupon_max_discount_price": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"coupon_category": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: s.couponResolver.UpdateCoupon(),
				},
				"UpdateCouponDuration": &graphql.Field{
					Type:        graphql.Boolean,
					Description: "Update coupon's duration by ID",
					Args: graphql.FieldConfigArgument{
						"coupon_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
						"coupon_start_date": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"coupon_end_date": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"coupon_duration": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: s.couponResolver.UpdateCouponDuration(),
				},
				"AddTargetUser": &graphql.Field{
					Type:        graphql.Boolean,
					Description: "Add target user to specific coupon",
					Args: graphql.FieldConfigArgument{
						"coupontarget_name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"coupon_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: s.couponResolver.AddTargetUser(),
				},
				"SetTargetUser": &graphql.Field{
					Type:        graphql.Boolean,
					Description: "Set / Update target user to specific coupon",
					Args: graphql.FieldConfigArgument{
						"coupontarget_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
						"coupontarget_name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"coupon_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: s.couponResolver.SetTargetUser(),
				},
			},
		}),
	})

	if err != nil {
		return err
	}

	s.Schema = schema

	return nil
}
