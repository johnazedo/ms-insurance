package phonelist

type Brand struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Model struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type ModelResponse struct {
	BrandCode string  `json:"brand-code"`
	BrandName string  `json:"brand-name"`
	ModelList []Model `json:"list-cellphone-model"`
}
