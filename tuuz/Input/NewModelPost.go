package Input

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/Calc"
	"html/template"
	"main.go/tuuz/Date"
	"time"
)

type ModelPost struct {
	reserv_col map[string]bool
	content    *gin.Context
	*modelConfig
	*modelData
}

type modelConfig struct {
	xss bool
}

type modelData struct {
	errMsgs []string
	errs    []error
	data    map[string]interface{}
}

func NewModelPost(c *gin.Context) *ModelPost {
	return &ModelPost{
		make(map[string]bool),
		c,
		&modelConfig{
			xss: false,
		},
		&modelData{
			errMsgs: nil,
			errs:    nil,
			data:    make(map[string]interface{}),
		},
	}
}

func (post *ModelPost) Errors() ([]error, []string) {
	return post.errs, post.errMsgs
}

func (post *ModelPost) Error() (err error, errMsg string) {
	if post.errMsgs != nil && len(post.errMsgs) > 0 {
		errMsg = post.errMsgs[0]
	}
	if post.errs != nil && len(post.errs) > 0 {
		err = post.errs[0]
	}
	return
}

func (post *ModelPost) IsComplete() bool {
	if post.errs != nil {
		return false
	}
	if len(post.errs) > 0 {
		return false
	}
	return true
}

func (post *ModelPost) GetPostMap() (data map[string]interface{}, err error, errMsg string) {
	err, errMsg = post.Error()
	data = post.data
	return
}

func (post *ModelPost) Select() (data map[string]interface{}) {
	data = post.data
	return
}

// Fields: 如果需要保证字段一定存在，则使用fields，否则默认允许所有字段均不传
func (post *ModelPost) Fields(fields ...string) *ModelPost {
	for _, field := range fields {
		post.reserv_col[field] = true
	}
	return post
}

func (post *ModelPost) Xss(anti_xss bool) *ModelPost {
	post.xss = anti_xss
	return post
}

func (post *ModelPost) Data(field string, value interface{}) *ModelPost {
	post.data[field] = value
	return post
}

func (post *ModelPost) Copy(from_field string, to_field string) *ModelPost {
	post.data[to_field] = post.data[from_field]
	return post
}

func (post *ModelPost) PostString(key string) *ModelPost {
	_, have := post.reserv_col[key]
	in, ok := post.content.GetPostForm(key)
	if !ok {
		if have {
			post.errMsgs = append(post.errMsgs, "POST-["+key+"]")
			post.errs = append(post.errs, errors.New("POST-["+key+"]"))
		}
	} else {
		if post.xss {
			post.data[key] = template.JSEscapeString(in)
		} else {
			post.data[key] = in
		}
	}
	return post
}

func (post *ModelPost) PostInt64(key string) *ModelPost {
	_, have := post.reserv_col[key]
	in, ok := post.content.GetPostForm(key)
	if !ok {
		if have {
			post.errMsgs = append(post.errMsgs, "POST-["+key+"]")
			post.errs = append(post.errs, errors.New("POST-["+key+"]"))
		}
	} else {
		i, e := Calc.String2Int64(in)
		if e != nil {
			post.errMsgs = append(post.errMsgs, key+" should be int64")
			post.errs = append(post.errs, e)
			return post
		}
		post.data[key] = i
	}
	return post
}

func (post *ModelPost) PostDateTime(key string) *ModelPost {
	_, have := post.reserv_col[key]
	in, ok := post.content.GetPostForm(key)
	if !ok {
		if have {
			post.errMsgs = append(post.errMsgs, "POST-["+key+"]")
			post.errs = append(post.errs, errors.New("POST-["+key+"]"))
		}
	} else {
		i, e := Date.Date_time_parser(in)
		if e != nil {
			post.errMsgs = append(post.errMsgs, key+" should only be a Date(+Time) or RFC3339")
			post.errs = append(post.errs, e)
			return post
		}
		post.data[key] = i
	}
	return post
}

func (post *ModelPost) PostTime(key string) *ModelPost {
	_, have := post.reserv_col[key]
	in, ok := post.content.GetPostForm(key)
	if !ok {
		if have {
			post.errMsgs = append(post.errMsgs, "POST-["+key+"]")
			post.errs = append(post.errs, errors.New("POST-["+key+"]"))
		}
	} else {
		i, e := Calc.String2Int64(in)
		if e != nil {
			post.errMsgs = append(post.errMsgs, key+" should be int64")
			post.errs = append(post.errs, e)
			return post
		}
		post.data[key] = time.Unix(i, 0)
	}
	return post
}

func (post *ModelPost) PostFloat64(key string) *ModelPost {
	_, have := post.reserv_col[key]
	in, ok := post.content.GetPostForm(key)
	if !ok {
		if have {
			post.errMsgs = append(post.errMsgs, "POST-["+key+"]")
			post.errs = append(post.errs, errors.New("POST-["+key+"]"))
		}
	} else {
		i, e := Calc.String2Float64(in)
		if e != nil {
			post.errMsgs = append(post.errMsgs, key+" should be float64")
			post.errs = append(post.errs, e)
			return post
		}
		post.data[key] = i
	}
	return post
}

func (post *ModelPost) PostBool(key string) *ModelPost {
	_, have := post.reserv_col[key]
	in, ok := post.content.GetPostForm(key)
	if !ok {
		if have {
			post.errMsgs = append(post.errMsgs, "POST-["+key+"]")
			post.errs = append(post.errs, errors.New("POST-["+key+"]"))
		}
	} else {
		switch in {
		case "1":
			post.data[key] = true
			break

		case "0":
			post.data[key] = false
			break

		case "true":
			post.data[key] = true
			break

		case "false":
			post.data[key] = false
			break

		default:
			post.errMsgs = append(post.errMsgs, key+" should be Boolean")
			post.errs = append(post.errs, errors.New(key+" should be boolean"))
		}
	}
	return post
}
