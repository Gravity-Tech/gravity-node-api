package response

import (
	"net/http"
	"strconv"
)

const (
	DefaultItemsPerPage = 16
)

func HandleParams(req *http.Request) (string, string, string) {
	q, page, items :=
		req.URL.Query().Get("q"),
		req.URL.Query().Get("page"),
		req.URL.Query().Get("items")

	return q, page, items
}

func RevealParams(page, items string) (int, int) {
	// Parse page
	currentPage, castPageErr := strconv.Atoi(page)
	if castPageErr != nil || currentPage < 1 { currentPage = 1 }
	// Parse items
	itemsPerPage, castItemsErr := strconv.Atoi(items)
	if castItemsErr != nil || itemsPerPage < 1 { itemsPerPage = DefaultItemsPerPage }

	return currentPage, itemsPerPage
}

func ComputeSliceRange(listLen, itemsPerPage, pageIndex int) (int, int) {
	firstItemIndex := itemsPerPage * pageIndex

	if listLen > itemsPerPage {
		// check for range existence
		if firstItemIndex < listLen {
			return firstItemIndex, firstItemIndex+itemsPerPage
		}
	}

	projectedMaxIndex := firstItemIndex+itemsPerPage

	if firstItemIndex > listLen { firstItemIndex = listLen }
	if projectedMaxIndex > listLen { projectedMaxIndex = listLen }

	return firstItemIndex, projectedMaxIndex
}
