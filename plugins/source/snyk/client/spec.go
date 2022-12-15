package client

import (
	"fmt"
	"os"

	"github.com/pavel-snyk/snyk-sdk-go/snyk"
)

type Spec struct {
	// APIKey can be set via spec or via "SNYK_API_KEY" environment variable
	APIKey string `json:"api_key,omitempty"`

	// Organizations is a list of organizations to fetch information from.
	// By default will fetch from all organizations available for user.
	Organizations []string `json:"organizations,omitempty"`

	// EndpointURL is optional parameter to override the API URL for snyk.Client.
	EndpointURL string `json:"endpoint_url,omitempty"`
}

const (
	EnvSnykApiKey = "SNYK_API_KEY"
)

func (s *Spec) getClient(version string) (*snyk.Client, error) {
	if len(s.APIKey) == 0 {
		s.APIKey = os.Getenv(EnvSnykApiKey)
		if len(s.APIKey) == 0 {
			return nil, fmt.Errorf("missing API key")
		}
	}

	options := []snyk.ClientOption{snyk.WithUserAgent("cloudquery/snyk/" + version)}
	if len(s.EndpointURL) > 0 {
		options = append(options, snyk.WithBaseURL(s.EndpointURL))
	}

	return snyk.NewClient(s.APIKey, options...), nil
}
