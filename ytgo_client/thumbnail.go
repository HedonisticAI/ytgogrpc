package main


type Thumbnail struct {
	VideoID           string
	link           string
	fileName       []string
	thumbnailsDir  string
	thumbnailsName string
}


func NewThumbnail(string link ) *Thumbnail {
	return &Thumbnail{
		VideoID:        "",
		link:           link,
		fileName:       []string{},
		thumbnailsDir:  "./pic",
		thumbnailsName: "",
	}
}

func (t *Thumbnail) findVideoID(urlVideo string) error {

	equalIndex := strings.Index(urlVideo, "=")
	ampIndex := strings.Index(urlVideo, "&")
	slash := strings.LastIndex(urlVideo, "/")
	questionIndex := strings.Index(urlVideo, "?")
	var id string

	if equalIndex != -1 {

		if ampIndex != -1 {
			id = urlVideo[equalIndex+1 : ampIndex]
		} else if questionIndex != -1 && strings.Contains(urlVideo, "?t=") {
			id = urlVideo[slash+1 : questionIndex]
		} else {
			id = urlVideo[equalIndex+1:]
		}

	} else {
		id = urlVideo[slash+1:]
	}

	t.VideoID = id

	if strings.ContainsAny(id, "?&/<%=") {
		return errors.New("invalid characters in video id")
	}
	if len(id) < 10 {
		return errors.New("the video id must be at least 10 characters long")
	}
	return nil
}
