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
// PaymentMethodsRequest struct for PaymentMethodsRequest
type PaymentMethodsRequest struct {
	// This field contains additional data, which may be required for a particular payment request.  The `additionalData` object consists of entries, each of which includes the key and value.
	AdditionalData interface{} `json:"additionalData,omitempty"`
	// List of payments methods to be presented to the shopper. To refer to payment methods, use their `paymentMethod.type` from [Payment methods overview](https://docs.adyen.com/payment-methods).  Example: `\"allowedPaymentMethods\":[\"ideal\",\"giropay\"]`
	AllowedPaymentMethods []string `json:"allowedPaymentMethods,omitempty"`
	Amount *Amount `json:"amount,omitempty"`
	// List of payments methods to be hidden from the shopper. To refer to payment methods, use their `paymentMethod.type` from [Payment methods overview](https://docs.adyen.com/payment-methods).  Example: `\"blockedPaymentMethods\":[\"ideal\",\"giropay\"]`
	BlockedPaymentMethods []string `json:"blockedPaymentMethods,omitempty"`
	// The platform where a payment transaction takes place. This field can be used for filtering out payment methods that are only available on specific platforms. Possible values: * iOS * Android * Web
	Channel string `json:"channel,omitempty"`
	// The shopper's country code.
	CountryCode string `json:"countryCode,omitempty"`
	// Choose if a specific transaction should use the Real-time Account Updater, regardless of other settings.
	EnableRealTimeUpdate bool `json:"enableRealTimeUpdate,omitempty"`
	// The merchant account identifier, with which you want to process the transaction.
	MerchantAccount string `json:"merchantAccount"`
	// The combination of a language code and a country code to specify the language to be used in the payment.
	ShopperLocale string `json:"shopperLocale,omitempty"`
	// The shopper's reference to uniquely identify this shopper (e.g. user ID or account ID). > This field is required for recurring payments.
	ShopperReference string `json:"shopperReference,omitempty"`
	// The physical store, for which this payment is processed.
	Store string `json:"store,omitempty"`
	// If set to true, you will only perform the [3D Secure 2 authentication](https://docs.adyen.com/checkout/3d-secure/other-3ds-flows/authentication-only), and not the payment authorisation.
	ThreeDSAuthenticationOnly bool `json:"threeDSAuthenticationOnly,omitempty"`
}
