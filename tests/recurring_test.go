/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package tests

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/adyen/adyen-go-api-library/src/common"

	adyen "github.com/adyen/adyen-go-api-library/src/api"
	"github.com/adyen/adyen-go-api-library/src/recurring"
	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Recurring(t *testing.T) {
	godotenv.Load("./../.env")

	var (
		APIKey          = os.Getenv("ADYEN_API_KEY")
		MerchantAccount = os.Getenv("ADYEN_MERCHANT")
	)

	client := adyen.NewAPIClientWithAPIKey(APIKey, "TEST")
	// client.GetConfig().Debug = true

	t.Run("ListRecurringDetails", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			res, httpRes, err := client.Recurring.ListRecurringDetailsPost(&recurring.RecurringDetailsRequest{})
			require.NotNil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, "403 Forbidden", err.Error())
			assert.Equal(t, 403, httpRes.StatusCode)
			assert.Equal(t, "403 Forbidden", httpRes.Status)
			require.NotNil(t, res)
			errBody := map[string]string{}
			json.Unmarshal(err.(common.GenericOpenAPIError).Body(), &errBody)
			assert.Equal(t, "Invalid Merchant Account", errBody["message"])
		})
		t.Run("Create an API request that should pass", func(t *testing.T) {
			res, httpRes, err := client.Recurring.ListRecurringDetailsPost(&recurring.RecurringDetailsRequest{
				MerchantAccount:  MerchantAccount,
				ShopperReference: "4343553GFGFYFY4654654675765",
				Recurring: &recurring.RecurringType{
					Contract: "RECURRING",
				},
			})
			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, "200 OK", httpRes.Status)
			require.NotNil(t, res)
			assert.Empty(t, res.ShopperReference)
			assert.Empty(t, res.Details)
		})
	})

	t.Run("Disable", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			res, httpRes, err := client.Recurring.DisablePost(&recurring.DisableRequest{
				MerchantAccount:          MerchantAccount,
				ShopperReference:         "4343553GFGFYFY4654654675765",
				RecurringDetailReference: "8314442372419167",
			})
			require.NotNil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, "422 Unprocessable Entity", err.Error())
			assert.Equal(t, 422, httpRes.StatusCode)
			assert.Equal(t, "422 Unprocessable Entity", httpRes.Status)
			require.NotNil(t, res)
			errBody := map[string]string{}
			json.Unmarshal(err.(common.GenericOpenAPIError).Body(), &errBody)
			assert.Equal(t, "Contract not found", errBody["message"])
		})
	})

	t.Run("ScheduleAccountUpdater", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			res, httpRes, err := client.Recurring.ScheduleAccountUpdaterPost(&recurring.ScheduleAccountUpdaterRequest{
				MerchantAccount: MerchantAccount,
				Reference:       "4343553GFGFYFY4654654675765",
				// RecurringDetailReference: "8314442372419167",
			})
			require.NotNil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, "422 Unprocessable Entity", err.Error())
			assert.Equal(t, 422, httpRes.StatusCode)
			assert.Equal(t, "422 Unprocessable Entity", httpRes.Status)
			require.NotNil(t, res)
			errBody := map[string]string{}
			json.Unmarshal(err.(common.GenericOpenAPIError).Body(), &errBody)
			assert.Equal(t, "Missing card or shopperReference & selectedRecurringDetailReference", errBody["message"])
		})
		t.Run("Create an API request that should fail with card", func(t *testing.T) {
			res, httpRes, err := client.Recurring.ScheduleAccountUpdaterPost(&recurring.ScheduleAccountUpdaterRequest{
				MerchantAccount: MerchantAccount,
				Reference:       "4343553GFGFYFY4654654675765",
				Card: &recurring.Card{
					ExpiryMonth: "03",
					ExpiryYear:  "2030",
					HolderName:  "Adyen test",
					Number:      "4111111111111111",
				},
			})
			require.NotNil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, "422 Unprocessable Entity", err.Error())
			assert.Equal(t, 422, httpRes.StatusCode)
			assert.Equal(t, "422 Unprocessable Entity", httpRes.Status)
			require.NotNil(t, res)
			errBody := map[string]string{}
			json.Unmarshal(err.(common.GenericOpenAPIError).Body(), &errBody)
			assert.Equal(t, "validation 000 No registered account for AccountUpdater", errBody["message"])
		})
	})
}