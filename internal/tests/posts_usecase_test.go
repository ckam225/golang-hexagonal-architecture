package tests

import (
	"clean-arch-hex/internal/cache/memcache"
	"clean-arch-hex/internal/controller/server/rest"
	mockdb "clean-arch-hex/internal/db/mocks"
	"clean-arch-hex/internal/domain/entity"
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := mockdb.NewMockDatabase(ctrl)

	// build stubs
	post := entity.Post{
		Title:   "Hello",
		Content: "Hi everyboby",
	}
	err := db.CreatePost(context.TODO(), &post)
	require.NoError(t, err)
	// db.EXPECT().CreatePost(context.TODO(), &post).
	// 	Return(nil)
	// require.NotEqual(t, post.ID, 0)
	log.Println("======PRINTED=====", post)

}

func TestHandler_createPosts(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *mockdb.MockDatabase, p entity.Post)

	tests := []struct {
		name                 string
		inputBody            string
		inputObj             *entity.Post
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			inputBody: `{
				"title": "hjin",
				"content": "898gjg. hjghjjhjg"
			  }`,
			inputObj: &entity.Post{
				Title:   "item 1",
				Content: "content 1",
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id": 1, "title": "item 1", "content": "content 1", "user_id": null}`,
		},
		// {
		// 	name:                 "Unprocessable Entity",
		// 	inputBody:            `{}`,
		// 	inputUser:            entity.Post{},
		// 	mockBehavior:         func(r *mockdb.MockDatabase, p entity.Post) {},
		// 	expectedStatusCode:   422,
		// 	expectedResponseBody: `{"code": 422, "message": "Unprocessable Entity"}`,
		// },
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			db := mockdb.NewMockDatabase(c)
			db.EXPECT().CreatePost(context.TODO(), test.inputObj).Return(nil)
			// db.EXPECT().CreatePost(gomock.Any(), gomock.Eq(test.inputObj)).Times(1).Return(nil)

			serv := rest.New(db, memcache.New())
			// Create Request
			// req := httptest.NewRequest(http.MethodPost, "/posts", bytes.NewBufferString(test.inputBody))
			req := httptest.NewRequest(http.MethodPost, "/posts", strings.NewReader(test.inputBody))

			// Make Request
			res, _ := serv.Test(req)
			// Read response
			defer res.Body.Close()
			body, _ := io.ReadAll(res.Body)
			// Assert
			log.Println("======PRINTED=====", string(body))
			assert.Equal(t, res.StatusCode, test.expectedStatusCode)
			// assert.Equal(t, string(body), test.expectedResponseBody)
		})
	}
}
