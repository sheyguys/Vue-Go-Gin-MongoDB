package repository_test

import (
	"log"
	"testing"

	"github.com/globalsign/mgo/bson"
	"github.com/sheyguys/gogin/model"
	"github.com/sheyguys/gogin/repository"

	"github.com/globalsign/mgo"
	"github.com/stretchr/testify/assert"
)

const mogoDBEnPint = "mongodb://localhost:27017"

func connectionDB() *mgo.Session {
	connectionDB, err := mgo.Dial(mogoDBEnPint)
	if err != nil {
		log.Panic("Con not connect to database", err.Error())
	}
	return connectionDB
}

func Test_AddMember_Shold_Be_Member(t *testing.T) {
	connectionDB := connectionDB()
	defer connectionDB.Close()
	memberRepository := repository.MemberRepositoryMongo{
		ConnectionDB: connectionDB,
	}
	member := model.Member{
		MemberID:       bson.ObjectIdHex("5beaf7bd62e63844ce22cc57"),
		MemberNameeng:  "Kasinan Rordthab",
		MemberNameth:   "กษิติ์นันท์ รอดทัพ",
		MemberIdcard:   1478523695461,
		MemberPhone:    "0925797702",
		MemberAddress:  "Suranaree University",
		MemberEmail:    "b5920914@gmail.com",
		MemberFacebook: "Ping Kasinan",
	}

	actual := memberRepository.AddMember(member)

	if actual != nil {
		t.Error("Con not Add Member", actual.Error())
	}
}

func Test_GetAllMember_Should_Be_Array_Member(t *testing.T) {
	connectionDB := connectionDB()
	defer connectionDB.Close()
	expected := []model.Member{
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
	}
	memberRepository := repository.MemberRepositoryMongo{
		ConnectionDB: connectionDB,
	}

	actual, _ := memberRepository.GetAllMember()

	assert.Equal(t, expected, actual)
}

func Test_EditMemberName_Input_Member_Name_Kasinan_Should_Be_Edited(t *testing.T) {
	connectionDB := connectionDB()
	defer connectionDB.Close()
	memberID := "5beaf7bd62e63844ce22cc57"
	member := model.Member{
		MemberNameth: "กษิติ์นันท์ รอดทัพ",
	}
	memberRepository := repository.MemberRepositoryMongo{
		ConnectionDB: connectionDB,
	}

	actual := memberRepository.EditMemberName(memberID, member)

	if actual != nil {
		t.Error("Con not Edit member", actual.Error())
	}
}

func Test_DeleteMemberByID_Input_Id_5befe40d9c71fe169a4341df_Should_Be_Deleted(t *testing.T) {
	connectionDB := connectionDB()
	defer connectionDB.Close()
	memberID := "5beaf7bd62e63844ce22cc57"
	memberRepository := repository.MemberRepositoryMongo{
		ConnectionDB: connectionDB,
	}

	actua := memberRepository.DeleteMemberByID(memberID)

	if actua != nil {
		t.Error("Con not delete member", actua.Error())
	}
}
