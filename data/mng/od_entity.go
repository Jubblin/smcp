package mng

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"smcp/data"
	"smcp/utils"
)

type DetectedObject struct {
	PredScore   float32 `json:"pred_score"  bson:"pred_score"`
	PredClsIdx  int     `json:"pred_cls_idx" bson:"pred_cls_idx"`
	PredClsName string  `json:"pred_cls_name" bson:"pred_cls_name"` //Index
}

type OdEntity struct {
	Id             primitive.ObjectID `json:"_id" bson:"_id"`
	GroupId        string             `json:"group_id" bson:"group_id"`   //Index
	SourceId       string             `json:"source_id" bson:"source_id"` //Index
	CreatedAt      string             `json:"created_at" bson:"created_at"`
	DetectedObject *DetectedObject    `json:"detected_object" bson:"detected_object"`
	ImageFileName  string             `json:"image_file_name" bson:"image_file_name"`

	VideoFile *VideoFile `json:"video_file" bson:"video_file"`

	AiClip *data.AiClip `json:"ai_clip" bson:"ai_clip"`

	CreatedDate primitive.DateTime `json:"created_date" bson:"created_date"`
}

func (o *OdEntity) SetupDates(createdAt string) {
	o.CreatedAt = createdAt
	t := utils.StringToTime(createdAt)
	o.CreatedDate = data.TimeToMongoDateTime(t)
}