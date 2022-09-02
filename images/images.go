package images

import (
	"fmt"
	"github.com/Bpazy/behappy/berrors"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func Test(name string) {
	templateFile, err := os.Open("assets/1.png")
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
	content.SetSrc(image.Black) // 设置字体颜色
	content.SetDPI(72)          // 设置字体分辨率

	content.SetFontSize(40) // 设置字体大小
	fontKai := berrors.Unwrap(loadFont("C:\\Windows\\Fonts\\simkai.ttf"))
	content.SetFont(fontKai) // 设置字体样式，就是我们上面加载的字体

	// TODO 考虑通过返回值自适应换行
	drawString, err := content.DrawString(name+"同志:", freetype.Pt(80, 320))
	berrors.Unwrap(drawString, err)
	berrors.Unwrap(content.DrawString("您在2022年度第30周中表现突出，参与比赛23局", freetype.Pt(120, 390)))
	berrors.Unwrap(content.DrawString("被评为", freetype.Pt(80, 460)))

	berrors.Unwrap(content.DrawString("特发此证，以资鼓励。", freetype.Pt(440, 460)))
	// 设置字体大小
	content.SetFontSize(48)
	// 设置字体颜色
	content.SetSrc(image.NewUniform(color.RGBA{R: 237, G: 39, B: 90, A: 255}))
	berrors.Unwrap(content.DrawString("最佳劳模，", freetype.Pt(210, 460)))

	content.SetFont(fontKai) // 设置字体样式
	// 设置字体大小
	content.SetFontSize(32)
	// 设置字体颜色
	content.SetSrc(image.Black)
	berrors.Unwrap(content.DrawString("黑哥TV DOTA分部", freetype.Pt(650, 550)))
	berrors.Unwrap(content.DrawString("二零二二年九月", freetype.Pt(650, 590)))
	saveFile(newTemplateImage)
}

func saveFile(pic *image.RGBA) {
	dstFile, err := os.Create("bin/2.png")
	if err != nil {
		fmt.Println(err)
	}
	defer dstFile.Close()
	png.Encode(dstFile, pic)
}

func loadFont(path string) (font *truetype.Font, err error) {
	fontBytes, err := os.ReadFile(path) // 读取字体文件
	if err != nil {
		err = fmt.Errorf("加载字体文件出错:%s", err.Error())
		return
	}
	font, err = freetype.ParseFont(fontBytes) // 解析字体文件
	if err != nil {
		err = fmt.Errorf("解析字体文件出错,%s", err.Error())
		return
	}
	return
}
