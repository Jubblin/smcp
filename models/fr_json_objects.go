package models

type FaceRecognitionJsonBaseObject struct {
	Id               string          `json:"id"`
	SourceId         string          `json:"source_id"`
	CreatedAt        string          `json:"created_at"`
	DetectedFaces    []*DetectedFace `json:"detected_faces"`
	VideoClipEnabled bool            `json:"video_clip_enabled"`
	ImageFileName    string          `json:"image_file_name"`
	DataFileName     string          `json:"data_file_name"`
}

type FaceRecognitionJsonObject struct {
	FaceRecognition *FaceRecognitionJsonBaseObject `json:"object_detection"`
	Video           *VideoClipJsonObject           `json:"video"`
}
