package main

var Users = []*User{}

type User struct {
    Id    string
    Username  string
    Follower int32
}

func GetUsers() []*User {
    return Users
}

func SelectUser(id string) *User {
    for _, each := range Users {
        if each.Id == id {
            return each
        }
    }

    return nil
}

func init() {
    Users = append(Users, &User{Id: "s001", Username: "via", Follower: 2})
    Users = append(Users, &User{Id: "s002", Username: "dwi", Follower: 2})
    Users = append(Users, &User{Id: "s003", Username: "miftakhur", Follower: 3})
}