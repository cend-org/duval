// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/nrfta/go-graphql-scalars"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNDateTime2timeᚐTime(ctx context.Context, v interface{}) (time.Time, error) {
	res, err := scalars.UnmarshalDateTime(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNDateTime2timeᚐTime(ctx context.Context, sel ast.SelectionSet, v time.Time) graphql.Marshaler {
	res := scalars.MarshalDateTime(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNUpload2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚐUpload(ctx context.Context, v interface{}) (graphql.Upload, error) {
	res, err := graphql.UnmarshalUpload(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNUpload2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚐUpload(ctx context.Context, sel ast.SelectionSet, v graphql.Upload) graphql.Marshaler {
	res := graphql.MarshalUpload(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
	}
	return res
}

func (ec *executionContext) unmarshalODateTime2ᚖtimeᚐTime(ctx context.Context, v interface{}) (*time.Time, error) {
	if v == nil {
		return nil, nil
	}
	res, err := scalars.UnmarshalDateTime(v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalODateTime2ᚖtimeᚐTime(ctx context.Context, sel ast.SelectionSet, v *time.Time) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	res := scalars.MarshalDateTime(*v)
	return res
}

// endregion ***************************** type.gotpl *****************************