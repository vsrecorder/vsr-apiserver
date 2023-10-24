package middlewares

import (
	"encoding/base64"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	ulid "github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/require"
	"github.com/vsrecorder/vsr-apiserver/pkg/controllers/helpers"
)

var (
	seed = time.Now().UnixNano()
)

func setup() {
	r := rand.New(rand.NewSource(seed))
	src := make([]byte, 32)
	r.Read(src)

	secretKey := base64.StdEncoding.EncodeToString(src)
	os.Setenv("VSRECORDER_JWT_SECRET", secretKey)
}

func TestRequiredAuthorization(t *testing.T) {
	setup()

	for scenario, fn := range map[string]func(
		t *testing.T,
	){
		"ValidRequiredAuthorization":   test_ValidRequiredAuthorization,
		"InvalidRequiredAuthorization": test_InvalidRequiredAuthorization,
	} {
		fn := fn // ↑引数の順番通りに実行するための設定(https://github.com/golang/go/wiki/CommonMistakes)
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}

func TestOptionalAuthorization(t *testing.T) {
	setup()

	for scenario, fn := range map[string]func(
		t *testing.T,
	){
		"ValidOptionalAuthorization":   test_ValidOptionalAuthorization,
		"InvalidOptionalAuthorization": test_InvalidOptionalAuthorization,
	} {
		fn := fn // ↑引数の順番通りに実行するための設定(https://github.com/golang/go/wiki/CommonMistakes)
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})

	}
}

func test_ValidRequiredAuthorization(t *testing.T) {
	entropy := rand.New(rand.NewSource(seed))
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, entropy)
	require.NoError(t, err)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	secretKey := os.Getenv("VSRECORDER_JWT_SECRET")
	tokenString, err := generateToken(id.String(), secretKey)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	require.NoError(t, err)

	req.Header.Add("Authorization", "Bearer "+tokenString)
	ctx.Request = req

	RequiredAuthorization(ctx)

	expectedUID := id.String()
	expectedExists := true

	actualUID, actualExists := helpers.GetUID(ctx)

	require.Equal(t, expectedUID, actualUID)
	require.Equal(t, expectedExists, actualExists)
}

func test_InvalidRequiredAuthorization(t *testing.T) {
	entropy := rand.New(rand.NewSource(seed))
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, entropy)
	require.NoError(t, err)

	{
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

		secretKey := os.Getenv("VSRECORDER_JWT_SECRET")
		tokenString, err := generateToken(id.String(), secretKey)
		require.NoError(t, err)

		invaliedTokenString := "invalid" + tokenString

		req, err := http.NewRequest(http.MethodGet, "/", nil)
		require.NoError(t, err)

		req.Header.Add("Authorization", "Bearer "+invaliedTokenString)
		ctx.Request = req

		RequiredAuthorization(ctx)

		expectedUID := ""
		expectedExists := false
		actualUID, actualExists := helpers.GetUID(ctx)
		require.Equal(t, expectedUID, actualUID)
		require.Equal(t, expectedExists, actualExists)

		expectedStatus := http.StatusUnauthorized
		actualStatus := ctx.Writer.Status()
		require.Equal(t, expectedStatus, actualStatus)
	}

	{
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

		req, err := http.NewRequest(http.MethodGet, "/", nil)
		require.NoError(t, err)
		ctx.Request = req

		RequiredAuthorization(ctx)

		expectedUID := ""
		expectedExists := false
		actualUID, actualExists := helpers.GetUID(ctx)
		require.Equal(t, expectedUID, actualUID)
		require.Equal(t, expectedExists, actualExists)

		expectedStatus := http.StatusUnauthorized
		actualStatus := ctx.Writer.Status()
		require.Equal(t, expectedStatus, actualStatus)

	}
}

func test_ValidOptionalAuthorization(t *testing.T) {
	entropy := rand.New(rand.NewSource(seed))
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, entropy)
	require.NoError(t, err)

	{

		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

		secretKey := os.Getenv("VSRECORDER_JWT_SECRET")
		tokenString, err := generateToken(id.String(), secretKey)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodGet, "/", nil)
		require.NoError(t, err)

		req.Header.Add("Authorization", "Bearer "+tokenString)
		ctx.Request = req

		RequiredAuthorization(ctx)

		expectedUID := id.String()
		expectedExists := true
		actualUID, actualExists := helpers.GetUID(ctx)
		require.Equal(t, expectedUID, actualUID)
		require.Equal(t, expectedExists, actualExists)
	}

	{

		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

		req, err := http.NewRequest(http.MethodGet, "/", nil)
		require.NoError(t, err)

		ctx.Request = req

		RequiredAuthorization(ctx)

		expectedUID := ""
		expectedExists := false
		actualUID, actualExists := helpers.GetUID(ctx)
		require.Equal(t, expectedUID, actualUID)
		require.Equal(t, expectedExists, actualExists)
	}
}

func test_InvalidOptionalAuthorization(t *testing.T) {
	entropy := rand.New(rand.NewSource(seed))
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, entropy)
	require.NoError(t, err)

	{
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

		secretKey := os.Getenv("VSRECORDER_JWT_SECRET")
		tokenString, err := generateToken(id.String(), secretKey)
		require.NoError(t, err)

		invaliedTokenString := "invalid" + tokenString

		req, err := http.NewRequest(http.MethodGet, "/", nil)
		require.NoError(t, err)

		req.Header.Add("Authorization", "Bearer "+invaliedTokenString)
		ctx.Request = req

		RequiredAuthorization(ctx)

		expectedUID := ""
		expectedExists := false
		actualUID, actualExists := helpers.GetUID(ctx)
		require.Equal(t, expectedUID, actualUID)
		require.Equal(t, expectedExists, actualExists)

		expectedStatus := http.StatusUnauthorized
		actualStatus := ctx.Writer.Status()
		require.Equal(t, expectedStatus, actualStatus)
	}
}
