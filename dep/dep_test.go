package dep

import (
	"testing"

	"go.uber.org/mock/gomock"
)

func TestServiceDo(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := NewMockDepService(ctrl)

	m.EXPECT().GetUserAttendanceInfo(gomock.Any()).
		DoAndReturn(func(userId int) (*UserAttendanceInfo, error) {
			if userId == 101 {
				return &UserAttendanceInfo{
					UserId:       101,
					UserName:     "张三",
					RealWorkTime: "6",
					VacationTime: "2",
				}, nil
			} else {
				return &UserAttendanceInfo{
					UserId:       999,
					UserName:     "其它人",
					RealWorkTime: "8",
					VacationTime: "0",
				}, nil
			}
		}).
		AnyTimes()

	m.EXPECT().ProcessOrder(gomock.Any()).Return(nil).AnyTimes()

	svc := NewDepServiceImpl()
	svc.Init()

	userId, orderId := 101, 101
	t.Log("一些前置工作")
	info, err := m.GetUserAttendanceInfo(userId)
	if err != nil {
		t.Fatal("访问依赖接口异常")
	}
	if info.RealWorkTime == info.VacationTime {
		t.Fatal("张三的测试用例中 RealWorkTime 与 VacationTime 不应该相等!")
	}
	if err := m.ProcessOrder(orderId); err != nil {
		t.Fatal("访问三方依赖接口异常")
	}
	t.Log("一些收尾工作")
	t.Log("完成业务流程测试")
}
