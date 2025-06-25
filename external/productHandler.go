package external

import (
	"bytes"
	"dummyProject/external/core"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

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
	defer res.Body.Close()

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
	defer res.Body.Close()

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
	defer res.Body.Close()

	err = json.Unmarshal(data, &pResponse)
	if err != nil {
		fmt.Println("Failed to Unmarshal Data")
		return p, err
	}

	p = pResponse
	return p, nil
}

// *fetching categories Data
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
	defer res.Body.Close()

	err = json.Unmarshal(data, &cgList)
	if err != nil {
		fmt.Println("Failed to Unmarshal Data")
		return cgList, err
	}
	return cgList, nil
}

// *fetching category List
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
	defer res.Body.Close()

	err = json.Unmarshal(data, &cgList)
	if err != nil {
		fmt.Println("Failed to Unmarshal Data")
		return cgList, err
	}
	return cgList, nil
}

// *fetching Search Product
func fetchSearchProduct(name string) (core.ProductsResponse, error) {
	p := core.ProductsResponse{}
	pResponse := core.ProductsResponse{}
	url := fmt.Sprintf("https://dummyjson.com/products/search?q=%s", name)
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
	defer res.Body.Close()

	err = json.Unmarshal(data, &pResponse)
	if err != nil {
		fmt.Println("Failed to Unmarshal Data")
		return p, err
	}

	p = pResponse
	return p, nil
}

// *deleting product
func makeDeleteProduct(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	return client.Do(req)
}

// *updating product
func updateProduct(url string, data []byte) (*http.Response, error) {
	res, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Unable to make request")
		return nil, err
	}
	res.Header.Set("Content-Type","application/json")
	client := &http.Client{}
	res1, err1 := client.Do(res)
	if err1 != nil {
		fmt.Println("Error Occured during Client do")
	}
	defer res.Body.Close()
	return res1, nil
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
	w.Header().Set("Content-Type", "application/json") //if we remove this line, then content will load in text instead of json
	//so we can not process this text response on frontEnd
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

func (p *ProductHandler) CategoryList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var categoryList []string
	categoryList, err := fetchCategoriesList()
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

func (p *ProductHandler) SearchProduct(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	w.Header().Set("Content-Type", "application/json")
	products, err := fetchSearchProduct(name)
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

// ^ Add Product
func (p *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp, err := http.Post("https://dummyjson.com/products/add", "application/json", r.Body)
	if err != nil {
		fmt.Println("Failed to add Product")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read Response Body")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(respData)
}

// ^ Delete Product
func (p *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	url := fmt.Sprintf("https://dummyjson.com/products/%s", id)
	res, err := makeDeleteProduct(url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Failed to Hit Products API")
		return
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to return from Body!")
		return
	}
	defer r.Body.Close()
	w.Write(data)
}

func (p *ProductHandler) Paging(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	limit, err := strconv.Atoi(r.URL.Query().Get("PerPageRecords"))
	if err != nil {
		fmt.Println("Failed to get Record per Page")
	}
	pageNo, err := strconv.Atoi(r.URL.Query().Get("PageNo"))
	if err != nil {
		fmt.Println("Failed get Page Number")
		return
	}
	skip := (pageNo - 1) * limit
	url := fmt.Sprintf("https://dummyjson.com/products?limit=%d&skip=%d", limit, skip)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to get Response")
		return
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to Read Data from Body")
		return
	}
	defer r.Body.Close()
	// products:= core.ProductsResponse{}
	w.Write(data)
}

//& Query parameters-> PerPageRecords, PageNo

func (p *ProductHandler) Sorting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	field := chi.URLParam(r, "field")
	sort := chi.URLParam(r, "sort")
	url := fmt.Sprintf("https://dummyjson.com/products?sortBy=%s&order=%s", field, sort)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to Get Response from API")
		return
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to read Response Body")
		return
	}
	r.Body.Close()
	w.Write(data)
}

// ^Update Product
func (p *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	url := fmt.Sprintf("https://dummyjson.com/products/%s", id)

	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Failed to read from Request body")
		return
	}
	defer r.Body.Close()

	res2, err := updateProduct(url, data)
	if err != nil {
		fmt.Println("Failed to fetch Data")
		return
	}
	data1, err := io.ReadAll(res2.Body)
	if err != nil {
		fmt.Println("Failed to read from Body")
		return
	}
	w.WriteHeader(res2.StatusCode)
	w.Write(data1)
}
