package main

import (
	"fmt"
	"os"
	"time"
)

var currSegments = []string{"segment0.ts", "segment0.ts", "segment1.ts"}

func createM3u8(segmentNumber int, sequenceNumer int) {

	currSegments = currSegments[1:]
	newSegment := fmt.Sprintf("segment%d.ts", segmentNumber)
	currSegments = append(currSegments, newSegment)
	fileText := ""
	for _, segment := range currSegments {
		duration := 10.0
		if segment == "segment63.ts" {
			duration = 4
		}
		fileText += fmt.Sprintf("#EXTINF:%.6f,\n %s\n", duration, segment)
	}
	fileText = fmt.Sprintf("#EXTM3U\n#EXT-X-VERSION:3\n#EXT-X-TARGETDURATION:10\n#EXT-X-MEDIA-SEQUENCE:%v\n", sequenceNumer) + fileText
	err := os.WriteFile("./src/zapp.m3u8", []byte(fileText), 0666)
	if err != nil {
		fmt.Println(err)
	}

}

func hlsRep() {

	segmentNum := 2
	sequenceNumer := 0
	createM3u8(segmentNum, sequenceNumer)
	segmentNum++
	sequenceNumer++
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				createM3u8(segmentNum, sequenceNumer)
				segmentNum++
				sequenceNumer++
				if segmentNum > 63 {
					segmentNum = 0
				}
			}
		}
	}()
	return
}
