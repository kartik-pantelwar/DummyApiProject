package external

import (
	"dummyProject/external/core"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct{}


func NewHandler() *ProductHandler {
	return &ProductHandler{}
}

// *fetching all product
func fetchAllProduct() (core.ProductsResponse, error) {
	p := core.ProductsResponse{}
	pResponse := core.ProductsResponse{}
	res, err := http.Get("https://dummyjson.com/products")
	if err != nil {
		fmt.Println("Failed to Hit Products API")
		return p, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to return from Body!")
		return p, err
	}

	err = json.Unmarshal(data, &pResponse)
	if err != nil {
		fmt.Println("Failed to Unmarshal Data")
		return p, err
	}

	p = pResponse
	return p, nil
}

// *fetching single product
func fetchSingleProduct(id string) (core.Product, error) {
	p := core.Product{}
	url := fmt.Sprintf("https://dummyjson.com/products/%s", id)
	fmt.Println("URL=", url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to Hit Products API")
		return p, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to return from Body!")
		return p, err
	}

	err = json.Unmarshal(data, &p)
	if err != nil {
		fmt.Println("Failed to Unmarshal Data")
		return p, err
	}

	return p, nil
}

// *fetching category product
func fetchProductByCategory(cg string) (core.ProductsResponse, error) {
	p := core.ProductsResponse{}
	pResponse := core.ProductsResponse{}
	url := fmt.Sprintf("https://dummyjson.com/products/category/%s", cg)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to Hit Products API")
		return p, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to return from Body!")
		return p, err
	}

	err = json.Unmarshal(data, &pResponse)
	if err != nil {
		fmt.Println("Failed to Unmarshal Data")
		return p, err
	}

	p = pResponse
	return p, nil
}

// *fetching category List
func fetchAllCategories() ([]core.Category, error) {
	var cgList []core.Category
	res, err := http.Get("https://dummyjson.com/products/categories")
	if err != nil {
		fmt.Println("Failed to Hit Products API")
		return cgList, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to return from Body!")
		return cgList, err
	}

	err = json.Unmarshal(data, &cgList)
	if err != nil {
		fmt.Println("Failed to Unmarshal Data")
		return cgList, err
	}
	return cgList, nil
}

func fetchCategoriesList() ([]string, error) {
	var cgList []string
	res, err := http.Get("https://dummyjson.com/products/category-list")
	if err != nil {
		fmt.Println("Failed to Hit Products API")
		return cgList, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to return from Body!")
		return cgList, err
	}

	err = json.Unmarshal(data, &cgList)
	if err != nil {
		fmt.Println("Failed to Unmarshal Data")
		return cgList, err
	}
	return cgList, nil
}


// ^ Methods here
func (p *ProductHandler) SingleProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")
	product, err := fetchSingleProduct(id)
	if err != nil {
		fmt.Println("Failed to Fetch Products!")
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&product)
	if err != nil {
		fmt.Println("Failed to Encode Products List")
		return
	}
}

func (p *ProductHandler) AllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	products, err := fetchAllProduct()
	if err != nil {
		fmt.Println("Failed to Fetch Products!")
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&products)
	if err != nil {
		fmt.Println("Failed to Encode Products List")
		return
	}
}

// TODO : response mei khali aayega to category do not exist Show krwana hai
func (p *ProductHandler) CategoryProduct(w http.ResponseWriter, r *http.Request) {
	cg := chi.URLParam(r, "cg")
	w.Header().Set("Content-Type", "application/json")
	product, err := fetchProductByCategory(cg)
	if err != nil {
		fmt.Println("Failed to Fetch Products!")
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&product)
	if err != nil {
		fmt.Println("Failed to Encode Products List")
		return
	}
}

func (p *ProductHandler) ProductCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	categoryList, err := fetchAllCategories()
	if err != nil {
		fmt.Println("Failed to Fetch Categories!")
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&categoryList)
	if err != nil {
		fmt.Println("Failed to Encode Categories")
		return
	}
}

func (p *ProductHandler) CategoryList(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var categoryList []string
	categoryList,err:=fetchCategoriesList()
	if err != nil {
		fmt.Println("Failed to Fetch Categories List!")
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&categoryList)
	if err != nil {
		fmt.Println("Failed to Encode Category List")
		return
	}
}