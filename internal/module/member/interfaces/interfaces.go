package interfaces

import "belajar-golang-dasar/internal/module/member/entity"

type MemberRepository interface {
	GetMember(id int) (*entity.Member, bool, error)
}

type MemberService interface {
	GetMember(req *entity.MemberReqByID) (*entity.MemberGet, error)
}
