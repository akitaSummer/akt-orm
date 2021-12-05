package schema

import (
	"aktorm/dialect"
	"go/ast"
	"reflect"
)

type Field struct {
	Name string // 字段名
	Type string // 类型
	Tag  string // 约束条件
}

type Schema struct {
	Model      interface{} // 被映射的对象
	Name       string      // 表名
	Fields     []*Field    // 字段
	FieldNames []string    //列名
	fieldMap   map[string]*Field
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

// 任意的对象解析为Schema实例

func Parse(dest interface{}, d dialect.Dialect) *Schema {
	// reflect.Indirect()获取指针指向的实例 reflect.ValueOf()获取值 reflect.Type()获取类型
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(), // 获取到结构体的名称
		fieldMap: make(map[string]*Field),
	}
	// modelType.NumField()获取实例的字段的个数
	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i) // 通过下标获取到特定字段
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("aktorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}
