package interfaces

import (
	memberEntity "belajar-golang-dasar/internal/module/member/entity"
	userEntity "belajar-golang-dasar/internal/module/user/entity"
)

type MemberRepository interface {
	GetMember(id int) (*memberEntity.Member, bool, error)
	CreateMember(user userEntity.User, member memberEntity.Member) (*memberEntity.Member, bool, error)
}

type MemberService interface {
	GetMember(req *memberEntity.MemberReqByID) (*memberEntity.MemberGet, error)
	CreateMember(req *memberEntity.MemberCreate) (*memberEntity.MemberGet, error)
}
