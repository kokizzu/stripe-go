//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"github.com/stripe/stripe-go/v82/form"
)

// Controls how prorations and invoices for subscriptions are calculated and orchestrated.
type SubscriptionScheduleBillingModeType string

// List of values that SubscriptionScheduleBillingModeType can take
const (
	SubscriptionScheduleBillingModeTypeClassic  SubscriptionScheduleBillingModeType = "classic"
	SubscriptionScheduleBillingModeTypeFlexible SubscriptionScheduleBillingModeType = "flexible"
)

// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
type SubscriptionScheduleDefaultSettingsBillingCycleAnchor string

// List of values that SubscriptionScheduleDefaultSettingsBillingCycleAnchor can take
const (
	SubscriptionScheduleDefaultSettingsBillingCycleAnchorAutomatic  SubscriptionScheduleDefaultSettingsBillingCycleAnchor = "automatic"
	SubscriptionScheduleDefaultSettingsBillingCycleAnchorPhaseStart SubscriptionScheduleDefaultSettingsBillingCycleAnchor = "phase_start"
)

// Type of the account referenced.
type SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerType string

// List of values that SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerType can take
const (
	SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerTypeAccount SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerType = "account"
	SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerTypeSelf    SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerType = "self"
)

// Behavior of the subscription schedule and underlying subscription when it ends. Possible values are `release` or `cancel` with the default being `release`. `release` will end the subscription schedule and keep the underlying subscription running. `cancel` will end the subscription schedule and cancel the underlying subscription.
type SubscriptionScheduleEndBehavior string

// List of values that SubscriptionScheduleEndBehavior can take
const (
	SubscriptionScheduleEndBehaviorCancel  SubscriptionScheduleEndBehavior = "cancel"
	SubscriptionScheduleEndBehaviorNone    SubscriptionScheduleEndBehavior = "none"
	SubscriptionScheduleEndBehaviorRelease SubscriptionScheduleEndBehavior = "release"
	SubscriptionScheduleEndBehaviorRenew   SubscriptionScheduleEndBehavior = "renew"
)

// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
type SubscriptionSchedulePhaseBillingCycleAnchor string

// List of values that SubscriptionSchedulePhaseBillingCycleAnchor can take
const (
	SubscriptionSchedulePhaseBillingCycleAnchorAutomatic  SubscriptionSchedulePhaseBillingCycleAnchor = "automatic"
	SubscriptionSchedulePhaseBillingCycleAnchorPhaseStart SubscriptionSchedulePhaseBillingCycleAnchor = "phase_start"
)

// Type of the account referenced.
type SubscriptionSchedulePhaseInvoiceSettingsIssuerType string

// List of values that SubscriptionSchedulePhaseInvoiceSettingsIssuerType can take
const (
	SubscriptionSchedulePhaseInvoiceSettingsIssuerTypeAccount SubscriptionSchedulePhaseInvoiceSettingsIssuerType = "account"
	SubscriptionSchedulePhaseInvoiceSettingsIssuerTypeSelf    SubscriptionSchedulePhaseInvoiceSettingsIssuerType = "self"
)

// When transitioning phases, controls how prorations are handled (if any). Possible values are `create_prorations`, `none`, and `always_invoice`.
type SubscriptionSchedulePhaseProrationBehavior string

// List of values that SubscriptionSchedulePhaseProrationBehavior can take
const (
	SubscriptionSchedulePhaseProrationBehaviorAlwaysInvoice    SubscriptionSchedulePhaseProrationBehavior = "always_invoice"
	SubscriptionSchedulePhaseProrationBehaviorCreateProrations SubscriptionSchedulePhaseProrationBehavior = "create_prorations"
	SubscriptionSchedulePhaseProrationBehaviorNone             SubscriptionSchedulePhaseProrationBehavior = "none"
)

// The present status of the subscription schedule. Possible values are `not_started`, `active`, `completed`, `released`, and `canceled`. You can read more about the different states in our [behavior guide](https://stripe.com/docs/billing/subscriptions/subscription-schedules).
type SubscriptionScheduleStatus string

// List of values that SubscriptionScheduleStatus can take
const (
	SubscriptionScheduleStatusActive     SubscriptionScheduleStatus = "active"
	SubscriptionScheduleStatusCanceled   SubscriptionScheduleStatus = "canceled"
	SubscriptionScheduleStatusCompleted  SubscriptionScheduleStatus = "completed"
	SubscriptionScheduleStatusNotStarted SubscriptionScheduleStatus = "not_started"
	SubscriptionScheduleStatusReleased   SubscriptionScheduleStatus = "released"
)

