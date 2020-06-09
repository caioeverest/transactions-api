package repository

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/caioeverest/transactions-api/util"
)

type Repository struct {
	c Conn
	m Model
}

func New(c Conn, m Model) *Repository { return &Repository{c, m} }

func (ar *Repository) Save(model Model) error {
	return ar.c.Create(model).Error
}

func (ar *Repository) Find(query, model Model) (err error) {
	err = ar.c.Where(query).Find(model).Error
	return
}

func (ar *Repository) FindAll(model Model) (err error) {
	err = ar.c.Find(model).Error
	return
}

func (ar *Repository) FindById(id int, model Model) (err error) {
	whereTempl := fmt.Sprintf("%s = ?", ar.getPK())
	err = ar.c.First(model, whereTempl, id).Error
	return
}

func (ar *Repository) Update(id int, changes Model) (err error) {
	whereTempl := fmt.Sprintf("%s = ?", ar.getPK())
	err = ar.c.Model(changes).Where(whereTempl, id).Update(changes).Error
	return
}

func (ar *Repository) Delete(id int) error {
	whereTempl := fmt.Sprintf("%s = ?", ar.getPK())
	rows := ar.c.Delete(ar.m, whereTempl, id).RowsAffected
	if rows == 0 {
		return util.Error("register not found")
	}
	return nil
}

func (ar *Repository) getPK() string {
	reelectedModel := reflect.TypeOf(ar.m)
	for i := 0; i < reelectedModel.NumField(); i++ {
		if ext, _, _ := readTag("primary_key", reelectedModel.Field(i)); ext {
			return getColumnName(reelectedModel.Field(i))
		}
	}
	return reelectedModel.Field(0).Name
}

func getColumnName(field reflect.StructField) string {
	ext, value, fieldName := readTag("Column", field)
	if ext {
		return value
	}
	return fieldName
}

func readTag(tagname string, field reflect.StructField) (ext bool, value string, fieldName string) {
	tagString := field.Tag.Get("gorm")
	tags := strings.Split(tagString, ";")
	for _, tag := range tags {
		if strings.Contains(tag, tagname) {
			return true, returnValue(tag), field.Name
		}
	}
	return false, "", field.Name
}

func returnValue(value string) string {
	slice := strings.SplitAfter(value, ":")
	if len(slice) > 1 {
		return slice[1]
	}
	return ""
}
