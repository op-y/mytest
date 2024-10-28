package dep

type UserAttendanceInfo struct {
	UserId       int    `json:"user_id"`
	UserName     string `json:"user_name"`
	RealWorkTime string `json:"real_work_time"`
	VacationTime string `json:"vacation_time"`
}

// 这个接口不一定要包含真实Service所有方法
type DepService interface {
	GetUserAttendanceInfo(userId int) (*UserAttendanceInfo, error)
	ProcessOrder(orderId int) error
}