// Retrieves the list of your subscription schedules.
type SubscriptionScheduleListParams struct {
	ListParams `form:"*"`
	// Only return subscription schedules that were created canceled the given date interval.
	CanceledAt *int64 `form:"canceled_at"`
	// Only return subscription schedules that were created canceled the given date interval.
	CanceledAtRange *RangeQueryParams `form:"canceled_at"`
	// Only return subscription schedules that completed during the given date interval.
	CompletedAt *int64 `form:"completed_at"`
	// Only return subscription schedules that completed during the given date interval.
	CompletedAtRange *RangeQueryParams `form:"completed_at"`
	// Only return subscription schedules that were created during the given date interval.
	Created *int64 `form:"created"`
	// Only return subscription schedules that were created during the given date interval.
	CreatedRange *RangeQueryParams `form:"created"`
	// Only return subscription schedules for the given customer.
	Customer *string `form:"customer"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Only return subscription schedules that were released during the given date interval.
	ReleasedAt *int64 `form:"released_at"`
	// Only return subscription schedules that were released during the given date interval.
	ReleasedAtRange *RangeQueryParams `form:"released_at"`
	// Only return subscription schedules that have not started yet.
	Scheduled *bool `form:"scheduled"`
}

// AddExpand appends a new field to expand.
func (p *SubscriptionScheduleListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Controls how prorations and invoices for subscriptions are calculated and orchestrated.
type SubscriptionScheduleBillingModeParams struct {
	Type *string `form:"type"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
type SubscriptionScheduleDefaultSettingsBillingThresholdsParams struct {
	// Monetary threshold that triggers the subscription to advance to a new billing period
	AmountGTE *int64 `form:"amount_gte"`
	// Indicates if the `billing_cycle_anchor` should be reset when a threshold is reached. If true, `billing_cycle_anchor` will be updated to the date/time the threshold was last reached; otherwise, the value will remain unchanged.
	ResetBillingCycleAnchor *bool `form:"reset_billing_cycle_anchor"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account"`
	// Type of the account referenced in the request.
	Type *string `form:"type"`
}

// All invoices will be billed using the specified settings.
type SubscriptionScheduleDefaultSettingsInvoiceSettingsParams struct {
	// The account tax IDs associated with the subscription schedule. Will be set on invoices generated by the subscription schedule.
	AccountTaxIDs []*string `form:"account_tax_ids"`
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `collection_method=charge_automatically`.
	DaysUntilDue *int64 `form:"days_until_due"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerParams `form:"issuer"`
}

// Object representing the subscription schedule's default settings.
type SubscriptionScheduleDefaultSettingsParams struct {
	Params `form:"*"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account. The request must be made by a platform account on a connected account in order to set an application fee percentage. For more information, see the application fees [documentation](https://stripe.com/docs/connect/subscriptions#collecting-fees-on-subscriptions).
	ApplicationFeePercent *float64 `form:"application_fee_percent,high_precision"`
	// Default settings for automatic tax computation.
	AutomaticTax *SubscriptionAutomaticTaxParams `form:"automatic_tax"`
	// Can be set to `phase_start` to set the anchor to the start of the phase or `automatic` to automatically change it if needed. Cannot be set to `phase_start` if this phase specifies a trial. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
	BillingCycleAnchor *string `form:"billing_cycle_anchor"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionScheduleDefaultSettingsBillingThresholdsParams `form:"billing_thresholds"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`. Defaults to `charge_automatically` on creation.
	CollectionMethod *string `form:"collection_method"`
	// ID of the default payment method for the subscription schedule. It must belong to the customer associated with the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *string `form:"default_payment_method"`
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.
	Description *string `form:"description"`
	// All invoices will be billed using the specified settings.
	InvoiceSettings *SubscriptionScheduleDefaultSettingsInvoiceSettingsParams `form:"invoice_settings"`
	// The account on behalf of which to charge, for each of the associated subscription's invoices.
	OnBehalfOf *string `form:"on_behalf_of"`
	// The data with which to automatically create a Transfer for each of the associated subscription's invoices.
	TransferData *SubscriptionTransferDataParams `form:"transfer_data"`
}

// The coupons to redeem into discounts for the item.
type SubscriptionSchedulePhaseAddInvoiceItemDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code"`
}

// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase. You may pass up to 20 items.
type SubscriptionSchedulePhaseAddInvoiceItemParams struct {
	// The coupons to redeem into discounts for the item.
	Discounts []*SubscriptionSchedulePhaseAddInvoiceItemDiscountParams `form:"discounts"`
	// The ID of the price object. One of `price` or `price_data` is required.
	Price *string `form:"price"`
	// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline. One of `price` or `price_data` is required.
	PriceData *InvoiceItemPriceDataParams `form:"price_data"`
	// Quantity for this item. Defaults to 1.
	Quantity *int64 `form:"quantity"`
	// The tax rates which apply to the item. When set, the `default_tax_rates` do not apply to this item.
	TaxRates []*string `form:"tax_rates"`
}

// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
type SubscriptionSchedulePhaseAutomaticTaxLiabilityParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account"`
	// Type of the account referenced in the request.
	Type *string `form:"type"`
}

// Automatic tax settings for this phase.
type SubscriptionSchedulePhaseAutomaticTaxParams struct {
	// Enabled automatic tax calculation which will automatically compute tax rates on all invoices generated by the subscription.
	Enabled *bool `form:"enabled"`
	// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
	Liability *SubscriptionSchedulePhaseAutomaticTaxLiabilityParams `form:"liability"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
type SubscriptionSchedulePhaseBillingThresholdsParams struct {
	// Monetary threshold that triggers the subscription to advance to a new billing period
	AmountGTE *int64 `form:"amount_gte"`
	// Indicates if the `billing_cycle_anchor` should be reset when a threshold is reached. If true, `billing_cycle_anchor` will be updated to the date/time the threshold was last reached; otherwise, the value will remain unchanged.
	ResetBillingCycleAnchor *bool `form:"reset_billing_cycle_anchor"`
}

// The coupons to redeem into discounts for the schedule phase. If not specified, inherits the discount from the subscription's customer. Pass an empty string to avoid inheriting any discounts.
type SubscriptionSchedulePhaseDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type SubscriptionSchedulePhaseInvoiceSettingsIssuerParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account"`
	// Type of the account referenced in the request.
	Type *string `form:"type"`
}

// All invoices will be billed using the specified settings.
type SubscriptionSchedulePhaseInvoiceSettingsParams struct {
	// The account tax IDs associated with this phase of the subscription schedule. Will be set on invoices generated by this phase of the subscription schedule.
	AccountTaxIDs []*string `form:"account_tax_ids"`
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.
	DaysUntilDue *int64 `form:"days_until_due"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *SubscriptionSchedulePhaseInvoiceSettingsIssuerParams `form:"issuer"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
type SubscriptionSchedulePhaseItemBillingThresholdsParams struct {
	// Number of units that meets the billing threshold to advance the subscription to a new billing period (e.g., it takes 10 $5 units to meet a $50 [monetary threshold](https://stripe.com/docs/api/subscriptions/update#update_subscription-billing_thresholds-amount_gte))
	UsageGTE *int64 `form:"usage_gte"`
}

// The coupons to redeem into discounts for the subscription item.
type SubscriptionSchedulePhaseItemDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code"`
}

// List of configuration items, each with an attached price, to apply during this phase of the subscription schedule.
type SubscriptionSchedulePhaseItemParams struct {
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionSchedulePhaseItemBillingThresholdsParams `form:"billing_thresholds"`
	// The coupons to redeem into discounts for the subscription item.
	Discounts []*SubscriptionSchedulePhaseItemDiscountParams `form:"discounts"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to a configuration item. Metadata on a configuration item will update the underlying subscription item's `metadata` when the phase is entered, adding new keys and replacing existing keys. Individual keys in the subscription item's `metadata` can be unset by posting an empty value to them in the configuration item's `metadata`. To unset all keys in the subscription item's `metadata`, update the subscription item directly or unset every key individually from the configuration item's `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The plan ID to subscribe to. You may specify the same ID in `plan` and `price`.
	Plan *string `form:"plan"`
	// The ID of the price object.
	Price *string `form:"price"`
	// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline.
	PriceData *SubscriptionItemPriceDataParams `form:"price_data"`
	// Quantity for the given price. Can be set only if the price's `usage_type` is `licensed` and not `metered`.
	Quantity *int64 `form:"quantity"`
	// A list of [Tax Rate](https://stripe.com/docs/api/tax_rates) ids. These Tax Rates will override the [`default_tax_rates`](https://stripe.com/docs/api/subscriptions/create#create_subscription-default_tax_rates) on the Subscription. When updating, pass an empty string to remove previously-defined tax rates.
	TaxRates []*string `form:"tax_rates"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionSchedulePhaseItemParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// List representing phases of the subscription schedule. Each phase can be customized to have different durations, plans, and coupons. If there are multiple phases, the `end_date` of one phase will always equal the `start_date` of the next phase.
type SubscriptionSchedulePhaseParams struct {
	// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase. You may pass up to 20 items.
	AddInvoiceItems []*SubscriptionSchedulePhaseAddInvoiceItemParams `form:"add_invoice_items"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account. The request must be made by a platform account on a connected account in order to set an application fee percentage. For more information, see the application fees [documentation](https://stripe.com/docs/connect/subscriptions#collecting-fees-on-subscriptions).
	ApplicationFeePercent *float64 `form:"application_fee_percent"`
	// Automatic tax settings for this phase.
	AutomaticTax *SubscriptionSchedulePhaseAutomaticTaxParams `form:"automatic_tax"`
	// Can be set to `phase_start` to set the anchor to the start of the phase or `automatic` to automatically change it if needed. Cannot be set to `phase_start` if this phase specifies a trial. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
	BillingCycleAnchor *string `form:"billing_cycle_anchor"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionSchedulePhaseBillingThresholdsParams `form:"billing_thresholds"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`. Defaults to `charge_automatically` on creation.
	CollectionMethod *string `form:"collection_method"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// ID of the default payment method for the subscription schedule. It must belong to the customer associated with the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *string `form:"default_payment_method"`
	// A list of [Tax Rate](https://stripe.com/docs/api/tax_rates) ids. These Tax Rates will set the Subscription's [`default_tax_rates`](https://stripe.com/docs/api/subscriptions/create#create_subscription-default_tax_rates), which means they will be the Invoice's [`default_tax_rates`](https://stripe.com/docs/api/invoices/create#create_invoice-default_tax_rates) for any Invoices issued by the Subscription during this Phase.
	DefaultTaxRates []*string `form:"default_tax_rates"`
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.
	Description *string `form:"description"`
	// The coupons to redeem into discounts for the schedule phase. If not specified, inherits the discount from the subscription's customer. Pass an empty string to avoid inheriting any discounts.
	Discounts []*SubscriptionSchedulePhaseDiscountParams `form:"discounts"`
	// The date at which this phase of the subscription schedule ends. If set, `iterations` must not be set.
	EndDate    *int64 `form:"end_date"`
	EndDateNow *bool  `form:"-"` // See custom AppendTo
	// All invoices will be billed using the specified settings.
	InvoiceSettings *SubscriptionSchedulePhaseInvoiceSettingsParams `form:"invoice_settings"`
	// List of configuration items, each with an attached price, to apply during this phase of the subscription schedule.
	Items []*SubscriptionSchedulePhaseItemParams `form:"items"`
	// Integer representing the multiplier applied to the price interval. For example, `iterations=2` applied to a price with `interval=month` and `interval_count=3` results in a phase of duration `2 * 3 months = 6 months`. If set, `end_date` must not be set.
	Iterations *int64 `form:"iterations"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to a phase. Metadata on a schedule's phase will update the underlying subscription's `metadata` when the phase is entered, adding new keys and replacing existing keys in the subscription's `metadata`. Individual keys in the subscription's `metadata` can be unset by posting an empty value to them in the phase's `metadata`. To unset all keys in the subscription's `metadata`, update the subscription directly or unset every key individually from the phase's `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The account on behalf of which to charge, for each of the associated subscription's invoices.
	OnBehalfOf *string `form:"on_behalf_of"`
	// Controls whether the subscription schedule should create [prorations](https://stripe.com/docs/billing/subscriptions/prorations) when transitioning to this phase if there is a difference in billing configuration. It's different from the request-level [proration_behavior](https://stripe.com/docs/api/subscription_schedules/update#update_subscription_schedule-proration_behavior) parameter which controls what happens if the update request affects the billing configuration (item price, quantity, etc.) of the current phase.
	ProrationBehavior *string `form:"proration_behavior"`
	// The date at which this phase of the subscription schedule starts or `now`. Must be set on the first phase.
	StartDate    *int64 `form:"start_date"`
	StartDateNow *bool  `form:"-"` // See custom AppendTo
	// The data with which to automatically create a Transfer for each of the associated subscription's invoices.
	TransferData *SubscriptionTransferDataParams `form:"transfer_data"`
	// If set to true the entire phase is counted as a trial and the customer will not be charged for any fees.
	Trial *bool `form:"trial"`
	// Sets the phase to trialing from the start date to this date. Must be before the phase end date, can not be combined with `trial`
	TrialEnd    *int64 `form:"trial_end"`
	TrialEndNow *bool  `form:"-"` // See custom AppendTo
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionSchedulePhaseParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// AppendTo implements custom encoding logic for SubscriptionSchedulePhaseParams.
func (p *SubscriptionSchedulePhaseParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.EndDateNow) {
		body.Add(form.FormatKey(append(keyParts, "end_date")), "now")
	}
	if BoolValue(p.TrialEndNow) {
		body.Add(form.FormatKey(append(keyParts, "trial_end")), "now")
	}
	if BoolValue(p.StartDateNow) {
		body.Add(form.FormatKey(append(keyParts, "start_date")), "now")
	}
}

// Creates a new subscription schedule object. Each customer can have up to 500 active or scheduled subscriptions.
type SubscriptionScheduleParams struct {
	Params `form:"*"`
	// Controls how prorations and invoices for subscriptions are calculated and orchestrated.
	BillingMode *SubscriptionScheduleBillingModeParams `form:"billing_mode"`
	// The identifier of the customer to create the subscription schedule for.
	Customer *string `form:"customer"`
	// Object representing the subscription schedule's default settings.
	DefaultSettings *SubscriptionScheduleDefaultSettingsParams `form:"default_settings"`
	// Behavior of the subscription schedule and underlying subscription when it ends. Possible values are `release` or `cancel` with the default being `release`. `release` will end the subscription schedule and keep the underlying subscription running. `cancel` will end the subscription schedule and cancel the underlying subscription.
	EndBehavior *string `form:"end_behavior"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Migrate an existing subscription to be managed by a subscription schedule. If this parameter is set, a subscription schedule will be created using the subscription's item(s), set to auto-renew using the subscription's interval. When using this parameter, other parameters (such as phase values) cannot be set. To create a subscription schedule with other modifications, we recommend making two separate API calls.
	FromSubscription *string `form:"from_subscription"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// List representing phases of the subscription schedule. Each phase can be customized to have different durations, plans, and coupons. If there are multiple phases, the `end_date` of one phase will always equal the `start_date` of the next phase. Note that past phases can be omitted.
	Phases []*SubscriptionSchedulePhaseParams `form:"phases"`
	// If the update changes the billing configuration (item price, quantity, etc.) of the current phase, indicates how prorations from this change should be handled. The default value is `create_prorations`.
	ProrationBehavior *string `form:"proration_behavior"`
	// When the subscription schedule starts. We recommend using `now` so that it starts the subscription immediately. You can also use a Unix timestamp to backdate the subscription so that it starts on a past date, or set a future date for the subscription to start on.
	StartDate    *int64 `form:"start_date"`
	StartDateNow *bool  `form:"-"` // See custom AppendTo
}

// AddExpand appends a new field to expand.
func (p *SubscriptionScheduleParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionScheduleParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// AppendTo implements custom encoding logic for SubscriptionScheduleParams.
func (p *SubscriptionScheduleParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.StartDateNow) {
		body.Add(form.FormatKey(append(keyParts, "start_date")), "now")
	}
}

// Cancels a subscription schedule and its associated subscription immediately (if the subscription schedule has an active subscription). A subscription schedule can only be canceled if its status is not_started or active.
type SubscriptionScheduleCancelParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// If the subscription schedule is `active`, indicates if a final invoice will be generated that contains any un-invoiced metered usage and new/pending proration invoice items. Defaults to `true`.
	InvoiceNow *bool `form:"invoice_now"`
	// If the subscription schedule is `active`, indicates if the cancellation should be prorated. Defaults to `true`.
	Prorate *bool `form:"prorate"`
}

// AddExpand appends a new field to expand.
func (p *SubscriptionScheduleCancelParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Releases the subscription schedule immediately, which will stop scheduling of its phases, but leave any existing subscription in place. A schedule can only be released if its status is not_started or active. If the subscription schedule is currently associated with a subscription, releasing it will remove its subscription property and set the subscription's ID to the released_subscription property.
type SubscriptionScheduleReleaseParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Keep any cancellation on the subscription that the schedule has set
	PreserveCancelDate *bool `form:"preserve_cancel_date"`
}

// AddExpand appends a new field to expand.
func (p *SubscriptionScheduleReleaseParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Controls how prorations and invoices for subscriptions are calculated and orchestrated.
type SubscriptionScheduleCreateBillingModeParams struct {
	Type *string `form:"type"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
type SubscriptionScheduleCreateDefaultSettingsBillingThresholdsParams struct {
	// Monetary threshold that triggers the subscription to advance to a new billing period
	AmountGTE *int64 `form:"amount_gte"`
	// Indicates if the `billing_cycle_anchor` should be reset when a threshold is reached. If true, `billing_cycle_anchor` will be updated to the date/time the threshold was last reached; otherwise, the value will remain unchanged.
	ResetBillingCycleAnchor *bool `form:"reset_billing_cycle_anchor"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type SubscriptionScheduleCreateDefaultSettingsInvoiceSettingsIssuerParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account"`
	// Type of the account referenced in the request.
	Type *string `form:"type"`
}

// All invoices will be billed using the specified settings.
type SubscriptionScheduleCreateDefaultSettingsInvoiceSettingsParams struct {
	// The account tax IDs associated with the subscription schedule. Will be set on invoices generated by the subscription schedule.
	AccountTaxIDs []*string `form:"account_tax_ids"`
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `collection_method=charge_automatically`.
	DaysUntilDue *int64 `form:"days_until_due"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *SubscriptionScheduleCreateDefaultSettingsInvoiceSettingsIssuerParams `form:"issuer"`
}

// Object representing the subscription schedule's default settings.
type SubscriptionScheduleCreateDefaultSettingsParams struct {
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account. The request must be made by a platform account on a connected account in order to set an application fee percentage. For more information, see the application fees [documentation](https://stripe.com/docs/connect/subscriptions#collecting-fees-on-subscriptions).
	ApplicationFeePercent *float64 `form:"application_fee_percent,high_precision"`
	// Default settings for automatic tax computation.
	AutomaticTax *SubscriptionAutomaticTaxParams `form:"automatic_tax"`
	// Can be set to `phase_start` to set the anchor to the start of the phase or `automatic` to automatically change it if needed. Cannot be set to `phase_start` if this phase specifies a trial. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
	BillingCycleAnchor *string `form:"billing_cycle_anchor"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionScheduleCreateDefaultSettingsBillingThresholdsParams `form:"billing_thresholds"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`. Defaults to `charge_automatically` on creation.
	CollectionMethod *string `form:"collection_method"`
	// ID of the default payment method for the subscription schedule. It must belong to the customer associated with the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *string `form:"default_payment_method"`
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.
	Description *string `form:"description"`
	// All invoices will be billed using the specified settings.
	InvoiceSettings *SubscriptionScheduleCreateDefaultSettingsInvoiceSettingsParams `form:"invoice_settings"`
	// The account on behalf of which to charge, for each of the associated subscription's invoices.
	OnBehalfOf *string `form:"on_behalf_of"`
	// The data with which to automatically create a Transfer for each of the associated subscription's invoices.
	TransferData *SubscriptionTransferDataParams `form:"transfer_data"`
}

// The coupons to redeem into discounts for the item.
type SubscriptionScheduleCreatePhaseAddInvoiceItemDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code"`
}

// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase. You may pass up to 20 items.
type SubscriptionScheduleCreatePhaseAddInvoiceItemParams struct {
	// The coupons to redeem into discounts for the item.
	Discounts []*SubscriptionScheduleCreatePhaseAddInvoiceItemDiscountParams `form:"discounts"`
	// The ID of the price object. One of `price` or `price_data` is required.
	Price *string `form:"price"`
	// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline. One of `price` or `price_data` is required.
	PriceData *InvoiceItemPriceDataParams `form:"price_data"`
	// Quantity for this item. Defaults to 1.
	Quantity *int64 `form:"quantity"`
	// The tax rates which apply to the item. When set, the `default_tax_rates` do not apply to this item.
	TaxRates []*string `form:"tax_rates"`
}

// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
type SubscriptionScheduleCreatePhaseAutomaticTaxLiabilityParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account"`
	// Type of the account referenced in the request.
	Type *string `form:"type"`
}

