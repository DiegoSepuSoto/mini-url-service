package miniurls

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/DiegoSepuSoto/mini-url-service/src/shared"
)

func TestGetMinifiedURLMongoDBIntegration(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image: "mongo:6.0.1",
		Env: map[string]string{
			"MONGO_INITDB_ROOT_USERNAME": "root",
			"MONGO_INITDB_ROOT_PASSWORD": "password",
			"MONGO_INITDB_DATABASE":      "marketingDB",
		},
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForLog("Waiting for connections"),
	}

	mongoDBContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Error(err)
	}

	endpoint, err := mongoDBContainer.Endpoint(ctx, "")
	if err != nil {
		t.Error(err)
	}

	mongoURI := fmt.Sprintf("mongodb://root:password@%s/marketingDB?authSource=admin", endpoint)

	t.Setenv("MONGODB_URI", mongoURI)

	mongoDBCollection := shared.CreateMongoDBCollection()
	mongoDBRepository := NewMongoDBMiniURLsRepository(mongoDBCollection)

	minifiedURL, err := mongoDBRepository.GetMinifiedURL(ctx, "abc123")

	assert.Equal(t, "", minifiedURL)
	assert.Equal(t, "[DB_NOT_FOUND]: mongo: no documents in result at miniURLsRepository - sending: 404", err.Error())

	defer func() {
		if err := mongoDBContainer.Terminate(ctx); err != nil {
			panic(err)
		}
	}()
}
