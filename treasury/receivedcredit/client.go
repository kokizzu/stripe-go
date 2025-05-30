//
//
// File generated from our OpenAPI spec
//
//

// Package receivedcredit provides the /v1/treasury/received_credits APIs
package receivedcredit

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/treasury/received_credits APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of an existing ReceivedCredit by passing the unique ReceivedCredit ID from the ReceivedCredit list.
func Get(id string, params *stripe.TreasuryReceivedCreditParams) (*stripe.TreasuryReceivedCredit, error) {
	return getC().Get(id, params)
}

// Retrieves the details of an existing ReceivedCredit by passing the unique ReceivedCredit ID from the ReceivedCredit list.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.TreasuryReceivedCreditParams) (*stripe.TreasuryReceivedCredit, error) {
	path := stripe.FormatURLPath("/v1/treasury/received_credits/%s", id)
	receivedcredit := &stripe.TreasuryReceivedCredit{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, receivedcredit)
	return receivedcredit, err
}

// Returns a list of ReceivedCredits.
func List(params *stripe.TreasuryReceivedCreditListParams) *Iter {
	return getC().List(params)
}

// Returns a list of ReceivedCredits.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.TreasuryReceivedCreditListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TreasuryReceivedCreditList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/treasury/received_credits", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for treasury received credits.
type Iter struct {
	*stripe.Iter
}

// TreasuryReceivedCredit returns the treasury received credit which the iterator is currently pointing to.
func (i *Iter) TreasuryReceivedCredit() *stripe.TreasuryReceivedCredit {
	return i.Current().(*stripe.TreasuryReceivedCredit)
}

// TreasuryReceivedCreditList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TreasuryReceivedCreditList() *stripe.TreasuryReceivedCreditList {
	return i.List().(*stripe.TreasuryReceivedCreditList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
