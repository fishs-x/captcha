package models


type ImageStyle struct {
	Width  int    `form:"width" json:"width"`
	Prefix string `form:"prefix" json:"prefix"`
	Height int    `form:"height" json:"height"`
	Length int    `form:"length" json:"length"`
}

type VerifyCode struct {
	Id   string `form:"id" json:"id" binding:"required"`
	Code string `json:"code" form:"code" binding:"required"`
}

type ImgCode struct {
	Val interface{}
	ValType string
}