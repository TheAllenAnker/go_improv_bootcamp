package go_bootcamp

import (
	"database/sql"
	"fmt"
)

func main() {

}

// testFirstWeekAssignment Error 章节作业
//  我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
// 答： 公司里的业务实现一般不会讲某个数据库找不到视为 error，因此不会 wrap 这个 error 并返回给上层，上游业务调用方视为记录不存在即可。
func testFirstWeekAssignment() {
	err := accessDatabase()
	if err != nil {
		// 因为调用的下层方法已经打印了日志，此处不再打印
		// log and handle
	}
}

func accessDatabase() error {
	// access database
	fmt.Println("assessing database...")
	err := sql.ErrNoRows // 数据库访问时返回了错误
	if err != nil {
		if err == sql.ErrNoRows {
			// do some logging
			fmt.Println("logging...")
			return nil
		}
		// do some logging
		fmt.Println("logging...")
		return fmt.Errorf("fail to access database, err=%w", err)
	}

	return nil
}
