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
	DoubleChaneRate  = 2 * ChangeRate
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

func checkQualityLimits(quality int) int {
	if quality < MinQuality {
		quality = MinQuality
	}

	if quality > MaxQuality {
		quality = MaxQuality
	}

	return quality
}

func getQualityChangeRate(item *Item) int {
	rate := 0

	isAgedBrie := item.Name == AgedBrie
	isBackstagePass := item.Name == BackstagePass
	if isAgedBrie || isBackstagePass {
		if isAgedBrie {
			rate = getAgedBrieQualityChangeRate(item)
		}

		if isBackstagePass {
			rate = getBackstagePassQualityChangeRate(item)
		}

	} else {
		rate = getRegularQualityChangeRate(item)
	}

	return rate
}

func getAgedBrieQualityChangeRate(item *Item) int {
	return doubleRateIfItemExpired(item)
}

func getBackstagePassQualityChangeRate(item *Item) int { // need to refactor
	rate := 0

	switch {
	case item.SellIn == -1:
		rate = -item.Quality
	case item.SellIn < 0:
		rate = 0
	case 4 < item.SellIn && item.SellIn < 10:
		rate = DoubleChaneRate
	case item.SellIn <= 4:
		rate = TripleChangeRate
	default:
		rate = ChangeRate
	}

	return rate
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
		rate = DoubleChaneRate
	}

	return rate
}

func isItemExpired(item *Item) bool {
	return item.SellIn < 0
}
