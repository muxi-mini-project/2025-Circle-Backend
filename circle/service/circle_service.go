package service
import (
    "circle/models"
	"circle/request"
	"circle/dao"
)
type CircleServices struct {
	ud *dao.CircleDao
}
func NewCircleServices(ud *dao.CircleDao) *CircleServices {
	return &CircleServices{
		ud: ud,
	}
}
func (us *CircleServices) CreateCircle(name string,get request.CreateCircle) string {
    userid,_:=us.ud.GetIdByUser(name)
    circle:=models.Circle{
        Name:get.Name,
        Discription:get.Discription,
        Imageurl:get.Imageurl,
        Userid:userid,
        Status:"pending",
    }
    _=us.ud.CreateCircle(&circle)
    return "等待审核"
}
func (us *CircleServices) PendingCircle(name string) (models.Circle,bool) {
    if name!="root" {
        return models.Circle{},false
    }
    circle,_:=us.ud.SelectPendingCircle()
    return circle,true
}
func (us *CircleServices) ApproveCircle(name string,get request.ApproveCircle) string{
    if name!="root" {
        return "权限不足"
    }
    circle,_:=us.ud.GetCircleByID(get.Circleid)
    if get.Decide=="false" {
        _=us.ud.DeleteCircleByID(get.Circleid)
    }else {
        circle.Status="approved"
        _=us.ud.UpdateCircle(&circle,get.Circleid)
    }
    return "审核结束"
}
func (us *CircleServices) GetCircle(get request.Circleid) models.Circle{
	circle,_:=us.ud.GetCircleByID(get.Circleid)
    return circle
}
func (us *CircleServices) SelectCircle() []models.Circle{
	circle,_:=us.ud.SelectCircle()
    return circle
}
func (us *CircleServices) FollowCircle(name string,get request.Circleid) string{
    id,_:=us.ud.GetIdByUser(name)
    _=us.ud.FollowCircle(get.Circleid,id)
    return "关注成功"
}