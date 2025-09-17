package cmd

import (
	commonutils "belajar-golang-dasar/common/utils"
	"belajar-golang-dasar/database"
	memberEntity "belajar-golang-dasar/internal/module/member/entity"
	memberRepo "belajar-golang-dasar/internal/module/member/repository"
	memberService "belajar-golang-dasar/internal/module/member/service"
	userEntity "belajar-golang-dasar/internal/module/user/entity"

	"fmt"
)

func App() {
	database.InitializeDB()
	db := database.GetDBInstance()

	memberRepo := memberRepo.NewMemberRepository(db)
	memberService := memberService.NewMemberService(memberRepo)

	// userRepo := userRepo.NewUserRepository(db)
	// userService := userService.NewUserService(userRepo)

	// userReq := &userEntity.UserReqByUUID{UUID: "201dc479-f464-4a74-8856-eae76577fba3"}
	// userReq := &userEntity.UserReqByUUID{UUID: "c87fcea2-acae-4a06-9651-56ae45dc8af1"}

	uuid := commonutils.GenerateUUID()

	createMemberStruct := &memberEntity.MemberCreate{
		User: userEntity.UserCreate{
			IsAdmin:  false,
			Email:    "Email4@gmail.com",
			Password: "12345678",
			Phone:    "088888888888",
		},
		Name:              "Orang Pertama",
		Major:             "Major Pertama",
		ProfilePictureUrl: "Profile Pertama",
	}

	createMemberReq, err := memberService.CreateMember(createMemberStruct)
	if err != nil {
		fmt.Print(err)
	}

	memberReq := &memberEntity.MemberReqByID{ID: createMemberReq.ID}

	updateMemberStruct := &memberEntity.MemberUpdate{
		ID:                createMemberReq.ID,
		UserID:            uuid,
		Name:              "Orang kedua",
		Major:             "Major Kedua",
		ProfilePictureUrl: "Profile Kedua",
	}

	deleteMemberStruct := &memberEntity.MemberReqByID{
		ID: createMemberReq.ID,
	}

	updateMemberReq, err := memberService.UpdateMember(updateMemberStruct)
	if err != nil {
		fmt.Print(err)
	}

	deleteMemberReq, err := memberService.DeleteMember(deleteMemberStruct)
	if err != nil {
		fmt.Print(err)
	}

	member, err := memberService.GetMember(memberReq)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("\n=========Hasil=========")
	fmt.Println(createMemberReq)
	fmt.Println(updateMemberReq)
	fmt.Println(deleteMemberReq)
	fmt.Println(member)
}
