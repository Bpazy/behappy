package images

import (
	"fmt"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
)

func main() {
	templateFile, err := os.Open("./1.png")
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
	fontKai, err := loadFont("C:\\Windows\\Fonts\\simkai.ttf")
	content.SetFont(fontKai) // 设置字体样式，就是我们上面加载的字体

	// 	正式写入文字
	// 参数1：要写入的文字
	// 参数2：文字坐标
	// TODO 考虑通过返回值自适应换行
	drawString, err := content.DrawString("某某同志:", freetype.Pt(120, 320))
	fmt.Println("", drawString)
	unwrap(drawString, err)
	unwrap(content.DrawString("您在2022年度中表现突出，忠诚奉献、认真负责，", freetype.Pt(160, 390)))
	unwrap(content.DrawString("被评为", freetype.Pt(120, 460)))

	unwrap(content.DrawString("特发此证，以资鼓励。", freetype.Pt(520, 460)))
	// 设置字体大小
	content.SetFontSize(48)
	// 设置字体颜色
	content.SetSrc(image.NewUniform(color.RGBA{R: 237, G: 39, B: 90, A: 255}))
	unwrap(content.DrawString("最佳奉献奖，", freetype.Pt(250, 460)))

	content.SetFont(fontKai) // 设置字体样式
	// 设置字体大小
	content.SetFontSize(32)
	// 设置字体颜色
	content.SetSrc(image.Black)
	unwrap(content.DrawString("黑哥TV DOTA分部", freetype.Pt(650, 550)))
	unwrap(content.DrawString("二零二二年九月", freetype.Pt(650, 590)))
	saveFile(newTemplateImage)
}

func saveFile(pic *image.RGBA) {
	dstFile, err := os.Create("./2.png")
	if err != nil {
		fmt.Println(err)
	}
	defer dstFile.Close()
	png.Encode(dstFile, pic)
}

func unwrap[T interface{}](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
func loadFont(path string) (font *truetype.Font, err error) {
	var fontBytes []byte
	fontBytes, err = ioutil.ReadFile(path) // 读取字体文件
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
