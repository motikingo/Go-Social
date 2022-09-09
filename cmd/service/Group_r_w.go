package service
import "github.com/motikingo/websocketproject/internal/pkg/entity"


type WSGroup struct {
	Group *entity.Group
	ActiveCount int 
	MembersID []string 
}