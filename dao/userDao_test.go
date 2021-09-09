package dao

import (
	"fmt"
	"testing"
)

func TestUserDao(t *testing.T) {
	fmt.Println("测试userDao的方法")
	//t.Run("保存用户",testSaveUser)
	t.Run("测试登录",testCheckUsernameAndPassword)
	t.Run("测试用户存在",testCheckUsername)

}

func testCheckUsernameAndPassword(t *testing.T) {
	user, _ := CheckUsernameAndPassword("admin", "123456")
	fmt.Println("获取用户信息是；",user)
}

func testCheckUsername(t *testing.T) {
	user, _ := CheckUsername("admin")
	fmt.Println("检查用户存在是：",user)
}

func testSaveUser(t *testing.T) {
	SaveUser("admin", "123456", "123456@123.com")
}
