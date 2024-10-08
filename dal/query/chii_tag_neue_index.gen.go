// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/bangumi/server/dal/dao"
)

func newTagIndex(db *gorm.DB, opts ...gen.DOOption) tagIndex {
	_tagIndex := tagIndex{}

	_tagIndex.tagIndexDo.UseDB(db, opts...)
	_tagIndex.tagIndexDo.UseModel(&dao.TagIndex{})

	tableName := _tagIndex.tagIndexDo.TableName()
	_tagIndex.ALL = field.NewAsterisk(tableName)
	_tagIndex.ID = field.NewUint32(tableName, "tag_id")
	_tagIndex.Name = field.NewString(tableName, "tag_name")
	_tagIndex.Cat = field.NewInt8(tableName, "tag_cat")
	_tagIndex.Type = field.NewUint8(tableName, "tag_type")
	_tagIndex.Results = field.NewUint32(tableName, "tag_results")
	_tagIndex.CreatedTime = field.NewUint32(tableName, "tag_dateline")
	_tagIndex.UpdatedTime = field.NewUint32(tableName, "tag_lasttouch")

	_tagIndex.fillFieldMap()

	return _tagIndex
}

type tagIndex struct {
	tagIndexDo tagIndexDo

	ALL         field.Asterisk
	ID          field.Uint32
	Name        field.String
	Cat         field.Int8 // 0=条目 1=日志 2=天窗
	Type        field.Uint8
	Results     field.Uint32
	CreatedTime field.Uint32
	UpdatedTime field.Uint32

	fieldMap map[string]field.Expr
}

func (t tagIndex) Table(newTableName string) *tagIndex {
	t.tagIndexDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tagIndex) As(alias string) *tagIndex {
	t.tagIndexDo.DO = *(t.tagIndexDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tagIndex) updateTableName(table string) *tagIndex {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewUint32(table, "tag_id")
	t.Name = field.NewString(table, "tag_name")
	t.Cat = field.NewInt8(table, "tag_cat")
	t.Type = field.NewUint8(table, "tag_type")
	t.Results = field.NewUint32(table, "tag_results")
	t.CreatedTime = field.NewUint32(table, "tag_dateline")
	t.UpdatedTime = field.NewUint32(table, "tag_lasttouch")

	t.fillFieldMap()

	return t
}

func (t *tagIndex) WithContext(ctx context.Context) *tagIndexDo { return t.tagIndexDo.WithContext(ctx) }

func (t tagIndex) TableName() string { return t.tagIndexDo.TableName() }

func (t tagIndex) Alias() string { return t.tagIndexDo.Alias() }

func (t tagIndex) Columns(cols ...field.Expr) gen.Columns { return t.tagIndexDo.Columns(cols...) }

func (t *tagIndex) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tagIndex) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 7)
	t.fieldMap["tag_id"] = t.ID
	t.fieldMap["tag_name"] = t.Name
	t.fieldMap["tag_cat"] = t.Cat
	t.fieldMap["tag_type"] = t.Type
	t.fieldMap["tag_results"] = t.Results
	t.fieldMap["tag_dateline"] = t.CreatedTime
	t.fieldMap["tag_lasttouch"] = t.UpdatedTime
}

func (t tagIndex) clone(db *gorm.DB) tagIndex {
	t.tagIndexDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tagIndex) replaceDB(db *gorm.DB) tagIndex {
	t.tagIndexDo.ReplaceDB(db)
	return t
}

type tagIndexDo struct{ gen.DO }

func (t tagIndexDo) Debug() *tagIndexDo {
	return t.withDO(t.DO.Debug())
}

func (t tagIndexDo) WithContext(ctx context.Context) *tagIndexDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tagIndexDo) ReadDB() *tagIndexDo {
	return t.Clauses(dbresolver.Read)
}

func (t tagIndexDo) WriteDB() *tagIndexDo {
	return t.Clauses(dbresolver.Write)
}

func (t tagIndexDo) Session(config *gorm.Session) *tagIndexDo {
	return t.withDO(t.DO.Session(config))
}

func (t tagIndexDo) Clauses(conds ...clause.Expression) *tagIndexDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tagIndexDo) Returning(value interface{}, columns ...string) *tagIndexDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tagIndexDo) Not(conds ...gen.Condition) *tagIndexDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tagIndexDo) Or(conds ...gen.Condition) *tagIndexDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tagIndexDo) Select(conds ...field.Expr) *tagIndexDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tagIndexDo) Where(conds ...gen.Condition) *tagIndexDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tagIndexDo) Order(conds ...field.Expr) *tagIndexDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tagIndexDo) Distinct(cols ...field.Expr) *tagIndexDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tagIndexDo) Omit(cols ...field.Expr) *tagIndexDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tagIndexDo) Join(table schema.Tabler, on ...field.Expr) *tagIndexDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tagIndexDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tagIndexDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tagIndexDo) RightJoin(table schema.Tabler, on ...field.Expr) *tagIndexDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tagIndexDo) Group(cols ...field.Expr) *tagIndexDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tagIndexDo) Having(conds ...gen.Condition) *tagIndexDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tagIndexDo) Limit(limit int) *tagIndexDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tagIndexDo) Offset(offset int) *tagIndexDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tagIndexDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tagIndexDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tagIndexDo) Unscoped() *tagIndexDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tagIndexDo) Create(values ...*dao.TagIndex) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tagIndexDo) CreateInBatches(values []*dao.TagIndex, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tagIndexDo) Save(values ...*dao.TagIndex) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tagIndexDo) First() (*dao.TagIndex, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.TagIndex), nil
	}
}

func (t tagIndexDo) Take() (*dao.TagIndex, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.TagIndex), nil
	}
}

func (t tagIndexDo) Last() (*dao.TagIndex, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.TagIndex), nil
	}
}

func (t tagIndexDo) Find() ([]*dao.TagIndex, error) {
	result, err := t.DO.Find()
	return result.([]*dao.TagIndex), err
}

func (t tagIndexDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.TagIndex, err error) {
	buf := make([]*dao.TagIndex, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tagIndexDo) FindInBatches(result *[]*dao.TagIndex, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tagIndexDo) Attrs(attrs ...field.AssignExpr) *tagIndexDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tagIndexDo) Assign(attrs ...field.AssignExpr) *tagIndexDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tagIndexDo) Joins(fields ...field.RelationField) *tagIndexDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tagIndexDo) Preload(fields ...field.RelationField) *tagIndexDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tagIndexDo) FirstOrInit() (*dao.TagIndex, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.TagIndex), nil
	}
}

func (t tagIndexDo) FirstOrCreate() (*dao.TagIndex, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.TagIndex), nil
	}
}

func (t tagIndexDo) FindByPage(offset int, limit int) (result []*dao.TagIndex, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t tagIndexDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tagIndexDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tagIndexDo) Delete(models ...*dao.TagIndex) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tagIndexDo) withDO(do gen.Dao) *tagIndexDo {
	t.DO = *do.(*gen.DO)
	return t
}
