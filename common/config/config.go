package config

import (
	"context"
	"os"
	"github.com/GoogleCloudPlatform/berglas/pkg/berglas"
)


func GetSecret(obj string) (string, error) {
	// Check environment variables
	val, exists := os.LookupEnv(obj)
	if exists {
		return val, nil
	}

	// Fetch from Google KMS
	ctx := context.Background()

	resp, err := berglas.Access(ctx, &berglas.AccessRequest{
		Bucket: "github-ntf-secrets",
		Object: obj,
	})
	if err != nil {
		return "", err
	}

	return string(resp), nil
}