// Automatic tax settings for this phase.
type SubscriptionScheduleCreatePhaseAutomaticTaxParams struct {
	// Enabled automatic tax calculation which will automatically compute tax rates on all invoices generated by the subscription.
	Enabled *bool `form:"enabled"`
	// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
	Liability *SubscriptionScheduleCreatePhaseAutomaticTaxLiabilityParams `form:"liability"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
type SubscriptionScheduleCreatePhaseBillingThresholdsParams struct {
	// Monetary threshold that triggers the subscription to advance to a new billing period
	AmountGTE *int64 `form:"amount_gte"`
	// Indicates if the `billing_cycle_anchor` should be reset when a threshold is reached. If true, `billing_cycle_anchor` will be updated to the date/time the threshold was last reached; otherwise, the value will remain unchanged.
	ResetBillingCycleAnchor *bool `form:"reset_billing_cycle_anchor"`
}

// The coupons to redeem into discounts for the schedule phase. If not specified, inherits the discount from the subscription's customer. Pass an empty string to avoid inheriting any discounts.
type SubscriptionScheduleCreatePhaseDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type SubscriptionScheduleCreatePhaseInvoiceSettingsIssuerParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account"`
	// Type of the account referenced in the request.
	Type *string `form:"type"`
}

// All invoices will be billed using the specified settings.
type SubscriptionScheduleCreatePhaseInvoiceSettingsParams struct {
	// The account tax IDs associated with this phase of the subscription schedule. Will be set on invoices generated by this phase of the subscription schedule.
	AccountTaxIDs []*string `form:"account_tax_ids"`
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.
	DaysUntilDue *int64 `form:"days_until_due"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *SubscriptionScheduleCreatePhaseInvoiceSettingsIssuerParams `form:"issuer"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
type SubscriptionScheduleCreatePhaseItemBillingThresholdsParams struct {
	// Number of units that meets the billing threshold to advance the subscription to a new billing period (e.g., it takes 10 $5 units to meet a $50 [monetary threshold](https://stripe.com/docs/api/subscriptions/update#update_subscription-billing_thresholds-amount_gte))
	UsageGTE *int64 `form:"usage_gte"`
}

// The coupons to redeem into discounts for the subscription item.
type SubscriptionScheduleCreatePhaseItemDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code"`
}

