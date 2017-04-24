package UserConfig

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"time"
	"encoding/json"
	"strconv"
)

const url string = "http://fg-69c8cbcd.herokuapp.com"
const querySchema = "/user/"

type UserDetail struct {
	Id      int `json:"id"`
	Name    string `json:"name"`
	Friends []int `json:"friends"`
}

type User struct {
	allUsers map[string]UserDetail
}

func (user *User) query(id string) UserDetail {
	finalUrl := fmt.Sprintf("%s%s%s", url, querySchema, id)
	result, _, err := gorequest.New().Timeout(time.Second * 3).Get(finalUrl).End()
	if err != nil {
		fmt.Printf("Error in querying data from %s : %s\n", finalUrl, err)
		return UserDetail{}
	}
	response := make([]byte, result.ContentLength)
	result.Body.Read(response)
	if len(response) <= 0 {
		return UserDetail{}
	}
	userDetail := UserDetail{}
	err1 := json.Unmarshal(response, &userDetail)
	if err1 != nil {
		fmt.Printf("Error in parsing data from %s : %s\n", finalUrl, err1)
		return UserDetail{}
	}
	//fmt.Println(string(response))
	//fmt.Println(userDetail)
	return userDetail
}

func (user *User) PopulateData(id string) {
	if user.allUsers == nil {
		user.allUsers = make(map[string]UserDetail)
	}
	_, found := user.allUsers[id]
	if found {
		return
	}

	userDetail := user.query(id)
	if userDetail.Name == "" {
		return
	}

	user.allUsers[strconv.Itoa(userDetail.Id)] = userDetail
	for _, item := range userDetail.Friends {
		newId := strconv.Itoa(item)
		_, ok := user.allUsers[newId]
		if ok {
			continue
		} else {
			user.PopulateData(newId)
		}
	}
}

func (user *User) PrintDetail() {
	for _, item := range user.allUsers {
		fmt.Printf("Id: %d\nName: %s\n", item.Id, item.Name)
		if len(item.Friends) <= 0 {
			fmt.Print("\n")
			continue
		}
		fmt.Println("Friends:")
		for _, friend := range item.Friends {
			fmt.Println("   +", friend)
		}
		fmt.Print("\n")
	}
}
