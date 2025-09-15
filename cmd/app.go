package cmd

import (
	"belajar-golang-dasar/database"
	memberEntity "belajar-golang-dasar/internal/module/member/entity"
	memberRepo "belajar-golang-dasar/internal/module/member/repository"
	memberService "belajar-golang-dasar/internal/module/member/service"
	"fmt"
)

func App() {
	database.InitializeDB()
	db := database.GetDBInstance()

	memberRepo := memberRepo.NewMemberRepository(db)
	memberService := memberService.NewMemberService(memberRepo)

	req := &memberEntity.MemberReqByID{ID: 1}

	member, err := memberService.GetMember(req)

	fmt.Println(member)
	fmt.Println(err)
}
