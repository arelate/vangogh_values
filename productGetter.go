package vangogh_values

import (
	"fmt"
	"github.com/arelate/gog_types"
	"github.com/arelate/vangogh_products"
)

func (vr *ValueReader) ProductGetter(
	page string,
	pt vangogh_products.ProductType) (productsGetter gog_types.ProductsGetter, err error) {
	switch pt {
	case vangogh_products.StorePage:
		productsGetter, err = vr.StorePage(page)
	case vangogh_products.AccountPage:
		productsGetter, err = vr.AccountStorePage(page)
	case vangogh_products.WishlistPage:
		productsGetter, err = vr.WishlistPage(page)
	default:
		err = fmt.Errorf("splitting page is not supported for type %s", pt)
	}
	return productsGetter, err
}
