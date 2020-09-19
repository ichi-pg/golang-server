package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	submoduleContexts "github.com/ichi-pg/golang-middleware/contexts"
	"github.com/ichi-pg/golang-server/internal/application/injection"
	"github.com/ichi-pg/golang-server/internal/infrastructure/mock"
	"github.com/ichi-pg/golang-server/internal/presentation/echo/contexts"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGraphQuery(t *testing.T) {
	t.Parallel()

	log := logrus.NewEntry(logrus.New())
	i := injection.MockInjector()
	user := &mock.User
	normalMutation := " mutation { updateUserName ( name : \\\"fuga\\\" ) { name } } "

	// Test case: Get user.
	{
		query := " query { user { name } } "

		body := strings.NewReader(fmt.Sprintf("{ \"query\" : \"%s\" }", query))

		req := httptest.NewRequest(http.MethodPost, "/", body)
		rec := httptest.NewRecorder()

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		c := echo.New().NewContext(req, rec)

		submoduleContexts.SetLogger(c, log)
		contexts.SetInjector(c, i)
		contexts.SetUser(c, user)

		if assert.NoError(t, GraphQuery(c)) {
			assert.Equal(t, rec.Code, http.StatusOK)
			assert.Equal(t, rec.Body.String(), fmt.Sprintf("{\"data\":{\"user\":{\"name\":\"hoge\"}}}"))
		}
	}

	// Test case: Update user.
	{
		body := strings.NewReader(fmt.Sprintf("{ \"query\" : \"%s\" }", normalMutation))

		req := httptest.NewRequest(http.MethodPost, "/", body)
		rec := httptest.NewRecorder()

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		c := echo.New().NewContext(req, rec)

		submoduleContexts.SetLogger(c, log)
		contexts.SetInjector(c, i)
		contexts.SetUser(c, user)

		if assert.NoError(t, GraphQuery(c)) {
			assert.Equal(t, rec.Code, http.StatusOK)
			assert.Equal(t, rec.Body.String(), "{\"data\":{\"updateUserName\":{\"name\":\"fuga\"}}}")
		}
	}

	// Test case: Bad request.
	{
		mutation := " mutation { updateUserName ( name : \\\"\\\" ) { name } } "

		body := strings.NewReader(fmt.Sprintf("{ \"query\" : \"%s\" }", mutation))

		req := httptest.NewRequest(http.MethodPost, "/", body)
		rec := httptest.NewRecorder()

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		c := echo.New().NewContext(req, rec)

		submoduleContexts.SetLogger(c, log)
		contexts.SetInjector(c, i)
		contexts.SetUser(c, user)

		if assert.NoError(t, GraphQuery(c)) {
			assert.Equal(t, rec.Code, http.StatusBadRequest)
			assert.Equal(t, rec.Body.String(), "{\"errors\":[{\"message\":\"ユーザー名が空です。\",\"path\":[\"updateUserName\"],\"extensions\":{\"status\":400}}],\"data\":null}")
		}
	}

	//TODO Test case: Internal server error.
	// {
	// 	query := " query { rankers( offset: 0, limit: 10 ) { userID, rank, score } } "

	// 	body := strings.NewReader(fmt.Sprintf("{ \"query\" : \"%s\" }", query))

	// 	req := httptest.NewRequest(http.MethodPost, "/", body)
	// 	rec := httptest.NewRecorder()

	// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// 	c := echo.New().NewContext(req, rec)

	// 	submoduleContexts.SetLogger(c, log)
	// 	contexts.SetInjector(c, injection.AppInjector())

	// 	if assert.NoError(t, GraphQuery(c)) {
	// 		assert.Equal(t, rec.Code, http.StatusInternalServerError)
	// 		assert.Equal(t, rec.Body.String(), fmt.Sprintf("{\"errors\":[{\"message\":\"dial tcp [::1]:6379: connect: connection refused\",\"path\":[\"rankers\"]}],\"data\":null}"))
	// 	}
	// }

	// Test case: Validation error.
	{
		mutation := " mutation { updateUserName { name } } "

		body := strings.NewReader(fmt.Sprintf("{ \"query\" : \"%s\" }", mutation))

		req := httptest.NewRequest(http.MethodPost, "/", body)
		rec := httptest.NewRecorder()

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		c := echo.New().NewContext(req, rec)

		submoduleContexts.SetLogger(c, log)
		contexts.SetInjector(c, i)
		contexts.SetUser(c, user)

		if assert.NoError(t, GraphQuery(c)) {
			assert.Equal(t, rec.Code, http.StatusUnprocessableEntity)
			assert.Equal(t, rec.Body.String(), "{\"errors\":[{\"message\":\"Field \\\"updateUserName\\\" argument \\\"name\\\" of type \\\"String!\\\" is required but not provided.\",\"locations\":[{\"line\":1,\"column\":13}],\"extensions\":{\"code\":\"GRAPHQL_VALIDATION_FAILED\"}}],\"data\":null}")
		}
	}

	// Test case: Parse error.
	{
		mutation := " mutation { { name } } "

		body := strings.NewReader(fmt.Sprintf("{ \"query\" : \"%s\" }", mutation))

		req := httptest.NewRequest(http.MethodPost, "/", body)
		rec := httptest.NewRecorder()

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		c := echo.New().NewContext(req, rec)

		submoduleContexts.SetLogger(c, log)
		contexts.SetInjector(c, i)
		contexts.SetUser(c, user)

		if assert.NoError(t, GraphQuery(c)) {
			assert.Equal(t, rec.Code, http.StatusUnprocessableEntity)
			assert.Equal(t, rec.Body.String(), "{\"errors\":[{\"message\":\"Expected Name, found {\",\"locations\":[{\"line\":1,\"column\":13}],\"extensions\":{\"code\":\"GRAPHQL_PARSE_FAILED\"}}],\"data\":null}")
		}
	}

	// Test case: Operation not found.
	{
		body := strings.NewReader(fmt.Sprintf("{ \"mutation\" : \"%s\" }", normalMutation))

		req := httptest.NewRequest(http.MethodPost, "/", body)
		rec := httptest.NewRecorder()

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		c := echo.New().NewContext(req, rec)

		submoduleContexts.SetLogger(c, log)
		contexts.SetInjector(c, i)
		contexts.SetUser(c, user)

		if assert.NoError(t, GraphQuery(c)) {
			assert.Equal(t, rec.Code, http.StatusBadRequest)
			assert.Equal(t, rec.Body.String(), "{\"errors\":[{\"message\":\"operation  not found\"}],\"data\":null}")
		}
	}

	// Test case: Transport not supported.
	{
		body := strings.NewReader(fmt.Sprintf("{ \"mutation\" : \"%s\" }", normalMutation))

		req := httptest.NewRequest(http.MethodPost, "/", body)
		rec := httptest.NewRecorder()

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

		c := echo.New().NewContext(req, rec)

		submoduleContexts.SetLogger(c, log)
		contexts.SetInjector(c, i)
		contexts.SetUser(c, user)

		if assert.NoError(t, GraphQuery(c)) {
			assert.Equal(t, rec.Code, http.StatusBadRequest)
			assert.Equal(t, rec.Body.String(), "{\"errors\":[{\"message\":\"transport not supported\"}],\"data\":null}")
		}
	}

	// Test case: Json body could not be decoded.
	{
		body := strings.NewReader(normalMutation)

		req := httptest.NewRequest(http.MethodPost, "/", body)
		rec := httptest.NewRecorder()

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		c := echo.New().NewContext(req, rec)

		submoduleContexts.SetLogger(c, log)
		contexts.SetInjector(c, i)
		contexts.SetUser(c, user)

		if assert.NoError(t, GraphQuery(c)) {
			assert.Equal(t, rec.Code, http.StatusBadRequest)
			assert.Equal(t, rec.Body.String(), "{\"errors\":[{\"message\":\"json body could not be decoded: invalid character 'm' looking for beginning of value\"}],\"data\":null}")
		}
	}
}
