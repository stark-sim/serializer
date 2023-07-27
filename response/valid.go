package response

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translation "github.com/go-playground/validator/v10/translations/zh"
	"strconv"
)

var (
	trans ut.Translator
)

func init() {
	trans, _ = ut.New(zh.New()).GetTranslator("zh")
	if err := translation.RegisterDefaultTranslations(binding.Validator.Engine().(*validator.Validate), trans); err != nil {
		panic(fmt.Sprintf("register default translations failed, err: %v", err))
	}
}

func validError(err error) (ret string) {
	switch err.(type) {
	case validator.ValidationErrors:
		for _, e := range err.(validator.ValidationErrors) {
			ret += e.Translate(trans) + ";"
		}
	case *strconv.NumError:
		errContent := err.(*strconv.NumError).Num
		ret = errContent + "不是数字"
	default:
		ret = err.Error()
	}
	return
}
