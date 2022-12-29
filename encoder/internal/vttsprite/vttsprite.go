package vttsprite

import (
	"fmt"
	"image/jpeg"
	"math"
	"os"
	"path"
	"path/filepath"
	"time"
	"yet-another-media-server/encoder/internal/vttsprite/wrapper"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

const (
	ROWS  = 100
	COLS  = 4
	WIDTH = 300
)

func GenerateVttSprite(videoFilePath string) (string, error) {
	dirPath := path.Dir(videoFilePath)
	videoFilename := path.Base(videoFilePath)
	spriteFilename := videoFilename + ".sprite.jpg"
	vttFilename := videoFilename + ".vtt"
	videoReader := wrapper.VideoReader{
		FileName: videoFilePath,
	}
	err := videoReader.Open()
	if err != nil {
		fmt.Printf("Failed to open video. %s", err.Error())
		return "", err
	}
	defer videoReader.Release()

	targetHeight := int(math.Round(float64(WIDTH) / float64(videoReader.VideoInfo().Width) * float64(videoReader.VideoInfo().Height)))
	everyNSeconds := videoReader.VideoInfo().Duration / ROWS / COLS
	everyNFrames := float64(videoReader.VideoInfo().FrameCount) / ROWS / COLS

	spriteCtx := gg.NewContext(WIDTH*COLS, targetHeight*ROWS)
	vttContent := "WEBVTT\n\n"

	curTs := 0.0
	curFrameIdx := 0.0
	idx := 0
	execTime := time.Now().Unix()
	perf := wrapper.Perf{}
	perf.Start()
	for idx < ROWS*COLS {
		T1 := wrapper.PerfTimer()
		videoReader.SeekSeconds(curTs)
		T2 := wrapper.PerfTimer()
		img, err := videoReader.Read()
		T3 := wrapper.PerfTimer()
		if err != nil {
			fmt.Printf("Failed to extract frame. %s", err.Error())
			break
		}
		row := idx / COLS
		col := idx % COLS
		x := col * WIDTH
		y := row * targetHeight
		scaled := resize.Resize(WIDTH, uint(targetHeight), img, resize.Bilinear)
		T4 := wrapper.PerfTimer()
		spriteCtx.DrawImage(scaled, x, y)
		T5 := wrapper.PerfTimer()
		perf.Record(int64(curTs * 1000))
		perf.RecordTiming("seek", T2-T1)
		perf.RecordTiming("read", T3-T2)
		perf.RecordTiming("resize", T4-T3)
		perf.RecordTiming("draw", T5-T4)
		now := time.Now().Unix()
		if now-execTime >= 1 {
			fmt.Printf(
				"Timestamp: %.3fs Speed: %.2fX Perf(ms) Seek: %.3f Read: %.3f Resize: %.3f Draw: %.3f\n",
				curTs,
				perf.GetSpeed(),
				perf.AvgPeriodTiming("seek"),
				perf.AvgPeriodTiming("read"),
				perf.AvgPeriodTiming("resize"),
				perf.AvgPeriodTiming("draw"),
			)
			execTime = now
		}

		vttContent += fmt.Sprintf(
			"%02d:%02d:%02d.%03d --> %02d:%02d:%02d.%03d\n%s#xywh=%d,%d,%d,%d\n\n",
			int(curTs)/3600,
			int(curTs)/60%60,
			int(curTs)%60,
			int(curTs*1000)%1000,
			int(curTs+everyNSeconds)/3600,
			int(curTs+everyNSeconds)/60%60,
			int(curTs+everyNSeconds)%60,
			int((curTs+everyNSeconds)*1000)%1000,
			spriteFilename,
			WIDTH*(col),
			targetHeight*row,
			WIDTH,
			targetHeight,
		)

		curTs += everyNSeconds
		curFrameIdx += everyNFrames
		idx += 1
	}
	perf.Stop()
	fmt.Printf(
		"Finished %d frames at %.2fX speed. Perf(ms) Seek: %.3f Read: %.3f Resize: %.3f Draw: %.3f\n",
		int64(curFrameIdx),
		perf.GetSpeed(),
		perf.AvgTiming("seek"),
		perf.AvgTiming("read"),
		perf.AvgTiming("resize"),
		perf.AvgTiming("draw"),
	)

	f, err := os.Create(path.Join(dirPath, spriteFilename))
	if err != nil {
		fmt.Printf("Failed to create sprite file.")
		return "", err
	}
	defer f.Close()

	jpeg.Encode(f, spriteCtx.Image(), &jpeg.Options{Quality: 80})

	vttAbsPath, _ := filepath.Abs(path.Join(dirPath, vttFilename))
	vttFile, err := os.Create(vttAbsPath)
	if err != nil {
		fmt.Printf("Failed to create vtt file.")
		return "", err
	}
	defer vttFile.Close()

	vttFile.WriteString(vttContent)

	return vttAbsPath, nil
}
