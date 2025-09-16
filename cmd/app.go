package cmd

import (
	"belajar-golang-dasar/database"
	memberEntity "belajar-golang-dasar/internal/module/member/entity"
	memberRepo "belajar-golang-dasar/internal/module/member/repository"
	memberService "belajar-golang-dasar/internal/module/member/service"

	userEntity "belajar-golang-dasar/internal/module/user/entity"
	userRepo "belajar-golang-dasar/internal/module/user/repository"
	userService "belajar-golang-dasar/internal/module/user/service"

	"fmt"
)

func App() {
	database.InitializeDB()
	db := database.GetDBInstance()

	memberRepo := memberRepo.NewMemberRepository(db)
	memberService := memberService.NewMemberService(memberRepo)

	userRepo := userRepo.NewUserRepository(db)
	userService := userService.NewUserService(userRepo)

	memberReq := &memberEntity.MemberReqByID{ID: 1}

	// userReq := &userEntity.UserReqByUUID{UUID: "201dc479-f464-4a74-8856-eae76577fba3"}
	userReq := &userEntity.UserReqByUUID{UUID: "c87fcea2-acae-4a06-9651-56ae45dc8af1"}

	member, err := memberService.GetMember(memberReq)
	if err != nil {
		fmt.Print(err)
	}
	user, err := userService.GetUser(userReq)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(member)
	fmt.Println(user)
}
