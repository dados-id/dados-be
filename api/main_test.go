package api

import (
	"os"
	"testing"

	"github.com/dados-id/dados-be/config"
	db "github.com/dados-id/dados-be/db/sqlc"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func newTestServer(t *testing.T, store db.Querier) *Server {
	configuration := config.Config{}

	server, err := NewServer(configuration, store)
	require.NoError(t, err)

	return server
}
