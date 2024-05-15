// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/cend-org/duval/graph/model"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type MutationResolver interface {
	NewStudent(ctx context.Context, email string) (*model.BearerToken, error)
	NewParent(ctx context.Context, email string) (*model.BearerToken, error)
	NewTutor(ctx context.Context, email string) (*model.BearerToken, error)
	NewProfessor(ctx context.Context, email string) (*model.BearerToken, error)
	NewPassword(ctx context.Context, password model.PasswordInput) (*bool, error)
	Login(ctx context.Context, email string, password string) (*model.BearerToken, error)
	UpdateMyProfile(ctx context.Context, profile model.UserInput) (*model.User, error)
	UpdateProfileAndPassword(ctx context.Context, profile model.UserInput, password model.PasswordInput) (*model.User, error)
	NewUserAcademicCourses(ctx context.Context, courses []*model.UserAcademicCourseInput) (*bool, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) field_Mutation_Login_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["email"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("email"))
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["email"] = arg0
	var arg1 string
	if tmp, ok := rawArgs["password"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("password"))
		arg1, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["password"] = arg1
	return args, nil
}

func (ec *executionContext) field_Mutation_NewParent_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["email"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("email"))
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["email"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_NewPassword_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 model.PasswordInput
	if tmp, ok := rawArgs["password"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("password"))
		arg0, err = ec.unmarshalNPasswordInput2githubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐPasswordInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["password"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_NewProfessor_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["email"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("email"))
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["email"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_NewStudent_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["email"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("email"))
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["email"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_NewTutor_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["email"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("email"))
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["email"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_NewUserAcademicCourses_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 []*model.UserAcademicCourseInput
	if tmp, ok := rawArgs["courses"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("courses"))
		arg0, err = ec.unmarshalNUserAcademicCourseInput2ᚕᚖgithubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐUserAcademicCourseInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["courses"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_UpdateMyProfile_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 model.UserInput
	if tmp, ok := rawArgs["profile"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("profile"))
		arg0, err = ec.unmarshalNUserInput2githubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐUserInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["profile"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_UpdateProfileAndPassword_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 model.UserInput
	if tmp, ok := rawArgs["profile"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("profile"))
		arg0, err = ec.unmarshalNUserInput2githubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐUserInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["profile"] = arg0
	var arg1 model.PasswordInput
	if tmp, ok := rawArgs["password"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("password"))
		arg1, err = ec.unmarshalNPasswordInput2githubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐPasswordInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["password"] = arg1
	return args, nil
}

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Mutation_NewStudent(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_NewStudent(ctx, field)
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
		return ec.resolvers.Mutation().NewStudent(rctx, fc.Args["email"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*model.BearerToken)
	fc.Result = res
	return ec.marshalOBearerToken2ᚖgithubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐBearerToken(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_NewStudent(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "T":
				return ec.fieldContext_BearerToken_T(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type BearerToken", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_NewStudent_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_NewParent(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_NewParent(ctx, field)
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
		return ec.resolvers.Mutation().NewParent(rctx, fc.Args["email"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*model.BearerToken)
	fc.Result = res
	return ec.marshalOBearerToken2ᚖgithubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐBearerToken(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_NewParent(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "T":
				return ec.fieldContext_BearerToken_T(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type BearerToken", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_NewParent_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_NewTutor(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_NewTutor(ctx, field)
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
		return ec.resolvers.Mutation().NewTutor(rctx, fc.Args["email"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*model.BearerToken)
	fc.Result = res
	return ec.marshalOBearerToken2ᚖgithubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐBearerToken(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_NewTutor(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "T":
				return ec.fieldContext_BearerToken_T(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type BearerToken", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_NewTutor_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_NewProfessor(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_NewProfessor(ctx, field)
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
		return ec.resolvers.Mutation().NewProfessor(rctx, fc.Args["email"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*model.BearerToken)
	fc.Result = res
	return ec.marshalOBearerToken2ᚖgithubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐBearerToken(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_NewProfessor(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "T":
				return ec.fieldContext_BearerToken_T(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type BearerToken", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_NewProfessor_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_NewPassword(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_NewPassword(ctx, field)
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
		return ec.resolvers.Mutation().NewPassword(rctx, fc.Args["password"].(model.PasswordInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*bool)
	fc.Result = res
	return ec.marshalOBoolean2ᚖbool(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_NewPassword(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Boolean does not have child fields")
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_NewPassword_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_Login(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_Login(ctx, field)
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
		return ec.resolvers.Mutation().Login(rctx, fc.Args["email"].(string), fc.Args["password"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*model.BearerToken)
	fc.Result = res
	return ec.marshalOBearerToken2ᚖgithubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐBearerToken(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_Login(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "T":
				return ec.fieldContext_BearerToken_T(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type BearerToken", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_Login_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_UpdateMyProfile(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_UpdateMyProfile(ctx, field)
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
		return ec.resolvers.Mutation().UpdateMyProfile(rctx, fc.Args["profile"].(model.UserInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*model.User)
	fc.Result = res
	return ec.marshalOUser2ᚖgithubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐUser(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_UpdateMyProfile(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "Id":
				return ec.fieldContext_User_Id(ctx, field)
			case "CreatedAt":
				return ec.fieldContext_User_CreatedAt(ctx, field)
			case "UpdatedAt":
				return ec.fieldContext_User_UpdatedAt(ctx, field)
			case "DeletedAt":
				return ec.fieldContext_User_DeletedAt(ctx, field)
			case "Name":
				return ec.fieldContext_User_Name(ctx, field)
			case "FamilyName":
				return ec.fieldContext_User_FamilyName(ctx, field)
			case "NickName":
				return ec.fieldContext_User_NickName(ctx, field)
			case "Email":
				return ec.fieldContext_User_Email(ctx, field)
			case "Matricule":
				return ec.fieldContext_User_Matricule(ctx, field)
			case "Age":
				return ec.fieldContext_User_Age(ctx, field)
			case "BirthDate":
				return ec.fieldContext_User_BirthDate(ctx, field)
			case "Sex":
				return ec.fieldContext_User_Sex(ctx, field)
			case "Lang":
				return ec.fieldContext_User_Lang(ctx, field)
			case "Status":
				return ec.fieldContext_User_Status(ctx, field)
			case "ProfileImageXid":
				return ec.fieldContext_User_ProfileImageXid(ctx, field)
			case "Description":
				return ec.fieldContext_User_Description(ctx, field)
			case "CoverText":
				return ec.fieldContext_User_CoverText(ctx, field)
			case "Profile":
				return ec.fieldContext_User_Profile(ctx, field)
			case "ExperienceDetail":
				return ec.fieldContext_User_ExperienceDetail(ctx, field)
			case "AdditionalDescription":
				return ec.fieldContext_User_AdditionalDescription(ctx, field)
			case "AddOnTitle":
				return ec.fieldContext_User_AddOnTitle(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type User", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_UpdateMyProfile_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_UpdateProfileAndPassword(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_UpdateProfileAndPassword(ctx, field)
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
		return ec.resolvers.Mutation().UpdateProfileAndPassword(rctx, fc.Args["profile"].(model.UserInput), fc.Args["password"].(model.PasswordInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*model.User)
	fc.Result = res
	return ec.marshalOUser2ᚖgithubᚗcomᚋcendᚑorgᚋduvalᚋgraphᚋmodelᚐUser(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_UpdateProfileAndPassword(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "Id":
				return ec.fieldContext_User_Id(ctx, field)
			case "CreatedAt":
				return ec.fieldContext_User_CreatedAt(ctx, field)
			case "UpdatedAt":
				return ec.fieldContext_User_UpdatedAt(ctx, field)
			case "DeletedAt":
				return ec.fieldContext_User_DeletedAt(ctx, field)
			case "Name":
				return ec.fieldContext_User_Name(ctx, field)
			case "FamilyName":
				return ec.fieldContext_User_FamilyName(ctx, field)
			case "NickName":
				return ec.fieldContext_User_NickName(ctx, field)
			case "Email":
				return ec.fieldContext_User_Email(ctx, field)
			case "Matricule":
				return ec.fieldContext_User_Matricule(ctx, field)
			case "Age":
				return ec.fieldContext_User_Age(ctx, field)
			case "BirthDate":
				return ec.fieldContext_User_BirthDate(ctx, field)
			case "Sex":
				return ec.fieldContext_User_Sex(ctx, field)
			case "Lang":
				return ec.fieldContext_User_Lang(ctx, field)
			case "Status":
				return ec.fieldContext_User_Status(ctx, field)
			case "ProfileImageXid":
				return ec.fieldContext_User_ProfileImageXid(ctx, field)
			case "Description":
				return ec.fieldContext_User_Description(ctx, field)
			case "CoverText":
				return ec.fieldContext_User_CoverText(ctx, field)
			case "Profile":
				return ec.fieldContext_User_Profile(ctx, field)
			case "ExperienceDetail":
				return ec.fieldContext_User_ExperienceDetail(ctx, field)
			case "AdditionalDescription":
				return ec.fieldContext_User_AdditionalDescription(ctx, field)
			case "AddOnTitle":
				return ec.fieldContext_User_AddOnTitle(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type User", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_UpdateProfileAndPassword_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_NewUserAcademicCourses(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_NewUserAcademicCourses(ctx, field)
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
		return ec.resolvers.Mutation().NewUserAcademicCourses(rctx, fc.Args["courses"].([]*model.UserAcademicCourseInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*bool)
	fc.Result = res
	return ec.marshalOBoolean2ᚖbool(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_NewUserAcademicCourses(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Boolean does not have child fields")
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_NewUserAcademicCourses_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var mutationImplementors = []string{"Mutation"}

func (ec *executionContext) _Mutation(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, mutationImplementors)
	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "Mutation",
	})

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		innerCtx := graphql.WithRootFieldContext(ctx, &graphql.RootFieldContext{
			Object: field.Name,
			Field:  field,
		})

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Mutation")
		case "NewStudent":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_NewStudent(ctx, field)
			})
		case "NewParent":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_NewParent(ctx, field)
			})
		case "NewTutor":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_NewTutor(ctx, field)
			})
		case "NewProfessor":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_NewProfessor(ctx, field)
			})
		case "NewPassword":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_NewPassword(ctx, field)
			})
		case "Login":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_Login(ctx, field)
			})
		case "UpdateMyProfile":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_UpdateMyProfile(ctx, field)
			})
		case "UpdateProfileAndPassword":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_UpdateProfileAndPassword(ctx, field)
			})
		case "NewUserAcademicCourses":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_NewUserAcademicCourses(ctx, field)
			})
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

// endregion ***************************** type.gotpl *****************************
