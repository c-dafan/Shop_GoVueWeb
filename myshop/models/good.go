package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Good struct {
	Id    int     `orm:"column(gid);auto"`
	Price float64 `orm:"column(price);digits(5);decimals(2)"`
	Num   int     `orm:"column(num)"`
	Name  string  `orm:"column(name);size(50)"`
	Img   string  `orm:"column(img);size(100);null"`
	Sid   *Seller `orm:"column(sid);rel(fk)"`
	Kind  int8    `orm:"column(kind)"`
}

func (t *Good) TableName() string {
	return "good"
}

func init() {
	orm.RegisterModel(new(Good))
}

// AddGood insert a new Good into database and returns
// last inserted Id on success.
func AddGood(m *Good) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGoodById retrieves Good by Id. Returns error if
// Id doesn't exist
func GetGoodById(id int) (v *Good, err error) {
	o := orm.NewOrm()
	v = &Good{Id: id}
	//o.LoadRelated()
	qs := o.QueryTable(new(Good)).Filter("Id", id).RelatedSel()
	//if err = o.Read(v); err == nil {
	//	return v, nil
	//}
	if err = qs.Limit(1, 0).One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllGood retrieves all Good matches certain condition. Returns empty list if
// no records exist
func GetAllGood(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, num int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Good))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, 0, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, 0, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Good
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		num, err = qs.Count()
		return ml, num, nil
	}
	return nil, 0, err
}

// UpdateGood updates Good by Id and returns error if
// the record to be updated doesn't exist
func UpdateGoodById(m *Good) (err error) {
	o := orm.NewOrm()
	v := Good{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGood deletes Good by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGood(id int) (err error) {
	o := orm.NewOrm()
	v := Good{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Good{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
