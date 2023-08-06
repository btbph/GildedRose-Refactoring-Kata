package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
	"github.com/stretchr/testify/assert"
)

func TestRegularItem(t *testing.T) {
	items := initItems(getRegularItem())
	const days = 2

	runShop(items, days)

	item := items[0]
	assert.Equal(t, "Some item", item.Name)
	assert.Equal(t, -1, item.SellIn)
	assert.Equal(t, 7, item.Quality)
}

func TestNegativeQuality(t *testing.T) {
	items := initItems(getRegularItem())
	const days = 11

	runShop(items, days)

	item := items[0]
	assert.Equal(t, "Some item", item.Name)
	assert.Equal(t, -10, item.SellIn)
	assert.Equal(t, 0, item.Quality)
}

func TestAgedBrie(t *testing.T) {
	items := initItems(getAgedBrie())
	const days = 3

	runShop(items, days)

	item := items[0]
	assert.Equal(t, "Aged Brie", item.Name)
	assert.Equal(t, -2, item.SellIn)
	assert.Equal(t, 15, item.Quality)
}

func TestAgedBrie_maxQuality(t *testing.T) {
	agedBrie := getAgedBrie()
	agedBrie.Quality = gildedrose.MaxQuality
	items := initItems(agedBrie)
	const days = 3

	runShop(items, days)

	item := items[0]
	assert.Equal(t, "Aged Brie", item.Name)
	assert.Equal(t, -2, item.SellIn)
	assert.Equal(t, gildedrose.MaxQuality, item.Quality)
}

func TestSulfras(t *testing.T) {
	items := initItems(getSulfras())
	const days = 10

	runShop(items, days)

	item := items[0]
	assert.Equal(t, "Sulfuras, Hand of Ragnaros", item.Name)
	assert.Equal(t, 1, item.SellIn)
	assert.Equal(t, 10, item.Quality)
}

func TestBackStagePass_before10Days(t *testing.T) {
	pass := getBackstagePass()
	pass.SellIn = 12
	items := initItems(pass)
	const days = 1

	runShop(items, days)

	item := items[0]
	assert.Equal(t, "Backstage passes to a TAFKAL80ETC concert", item.Name)
	assert.Equal(t, 11, item.SellIn)
	assert.Equal(t, 11, item.Quality)
}

func TestBackStagePass_after10Before5Days(t *testing.T) {
	pass := getBackstagePass()
	pass.SellIn = 9
	items := initItems(pass)
	const days = 1

	runShop(items, days)

	item := items[0]
	assert.Equal(t, "Backstage passes to a TAFKAL80ETC concert", item.Name)
	assert.Equal(t, 8, item.SellIn)
	assert.Equal(t, 12, item.Quality)
}

func TestBackStagePass_after5Days(t *testing.T) {
	pass := getBackstagePass()
	pass.SellIn = 5
	items := initItems(pass)
	const days = 1

	runShop(items, days)

	item := items[0]
	assert.Equal(t, "Backstage passes to a TAFKAL80ETC concert", item.Name)
	assert.Equal(t, 4, item.SellIn)
	assert.Equal(t, 13, item.Quality)
}

func TestBackStagePass_afterSellIn(t *testing.T) {
	pass := getBackstagePass()
	items := initItems(pass)
	const days = 2

	runShop(items, days)

	item := items[0]
	assert.Equal(t, "Backstage passes to a TAFKAL80ETC concert", item.Name)
	assert.Equal(t, -1, item.SellIn)
	assert.Equal(t, 0, item.Quality)
}

func initItems(items ...*gildedrose.Item) []*gildedrose.Item {
	return items
}

func getRegularItem() *gildedrose.Item {
	return &gildedrose.Item{
		Name:    "Some item",
		SellIn:  1,
		Quality: 10,
	}
}

func runShop(items []*gildedrose.Item, days int) {
	for i := 0; i < days; i++ {
		gildedrose.UpdateQuality(items)
	}
}

func getAgedBrie() *gildedrose.Item {
	return &gildedrose.Item{
		Name:    "Aged Brie",
		SellIn:  1,
		Quality: 10,
	}
}

func getSulfras() *gildedrose.Item {
	return &gildedrose.Item{
		Name:    "Sulfuras, Hand of Ragnaros",
		SellIn:  1,
		Quality: 10,
	}
}

func getBackstagePass() *gildedrose.Item {
	return &gildedrose.Item{
		Name:    "Backstage passes to a TAFKAL80ETC concert",
		SellIn:  1,
		Quality: 10,
	}
}
