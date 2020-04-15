/*
 * Adyen Payment API
 *
 * A set of API endpoints that allow you to initiate, settle, and modify payments on the Adyen payments platform. You can use the API to accept card payments (including One-Click and 3D Secure), bank transfers, ewallets, and many other payment methods.  To learn more about the API, visit [Classic integration](https://docs.adyen.com/classic-integration).  ## Authentication To connect to the Payments API, you must use your basic authentication credentials. For this, create your web service user, as described in [How to get the WS user password](https://docs.adyen.com/user-management/how-to-get-the-web-service-ws-user-password). Then use its credentials to authenticate your request, for example:  ``` curl -U \"ws@Company.YourCompany\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Payments API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://pal-test.adyen.com/pal/servlet/Payment/v52/authorise ```
 *
 * API version: 52
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package payments

// AdditionalDataRatepay struct for AdditionalDataRatepay
type AdditionalDataRatepay struct {
	// Amount the customer has to pay each month.
	RatepayInstallmentAmount int32 `json:"ratepay.installmentAmount,omitempty"`
	// Amount of the last installment.
	RatepayLastInstallmentAmount int32 `json:"ratepay.lastInstallmentAmount,omitempty"`
	// Interest rate of this installment. Double
	RatepayInterestRate int32 `json:"ratepay.interestRate,omitempty"`
	// Calendar day of the first payment.
	RatepayPaymentFirstday int32 `json:"ratepay.paymentFirstday,omitempty"`
	// Identification name or number for the invoice, defined by the merchant.
	RatepaydataInvoiceId string `json:"ratepaydata.invoiceId,omitempty"`
	// Invoice date, defined by the merchant. If not included, the invoice date is set to the delivery date.
	RatepaydataInvoiceDate string `json:"ratepaydata.invoiceDate,omitempty"`
	// Date the merchant delivered the goods to the customer.
	RatepaydataDeliveryDate string `json:"ratepaydata.deliveryDate,omitempty"`
	// Date by which the customer must settle the payment.
	RatepaydataDueDate string `json:"ratepaydata.dueDate,omitempty"`
}