// List of configuration items, each with an attached price, to apply during this phase of the subscription schedule.
type SubscriptionScheduleCreatePhaseItemParams struct {
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionScheduleCreatePhaseItemBillingThresholdsParams `form:"billing_thresholds"`
	// The coupons to redeem into discounts for the subscription item.
	Discounts []*SubscriptionScheduleCreatePhaseItemDiscountParams `form:"discounts"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to a configuration item. Metadata on a configuration item will update the underlying subscription item's `metadata` when the phase is entered, adding new keys and replacing existing keys. Individual keys in the subscription item's `metadata` can be unset by posting an empty value to them in the configuration item's `metadata`. To unset all keys in the subscription item's `metadata`, update the subscription item directly or unset every key individually from the configuration item's `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The plan ID to subscribe to. You may specify the same ID in `plan` and `price`.
	Plan *string `form:"plan"`
	// The ID of the price object.
	Price *string `form:"price"`
	// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline.
	PriceData *SubscriptionItemPriceDataParams `form:"price_data"`
	// Quantity for the given price. Can be set only if the price's `usage_type` is `licensed` and not `metered`.
	Quantity *int64 `form:"quantity"`
	// A list of [Tax Rate](https://stripe.com/docs/api/tax_rates) ids. These Tax Rates will override the [`default_tax_rates`](https://stripe.com/docs/api/subscriptions/create#create_subscription-default_tax_rates) on the Subscription. When updating, pass an empty string to remove previously-defined tax rates.
	TaxRates []*string `form:"tax_rates"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionScheduleCreatePhaseItemParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// List representing phases of the subscription schedule. Each phase can be customized to have different durations, plans, and coupons. If there are multiple phases, the `end_date` of one phase will always equal the `start_date` of the next phase.
type SubscriptionScheduleCreatePhaseParams struct {
	// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase. You may pass up to 20 items.
	AddInvoiceItems []*SubscriptionScheduleCreatePhaseAddInvoiceItemParams `form:"add_invoice_items"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account. The request must be made by a platform account on a connected account in order to set an application fee percentage. For more information, see the application fees [documentation](https://stripe.com/docs/connect/subscriptions#collecting-fees-on-subscriptions).
	ApplicationFeePercent *float64 `form:"application_fee_percent"`
	// Automatic tax settings for this phase.
	AutomaticTax *SubscriptionScheduleCreatePhaseAutomaticTaxParams `form:"automatic_tax"`
	// Can be set to `phase_start` to set the anchor to the start of the phase or `automatic` to automatically change it if needed. Cannot be set to `phase_start` if this phase specifies a trial. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
	BillingCycleAnchor *string `form:"billing_cycle_anchor"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionScheduleCreatePhaseBillingThresholdsParams `form:"billing_thresholds"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`. Defaults to `charge_automatically` on creation.
	CollectionMethod *string `form:"collection_method"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// ID of the default payment method for the subscription schedule. It must belong to the customer associated with the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *string `form:"default_payment_method"`
	// A list of [Tax Rate](https://stripe.com/docs/api/tax_rates) ids. These Tax Rates will set the Subscription's [`default_tax_rates`](https://stripe.com/docs/api/subscriptions/create#create_subscription-default_tax_rates), which means they will be the Invoice's [`default_tax_rates`](https://stripe.com/docs/api/invoices/create#create_invoice-default_tax_rates) for any Invoices issued by the Subscription during this Phase.
	DefaultTaxRates []*string `form:"default_tax_rates"`
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.
	Description *string `form:"description"`
	// The coupons to redeem into discounts for the schedule phase. If not specified, inherits the discount from the subscription's customer. Pass an empty string to avoid inheriting any discounts.
	Discounts []*SubscriptionScheduleCreatePhaseDiscountParams `form:"discounts"`
	// The date at which this phase of the subscription schedule ends. If set, `iterations` must not be set.
	EndDate *int64 `form:"end_date"`
	// All invoices will be billed using the specified settings.
	InvoiceSettings *SubscriptionScheduleCreatePhaseInvoiceSettingsParams `form:"invoice_settings"`
	// List of configuration items, each with an attached price, to apply during this phase of the subscription schedule.
	Items []*SubscriptionScheduleCreatePhaseItemParams `form:"items"`
	// Integer representing the multiplier applied to the price interval. For example, `iterations=2` applied to a price with `interval=month` and `interval_count=3` results in a phase of duration `2 * 3 months = 6 months`. If set, `end_date` must not be set.
	Iterations *int64 `form:"iterations"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to a phase. Metadata on a schedule's phase will update the underlying subscription's `metadata` when the phase is entered, adding new keys and replacing existing keys in the subscription's `metadata`. Individual keys in the subscription's `metadata` can be unset by posting an empty value to them in the phase's `metadata`. To unset all keys in the subscription's `metadata`, update the subscription directly or unset every key individually from the phase's `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The account on behalf of which to charge, for each of the associated subscription's invoices.
	OnBehalfOf *string `form:"on_behalf_of"`
	// Controls whether the subscription schedule should create [prorations](https://stripe.com/docs/billing/subscriptions/prorations) when transitioning to this phase if there is a difference in billing configuration. It's different from the request-level [proration_behavior](https://stripe.com/docs/api/subscription_schedules/update#update_subscription_schedule-proration_behavior) parameter which controls what happens if the update request affects the billing configuration (item price, quantity, etc.) of the current phase.
	ProrationBehavior *string `form:"proration_behavior"`
	// The data with which to automatically create a Transfer for each of the associated subscription's invoices.
	TransferData *SubscriptionTransferDataParams `form:"transfer_data"`
	// If set to true the entire phase is counted as a trial and the customer will not be charged for any fees.
	Trial *bool `form:"trial"`
	// Sets the phase to trialing from the start date to this date. Must be before the phase end date, can not be combined with `trial`
	TrialEnd *int64 `form:"trial_end"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionScheduleCreatePhaseParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Creates a new subscription schedule object. Each customer can have up to 500 active or scheduled subscriptions.
type SubscriptionScheduleCreateParams struct {
	Params `form:"*"`
	// Controls how prorations and invoices for subscriptions are calculated and orchestrated.
	BillingMode *SubscriptionScheduleCreateBillingModeParams `form:"billing_mode"`
	// The identifier of the customer to create the subscription schedule for.
	Customer *string `form:"customer"`
	// Object representing the subscription schedule's default settings.
	DefaultSettings *SubscriptionScheduleCreateDefaultSettingsParams `form:"default_settings"`
	// Behavior of the subscription schedule and underlying subscription when it ends. Possible values are `release` or `cancel` with the default being `release`. `release` will end the subscription schedule and keep the underlying subscription running. `cancel` will end the subscription schedule and cancel the underlying subscription.
	EndBehavior *string `form:"end_behavior"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Migrate an existing subscription to be managed by a subscription schedule. If this parameter is set, a subscription schedule will be created using the subscription's item(s), set to auto-renew using the subscription's interval. When using this parameter, other parameters (such as phase values) cannot be set. To create a subscription schedule with other modifications, we recommend making two separate API calls.
	FromSubscription *string `form:"from_subscription"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// List representing phases of the subscription schedule. Each phase can be customized to have different durations, plans, and coupons. If there are multiple phases, the `end_date` of one phase will always equal the `start_date` of the next phase.
	Phases []*SubscriptionScheduleCreatePhaseParams `form:"phases"`
	// When the subscription schedule starts. We recommend using `now` so that it starts the subscription immediately. You can also use a Unix timestamp to backdate the subscription so that it starts on a past date, or set a future date for the subscription to start on.
	StartDate    *int64 `form:"start_date"`
	StartDateNow *bool  `form:"-"` // See custom AppendTo
}

// AddExpand appends a new field to expand.
func (p *SubscriptionScheduleCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionScheduleCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// AppendTo implements custom encoding logic for SubscriptionScheduleCreateParams.
func (p *SubscriptionScheduleCreateParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.StartDateNow) {
		body.Add(form.FormatKey(append(keyParts, "start_date")), "now")
	}
}

// Retrieves the details of an existing subscription schedule. You only need to supply the unique subscription schedule identifier that was returned upon subscription schedule creation.
type SubscriptionScheduleRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *SubscriptionScheduleRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
type SubscriptionScheduleUpdateDefaultSettingsBillingThresholdsParams struct {
	// Monetary threshold that triggers the subscription to advance to a new billing period
	AmountGTE *int64 `form:"amount_gte"`
	// Indicates if the `billing_cycle_anchor` should be reset when a threshold is reached. If true, `billing_cycle_anchor` will be updated to the date/time the threshold was last reached; otherwise, the value will remain unchanged.
	ResetBillingCycleAnchor *bool `form:"reset_billing_cycle_anchor"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type SubscriptionScheduleUpdateDefaultSettingsInvoiceSettingsIssuerParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account"`
	// Type of the account referenced in the request.
	Type *string `form:"type"`
}

// All invoices will be billed using the specified settings.
type SubscriptionScheduleUpdateDefaultSettingsInvoiceSettingsParams struct {
	// The account tax IDs associated with the subscription schedule. Will be set on invoices generated by the subscription schedule.
	AccountTaxIDs []*string `form:"account_tax_ids"`
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `collection_method=charge_automatically`.
	DaysUntilDue *int64 `form:"days_until_due"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *SubscriptionScheduleUpdateDefaultSettingsInvoiceSettingsIssuerParams `form:"issuer"`
}

// Object representing the subscription schedule's default settings.
type SubscriptionScheduleUpdateDefaultSettingsParams struct {
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account. The request must be made by a platform account on a connected account in order to set an application fee percentage. For more information, see the application fees [documentation](https://stripe.com/docs/connect/subscriptions#collecting-fees-on-subscriptions).
	ApplicationFeePercent *float64 `form:"application_fee_percent,high_precision"`
	// Default settings for automatic tax computation.
	AutomaticTax *SubscriptionAutomaticTaxParams `form:"automatic_tax"`
	// Can be set to `phase_start` to set the anchor to the start of the phase or `automatic` to automatically change it if needed. Cannot be set to `phase_start` if this phase specifies a trial. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
	BillingCycleAnchor *string `form:"billing_cycle_anchor"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionScheduleUpdateDefaultSettingsBillingThresholdsParams `form:"billing_thresholds"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`. Defaults to `charge_automatically` on creation.
	CollectionMethod *string `form:"collection_method"`
	// ID of the default payment method for the subscription schedule. It must belong to the customer associated with the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *string `form:"default_payment_method"`
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.
	Description *string `form:"description"`
	// All invoices will be billed using the specified settings.
	InvoiceSettings *SubscriptionScheduleUpdateDefaultSettingsInvoiceSettingsParams `form:"invoice_settings"`
	// The account on behalf of which to charge, for each of the associated subscription's invoices.
	OnBehalfOf *string `form:"on_behalf_of"`
	// The data with which to automatically create a Transfer for each of the associated subscription's invoices.
	TransferData *SubscriptionTransferDataParams `form:"transfer_data"`
}

// The coupons to redeem into discounts for the item.
type SubscriptionScheduleUpdatePhaseAddInvoiceItemDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code"`
}

// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase. You may pass up to 20 items.
type SubscriptionScheduleUpdatePhaseAddInvoiceItemParams struct {
	// The coupons to redeem into discounts for the item.
	Discounts []*SubscriptionScheduleUpdatePhaseAddInvoiceItemDiscountParams `form:"discounts"`
	// The ID of the price object. One of `price` or `price_data` is required.
	Price *string `form:"price"`
	// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline. One of `price` or `price_data` is required.
	PriceData *InvoiceItemPriceDataParams `form:"price_data"`
	// Quantity for this item. Defaults to 1.
	Quantity *int64 `form:"quantity"`
	// The tax rates which apply to the item. When set, the `default_tax_rates` do not apply to this item.
	TaxRates []*string `form:"tax_rates"`
}

// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
type SubscriptionScheduleUpdatePhaseAutomaticTaxLiabilityParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account"`
	// Type of the account referenced in the request.
	Type *string `form:"type"`
}

// Automatic tax settings for this phase.
type SubscriptionScheduleUpdatePhaseAutomaticTaxParams struct {
	// Enabled automatic tax calculation which will automatically compute tax rates on all invoices generated by the subscription.
	Enabled *bool `form:"enabled"`
	// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
	Liability *SubscriptionScheduleUpdatePhaseAutomaticTaxLiabilityParams `form:"liability"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
type SubscriptionScheduleUpdatePhaseBillingThresholdsParams struct {
	// Monetary threshold that triggers the subscription to advance to a new billing period
	AmountGTE *int64 `form:"amount_gte"`
	// Indicates if the `billing_cycle_anchor` should be reset when a threshold is reached. If true, `billing_cycle_anchor` will be updated to the date/time the threshold was last reached; otherwise, the value will remain unchanged.
	ResetBillingCycleAnchor *bool `form:"reset_billing_cycle_anchor"`
}

// The coupons to redeem into discounts for the schedule phase. If not specified, inherits the discount from the subscription's customer. Pass an empty string to avoid inheriting any discounts.
type SubscriptionScheduleUpdatePhaseDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type SubscriptionScheduleUpdatePhaseInvoiceSettingsIssuerParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account"`
	// Type of the account referenced in the request.
	Type *string `form:"type"`
}

