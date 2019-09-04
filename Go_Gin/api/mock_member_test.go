package api_test

import (
	"github.com/sheyguys/gogin/model"

	"github.com/globalsign/mgo/bson"
)

type mockMemberRepository struct{}

const (
	listMember = `{"member":[{"member_id":"5beaf7bd62e63844ce22cc57","member_name_eng":"Kasinan Rordthab","member_name_th":"กษิติ์นันท์ รอดทัพ","member_idcard":1478523695461,"member_phone":"0925797702",
	"member_address":"Suranaree University","member_email":"b5920914@gmail.com","member_facebook":"Ping Kasinan"},{"member_id":"5beaf7bd62e63844ce22cc58","member_name_eng":"Chatcha Kummoon",
	"member_name_th":"ฌัชชา คำมูล","member_idcard":1234567890123,"member_phone":"0923456789","member_address":"Suranaree University","member_email":"b5922491@gmail.com","member_facebook":"Gate Up"}]}`
)

func (memberRepository mockMemberRepository) GetAllMember() ([]model.Member, error) {
	return []model.Member{
		{
			MemberID:       bson.ObjectIdHex("5beaf7bd62e63844ce22cc57"),
			MemberNameeng:  "Kasinan Rordthab",
			MemberNameth:   "กษิติ์นันท์ รอดทัพ",
			MemberIdcard:   1478523695461,
			MemberPhone:    "0925797702",
			MemberAddress:  "Suranaree University",
			MemberEmail:    "b5920914@gmail.com",
			MemberFacebook: "Ping Kasinan",
		},
		{
			MemberID:       bson.ObjectIdHex("5beaf7bd62e63844ce22cc58"),
			MemberNameeng:  "Chatcha Kummoon",
			MemberNameth:   "ฌัชชา คำมูล",
			MemberIdcard:   1234567890123,
			MemberPhone:    "0923456789",
			MemberAddress:  "Suranaree University",
			MemberEmail:    "b5922491@gmail.com",
			MemberFacebook: "Gate Up",
		},
	}, nil
}

func (memberRepository mockMemberRepository) AddMember(member model.Member) error {
	return nil
}

func (memberRepository mockMemberRepository) EditMemberName(memberID string, member model.Member) error {
	return nil
}

func (memberRepository mockMemberRepository) DeleteMemberByID(memberID string) error {
	return nil
}
