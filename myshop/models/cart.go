package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Cart struct {
	Id int  `orm:"column(cartid);auto"`
	Num int32 `orm:"column(num)"`
	Gid *Good `orm:"column(gid);rel(fk)"`
	Uid *User `orm:"column(uid);rel(fk)"`
}

// UpdateGood updates Good by Id and returns error if
// the record to be updated doesn't exist
func UpdateCartById(m *Cart) (err error) {
	o := orm.NewOrm()
	v := Cart{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		m.Gid = v.Gid
		m.Uid = v.Uid
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func (t *Cart) TableName() string {
	return "cart"
}

func init() {
	orm.RegisterModel(new(Cart))
}

// AddGood insert a new Good into database and returns
// last inserted Id on success.
func AddCart(m *Cart) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func FindCardByUidAndGid(m *Cart) (v *Cart, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Cart)).Filter("Gid", m.Gid.Id).Filter("Uid", m.Uid.Id)
	v = &Cart{Id: m.Id}
	qs = qs.Limit(1, 0)
	if err = qs.One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetGoodById retrieves Good by Id. Returns error if
// Id doesn't exist
func GetCartById(id int) (v *Cart, err error) {
	o := orm.NewOrm()
	v = &Cart{Id: id}
	qs := o.QueryTable(new(Cart)).Filter("Id", id).RelatedSel()
	if err = qs.Limit(1, 0).One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetGoodById retrieves Good by Id. Returns error if
// Id doesn't exist
func GetCartByUid(id int) (l []Cart, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Cart)).Filter("Uid", id).RelatedSel()
	//var l []Cart
	if _,err := qs.All(&l); err == nil {
		return l, nil
	}
	return nil, err
}

// DeleteGood deletes Good by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCart(id int) (err error) {
	o := orm.NewOrm()
	v := Cart{Id: id}
	// ascertain id exists in the database
	fmt.Println(v)
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Cart{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
