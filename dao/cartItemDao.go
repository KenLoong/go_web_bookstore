package dao

import (
	"go_web_project/bookstore/model"
	"go_web_project/bookstore/utils"
)

//AddCartItem 向购物项表中插入购物项
func AddCartItem(cartItem *model.CartItem) error {
	//写sql
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values(?,?,?,?)"
	//执行sql
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		return err
	}
	return nil
}