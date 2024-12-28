package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Password string
	Email    string
	Academy  string
	Major    string
	Building []Building `gorm:"foreignKey:UserID"`
}

type Building struct {
	gorm.Model
	Name   string
	Zone   string
	UserID uint
}
