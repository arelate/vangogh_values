package vangogh_values

import (
	"encoding/json"
	"fmt"
	"github.com/arelate/gog_media"
	"github.com/arelate/gog_types"
	"github.com/arelate/vangogh_products"
	"github.com/arelate/vangogh_urls"
	"github.com/boggydigital/kvas"
)

type ValueReader struct {
	productType vangogh_products.ProductType
	mediaType   gog_media.Media
	valueSet    *kvas.ValueSet
}

func NewReader(pt vangogh_products.ProductType, mt gog_media.Media) (*ValueReader, error) {
	dst, err := vangogh_urls.LocalProductsDir(pt, mt)
	if err != nil {
		return nil, err
	}

	vs, err := kvas.NewJsonLocal(dst)
	if err != nil {
		return nil, err
	}

	vr := &ValueReader{
		productType: pt,
		mediaType:   mt,
		valueSet:    vs,
	}

	return vr, nil
}

func (vr *ValueReader) readValue(id string, val interface{}) error {
	spReadCloser, err := vr.valueSet.Get(id)
	if err != nil {
		return err
	}

	if spReadCloser == nil {
		return nil
	}

	defer spReadCloser.Close()

	if err := json.NewDecoder(spReadCloser).Decode(val); err != nil {
		return err
	}

	return nil
}

func (vr *ValueReader) All() []string {
	return vr.valueSet.All()
}

func (vr *ValueReader) Contains(id string) bool {
	return vr.valueSet.Contains(id)
}

func (vr *ValueReader) CreatedAfter(timestamp int64) []string {
	return vr.valueSet.CreatedAfter(timestamp)
}

func (vr *ValueReader) ModifiedAfter(timestamp int64) []string {
	return vr.valueSet.ModifiedAfter(timestamp)
}

func (vr *ValueReader) StoreProduct(id string) (storeProduct *gog_types.StoreProduct, err error) {
	err = vr.readValue(id, &storeProduct)
	return storeProduct, err
}

func (vr *ValueReader) AccountProduct(id string) (accountProduct *gog_types.AccountProduct, err error) {
	err = vr.readValue(id, &accountProduct)
	return accountProduct, err
}

func (vr *ValueReader) WishlistProduct(id string) (wishlistProduct *gog_types.StoreProduct, err error) {
	err = vr.readValue(id, &wishlistProduct)
	return wishlistProduct, err
}

func (vr *ValueReader) Details(id string) (details *gog_types.Details, err error) {
	err = vr.readValue(id, &details)
	return details, err
}

func (vr *ValueReader) ApiProductV1(id string) (apiProductV1 *gog_types.ApiProductV1, err error) {
	err = vr.readValue(id, &apiProductV1)
	return apiProductV1, err
}

func (vr *ValueReader) ApiProductV2(id string) (apiProductV2 *gog_types.ApiProductV2, err error) {
	err = vr.readValue(id, &apiProductV2)
	return apiProductV2, err
}

func (vr *ValueReader) StorePage(page string) (storePage *gog_types.StorePage, err error) {
	err = vr.readValue(page, &storePage)
	return storePage, err
}

func (vr *ValueReader) AccountStorePage(page string) (accountStorePage *gog_types.AccountPage, err error) {
	err = vr.readValue(page, &accountStorePage)
	return accountStorePage, err
}

func (vr *ValueReader) WishlistPage(page string) (wishlistPage *gog_types.WishlistPage, err error) {
	err = vr.readValue(page, &wishlistPage)
	return wishlistPage, err
}

func (vr *ValueReader) ProductType(key string) (interface{}, error) {
	switch vr.productType {
	case vangogh_products.StoreProducts:
		return vr.StoreProduct(key)
	case vangogh_products.AccountProducts:
		return vr.AccountProduct(key)
	case vangogh_products.WishlistProducts:
		return vr.WishlistProduct(key)
	case vangogh_products.Details:
		return vr.Details(key)
	case vangogh_products.ApiProductsV1:
		return vr.ApiProductV1(key)
	case vangogh_products.ApiProductsV2:
		return vr.ApiProductV2(key)
	case vangogh_products.StorePage:
		return vr.StorePage(key)
	case vangogh_products.AccountPage:
		return vr.AccountStorePage(key)
	case vangogh_products.WishlistPage:
		return vr.WishlistPage(key)
	default:
		return nil, fmt.Errorf("vangogh_values: cannot create %s value", vr.productType)
	}
}

func (vr *ValueReader) ProductGetter(page string) (productsGetter gog_types.ProductsGetter, err error) {
	switch vr.productType {
	case vangogh_products.StorePage:
		productsGetter, err = vr.StorePage(page)
	case vangogh_products.AccountPage:
		productsGetter, err = vr.AccountStorePage(page)
	case vangogh_products.WishlistPage:
		productsGetter, err = vr.WishlistPage(page)
	default:
		err = fmt.Errorf("%s doesn't implement ProductGetter interface", vr.productType)
	}
	return productsGetter, err
}