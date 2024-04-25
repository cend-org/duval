// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/cend-org/duval/graph/model"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Authorization_id(ctx context.Context, field graphql.CollectedField, obj *model.Authorization) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Authorization_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNID2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Authorization_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Authorization",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Authorization_createdAt(ctx context.Context, field graphql.CollectedField, obj *model.Authorization) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Authorization_createdAt(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.CreatedAt, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(time.Time)
	fc.Result = res
	return ec.marshalNDateTime2timeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Authorization_createdAt(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Authorization",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type DateTime does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Authorization_updatedAt(ctx context.Context, field graphql.CollectedField, obj *model.Authorization) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Authorization_updatedAt(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.UpdatedAt, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(time.Time)
	fc.Result = res
	return ec.marshalNDateTime2timeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Authorization_updatedAt(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Authorization",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type DateTime does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Authorization_deletedAt(ctx context.Context, field graphql.CollectedField, obj *model.Authorization) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Authorization_deletedAt(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DeletedAt, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*time.Time)
	fc.Result = res
	return ec.marshalODateTime2ᚖtimeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Authorization_deletedAt(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Authorization",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type DateTime does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Authorization_userId(ctx context.Context, field graphql.CollectedField, obj *model.Authorization) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Authorization_userId(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.UserID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Authorization_userId(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Authorization",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Authorization_AccessLevel(ctx context.Context, field graphql.CollectedField, obj *model.Authorization) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Authorization_AccessLevel(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.AccessLevel, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Authorization_AccessLevel(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Authorization",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _UserAuthorizationLink_Id(ctx context.Context, field graphql.CollectedField, obj *model.UserAuthorizationLink) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UserAuthorizationLink_Id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNID2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UserAuthorizationLink_Id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UserAuthorizationLink",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _UserAuthorizationLink_CreatedAt(ctx context.Context, field graphql.CollectedField, obj *model.UserAuthorizationLink) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UserAuthorizationLink_CreatedAt(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.CreatedAt, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(time.Time)
	fc.Result = res
	return ec.marshalNDateTime2timeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UserAuthorizationLink_CreatedAt(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UserAuthorizationLink",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type DateTime does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _UserAuthorizationLink_UpdatedAt(ctx context.Context, field graphql.CollectedField, obj *model.UserAuthorizationLink) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UserAuthorizationLink_UpdatedAt(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.UpdatedAt, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(time.Time)
	fc.Result = res
	return ec.marshalNDateTime2timeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UserAuthorizationLink_UpdatedAt(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UserAuthorizationLink",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type DateTime does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _UserAuthorizationLink_DeletedAt(ctx context.Context, field graphql.CollectedField, obj *model.UserAuthorizationLink) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UserAuthorizationLink_DeletedAt(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DeletedAt, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*time.Time)
	fc.Result = res
	return ec.marshalODateTime2ᚖtimeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UserAuthorizationLink_DeletedAt(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UserAuthorizationLink",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type DateTime does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _UserAuthorizationLink_LinkType(ctx context.Context, field graphql.CollectedField, obj *model.UserAuthorizationLink) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UserAuthorizationLink_LinkType(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.LinkType, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UserAuthorizationLink_LinkType(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UserAuthorizationLink",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _UserAuthorizationLink_Actors(ctx context.Context, field graphql.CollectedField, obj *model.UserAuthorizationLink) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UserAuthorizationLink_Actors(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Actors, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]*model.UserAuthorizationLinkActor)
	fc.Result = res
	return ec.marshalNUserAuthorizationLinkActor2ᚕᚖgithubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐUserAuthorizationLinkActor(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UserAuthorizationLink_Actors(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UserAuthorizationLink",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "Id":
				return ec.fieldContext_UserAuthorizationLinkActor_Id(ctx, field)
			case "CreatedAt":
				return ec.fieldContext_UserAuthorizationLinkActor_CreatedAt(ctx, field)
			case "UpdatedAt":
				return ec.fieldContext_UserAuthorizationLinkActor_UpdatedAt(ctx, field)
			case "DeletedAt":
				return ec.fieldContext_UserAuthorizationLinkActor_DeletedAt(ctx, field)
			case "UserAuthorizationLink":
				return ec.fieldContext_UserAuthorizationLinkActor_UserAuthorizationLink(ctx, field)
			case "AuthorizationId":
				return ec.fieldContext_UserAuthorizationLinkActor_AuthorizationId(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type UserAuthorizationLinkActor", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _UserAuthorizationLinkActor_Id(ctx context.Context, field graphql.CollectedField, obj *model.UserAuthorizationLinkActor) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UserAuthorizationLinkActor_Id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNID2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UserAuthorizationLinkActor_Id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UserAuthorizationLinkActor",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _UserAuthorizationLinkActor_CreatedAt(ctx context.Context, field graphql.CollectedField, obj *model.UserAuthorizationLinkActor) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UserAuthorizationLinkActor_CreatedAt(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.CreatedAt, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(time.Time)
	fc.Result = res
	return ec.marshalNDateTime2timeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UserAuthorizationLinkActor_CreatedAt(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UserAuthorizationLinkActor",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type DateTime does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _UserAuthorizationLinkActor_UpdatedAt(ctx context.Context, field graphql.CollectedField, obj *model.UserAuthorizationLinkActor) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UserAuthorizationLinkActor_UpdatedAt(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.UpdatedAt, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(time.Time)
	fc.Result = res
	return ec.marshalNDateTime2timeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UserAuthorizationLinkActor_UpdatedAt(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UserAuthorizationLinkActor",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type DateTime does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _UserAuthorizationLinkActor_DeletedAt(ctx context.Context, field graphql.CollectedField, obj *model.UserAuthorizationLinkActor) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UserAuthorizationLinkActor_DeletedAt(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DeletedAt, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*time.Time)
	fc.Result = res
	return ec.marshalODateTime2ᚖtimeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UserAuthorizationLinkActor_DeletedAt(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UserAuthorizationLinkActor",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type DateTime does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _UserAuthorizationLinkActor_UserAuthorizationLink(ctx context.Context, field graphql.CollectedField, obj *model.UserAuthorizationLinkActor) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UserAuthorizationLinkActor_UserAuthorizationLink(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.UserAuthorizationLink, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.UserAuthorizationLink)
	fc.Result = res
	return ec.marshalNUserAuthorizationLink2ᚖgithubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐUserAuthorizationLink(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UserAuthorizationLinkActor_UserAuthorizationLink(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UserAuthorizationLinkActor",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "Id":
				return ec.fieldContext_UserAuthorizationLink_Id(ctx, field)
			case "CreatedAt":
				return ec.fieldContext_UserAuthorizationLink_CreatedAt(ctx, field)
			case "UpdatedAt":
				return ec.fieldContext_UserAuthorizationLink_UpdatedAt(ctx, field)
			case "DeletedAt":
				return ec.fieldContext_UserAuthorizationLink_DeletedAt(ctx, field)
			case "LinkType":
				return ec.fieldContext_UserAuthorizationLink_LinkType(ctx, field)
			case "Actors":
				return ec.fieldContext_UserAuthorizationLink_Actors(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type UserAuthorizationLink", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _UserAuthorizationLinkActor_AuthorizationId(ctx context.Context, field graphql.CollectedField, obj *model.UserAuthorizationLinkActor) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UserAuthorizationLinkActor_AuthorizationId(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.AuthorizationID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UserAuthorizationLinkActor_AuthorizationId(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UserAuthorizationLinkActor",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var authorizationImplementors = []string{"Authorization"}

func (ec *executionContext) _Authorization(ctx context.Context, sel ast.SelectionSet, obj *model.Authorization) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, authorizationImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Authorization")
		case "id":
			out.Values[i] = ec._Authorization_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "createdAt":
			out.Values[i] = ec._Authorization_createdAt(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "updatedAt":
			out.Values[i] = ec._Authorization_updatedAt(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "deletedAt":
			out.Values[i] = ec._Authorization_deletedAt(ctx, field, obj)
		case "userId":
			out.Values[i] = ec._Authorization_userId(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "AccessLevel":
			out.Values[i] = ec._Authorization_AccessLevel(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var userAuthorizationLinkImplementors = []string{"UserAuthorizationLink"}

func (ec *executionContext) _UserAuthorizationLink(ctx context.Context, sel ast.SelectionSet, obj *model.UserAuthorizationLink) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, userAuthorizationLinkImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("UserAuthorizationLink")
		case "Id":
			out.Values[i] = ec._UserAuthorizationLink_Id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "CreatedAt":
			out.Values[i] = ec._UserAuthorizationLink_CreatedAt(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "UpdatedAt":
			out.Values[i] = ec._UserAuthorizationLink_UpdatedAt(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "DeletedAt":
			out.Values[i] = ec._UserAuthorizationLink_DeletedAt(ctx, field, obj)
		case "LinkType":
			out.Values[i] = ec._UserAuthorizationLink_LinkType(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "Actors":
			out.Values[i] = ec._UserAuthorizationLink_Actors(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var userAuthorizationLinkActorImplementors = []string{"UserAuthorizationLinkActor"}

func (ec *executionContext) _UserAuthorizationLinkActor(ctx context.Context, sel ast.SelectionSet, obj *model.UserAuthorizationLinkActor) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, userAuthorizationLinkActorImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("UserAuthorizationLinkActor")
		case "Id":
			out.Values[i] = ec._UserAuthorizationLinkActor_Id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "CreatedAt":
			out.Values[i] = ec._UserAuthorizationLinkActor_CreatedAt(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "UpdatedAt":
			out.Values[i] = ec._UserAuthorizationLinkActor_UpdatedAt(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "DeletedAt":
			out.Values[i] = ec._UserAuthorizationLinkActor_DeletedAt(ctx, field, obj)
		case "UserAuthorizationLink":
			out.Values[i] = ec._UserAuthorizationLinkActor_UserAuthorizationLink(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "AuthorizationId":
			out.Values[i] = ec._UserAuthorizationLinkActor_AuthorizationId(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNUserAuthorizationLink2githubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐUserAuthorizationLink(ctx context.Context, sel ast.SelectionSet, v model.UserAuthorizationLink) graphql.Marshaler {
	return ec._UserAuthorizationLink(ctx, sel, &v)
}

func (ec *executionContext) marshalNUserAuthorizationLink2ᚕgithubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐUserAuthorizationLinkᚄ(ctx context.Context, sel ast.SelectionSet, v []model.UserAuthorizationLink) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNUserAuthorizationLink2githubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐUserAuthorizationLink(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalNUserAuthorizationLink2ᚖgithubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐUserAuthorizationLink(ctx context.Context, sel ast.SelectionSet, v *model.UserAuthorizationLink) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._UserAuthorizationLink(ctx, sel, v)
}

func (ec *executionContext) marshalNUserAuthorizationLinkActor2ᚕᚖgithubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐUserAuthorizationLinkActor(ctx context.Context, sel ast.SelectionSet, v []*model.UserAuthorizationLinkActor) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalOUserAuthorizationLinkActor2ᚖgithubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐUserAuthorizationLinkActor(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	return ret
}

func (ec *executionContext) marshalOUserAuthorizationLink2ᚖgithubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐUserAuthorizationLink(ctx context.Context, sel ast.SelectionSet, v *model.UserAuthorizationLink) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._UserAuthorizationLink(ctx, sel, v)
}

func (ec *executionContext) marshalOUserAuthorizationLinkActor2ᚖgithubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐUserAuthorizationLinkActor(ctx context.Context, sel ast.SelectionSet, v *model.UserAuthorizationLinkActor) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._UserAuthorizationLinkActor(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
