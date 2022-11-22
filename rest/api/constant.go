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
	PosSideNet   = "net"

	OrdTypeMarket          = "market"
	OrdTypeLimit           = "limit"
	OrdTypePostOnly        = "post_only"
	OrdTypeFok             = "fok"
	OrdTypeIoc             = "ioc"
	OrdTypeOptimalLimitIoc = "optimal_limit_ioc"

	OrdStateCanceled        = "canceled"
	OrdStateLive            = "live"
	OrdStatePartiallyFilled = "partially_filled"
	OrdStateFilled          = "filled"

	PosModeLongShort = "long_short_mode"
	PosModeNet       = "net_mode"

	MgnModeIsolated = "isolated"
	MgnModeCross    = "cross"

	TypeAdd    = "add"
	TypeReduce = "reduce"

	OrdCategoryTwap               = "twap"
	OrdCategoryAdl                = "adl"
	OrdCategoryFullLiquidation    = "full_liquidation"
	OrdCategoryPartialLiquidation = "partial_liquidation"
	OrdCategoryDelivery           = "delivery"
	OrdCategoryDdh                = "ddh"
)
