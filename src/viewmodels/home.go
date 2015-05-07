package viewmodels

import (

)

type Home struct {
    Title string
    Active string
}

func GetHome() Home {
    result := Home{
        Title: "Stories",
        Active: "home",
    }
    return result
}