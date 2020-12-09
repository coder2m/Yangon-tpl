package {{tableName}}

import (
	_map "{{ProjectName}}/internal/{{appName}}/map"
	"{{ProjectName}}/internal/{{appName}}/model/{{tableName}}"
	"github.com/jinzhu/gorm"
)

func GetAllServer(page _map.PageList) (data []{{tableName}}.{{TableName}}, total int64, err error) {
	data = make([]{{tableName}}.{{TableName}}, 0)
	total, err = new({{tableName}}.{{TableName}}).Get(page.PageSize*(page.Page-1), page.PageSize, &data, nil, page.IsDelete)
	if !gorm.IsRecordNotFoundError(err) {
		log.Info("GetAllServer", err.Error())
	}
	return
}

func AddServer(add _map.{{TableName}}AddServer) (err error) {
	//todo 添加手动赋值
	data := &{{tableName}}.{{TableName}}{

	}
	err = data.Add()
	return
}

func GetByIdServer(idMap _map.IdMap) (data *{{tableName}}.{{TableName}}, err error) {
	data = new({{tableName}}.{{TableName}})
	data.{{Id}} = idMap.Id
	err = data.GetById()
	if !gorm.IsRecordNotFoundError(err) {
		log.Info("GetByIdServer", err.Error())
	}
	return
}

func PutByIdServer(put _map.{{TableName}}PutServer) (err error) {
	//todo 修改手动赋值
	data := &{{tableName}}.{{TableName}}{

	}
	err = data.UpdateById()
	if !gorm.IsRecordNotFoundError(err) {
		log.Info("PutByIdServer", err.Error())
	}
	return
}

func DelServer(idMap _map.IdMap) (err error) {
	err = new({{tableName}}.{{TableName}}).Del(map[string]interface{}{"{{id}}=?": idMap.Id})
	if !gorm.IsRecordNotFoundError(err) {
		log.Info("DelServer", err.Error())
	}
	return
}

func RecDelServer(idMap _map.IdMap) (err error) {
	data = new({{tableName}}.{{TableName}})
	data.{{Id}} = idMap.Id
	err = data.DelRes()
	if !gorm.IsRecordNotFoundError(err) {
		log.Info("RecDelServer", err.Error())
	}
	return
}
