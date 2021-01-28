package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type errResponse struct {
	ErrorCode int32  `json:"error_code"`
	Error     string `json:"error"`
}

// 返回JSON体。
func ResponseJSON(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, obj)
}

// 将Proto结构按照JSON协议序列化，然后返回。
func ResponseProto(c *gin.Context, p proto.Message) {
	c.Render(http.StatusOK, ProtoBuf{p})
}

// ProtoBuf contains the given interface object.
type ProtoBuf struct {
	Data proto.Message
}

// Render (ProtoBuf) marshals the given interface object and writes data with custom ContentType.
func (r ProtoBuf) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)

	m := protojson.MarshalOptions{EmitUnpopulated: true}
	s, err := m.Marshal(r.Data)
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(s))
	return err
}

// WriteContentType (ProtoBuf) writes ProtoBuf ContentType.
func (r ProtoBuf) WriteContentType(w http.ResponseWriter) {
	value := []string{"application/json; charset=utf-8"}
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}
