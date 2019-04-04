package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type OrderItem struct {
	Id  int      `orm:"column(otid);auto"`
	Num int      `orm:"column(num)"`
	Oid *Order   `orm:"column(oid);rel(fk)"`
	Gid *Good    `orm:"column(gid);rel(fk)"`
	Cid *Comment `orm:"column(cid);rel(fk);null"`
}

func (t *OrderItem) TableName() string {
	return "order_item"
}

func init() {
	orm.RegisterModel(new(OrderItem))
}

// AddOrderItem insert a new OrderItem into database and returns
// last inserted Id on success.
func AddOrderItem(m *OrderItem) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOrderItemById retrieves OrderItem by Id. Returns error if
// Id doesn't exist
func GetOrderItemById(id int) (v *OrderItem, err error) {
	o := orm.NewOrm()
	v = &OrderItem{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetOrderItemByOId(oid int)(ml []OrderItem, err error)  {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OrderItem)).Filter("Oid",oid).RelatedSel("Gid","Cid","Gid__Sid")
	_, err = qs.All(&ml)
	if err == nil{
		return ml, nil
	}
	return nil,err
}

// GetAllOrderItem retrieves all OrderItem matches certain condition. Returns empty list if
// no records exist
func GetAllOrderItem(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OrderItem))
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
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
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
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []OrderItem
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
		return ml, nil
	}
	return nil, err
}

// UpdateOrderItem updates OrderItem by Id and returns error if
// the record to be updated doesn't exist
func UpdateOrderItemById(m *OrderItem) (err error) {
	o := orm.NewOrm()
	v := OrderItem{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOrderItem deletes OrderItem by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOrderItem(id int) (err error) {
	o := orm.NewOrm()
	v := OrderItem{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OrderItem{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
