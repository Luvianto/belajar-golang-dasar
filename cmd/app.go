package cmd

import (
	commonutils "belajar-golang-dasar/common/utils"
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

	// userRepo := userRepo.NewUserRepository(db)
	// userService := userService.NewUserService(userRepo)

	memberReq := &memberEntity.MemberReqByID{ID: 1}

	// userReq := &userEntity.UserReqByUUID{UUID: "201dc479-f464-4a74-8856-eae76577fba3"}
	// userReq := &userEntity.UserReqByUUID{UUID: "c87fcea2-acae-4a06-9651-56ae45dc8af1"}

	member, err := memberService.GetMember(memberReq)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(member)

	uuid, success := commonutils.ParseUUID("201dc479-f464-4a74-8856-eae76577fba3")
	if !success {
		fmt.Print(success)
	}

	updateMemberStruct := &memberEntity.MemberUpdate{
		ID:                1,
		UserID:            uuid,
		Name:              "Orang kedua",
		Major:             "Bukan Major",
		ProfilePictureUrl: "Bukan Profile",
	}

	updateMemberReq, err := memberService.UpdateMember(updateMemberStruct)
	if err != nil {
		fmt.Print(err)
	}

	member, err = memberService.GetMember(memberReq)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(updateMemberReq)
	fmt.Println(member)
}
