package handle

import (
	"{{ProjectName}}/internal/{{appName}}/error/httpError"
	_map "{{ProjectName}}/internal/{{appName}}/map"
	"{{ProjectName}}/internal/{{appName}}/server/{{tableName}}"
	R "{{ProjectName}}/pkg/response"
	"{{ProjectName}}/pkg/validator"
	"github.com/gin-gonic/gin"
)

type {{TableName}} struct{}

func ({{TableName}}) GetAll{{TableName}}(ctx *gin.Context) {
	var page = _map.DefaultPageRequest
	if err := ctx.ShouldBind(&page); err != nil {
		httpError.HandleBadRequest(ctx, nil)
		return
	}
	if err := validator.Struct(page); err != nil {
		httpError.HandleBadRequest(ctx, validator.GetMsg(err))
		return
	}
	if data, total, err := {{tableName}}.GetAllServer(page); err != nil {
		R.Error(ctx, err.Error(), nil)
	} else {
		R.Ok(ctx, R.MSG_OK, R.Page(total, page.Page, page.PageSize, data))
	}
	return
}

func ({{TableName}}) Post{{TableName}}(ctx *gin.Context) {
	var addMap _map.{{TableName}}AddServer
	if err := ctx.ShouldBind(&addMap); err != nil {
		httpError.HandleBadRequest(ctx, nil)
		return
	}
	if err := validator.Struct(addMap); err != nil {
		httpError.HandleBadRequest(ctx, validator.GetMsg(err))
		return
	}
	if err := {{tableName}}.AddServer(addMap); err != nil {
		R.Error(ctx, err.Error(), nil)
		return
	}
	R.Ok(ctx, R.MSG_OK, nil)
	return
}

func ({{TableName}}) Get{{TableName}}(ctx *gin.Context) {
	var id _map.IdMap
	if err := ctx.ShouldBindUri(&id); err != nil {
		httpError.HandleBadRequest(ctx, nil)
		return
	}
	if err := validator.Struct(id); err != nil {
		httpError.HandleBadRequest(ctx, validator.GetMsg(err))
		return
	}
	if data, err := {{tableName}}.GetByIdServer(id); err != nil {
		R.Error(ctx, err.Error(), nil)
	} else {
		R.Ok(ctx, R.MSG_OK, data)
	}
	return
}

func ({{TableName}}) Put{{TableName}}(ctx *gin.Context) {
	var put _map.{{TableName}}PutServer
	if err := ctx.ShouldBind(&put); err != nil {
		httpError.HandleBadRequest(ctx, nil)
		return
	}
	if err := validator.Struct(put); err != nil {
		httpError.HandleBadRequest(ctx, validator.GetMsg(err))
		return
	}
	if err := {{tableName}}.PutByIdServer(put); err != nil {
		R.Error(ctx, err.Error(), nil)
	} else {
		R.Ok(ctx, R.MSG_OK, nil)
	}
	return
}

func ({{TableName}}) Del{{TableName}}(ctx *gin.Context) {
	var del _map.IdMap
	if err := ctx.ShouldBind(&del); err != nil {
		httpError.HandleBadRequest(ctx, nil)
		return
	}
	if err := validator.Struct(del); err != nil {
		httpError.HandleBadRequest(ctx, validator.GetMsg(err))
		return
	}
	if err := {{tableName}}.DelServer(del); err != nil {
		R.Error(ctx, err.Error(), nil)
	} else {
		R.Ok(ctx, R.MSG_OK, nil)
	}
	return
}

func ({{TableName}}) Rec{{TableName}}(ctx *gin.Context) {
	var recDel _map.IdMap
	if err := ctx.ShouldBind(&recDel); err != nil {
		httpError.HandleBadRequest(ctx, nil)
		return
	}
	if err := validator.Struct(recDel); err != nil {
		httpError.HandleBadRequest(ctx, validator.GetMsg(err))
		return
	}
	if err := {{tableName}}.RecDelServer(recDel); err != nil {
		R.Error(ctx, err.Error(), nil)
	} else {
		R.Ok(ctx, R.MSG_OK, nil)
	}
	return
}
