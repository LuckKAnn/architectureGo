package article_service

import (
	"ginDemo/pkg/file"
	"ginDemo/pkg/qrcode"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

// 海报
type ArticlePoster struct {
	PosterName string
	*Article
	Qr *qrcode.QrCode
}

func NewArticlePoster(postName string, article *Article, code *qrcode.QrCode) *ArticlePoster {
	return &ArticlePoster{PosterName: postName, Article: article, Qr: code}
}
func GetPosterFlag() string {
	return "poster"
}

func (a *ArticlePoster) CheckMergedImage(path string) bool {
	if file.CheckExist(path+a.PosterName) == false {
		return false
	}
	return true
}

func (a *ArticlePoster) OpenMergedImage(path string) (*os.File, error) {
	f, err := file.MustOpen(a.PosterName, path)
	if err != nil {
		return nil, err
	}

	return f, nil
}

type ArticlePosterBg struct {
	Name string
	*ArticlePoster
	*Rect
	*Pt
}

func NewArticlePosterBg(name string, poster *ArticlePoster, rect *Rect, pt *Pt) *ArticlePosterBg {
	return &ArticlePosterBg{Name: name, ArticlePoster: poster, Rect: rect, Pt: pt}
}

// 图片大小
type Rect struct {
	Name string
	X0   int
	Y0   int
	X1   int
	Y1   int
}

// 插入图所处的位置
type Pt struct {
	X int
	Y int
}

func (a *ArticlePosterBg) Generate() (string, string, error) {
	fullPath := qrcode.GetQrCodeFullSavePath()
	fileName, path, err := a.Qr.Encode(fullPath)
	if err != nil {
		return "", "", err
	}
	if !a.CheckMergedImage(path) {
		// 如果合并后的图不在
		mergedF, err := a.OpenMergedImage(path)
		defer mergedF.Close()
		bgF, err := file.MustOpen(a.Name, path)
		if err != nil {
			return "", "", err
		}
		defer bgF.Close()

		qrF, err := file.MustOpen(fileName, path)
		if err != nil {
			return "", "", err
		}
		defer qrF.Close()
		bgImage, err := jpeg.Decode(bgF)
		if err != nil {
			return "", "", err
		}
		qrImage, err := jpeg.Decode(qrF)
		if err != nil {
			return "", "", err
		}
		// 创建新的图像
		jpg := image.NewRGBA(image.Rect(a.Rect.X0, a.Rect.Y0, a.Rect.X1, a.Rect.Y1))

		draw.Draw(jpg, jpg.Bounds(), bgImage, bgImage.Bounds().Min, draw.Over)
		draw.Draw(jpg, jpg.Bounds(), qrImage, qrImage.Bounds().Min.Sub(image.Pt(a.Pt.X, a.Pt.Y)), draw.Over)

		// 在mergedF的基础上去建造jpg
		jpeg.Encode(mergedF, jpg, nil)
	}

	return fileName, path, nil

}
