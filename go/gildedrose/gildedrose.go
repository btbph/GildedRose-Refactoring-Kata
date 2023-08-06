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
	MaxQuality = 50
)

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		if items[i].Name != AgedBrie && items[i].Name != BackstagePass {
			if items[i].Quality > 0 {
				if items[i].Name != Sulfras {
					items[i].Quality = items[i].Quality - 1
				}
			}
		} else {
			if items[i].Quality < MaxQuality {
				items[i].Quality = items[i].Quality + 1
				if items[i].Name == BackstagePass {
					if items[i].SellIn < 11 {
						if items[i].Quality < MaxQuality {
							items[i].Quality = items[i].Quality + 1
						}
					}
					if items[i].SellIn < 6 {
						if items[i].Quality < MaxQuality {
							items[i].Quality = items[i].Quality + 1
						}
					}
				}
			}
		}

		if items[i].Name != Sulfras {
			items[i].SellIn = items[i].SellIn - 1
		}

		if items[i].SellIn < 0 {
			if items[i].Name != AgedBrie {
				if items[i].Name != BackstagePass {
					if items[i].Quality > 0 {
						if items[i].Name != Sulfras {
							items[i].Quality = items[i].Quality - 1
						}
					}
				} else {
					items[i].Quality = items[i].Quality - items[i].Quality
				}
			} else {
				if items[i].Quality < MaxQuality {
					items[i].Quality = items[i].Quality + 1
				}
			}
		}
	}
}
