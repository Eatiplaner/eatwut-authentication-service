package serializers

import (
	"Eatiplan-Auth/app/models"

	"github.com/danhper/structomap"
)

func UserSerializer(user models.User) map[string]interface{} {
	userSerializer := structomap.New().UseSnakeCase().Pick("ID", "FirstName", "LastName", "Email", "Phone", "UserName")
	userMap := userSerializer.Transform(user)

	return userMap
}
