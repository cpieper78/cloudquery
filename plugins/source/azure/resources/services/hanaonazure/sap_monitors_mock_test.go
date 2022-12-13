// Code generated by codegen2; DO NOT EDIT.
package hanaonazure

import (
	"encoding/json"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hanaonazure/armhanaonazure"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
)

func createSapMonitors(router *mux.Router) error {
	var item armhanaonazure.SapMonitorsClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyStr := ""
	item.NextLink = &emptyStr

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.HanaOnAzure/sapMonitors", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&item)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})
	return nil
}

func TestSapMonitors(t *testing.T) {
	client.MockTestHelper(t, SapMonitors(), createSapMonitors)
}