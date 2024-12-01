package usermodel

type UserPronounce int8

const (
	He UserPronounce = iota
	She
	They
	Other
	PreferNotToSay
)

const (
	SubscriptionTypeMonthly = "monthly"
	SubscriptionTypeYearly  = "yearly"

	SubscriptionStatusActive        = "active"
	SubscriptionStatusPaymentFailed = "payment_failed"
)

type SubscriptionType string

const (
	Monthly SubscriptionType = "highperformanceplusmonthly"
	Yearly  SubscriptionType = "highperformanceplusanual"
)
