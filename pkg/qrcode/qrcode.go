package qrcode

import (
	"ginDemo/pkg/file"
	"ginDemo/pkg/setting"
	"ginDemo/pkg/util"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/jpeg"
)

type QrCode struct {
	URL    string
	Width  int64
	Height int64
	Ext    string
	Level  qr.ErrorCorrectionLevel
	Mode   qr.Encoding
}

// 生成的结果
const EXT_JPG = ".jpg"

func GetQrCodeSavePath() string {
	return setting.AppSetting.QrCodeSavePath
}

func GetQrCodeFullSavePath() string {
	return setting.AppSetting.RuntimeRootPath + GetQrCodeSavePath()
}
func GetQrCodeFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetQrCodeSavePath() + name
}

func NewQrCode(url string, width, height int64, level qr.ErrorCorrectionLevel, mode qr.Encoding) *QrCode {
	return &QrCode{
		URL:    url,
		Width:  width,
		Height: height,
		Level:  level,
		Mode:   mode,
		Ext:    EXT_JPG,
	}
}

func GetQrCodeFileName(name string) string {
	return util.EncodeMD5(name)
}
func (qr *QrCode) GetQrCodeExt() string {
	return qr.Ext
}

func (q *QrCode) CheckEncode(path string) bool {
	src := path + GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	if file.CheckExist(src) == false {
		return false
	}
	return true
}
func (q *QrCode) Encode(path string) (string, string, error) {
	// 这里的需求是把path 进行编码？
	name := GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	src := path + name
	if file.CheckExist(src) == false {
		// not exist
		code, err := qr.Encode(q.URL, q.Level, q.Mode)
		if err != nil {
			return "", "", err
		}

		code, err = barcode.Scale(code, int(q.Width), int(q.Height))
		if err != nil {
			return "", "", err
		}

		f, err := file.MustOpen(name, path)
		if err != nil {
			return "", "", err
		}
		defer f.Close()

		err = jpeg.Encode(f, code, nil)
		if err != nil {
			return "", "", err
		}

	}
	return name, path, nil
}
