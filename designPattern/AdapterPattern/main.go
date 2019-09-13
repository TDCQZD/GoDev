package main

import "fmt"

type MediaPlayer interface {
	play(audioType, fileName string)
}
type AdvancedMediaPlayer interface {
	playVlc(fileName string)
	playMp4(fileName string)
}
type VlcPlayer struct{}

func (pv *VlcPlayer) playVlc(fileName string) {
	fmt.Println("Playing vlc file. Name: " + fileName)
}
func (pm *VlcPlayer) playMp4(fileName string) {

}

type Mp4Player struct{}

func (pv *Mp4Player) playVlc(fileName string) {

}
func (pm *Mp4Player) playMp4(fileName string) {
	fmt.Println("Playing Mp4 file. Name: " + fileName)
}

type MediaAdapter struct {
	Mp4Player
	VlcPlayer
}

func (ap *MediaAdapter) play(audioType, fileName string) {
	if audioType == "vlc" {
		ap.VlcPlayer.playVlc(fileName)
	} else if audioType == "mp4" {
		ap.Mp4Player.playMp4(fileName)
	}
}

type AudioPlayer struct {
	MediaAdapter
}

func (ap *AudioPlayer) play(audioType, fileName string) {
	if audioType == "mp3" {
		fmt.Println("Playing mp3 file. Name: " + fileName)
	} else if audioType == "mp4" || audioType == "vlc" {
		ap.MediaAdapter.play(audioType, fileName)
	} else {
		fmt.Println("Invalid media. " + audioType + " format not supported")
	}
}
func main() {
	ap := &AudioPlayer{}
	ap.play("mp3", "beyond the horizon.mp3")
	ap.play("mp4", "alone.mp4")
	ap.play("vlc", "far far away.vlc")
	ap.play("avi", "mind me.avi")
}