// All invoices will be billed using the specified settings.
type SubscriptionScheduleUpdatePhaseInvoiceSettingsParams struct {
	// The account tax IDs associated with this phase of the subscription schedule. Will be set on invoices generated by this phase of the subscription schedule.
	AccountTaxIDs []*string `form:"account_tax_ids"`
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.
	DaysUntilDue *int64 `form:"days_until_due"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *SubscriptionScheduleUpdatePhaseInvoiceSettingsIssuerParams `form:"issuer"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
type SubscriptionScheduleUpdatePhaseItemBillingThresholdsParams struct {
	// Number of units that meets the billing threshold to advance the subscription to a new billing period (e.g., it takes 10 $5 units to meet a $50 [monetary threshold](https://stripe.com/docs/api/subscriptions/update#update_subscription-billing_thresholds-amount_gte))
	UsageGTE *int64 `form:"usage_gte"`
}

// The coupons to redeem into discounts for the subscription item.
type SubscriptionScheduleUpdatePhaseItemDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code"`
}

// List of configuration items, each with an attached price, to apply during this phase of the subscription schedule.
type SubscriptionScheduleUpdatePhaseItemParams struct {
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionScheduleUpdatePhaseItemBillingThresholdsParams `form:"billing_thresholds"`
	// The coupons to redeem into discounts for the subscription item.
	Discounts []*SubscriptionScheduleUpdatePhaseItemDiscountParams `form:"discounts"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to a configuration item. Metadata on a configuration item will update the underlying subscription item's `metadata` when the phase is entered, adding new keys and replacing existing keys. Individual keys in the subscription item's `metadata` can be unset by posting an empty value to them in the configuration item's `metadata`. To unset all keys in the subscription item's `metadata`, update the subscription item directly or unset every key individually from the configuration item's `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The plan ID to subscribe to. You may specify the same ID in `plan` and `price`.
	Plan *string `form:"plan"`
	// The ID of the price object.
	Price *string `form:"price"`
	// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline.
	PriceData *SubscriptionItemPriceDataParams `form:"price_data"`
	// Quantity for the given price. Can be set only if the price's `usage_type` is `licensed` and not `metered`.
	Quantity *int64 `form:"quantity"`
	// A list of [Tax Rate](https://stripe.com/docs/api/tax_rates) ids. These Tax Rates will override the [`default_tax_rates`](https://stripe.com/docs/api/subscriptions/create#create_subscription-default_tax_rates) on the Subscription. When updating, pass an empty string to remove previously-defined tax rates.
	TaxRates []*string `form:"tax_rates"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionScheduleUpdatePhaseItemParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// List representing phases of the subscription schedule. Each phase can be customized to have different durations, plans, and coupons. If there are multiple phases, the `end_date` of one phase will always equal the `start_date` of the next phase. Note that past phases can be omitted.
type SubscriptionScheduleUpdatePhaseParams struct {
	// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase. You may pass up to 20 items.
	AddInvoiceItems []*SubscriptionScheduleUpdatePhaseAddInvoiceItemParams `form:"add_invoice_items"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account. The request must be made by a platform account on a connected account in order to set an application fee percentage. For more information, see the application fees [documentation](https://stripe.com/docs/connect/subscriptions#collecting-fees-on-subscriptions).
	ApplicationFeePercent *float64 `form:"application_fee_percent"`
	// Automatic tax settings for this phase.
	AutomaticTax *SubscriptionScheduleUpdatePhaseAutomaticTaxParams `form:"automatic_tax"`
	// Can be set to `phase_start` to set the anchor to the start of the phase or `automatic` to automatically change it if needed. Cannot be set to `phase_start` if this phase specifies a trial. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
	BillingCycleAnchor *string `form:"billing_cycle_anchor"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionScheduleUpdatePhaseBillingThresholdsParams `form:"billing_thresholds"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`. Defaults to `charge_automatically` on creation.
	CollectionMethod *string `form:"collection_method"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// ID of the default payment method for the subscription schedule. It must belong to the customer associated with the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *string `form:"default_payment_method"`
	// A list of [Tax Rate](https://stripe.com/docs/api/tax_rates) ids. These Tax Rates will set the Subscription's [`default_tax_rates`](https://stripe.com/docs/api/subscriptions/create#create_subscription-default_tax_rates), which means they will be the Invoice's [`default_tax_rates`](https://stripe.com/docs/api/invoices/create#create_invoice-default_tax_rates) for any Invoices issued by the Subscription during this Phase.
	DefaultTaxRates []*string `form:"default_tax_rates"`
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.
	Description *string `form:"description"`
	// The coupons to redeem into discounts for the schedule phase. If not specified, inherits the discount from the subscription's customer. Pass an empty string to avoid inheriting any discounts.
	Discounts []*SubscriptionScheduleUpdatePhaseDiscountParams `form:"discounts"`
	// The date at which this phase of the subscription schedule ends. If set, `iterations` must not be set.
	EndDate    *int64 `form:"end_date"`
	EndDateNow *bool  `form:"-"` // See custom AppendTo
	// All invoices will be billed using the specified settings.
	InvoiceSettings *SubscriptionScheduleUpdatePhaseInvoiceSettingsParams `form:"invoice_settings"`
	// List of configuration items, each with an attached price, to apply during this phase of the subscription schedule.
	Items []*SubscriptionScheduleUpdatePhaseItemParams `form:"items"`
	// Integer representing the multiplier applied to the price interval. For example, `iterations=2` applied to a price with `interval=month` and `interval_count=3` results in a phase of duration `2 * 3 months = 6 months`. If set, `end_date` must not be set.
	Iterations *int64 `form:"iterations"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to a phase. Metadata on a schedule's phase will update the underlying subscription's `metadata` when the phase is entered, adding new keys and replacing existing keys in the subscription's `metadata`. Individual keys in the subscription's `metadata` can be unset by posting an empty value to them in the phase's `metadata`. To unset all keys in the subscription's `metadata`, update the subscription directly or unset every key individually from the phase's `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The account on behalf of which to charge, for each of the associated subscription's invoices.
	OnBehalfOf *string `form:"on_behalf_of"`
	// Controls whether the subscription schedule should create [prorations](https://stripe.com/docs/billing/subscriptions/prorations) when transitioning to this phase if there is a difference in billing configuration. It's different from the request-level [proration_behavior](https://stripe.com/docs/api/subscription_schedules/update#update_subscription_schedule-proration_behavior) parameter which controls what happens if the update request affects the billing configuration (item price, quantity, etc.) of the current phase.
	ProrationBehavior *string `form:"proration_behavior"`
	// The date at which this phase of the subscription schedule starts or `now`. Must be set on the first phase.
	StartDate    *int64 `form:"start_date"`
	StartDateNow *bool  `form:"-"` // See custom AppendTo
	// The data with which to automatically create a Transfer for each of the associated subscription's invoices.
	TransferData *SubscriptionTransferDataParams `form:"transfer_data"`
	// If set to true the entire phase is counted as a trial and the customer will not be charged for any fees.
	Trial *bool `form:"trial"`
	// Sets the phase to trialing from the start date to this date. Must be before the phase end date, can not be combined with `trial`
	TrialEnd    *int64 `form:"trial_end"`
	TrialEndNow *bool  `form:"-"` // See custom AppendTo
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionScheduleUpdatePhaseParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// AppendTo implements custom encoding logic for SubscriptionScheduleUpdatePhaseParams.
func (p *SubscriptionScheduleUpdatePhaseParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.EndDateNow) {
		body.Add(form.FormatKey(append(keyParts, "end_date")), "now")
	}
	if BoolValue(p.StartDateNow) {
		body.Add(form.FormatKey(append(keyParts, "start_date")), "now")
	}
	if BoolValue(p.TrialEndNow) {
		body.Add(form.FormatKey(append(keyParts, "trial_end")), "now")
	}
}

// Updates an existing subscription schedule.
type SubscriptionScheduleUpdateParams struct {
	Params `form:"*"`
	// Object representing the subscription schedule's default settings.
	DefaultSettings *SubscriptionScheduleUpdateDefaultSettingsParams `form:"default_settings"`
	// Behavior of the subscription schedule and underlying subscription when it ends. Possible values are `release` or `cancel` with the default being `release`. `release` will end the subscription schedule and keep the underlying subscription running. `cancel` will end the subscription schedule and cancel the underlying subscription.
	EndBehavior *string `form:"end_behavior"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// List representing phases of the subscription schedule. Each phase can be customized to have different durations, plans, and coupons. If there are multiple phases, the `end_date` of one phase will always equal the `start_date` of the next phase. Note that past phases can be omitted.
	Phases []*SubscriptionScheduleUpdatePhaseParams `form:"phases"`
	// If the update changes the billing configuration (item price, quantity, etc.) of the current phase, indicates how prorations from this change should be handled. The default value is `create_prorations`.
	ProrationBehavior *string `form:"proration_behavior"`
}

// AddExpand appends a new field to expand.
func (p *SubscriptionScheduleUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionScheduleUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The billing mode of the subscription.
type SubscriptionScheduleBillingMode struct {
	// Controls how prorations and invoices for subscriptions are calculated and orchestrated.
	Type SubscriptionScheduleBillingModeType `json:"type"`
	// Details on when the current billing_mode was adopted.
	UpdatedAt int64 `json:"updated_at"`
}

// Object representing the start and end dates for the current phase of the subscription schedule, if it is `active`.
type SubscriptionScheduleCurrentPhase struct {
	// The end of this phase of the subscription schedule.
	EndDate int64 `json:"end_date"`
	// The start of this phase of the subscription schedule.
	StartDate int64 `json:"start_date"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period
type SubscriptionScheduleDefaultSettingsBillingThresholds struct {
	// Monetary threshold that triggers the subscription to create an invoice
	AmountGTE int64 `json:"amount_gte"`
	// Indicates if the `billing_cycle_anchor` should be reset when a threshold is reached. If true, `billing_cycle_anchor` will be updated to the date/time the threshold was last reached; otherwise, the value will remain unchanged. This value may not be `true` if the subscription contains items with plans that have `aggregate_usage=last_ever`.
	ResetBillingCycleAnchor bool `json:"reset_billing_cycle_anchor"`
}
type SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuer struct {
	// The connected account being referenced when `type` is `account`.
	Account *Account `json:"account"`
	// Type of the account referenced.
	Type SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerType `json:"type"`
}
type SubscriptionScheduleDefaultSettingsInvoiceSettings struct {
	// The account tax IDs associated with the subscription schedule. Will be set on invoices generated by the subscription schedule.
	AccountTaxIDs []*TaxID `json:"account_tax_ids"`
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.
	DaysUntilDue int64                                                     `json:"days_until_due"`
	Issuer       *SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuer `json:"issuer"`
}
type SubscriptionScheduleDefaultSettings struct {
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account during this phase of the schedule.
	ApplicationFeePercent float64                   `json:"application_fee_percent"`
	AutomaticTax          *SubscriptionAutomaticTax `json:"automatic_tax"`
	// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
	BillingCycleAnchor SubscriptionScheduleDefaultSettingsBillingCycleAnchor `json:"billing_cycle_anchor"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period
	BillingThresholds *SubscriptionScheduleDefaultSettingsBillingThresholds `json:"billing_thresholds"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`.
	CollectionMethod *SubscriptionCollectionMethod `json:"collection_method"`
	// ID of the default payment method for the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *PaymentMethod `json:"default_payment_method"`
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.
	Description     string                                              `json:"description"`
	InvoiceSettings *SubscriptionScheduleDefaultSettingsInvoiceSettings `json:"invoice_settings"`
	// The account (if any) the charge was made on behalf of for charges associated with the schedule's subscription. See the Connect documentation for details.
	OnBehalfOf *Account `json:"on_behalf_of"`
	// The account (if any) the associated subscription's payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the subscription's invoices.
	TransferData *SubscriptionTransferData `json:"transfer_data"`
}

// The stackable discounts that will be applied to the item.
type SubscriptionSchedulePhaseAddInvoiceItemDiscount struct {
	// ID of the coupon to create a new discount for.
	Coupon *Coupon `json:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *Discount `json:"discount"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *PromotionCode `json:"promotion_code"`
}

// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase.
type SubscriptionSchedulePhaseAddInvoiceItem struct {
	// The stackable discounts that will be applied to the item.
	Discounts []*SubscriptionSchedulePhaseAddInvoiceItemDiscount `json:"discounts"`
	// ID of the price used to generate the invoice item.
	Price *Price `json:"price"`
	// The quantity of the invoice item.
	Quantity int64 `json:"quantity"`
	// The tax rates which apply to the item. When set, the `default_tax_rates` do not apply to this item.
	TaxRates []*TaxRate `json:"tax_rates"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period
type SubscriptionSchedulePhaseBillingThresholds struct {
	// Monetary threshold that triggers the subscription to create an invoice
	AmountGTE int64 `json:"amount_gte"`
	// Indicates if the `billing_cycle_anchor` should be reset when a threshold is reached. If true, `billing_cycle_anchor` will be updated to the date/time the threshold was last reached; otherwise, the value will remain unchanged. This value may not be `true` if the subscription contains items with plans that have `aggregate_usage=last_ever`.
	ResetBillingCycleAnchor bool `json:"reset_billing_cycle_anchor"`
}

// The stackable discounts that will be applied to the subscription on this phase. Subscription item discounts are applied before subscription discounts.
type SubscriptionSchedulePhaseDiscount struct {
	// ID of the coupon to create a new discount for.
	Coupon *Coupon `json:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *Discount `json:"discount"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *PromotionCode `json:"promotion_code"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type SubscriptionSchedulePhaseInvoiceSettingsIssuer struct {
	// The connected account being referenced when `type` is `account`.
	Account *Account `json:"account"`
	// Type of the account referenced.
	Type SubscriptionSchedulePhaseInvoiceSettingsIssuerType `json:"type"`
}

// The invoice settings applicable during this phase.
type SubscriptionSchedulePhaseInvoiceSettings struct {
	// The account tax IDs associated with this phase of the subscription schedule. Will be set on invoices generated by this phase of the subscription schedule.
	AccountTaxIDs []*TaxID `json:"account_tax_ids"`
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.
	DaysUntilDue int64 `json:"days_until_due"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *SubscriptionSchedulePhaseInvoiceSettingsIssuer `json:"issuer"`
}

// Define thresholds at which an invoice will be sent, and the related subscription advanced to a new billing period
type SubscriptionSchedulePhaseItemBillingThresholds struct {
	// Usage threshold that triggers the subscription to create an invoice
	UsageGTE int64 `json:"usage_gte"`
}

// The discounts applied to the subscription item. Subscription item discounts are applied before subscription discounts. Use `expand[]=discounts` to expand each discount.
type SubscriptionSchedulePhaseItemDiscount struct {
	// ID of the coupon to create a new discount for.
	Coupon *Coupon `json:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *Discount `json:"discount"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *PromotionCode `json:"promotion_code"`
}

// Subscription items to configure the subscription to during this phase of the subscription schedule.
type SubscriptionSchedulePhaseItem struct {
	// Define thresholds at which an invoice will be sent, and the related subscription advanced to a new billing period
	BillingThresholds *SubscriptionSchedulePhaseItemBillingThresholds `json:"billing_thresholds"`
	// The discounts applied to the subscription item. Subscription item discounts are applied before subscription discounts. Use `expand[]=discounts` to expand each discount.
	Discounts []*SubscriptionSchedulePhaseItemDiscount `json:"discounts"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an item. Metadata on this item will update the underlying subscription item's `metadata` when the phase is entered.
	Metadata map[string]string `json:"metadata"`
	// ID of the plan to which the customer should be subscribed.
	Plan *Plan `json:"plan"`
	// ID of the price to which the customer should be subscribed.
	Price *Price `json:"price"`
	// Quantity of the plan to which the customer should be subscribed.
	Quantity int64 `json:"quantity"`
	// The tax rates which apply to this `phase_item`. When set, the `default_tax_rates` on the phase do not apply to this `phase_item`.
	TaxRates []*TaxRate `json:"tax_rates"`
}

// Configuration for the subscription schedule's phases.
type SubscriptionSchedulePhase struct {
	// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase.
	AddInvoiceItems []*SubscriptionSchedulePhaseAddInvoiceItem `json:"add_invoice_items"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account during this phase of the schedule.
	ApplicationFeePercent float64                   `json:"application_fee_percent"`
	AutomaticTax          *SubscriptionAutomaticTax `json:"automatic_tax"`
	// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
	BillingCycleAnchor SubscriptionSchedulePhaseBillingCycleAnchor `json:"billing_cycle_anchor"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period
	BillingThresholds *SubscriptionSchedulePhaseBillingThresholds `json:"billing_thresholds"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`.
	CollectionMethod *SubscriptionCollectionMethod `json:"collection_method"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// ID of the default payment method for the subscription schedule. It must belong to the customer associated with the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *PaymentMethod `json:"default_payment_method"`
	// The default tax rates to apply to the subscription during this phase of the subscription schedule.
	DefaultTaxRates []*TaxRate `json:"default_tax_rates"`
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.
	Description string `json:"description"`
	// The stackable discounts that will be applied to the subscription on this phase. Subscription item discounts are applied before subscription discounts.
	Discounts []*SubscriptionSchedulePhaseDiscount `json:"discounts"`
	// The end of this phase of the subscription schedule.
	EndDate int64 `json:"end_date"`
	// The invoice settings applicable during this phase.
	InvoiceSettings *SubscriptionSchedulePhaseInvoiceSettings `json:"invoice_settings"`
	// Subscription items to configure the subscription to during this phase of the subscription schedule.
	Items []*SubscriptionSchedulePhaseItem `json:"items"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to a phase. Metadata on a schedule's phase will update the underlying subscription's `metadata` when the phase is entered. Updating the underlying subscription's `metadata` directly will not affect the current phase's `metadata`.
	Metadata map[string]string `json:"metadata"`
	// The account (if any) the charge was made on behalf of for charges associated with the schedule's subscription. See the Connect documentation for details.
	OnBehalfOf *Account `json:"on_behalf_of"`
	// When transitioning phases, controls how prorations are handled (if any). Possible values are `create_prorations`, `none`, and `always_invoice`.
	ProrationBehavior SubscriptionSchedulePhaseProrationBehavior `json:"proration_behavior"`
	// The start of this phase of the subscription schedule.
	StartDate int64 `json:"start_date"`
	// The account (if any) the associated subscription's payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the subscription's invoices.
	TransferData *SubscriptionTransferData `json:"transfer_data"`
	// When the trial ends within the phase.
	TrialEnd int64 `json:"trial_end"`
}

// A subscription schedule allows you to create and manage the lifecycle of a subscription by predefining expected changes.
//
// Related guide: [Subscription schedules](https://stripe.com/docs/billing/subscriptions/subscription-schedules)
type SubscriptionSchedule struct {
	APIResource
	// ID of the Connect Application that created the schedule.
	Application *Application `json:"application"`
	// The billing mode of the subscription.
	BillingMode *SubscriptionScheduleBillingMode `json:"billing_mode"`
	// Time at which the subscription schedule was canceled. Measured in seconds since the Unix epoch.
	CanceledAt int64 `json:"canceled_at"`
	// Time at which the subscription schedule was completed. Measured in seconds since the Unix epoch.
	CompletedAt int64 `json:"completed_at"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Object representing the start and end dates for the current phase of the subscription schedule, if it is `active`.
	CurrentPhase *SubscriptionScheduleCurrentPhase `json:"current_phase"`
	// ID of the customer who owns the subscription schedule.
	Customer        *Customer                            `json:"customer"`
	DefaultSettings *SubscriptionScheduleDefaultSettings `json:"default_settings"`
	// Behavior of the subscription schedule and underlying subscription when it ends. Possible values are `release` or `cancel` with the default being `release`. `release` will end the subscription schedule and keep the underlying subscription running. `cancel` will end the subscription schedule and cancel the underlying subscription.
	EndBehavior SubscriptionScheduleEndBehavior `json:"end_behavior"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Configuration for the subscription schedule's phases.
	Phases []*SubscriptionSchedulePhase `json:"phases"`
	// Time at which the subscription schedule was released. Measured in seconds since the Unix epoch.
	ReleasedAt int64 `json:"released_at"`
	// ID of the subscription once managed by the subscription schedule (if it is released).
	ReleasedSubscription *Subscription `json:"released_subscription"`
	// The present status of the subscription schedule. Possible values are `not_started`, `active`, `completed`, `released`, and `canceled`. You can read more about the different states in our [behavior guide](https://stripe.com/docs/billing/subscriptions/subscription-schedules).
	Status SubscriptionScheduleStatus `json:"status"`
	// ID of the subscription managed by the subscription schedule.
	Subscription *Subscription `json:"subscription"`
	// ID of the test clock this subscription schedule belongs to.
	TestClock *TestHelpersTestClock `json:"test_clock"`
}

// SubscriptionScheduleList is a list of SubscriptionSchedules as retrieved from a list endpoint.
type SubscriptionScheduleList struct {
	APIResource
	ListMeta
	Data []*SubscriptionSchedule `json:"data"`
}

// UnmarshalJSON handles deserialization of a SubscriptionSchedule.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *SubscriptionSchedule) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = id
		return nil
	}

	type subscriptionSchedule SubscriptionSchedule
	var v subscriptionSchedule
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = SubscriptionSchedule(v)
	return nil
}
