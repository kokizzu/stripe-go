//
//
// File generated from our OpenAPI spec
//
//

// Package customercashbalancetransaction provides the /v1/customers/{customer}/cash_balance_transactions APIs
package customercashbalancetransaction

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/customers/{customer}/cash_balance_transactions APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a specific cash balance transaction, which updated the customer's [cash balance](https://docs.stripe.com/docs/payments/customer-balance).
func Get(id string, params *stripe.CustomerCashBalanceTransactionParams) (*stripe.CustomerCashBalanceTransaction, error) {
	return getC().Get(id, params)
}

// Retrieves a specific cash balance transaction, which updated the customer's [cash balance](https://docs.stripe.com/docs/payments/customer-balance).
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.CustomerCashBalanceTransactionParams) (*stripe.CustomerCashBalanceTransaction, error) {
	path := stripe.FormatURLPath(
		"/v1/customers/%s/cash_balance_transactions/%s", stripe.StringValue(
			params.Customer), id)
	customercashbalancetransaction := &stripe.CustomerCashBalanceTransaction{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, customercashbalancetransaction)
	return customercashbalancetransaction, err
}

// Returns a list of transactions that modified the customer's [cash balance](https://docs.stripe.com/docs/payments/customer-balance).
func List(params *stripe.CustomerCashBalanceTransactionListParams) *Iter {
	return getC().List(params)
}

// Returns a list of transactions that modified the customer's [cash balance](https://docs.stripe.com/docs/payments/customer-balance).
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.CustomerCashBalanceTransactionListParams) *Iter {
	path := stripe.FormatURLPath(
		"/v1/customers/%s/cash_balance_transactions", stripe.StringValue(
			listParams.Customer))
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.CustomerCashBalanceTransactionList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for customer cash balance transactions.
type Iter struct {
	*stripe.Iter
}

// CustomerCashBalanceTransaction returns the customer cash balance transaction which the iterator is currently pointing to.
func (i *Iter) CustomerCashBalanceTransaction() *stripe.CustomerCashBalanceTransaction {
	return i.Current().(*stripe.CustomerCashBalanceTransaction)
}

// CustomerCashBalanceTransactionList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) CustomerCashBalanceTransactionList() *stripe.CustomerCashBalanceTransactionList {
	return i.List().(*stripe.CustomerCashBalanceTransactionList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
