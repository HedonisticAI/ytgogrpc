package main 

import(
	"log"
	"strings"
)


func ParseID ( urlVideo string) string{

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
	if strings.ContainsAny(id, "?&/<%=") {
		log.Printf("invalid characters in video id")
		return ""
	}
	if len(id) < 10 {
		log.Printf("the video id must be at least 10 characters long")
		return ""
	}
	return id
}