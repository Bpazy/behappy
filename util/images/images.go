package images

import (
	"embed"
	"fmt"
	"github.com/Bpazy/berrors"
	"github.com/flopp/go-findfont"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

//go:embed assets/*
var f embed.FS

func HonorTemplate(name string, year int, week int, month int, times int64) (string, error) {
	templateFile, err := f.Open("assets/honor.png")
	if err != nil {
		panic(err)
	}
	defer templateFile.Close()

	templateFileImage, err := png.Decode(templateFile)
	if err != nil {
		panic(err)
	}

	newTemplateImage := image.NewRGBA(templateFileImage.Bounds())
	draw.Draw(newTemplateImage, templateFileImage.Bounds(), templateFileImage, templateFileImage.Bounds().Min, draw.Over)

	content := freetype.NewContext()
	content.SetClip(newTemplateImage.Bounds())
	content.SetDst(newTemplateImage)
	content.SetSrc(image.Black)
	content.SetDPI(72)

	content.SetFontSize(40) // 设置字体大小
	fontKai := berrors.Unwrap(loadFont())
	content.SetFont(fontKai)

	// TODO 考虑通过返回值自适应换行
	berrors.Unwrap(content.DrawString(name+"同志:", freetype.Pt(80, 320)))
	berrors.Unwrap(content.DrawString(fmt.Sprintf("您在%d年度第%d周中表现突出，参与比赛%d局", year, week, times), freetype.Pt(120, 390)))
	berrors.Unwrap(content.DrawString("被评为", freetype.Pt(80, 460)))

	berrors.Unwrap(content.DrawString("特发此证，以资鼓励。", freetype.Pt(440, 460)))
	content.SetFontSize(48)
	content.SetSrc(image.NewUniform(color.RGBA{R: 237, G: 39, B: 90, A: 255}))
	berrors.Unwrap(content.DrawString("最佳劳模，", freetype.Pt(210, 460)))

	content.SetFont(fontKai) // 设置字体样式
	content.SetFontSize(32)
	content.SetSrc(image.Black)
	berrors.Unwrap(content.DrawString("黑哥TV刀塔分部", freetype.Pt(650, 550)))
	berrors.Unwrap(content.DrawString(fmt.Sprintf("%d年%d月", year, month), freetype.Pt(730, 590)))
	return saveFile(newTemplateImage)
}

func saveFile(pic *image.RGBA) (string, error) {
	dstFile, err := os.CreateTemp("", "honor.*.png")
	if err != nil {
		return "", err
	}
	defer dstFile.Close()
	err = png.Encode(dstFile, pic)
	if err != nil {
		return "", err
	}
	return dstFile.Name(), nil
}

func loadFont() (font *truetype.Font, err error) {
	fontPath, err := findfont.Find("simkai.ttf")
	if err != nil {
		err = fmt.Errorf("寻找字体文件出错: %s", err.Error())
		return
	}
	fontBytes, err := os.ReadFile(fontPath)
	if err != nil {
		err = fmt.Errorf("加载字体文件出错: %s", err.Error())
		return
	}
	font, err = freetype.ParseFont(fontBytes)
	if err != nil {
		err = fmt.Errorf("解析字体文件出错: %s", err.Error())
		return
	}
	return
}
