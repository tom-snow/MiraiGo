// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/profilecard/accountsearch.proto

package profilecard

import (
	proto "github.com/RomiChan/protobuf/proto"
)

type Location struct {
	Latitude  proto.Option[float64] `protobuf:"fixed64,1,opt"`
	Longitude proto.Option[float64] `protobuf:"fixed64,2,opt"`
	_         [0]func()
}

type ResultItem struct {
	FeedId      []byte               `protobuf:"bytes,1,opt"`
	Name        []byte               `protobuf:"bytes,2,opt"`
	PicUrl      []byte               `protobuf:"bytes,3,opt"`
	JmpUrl      []byte               `protobuf:"bytes,4,opt"`
	FeedType    []byte               `protobuf:"bytes,5,opt"`
	Summary     []byte               `protobuf:"bytes,6,opt"`
	HasVideo    []byte               `protobuf:"bytes,7,opt"`
	PhtotUpdate []byte               `protobuf:"bytes,8,opt"`
	Uin         proto.Option[uint64] `protobuf:"varint,9,opt"`
	ResultId    []byte               `protobuf:"bytes,10,opt"`
	Ftime       proto.Option[uint32] `protobuf:"varint,11,opt"`
	NickName    []byte               `protobuf:"bytes,12,opt"`
	PicUrlList  [][]byte             `protobuf:"bytes,13,rep"`
	TotalPicNum proto.Option[uint32] `protobuf:"varint,14,opt"`
}

type Hotwordrecord struct {
	Hotword            proto.Option[string] `protobuf:"bytes,1,opt"`
	HotwordType        proto.Option[uint32] `protobuf:"varint,2,opt"`
	HotwordCoverUrl    proto.Option[string] `protobuf:"bytes,3,opt"`
	HotwordTitle       proto.Option[string] `protobuf:"bytes,4,opt"`
	HotwordDescription proto.Option[string] `protobuf:"bytes,5,opt"`
	_                  [0]func()
}

type AccountSearchRecord struct {
	Uin               proto.Option[uint64] `protobuf:"varint,1,opt"`
	Code              proto.Option[uint64] `protobuf:"varint,2,opt"`
	Source            proto.Option[uint32] `protobuf:"varint,3,opt"`
	Name              proto.Option[string] `protobuf:"bytes,4,opt"`
	Sex               proto.Option[uint32] `protobuf:"varint,5,opt"`
	Age               proto.Option[uint32] `protobuf:"varint,6,opt"`
	Accout            proto.Option[string] `protobuf:"bytes,7,opt"`
	Brief             proto.Option[string] `protobuf:"bytes,8,opt"`
	Number            proto.Option[uint32] `protobuf:"varint,9,opt"`
	Flag              proto.Option[uint64] `protobuf:"varint,10,opt"`
	Relation          proto.Option[uint64] `protobuf:"varint,11,opt"`
	Mobile            proto.Option[string] `protobuf:"bytes,12,opt"`
	Sign              []byte               `protobuf:"bytes,13,opt"`
	Country           proto.Option[uint32] `protobuf:"varint,14,opt"`
	Province          proto.Option[uint32] `protobuf:"varint,15,opt"`
	City              proto.Option[uint32] `protobuf:"varint,16,opt"`
	ClassIndex        proto.Option[uint32] `protobuf:"varint,17,opt"`
	ClassName         proto.Option[string] `protobuf:"bytes,18,opt"`
	CountryName       proto.Option[string] `protobuf:"bytes,19,opt"`
	ProvinceName      proto.Option[string] `protobuf:"bytes,20,opt"`
	CityName          proto.Option[string] `protobuf:"bytes,21,opt"`
	AccountFlag       proto.Option[uint32] `protobuf:"varint,22,opt"`
	TitleImage        proto.Option[string] `protobuf:"bytes,23,opt"`
	ArticleShortUrl   proto.Option[string] `protobuf:"bytes,24,opt"`
	ArticleCreateTime proto.Option[string] `protobuf:"bytes,25,opt"`
	ArticleAuthor     proto.Option[string] `protobuf:"bytes,26,opt"`
	AccountId         proto.Option[uint64] `protobuf:"varint,27,opt"`
	// repeated Label groupLabels = 30;
	VideoAccount  proto.Option[uint32] `protobuf:"varint,31,opt"`
	VideoArticle  proto.Option[uint32] `protobuf:"varint,32,opt"`
	UinPrivilege  proto.Option[int32]  `protobuf:"varint,33,opt"`
	JoinGroupAuth []byte               `protobuf:"bytes,34,opt"`
	Token         []byte               `protobuf:"bytes,500,opt"`
	Richflag1_59  proto.Option[uint32] `protobuf:"varint,40603,opt"`
	Richflag4_409 proto.Option[uint32] `protobuf:"varint,42409,opt"`
}

type AccountSearch struct {
	Start         proto.Option[int32]    `protobuf:"varint,1,opt"`
	Count         proto.Option[uint32]   `protobuf:"varint,2,opt"`
	End           proto.Option[uint32]   `protobuf:"varint,3,opt"`
	Keyword       proto.Option[string]   `protobuf:"bytes,4,opt"`
	List          []*AccountSearchRecord `protobuf:"bytes,5,rep"`
	Highlight     []string               `protobuf:"bytes,6,rep"`
	UserLocation  *Location              `protobuf:"bytes,10,opt"`
	LocationGroup proto.Option[bool]     `protobuf:"varint,11,opt"`
	Filtertype    proto.Option[int32]    `protobuf:"varint,12,opt"`
	// repeated C33304record recommendList = 13;
	HotwordRecord  *Hotwordrecord       `protobuf:"bytes,14,opt"`
	ArticleMoreUrl proto.Option[string] `protobuf:"bytes,15,opt"`
	ResultItems    []*ResultItem        `protobuf:"bytes,16,rep"`
	KeywordSuicide proto.Option[bool]   `protobuf:"varint,17,opt"`
	ExactSearch    proto.Option[bool]   `protobuf:"varint,18,opt"`
}
