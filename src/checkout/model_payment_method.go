/*
 * Adyen Checkout API
 *
 * Adyen Checkout API provides a simple and flexible way to initiate and authorise online payments. You can use the same integration for payments made with cards (including 3D Secure), mobile wallets, and local payment methods (for example, iDEAL and Sofort).  This API reference provides information on available endpoints and how to interact with them. To learn more about the API, visit [Checkout documentation](https://docs.adyen.com/checkout).  ## Authentication Each request to the Checkout API must be signed with an API key. For this, obtain an API Key from your Customer Area, as described in [How to get the API key](https://docs.adyen.com/user-management/how-to-get-the-api-key). Then set this key to the `X-API-Key` header value, for example:  ``` curl -H \"Content-Type: application/json\" \\ -H \"X-API-Key: Your_Checkout_API_key\" \\ ... ``` Note that when going live, you need to generate a new API Key to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Checkout API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://checkout-test.adyen.com/v52/payments ```
 *
 * API version: 52
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package checkout
// PaymentMethod struct for PaymentMethod
type PaymentMethod struct {
	// List of possible brands. For example: visa, mc.
	Brands []string `json:"brands,omitempty"`
	// The configuration of the payment method.
	Configuration *map[string]interface{} `json:"configuration,omitempty"`
	// All input details to be provided to complete the payment with this payment method.
	Details *[]InputDetail `json:"details,omitempty"`
	Group *PaymentMethodGroup `json:"group,omitempty"`
	// All input details to be provided to complete the payment with this payment method.
	InputDetails *[]InputDetail `json:"inputDetails,omitempty"`
	// The displayable name of this payment method.
	Name string `json:"name,omitempty"`
	// Echo data required to send in next calls.
	PaymentMethodData string `json:"paymentMethodData,omitempty"`
	// Indicates whether this payment method supports tokenization or not.
	SupportsRecurring bool `json:"supportsRecurring,omitempty"`
	// The unique payment method code.
	Type string `json:"type,omitempty"`
}
