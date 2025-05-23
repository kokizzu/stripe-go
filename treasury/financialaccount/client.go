//
//
// File generated from our OpenAPI spec
//
//

// Package financialaccount provides the /v1/treasury/financial_accounts APIs
package financialaccount

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/treasury/financial_accounts APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new FinancialAccount. Each connected account can have up to three FinancialAccounts by default.
func New(params *stripe.TreasuryFinancialAccountParams) (*stripe.TreasuryFinancialAccount, error) {
	return getC().New(params)
}

// Creates a new FinancialAccount. Each connected account can have up to three FinancialAccounts by default.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.TreasuryFinancialAccountParams) (*stripe.TreasuryFinancialAccount, error) {
	financialaccount := &stripe.TreasuryFinancialAccount{}
	err := c.B.Call(
		http.MethodPost, "/v1/treasury/financial_accounts", c.Key, params, financialaccount)
	return financialaccount, err
}

// Retrieves the details of a FinancialAccount.
func Get(id string, params *stripe.TreasuryFinancialAccountParams) (*stripe.TreasuryFinancialAccount, error) {
	return getC().Get(id, params)
}

// Retrieves the details of a FinancialAccount.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.TreasuryFinancialAccountParams) (*stripe.TreasuryFinancialAccount, error) {
	path := stripe.FormatURLPath("/v1/treasury/financial_accounts/%s", id)
	financialaccount := &stripe.TreasuryFinancialAccount{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, financialaccount)
	return financialaccount, err
}

// Updates the details of a FinancialAccount.
func Update(id string, params *stripe.TreasuryFinancialAccountParams) (*stripe.TreasuryFinancialAccount, error) {
	return getC().Update(id, params)
}

// Updates the details of a FinancialAccount.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.TreasuryFinancialAccountParams) (*stripe.TreasuryFinancialAccount, error) {
	path := stripe.FormatURLPath("/v1/treasury/financial_accounts/%s", id)
	financialaccount := &stripe.TreasuryFinancialAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, financialaccount)
	return financialaccount, err
}

// Closes a FinancialAccount. A FinancialAccount can only be closed if it has a zero balance, has no pending InboundTransfers, and has canceled all attached Issuing cards.
func Close(id string, params *stripe.TreasuryFinancialAccountCloseParams) (*stripe.TreasuryFinancialAccount, error) {
	return getC().Close(id, params)
}

// Closes a FinancialAccount. A FinancialAccount can only be closed if it has a zero balance, has no pending InboundTransfers, and has canceled all attached Issuing cards.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Close(id string, params *stripe.TreasuryFinancialAccountCloseParams) (*stripe.TreasuryFinancialAccount, error) {
	path := stripe.FormatURLPath("/v1/treasury/financial_accounts/%s/close", id)
	financialaccount := &stripe.TreasuryFinancialAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, financialaccount)
	return financialaccount, err
}

// Retrieves Features information associated with the FinancialAccount.
func RetrieveFeatures(id string, params *stripe.TreasuryFinancialAccountRetrieveFeaturesParams) (*stripe.TreasuryFinancialAccountFeatures, error) {
	return getC().RetrieveFeatures(id, params)
}

// Retrieves Features information associated with the FinancialAccount.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) RetrieveFeatures(id string, params *stripe.TreasuryFinancialAccountRetrieveFeaturesParams) (*stripe.TreasuryFinancialAccountFeatures, error) {
	path := stripe.FormatURLPath(
		"/v1/treasury/financial_accounts/%s/features", id)
	financialaccountfeatures := &stripe.TreasuryFinancialAccountFeatures{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, financialaccountfeatures)
	return financialaccountfeatures, err
}

// Updates the Features associated with a FinancialAccount.
func UpdateFeatures(id string, params *stripe.TreasuryFinancialAccountUpdateFeaturesParams) (*stripe.TreasuryFinancialAccountFeatures, error) {
	return getC().UpdateFeatures(id, params)
}

// Updates the Features associated with a FinancialAccount.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) UpdateFeatures(id string, params *stripe.TreasuryFinancialAccountUpdateFeaturesParams) (*stripe.TreasuryFinancialAccountFeatures, error) {
	path := stripe.FormatURLPath(
		"/v1/treasury/financial_accounts/%s/features", id)
	financialaccountfeatures := &stripe.TreasuryFinancialAccountFeatures{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, financialaccountfeatures)
	return financialaccountfeatures, err
}

// Returns a list of FinancialAccounts.
func List(params *stripe.TreasuryFinancialAccountListParams) *Iter {
	return getC().List(params)
}

// Returns a list of FinancialAccounts.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.TreasuryFinancialAccountListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TreasuryFinancialAccountList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/treasury/financial_accounts", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for treasury financial accounts.
type Iter struct {
	*stripe.Iter
}

// TreasuryFinancialAccount returns the treasury financial account which the iterator is currently pointing to.
func (i *Iter) TreasuryFinancialAccount() *stripe.TreasuryFinancialAccount {
	return i.Current().(*stripe.TreasuryFinancialAccount)
}

// TreasuryFinancialAccountList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TreasuryFinancialAccountList() *stripe.TreasuryFinancialAccountList {
	return i.List().(*stripe.TreasuryFinancialAccountList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
