package dep

import (
	"errors"
	"fmt"
)

type DepServiceImpl struct{}

func NewDepServiceImpl() *DepServiceImpl {
	return &DepServiceImpl{}
}

func (svc *DepServiceImpl) Init() {}

func (svc *DepServiceImpl) GetUserAttendanceInfo(userId int) (*UserAttendanceInfo, error) {
	fmt.Println("不用想了, 依赖的接口不会返回数据!")
	return nil, nil
}

func (svc *DepServiceImpl) ProcessOrder(orderId int) error {
	fmt.Println("不用想了, 三方系统无法正常处理测试数据!")
	return errors.New("参数异常")
}

func (svc *DepServiceImpl) Do(userId, orderId int) error {
	fmt.Println("一些前置工作")
	info, err := svc.GetUserAttendanceInfo(userId)
	if err != nil {
		fmt.Println("获取用户出勤信息异常")
		return err
	}
	if info.VacationTime != "0" && info.RealWorkTime < "8" {
		if err := svc.ProcessOrder(orderId); err != nil {
			fmt.Println("处理订单异常")
			return err
		}
	}
	fmt.Println("一些收尾工作")
	return nil
}
