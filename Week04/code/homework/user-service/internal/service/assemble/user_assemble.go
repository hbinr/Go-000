package assemble

import (
	pb "user.service/api/user/v1"
	"user.service/internal/biz/entity"
)

// Dto2Do dto -> do，在此处，相当于request转为了do
func Dto2Do(dto *pb.RegisterReq) (do *entity.User) {
	do = new(entity.User)
	do.Passport = dto.Passport
	do.Password = dto.Password
	do.Email = dto.Email
	do.Nickname = dto.Nickname
	return
}
