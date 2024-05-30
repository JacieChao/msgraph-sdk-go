package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappsItemGraphwindowsmobilemsiCategoriesMobileAppCategoryItemRequestBuilder provides operations to manage the categories property of the microsoft.graph.mobileApp entity.
type MobileappsItemGraphwindowsmobilemsiCategoriesMobileAppCategoryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsItemGraphwindowsmobilemsiCategoriesMobileAppCategoryItemRequestBuilderGetQueryParameters the list of categories for this app.
type MobileappsItemGraphwindowsmobilemsiCategoriesMobileAppCategoryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileappsItemGraphwindowsmobilemsiCategoriesMobileAppCategoryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphwindowsmobilemsiCategoriesMobileAppCategoryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappsItemGraphwindowsmobilemsiCategoriesMobileAppCategoryItemRequestBuilderGetQueryParameters
}
// NewMobileappsItemGraphwindowsmobilemsiCategoriesMobileAppCategoryItemRequestBuilderInternal instantiates a new MobileappsItemGraphwindowsmobilemsiCategoriesMobileAppCategoryItemRequestBuilder and sets the default values.
func NewMobileappsItemGraphwindowsmobilemsiCategoriesMobileAppCategoryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphwindowsmobilemsiCategoriesMobileAppCategoryItemRequestBuilder) {
    m := &MobileappsItemGraphwindowsmobilemsiCategoriesMobileAppCategoryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.windowsMobileMSI/categories/{mobileAppCategory%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileappsItemGraphwindowsmobilemsiCategoriesMobileAppCategoryItemRequestBuilder instantiates a new MobileappsItemGraphwindowsmobilemsiCategoriesMobileAppCategoryItemRequestBuilder and sets the default values.
func NewMobileappsItemGraphwindowsmobilemsiCategoriesMobileAppCategoryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphwindowsmobilemsiCategoriesMobileAppCategoryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsItemGraphwindowsmobilemsiCategoriesMobileAppCategoryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the list of categories for this app.
// returns a MobileAppCategoryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphwindowsmobilemsiCategoriesMobileAppCategoryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappsItemGraphwindowsmobilemsiCategoriesMobileAppCategoryItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppCategoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMobileAppCategoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppCategoryable), nil
}
// ToGetRequestInformation the list of categories for this app.
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphwindowsmobilemsiCategoriesMobileAppCategoryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphwindowsmobilemsiCategoriesMobileAppCategoryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MobileappsItemGraphwindowsmobilemsiCategoriesMobileAppCategoryItemRequestBuilder when successful
func (m *MobileappsItemGraphwindowsmobilemsiCategoriesMobileAppCategoryItemRequestBuilder) WithUrl(rawUrl string)(*MobileappsItemGraphwindowsmobilemsiCategoriesMobileAppCategoryItemRequestBuilder) {
    return NewMobileappsItemGraphwindowsmobilemsiCategoriesMobileAppCategoryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
