package auroragee

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
)

type H map[string]interface{}

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request

	Path       string
	Method     string
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

/****** 请求解析 *******/

// TODO: 解析路由参数
// func (c *Context) Param(key string) string {

// }

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) ShouldBindJSON(j interface{}) error {
	body, err := io.ReadAll(c.Req.Body)
	if err != nil {
		return err
	}

	defer func() {
		c.Req.Body.Close()
	}()

	err = json.Unmarshal(body, j)
	if err != nil {
		return err
	}

	return nil
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// 解析单文件
func (c *Context) FormFile(key string) (multipart.File, *multipart.FileHeader, error) {
	file, header, err := c.Req.FormFile(key)
	if err != nil {
		return nil, nil, err
	}
	return file, header, nil
}

// 解析多文件
func (c *Context) MultipartForm(key string) ([]*multipart.FileHeader, error) {
	err := c.Req.ParseMultipartForm(1024 * 1024 * 10)

	if err != nil {
		return nil, err
	}

	files := c.Req.MultipartForm.File[key]
	if len(files) == 0 {
		return nil, fmt.Errorf("no files uploaded")
	}

	return files, nil

}

/****** 返回 *******/
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.SetHeader("Content-Type", "application/octet-stream")
	c.SetHeader("Content-Length", strconv.Itoa(len(data)))
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
