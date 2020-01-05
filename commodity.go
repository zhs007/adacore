package adacore

// CommodityShop - commodity shop
type CommodityShop struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

// CommodityItem - commodity item
type CommodityItem struct {
	Title    string        `yaml:"title"`
	CurPrice float32       `yaml:"curprice"`
	Img      string        `yaml:"img"`
	URL      string        `yaml:"url"`
	Shop     CommodityShop `yaml:"shop"`
}

// Commodity - commodity
type Commodity struct {
	ID    string           `yaml:"id"`
	Items []*CommodityItem `yaml:"items"`
}

// LoadImageMap - load ImageMap
func (c *Commodity) LoadImageMap() (*ImageMap, error) {
	if len(c.Items) <= 0 {
		return nil, nil
	}

	im := NewImageMap()

	for _, v := range c.Items {
		err := im.AddImage(v.Img)
		if err != nil {
			return nil, err
		}
	}

	return im, nil
}
