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
					Args:        graphql.FieldConfigArgument{},
					Resolve:     s.couponResolver.CreateCoupons(),
				},
				"UpdateCouponDuration": &graphql.Field{
					Type:        graphql.Boolean,
					Description: "Update coupon's duration by ID",
					Args: graphql.FieldConfigArgument{
						"coupon_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
						"start_date": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"end_date": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"duration": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: s.couponResolver.UpdateCouponDuration(),
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
