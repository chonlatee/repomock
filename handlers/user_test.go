package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chonlatee/repomock/models"
	mock_repository "github.com/chonlatee/repomock/repository/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_GetUserByID(t *testing.T) {

	t.Run("http status ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		um := mock_repository.NewMockUsers(ctrl)
		um.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Return(&models.User{
			UserID:    "1234",
			FirstName: "foo",
			LastName:  "bar",
		}, nil)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		u := NewUserAPI(um)
		u.GetUserByID(c)

		var res UserRes

		err := json.Unmarshal(w.Body.Bytes(), &res)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, "1234", res.UserID)
		require.Equal(t, "foo", res.FirstName)
		require.Equal(t, "bar", res.LastName)
	})

	t.Run("http badgate way", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		um := mock_repository.NewMockUsers(ctrl)
		um.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("find user error"))
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		u := NewUserAPI(um)
		u.GetUserByID(c)

		require.Equal(t, http.StatusInternalServerError, w.Code)
	})

}
