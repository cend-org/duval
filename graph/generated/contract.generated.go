// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
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

func (ec *executionContext) _Contract_Id(ctx context.Context, field graphql.CollectedField, obj *model.Contract) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Contract_Id(ctx, field)
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
		return obj.Id, nil
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

func (ec *executionContext) fieldContext_Contract_Id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Contract",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Contract_CreatedAt(ctx context.Context, field graphql.CollectedField, obj *model.Contract) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Contract_CreatedAt(ctx, field)
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

func (ec *executionContext) fieldContext_Contract_CreatedAt(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Contract",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type DateTime does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Contract_UpdatedAt(ctx context.Context, field graphql.CollectedField, obj *model.Contract) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Contract_UpdatedAt(ctx, field)
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

func (ec *executionContext) fieldContext_Contract_UpdatedAt(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Contract",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type DateTime does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Contract_DeletedAt(ctx context.Context, field graphql.CollectedField, obj *model.Contract) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Contract_DeletedAt(ctx, field)
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

func (ec *executionContext) fieldContext_Contract_DeletedAt(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Contract",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type DateTime does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Contract_TutorId(ctx context.Context, field graphql.CollectedField, obj *model.Contract) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Contract_TutorId(ctx, field)
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
		return obj.TutorId, nil
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

func (ec *executionContext) fieldContext_Contract_TutorId(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Contract",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Contract_ParentId(ctx context.Context, field graphql.CollectedField, obj *model.Contract) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Contract_ParentId(ctx, field)
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
		return obj.ParentId, nil
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

func (ec *executionContext) fieldContext_Contract_ParentId(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Contract",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Contract_StudentId(ctx context.Context, field graphql.CollectedField, obj *model.Contract) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Contract_StudentId(ctx, field)
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
		return obj.StudentId, nil
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

func (ec *executionContext) fieldContext_Contract_StudentId(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Contract",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Contract_StartDate(ctx context.Context, field graphql.CollectedField, obj *model.Contract) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Contract_StartDate(ctx, field)
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
		return obj.StartDate, nil
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

func (ec *executionContext) fieldContext_Contract_StartDate(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Contract",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type DateTime does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Contract_EndDate(ctx context.Context, field graphql.CollectedField, obj *model.Contract) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Contract_EndDate(ctx, field)
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
		return obj.EndDate, nil
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

func (ec *executionContext) fieldContext_Contract_EndDate(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Contract",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type DateTime does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Contract_PaymentType(ctx context.Context, field graphql.CollectedField, obj *model.Contract) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Contract_PaymentType(ctx, field)
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
		return obj.PaymentType, nil
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

func (ec *executionContext) fieldContext_Contract_PaymentType(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Contract",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Contract_SalaryValue(ctx context.Context, field graphql.CollectedField, obj *model.Contract) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Contract_SalaryValue(ctx, field)
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
		return obj.SalaryValue, nil
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
	res := resTmp.(float64)
	fc.Result = res
	return ec.marshalNFloat2float64(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Contract_SalaryValue(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Contract",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Float does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Contract_PaymentMethod(ctx context.Context, field graphql.CollectedField, obj *model.Contract) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Contract_PaymentMethod(ctx, field)
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
		return obj.PaymentMethod, nil
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

func (ec *executionContext) fieldContext_Contract_PaymentMethod(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Contract",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ContractTimesheetDetail_Id(ctx context.Context, field graphql.CollectedField, obj *model.ContractTimesheetDetail) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ContractTimesheetDetail_Id(ctx, field)
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
		return obj.Id, nil
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

func (ec *executionContext) fieldContext_ContractTimesheetDetail_Id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ContractTimesheetDetail",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ContractTimesheetDetail_CreatedAt(ctx context.Context, field graphql.CollectedField, obj *model.ContractTimesheetDetail) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ContractTimesheetDetail_CreatedAt(ctx, field)
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

func (ec *executionContext) fieldContext_ContractTimesheetDetail_CreatedAt(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ContractTimesheetDetail",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type DateTime does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ContractTimesheetDetail_UpdatedAt(ctx context.Context, field graphql.CollectedField, obj *model.ContractTimesheetDetail) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ContractTimesheetDetail_UpdatedAt(ctx, field)
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

func (ec *executionContext) fieldContext_ContractTimesheetDetail_UpdatedAt(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ContractTimesheetDetail",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type DateTime does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ContractTimesheetDetail_DeletedAt(ctx context.Context, field graphql.CollectedField, obj *model.ContractTimesheetDetail) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ContractTimesheetDetail_DeletedAt(ctx, field)
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

func (ec *executionContext) fieldContext_ContractTimesheetDetail_DeletedAt(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ContractTimesheetDetail",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type DateTime does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ContractTimesheetDetail_ContractId(ctx context.Context, field graphql.CollectedField, obj *model.ContractTimesheetDetail) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ContractTimesheetDetail_ContractId(ctx, field)
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
		return obj.ContractId, nil
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

func (ec *executionContext) fieldContext_ContractTimesheetDetail_ContractId(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ContractTimesheetDetail",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ContractTimesheetDetail_Date(ctx context.Context, field graphql.CollectedField, obj *model.ContractTimesheetDetail) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ContractTimesheetDetail_Date(ctx, field)
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
		return obj.Date, nil
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
	return ec.marshalNDate2timeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ContractTimesheetDetail_Date(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ContractTimesheetDetail",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Date does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ContractTimesheetDetail_Hours(ctx context.Context, field graphql.CollectedField, obj *model.ContractTimesheetDetail) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ContractTimesheetDetail_Hours(ctx, field)
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
		return obj.Hours, nil
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
	res := resTmp.(float64)
	fc.Result = res
	return ec.marshalNFloat2float64(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ContractTimesheetDetail_Hours(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ContractTimesheetDetail",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Float does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputContractInput(ctx context.Context, obj interface{}) (model.ContractInput, error) {
	var it model.ContractInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"TutorId", "ParentId", "StudentId", "StartDate", "EndDate", "PaymentType", "SalaryValue", "PaymentMethod"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "TutorId":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("TutorId"))
			data, err := ec.unmarshalOInt2ᚖint(ctx, v)
			if err != nil {
				return it, err
			}
			it.TutorId = data
		case "ParentId":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("ParentId"))
			data, err := ec.unmarshalOInt2ᚖint(ctx, v)
			if err != nil {
				return it, err
			}
			it.ParentId = data
		case "StudentId":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("StudentId"))
			data, err := ec.unmarshalOInt2ᚖint(ctx, v)
			if err != nil {
				return it, err
			}
			it.StudentId = data
		case "StartDate":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("StartDate"))
			data, err := ec.unmarshalODateTime2ᚖtimeᚐTime(ctx, v)
			if err != nil {
				return it, err
			}
			it.StartDate = data
		case "EndDate":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("EndDate"))
			data, err := ec.unmarshalODateTime2ᚖtimeᚐTime(ctx, v)
			if err != nil {
				return it, err
			}
			it.EndDate = data
		case "PaymentType":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("PaymentType"))
			data, err := ec.unmarshalOInt2ᚖint(ctx, v)
			if err != nil {
				return it, err
			}
			it.PaymentType = data
		case "SalaryValue":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("SalaryValue"))
			data, err := ec.unmarshalOFloat2ᚖfloat64(ctx, v)
			if err != nil {
				return it, err
			}
			it.SalaryValue = data
		case "PaymentMethod":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("PaymentMethod"))
			data, err := ec.unmarshalOInt2ᚖint(ctx, v)
			if err != nil {
				return it, err
			}
			it.PaymentMethod = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputContractTimesheetDetailInput(ctx context.Context, obj interface{}) (model.ContractTimesheetDetailInput, error) {
	var it model.ContractTimesheetDetailInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"ContractId", "Date", "Hours"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "ContractId":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("ContractId"))
			data, err := ec.unmarshalOInt2ᚖint(ctx, v)
			if err != nil {
				return it, err
			}
			it.ContractId = data
		case "Date":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("Date"))
			data, err := ec.unmarshalODate2ᚖtimeᚐTime(ctx, v)
			if err != nil {
				return it, err
			}
			it.Date = data
		case "Hours":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("Hours"))
			data, err := ec.unmarshalOFloat2ᚖfloat64(ctx, v)
			if err != nil {
				return it, err
			}
			it.Hours = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var contractImplementors = []string{"Contract"}

func (ec *executionContext) _Contract(ctx context.Context, sel ast.SelectionSet, obj *model.Contract) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, contractImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Contract")
		case "Id":
			out.Values[i] = ec._Contract_Id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "CreatedAt":
			out.Values[i] = ec._Contract_CreatedAt(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "UpdatedAt":
			out.Values[i] = ec._Contract_UpdatedAt(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "DeletedAt":
			out.Values[i] = ec._Contract_DeletedAt(ctx, field, obj)
		case "TutorId":
			out.Values[i] = ec._Contract_TutorId(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "ParentId":
			out.Values[i] = ec._Contract_ParentId(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "StudentId":
			out.Values[i] = ec._Contract_StudentId(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "StartDate":
			out.Values[i] = ec._Contract_StartDate(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "EndDate":
			out.Values[i] = ec._Contract_EndDate(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "PaymentType":
			out.Values[i] = ec._Contract_PaymentType(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "SalaryValue":
			out.Values[i] = ec._Contract_SalaryValue(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "PaymentMethod":
			out.Values[i] = ec._Contract_PaymentMethod(ctx, field, obj)
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

var contractTimesheetDetailImplementors = []string{"ContractTimesheetDetail"}

func (ec *executionContext) _ContractTimesheetDetail(ctx context.Context, sel ast.SelectionSet, obj *model.ContractTimesheetDetail) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, contractTimesheetDetailImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("ContractTimesheetDetail")
		case "Id":
			out.Values[i] = ec._ContractTimesheetDetail_Id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "CreatedAt":
			out.Values[i] = ec._ContractTimesheetDetail_CreatedAt(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "UpdatedAt":
			out.Values[i] = ec._ContractTimesheetDetail_UpdatedAt(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "DeletedAt":
			out.Values[i] = ec._ContractTimesheetDetail_DeletedAt(ctx, field, obj)
		case "ContractId":
			out.Values[i] = ec._ContractTimesheetDetail_ContractId(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "Date":
			out.Values[i] = ec._ContractTimesheetDetail_Date(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "Hours":
			out.Values[i] = ec._ContractTimesheetDetail_Hours(ctx, field, obj)
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

func (ec *executionContext) marshalNContract2githubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐContract(ctx context.Context, sel ast.SelectionSet, v model.Contract) graphql.Marshaler {
	return ec._Contract(ctx, sel, &v)
}

func (ec *executionContext) marshalNContract2ᚕgithubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐContractᚄ(ctx context.Context, sel ast.SelectionSet, v []model.Contract) graphql.Marshaler {
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
			ret[i] = ec.marshalNContract2githubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐContract(ctx, sel, v[i])
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

func (ec *executionContext) marshalNContract2ᚖgithubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐContract(ctx context.Context, sel ast.SelectionSet, v *model.Contract) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._Contract(ctx, sel, v)
}

func (ec *executionContext) unmarshalNContractInput2githubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐContractInput(ctx context.Context, v interface{}) (model.ContractInput, error) {
	res, err := ec.unmarshalInputContractInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNContractTimesheetDetail2githubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐContractTimesheetDetail(ctx context.Context, sel ast.SelectionSet, v model.ContractTimesheetDetail) graphql.Marshaler {
	return ec._ContractTimesheetDetail(ctx, sel, &v)
}

func (ec *executionContext) marshalNContractTimesheetDetail2ᚕgithubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐContractTimesheetDetailᚄ(ctx context.Context, sel ast.SelectionSet, v []model.ContractTimesheetDetail) graphql.Marshaler {
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
			ret[i] = ec.marshalNContractTimesheetDetail2githubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐContractTimesheetDetail(ctx, sel, v[i])
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

func (ec *executionContext) marshalNContractTimesheetDetail2ᚖgithubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐContractTimesheetDetail(ctx context.Context, sel ast.SelectionSet, v *model.ContractTimesheetDetail) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._ContractTimesheetDetail(ctx, sel, v)
}

func (ec *executionContext) unmarshalNContractTimesheetDetailInput2githubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐContractTimesheetDetailInput(ctx context.Context, v interface{}) (model.ContractTimesheetDetailInput, error) {
	res, err := ec.unmarshalInputContractTimesheetDetailInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

// endregion ***************************** type.gotpl *****************************