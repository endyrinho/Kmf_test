package model

import (
	"time"
)

type R_CURRENCY struct {
	ID   	int
	TITLE	string
	CODE    string
	VALUE   float32
	A_DATE  time.Time
}
type SearchClientAnswerModel struct { 
	XMLName xml.Name xml:"Envelope" json:"-" 
	Body    struct { 
	 XMLName                   xml.Name xml:"Body" json:"-" 
	 GetClientInfoByIdResponse struct { 
	  XMLName                 xml.Name xml:"GetClientInfoByIdResponse" json:"-" 
	  GetClientInfoByIdResult struct { 
	   XMLName     xml.Name xml:"GetClientInfoByIdResult" json:"-" 
	   IdClientCRM string   xml:"IdClientCRM" json:"id" 
	   Error       string   xml:"Error" json:"Error" 
	   ErrorMsg    string   xml:"ErrorMsg" json:"ErrorMsg" 
	   Doc         struct { 
		XMLName    xml.Name xml:"doc" json:"-" 
		NewDataSet struct { 
		 XMLName xml.Name xml:"NewDataSet" json:"-" 
		 Table   struct { 
		  XMLName                xml.Name           xml:"Table" json:"-" 
		  ID                     string             xml:"ID" json:"ID" 
		  IsBlockBP              int                xml:"ISBLOCKBP" json:"isBlockBP" 
		  BlockReason            *string             xml:"BLOCKREASON" json:"blockReason"

type money struct {
	item struct	{ 
	fullname string
	title string
	desc float32
	quant int
	index string
	change string
				}
}