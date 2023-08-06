package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

const (
	Sulfras       = "Sulfuras, Hand of Ragnaros"
	BackstagePass = "Backstage passes to a TAFKAL80ETC concert"
	AgedBrie      = "Aged Brie"
)

const (
	MinQuality = 0
	MaxQuality = 50
)

const (
	ChangeRate       = 1
	DoubleChangeRate = 2 * ChangeRate
	TripleChangeRate = 3 * ChangeRate
)

func UpdateQuality(items []*Item) {
	for _, item := range items {
		if item.Name != Sulfras {
			setItemSellIn(item)
			setItemQuality(item)
		}
	}
}

func setItemSellIn(item *Item) {
	item.SellIn -= ChangeRate
}

func setItemQuality(item *Item) {
	item.Quality += getQualityChangeRate(item)
	item.Quality = checkQualityLimits(item.Quality)
}

func getQualityChangeRate(item *Item) int {
	rate := 0

	if isSpecialItem(item) {
		rate = getSpecialItemQualityChangeRate(item)
	} else {
		rate = getRegularQualityChangeRate(item)
	}

	return rate
}

func isSpecialItem(item *Item) bool {
	isAgedBrie := item.Name == AgedBrie
	isBackstagePass := item.Name == BackstagePass

	return isAgedBrie || isBackstagePass
}

func getSpecialItemQualityChangeRate(item *Item) int {
	rate := 0
	isAgedBrie := item.Name == AgedBrie
	isBackstagePass := item.Name == BackstagePass

	if isAgedBrie {
		rate = getAgedBrieQualityChangeRate(item)
	}

	if isBackstagePass {
		rate = getBackstagePassQualityChangeRate(item)
	}

	return rate
}

func getAgedBrieQualityChangeRate(item *Item) int {
	return doubleRateIfItemExpired(item)
}

func getRegularQualityChangeRate(item *Item) int {
	rate := 0
	if item.Quality > 0 {
		rate = doubleRateIfItemExpired(item)
	}

	return -rate
}

func doubleRateIfItemExpired(item *Item) int {
	rate := ChangeRate
	if isItemExpired(item) {
		rate = DoubleChangeRate
	}

	return rate
}

func getBackstagePassQualityChangeRate(item *Item) int {
	const (
		firthThreshold  = 10
		secondThreshold = 4
	)

	rate := 0
	concertDay := item.SellIn == -1
	afterConcert := isItemExpired(item)
	applyFirtThresholdRate := secondThreshold < item.SellIn && item.SellIn < firthThreshold
	applySecondThreholdRate := item.SellIn <= secondThreshold

	switch {
	case concertDay:
		rate = -item.Quality
	case afterConcert:
		rate = 0
	case applyFirtThresholdRate:
		rate = DoubleChangeRate
	case applySecondThreholdRate:
		rate = TripleChangeRate
	default:
		rate = ChangeRate
	}

	return rate
}

func isItemExpired(item *Item) bool {
	return item.SellIn < 0
}

func checkQualityLimits(quality int) int {
	if quality < MinQuality {
		quality = MinQuality
	}

	if quality > MaxQuality {
		quality = MaxQuality
	}

	return quality
}
