package api

const (
	InstTypeSPOT    = "SPOT"
	InstTypeMARGIN  = "MARGIN"
	InstTypeSWAP    = "SWAP"
	InstTypeFUTURES = "FUTURES"
	InstTypeOPTION  = "OPTION"

	TdModeIsolated = "isolated"
	TdModeCross    = "cross"
	TdModeCash     = "cash"

	SideBuy  = "buy"
	SideSell = "sell"

	PosSideLong  = "long"
	PosSideShort = "short"

	OrdTypeMarket          = "market"
	OrdTypeLimit           = "limit"
	OrdTypePostOnly        = "post_only"
	OrdTypeFok             = "fok"
	OrdTypeIoc             = "ioc"
	OrdTypeOptimalLimitIoc = "optimal_limit_ioc"
)
