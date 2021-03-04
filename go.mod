module github.com/arelate/vangogh_values

go 1.16

require (
	github.com/arelate/gog_types v0.1.8-alpha
	github.com/arelate/vangogh_types v0.1.3-alpha
	github.com/arelate/vangogh_urls v0.1.0-alpha
	github.com/boggydigital/kvas v0.1.9-alpha
)

replace (
	github.com/arelate/gog_types => ../gog_types
	github.com/arelate/vangogh_types => ../vangogh_types
	github.com/arelate/vangogh_urls => ../vangogh_urls
	github.com/boggydigital/kvas => ../kvas
)
