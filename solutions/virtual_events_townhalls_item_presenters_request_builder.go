package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// VirtualEventsTownhallsItemPresentersRequestBuilder provides operations to manage the presenters property of the microsoft.graph.virtualEvent entity.
type VirtualEventsTownhallsItemPresentersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEventsTownhallsItemPresentersRequestBuilderGetQueryParameters get the list of all virtualEventPresenter objects associated with a virtual event. Currently the supported virtual event types are:- virtualEventTownhall- virtualEventWebinar
type VirtualEventsTownhallsItemPresentersRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// VirtualEventsTownhallsItemPresentersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsTownhallsItemPresentersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEventsTownhallsItemPresentersRequestBuilderGetQueryParameters
}
// VirtualEventsTownhallsItemPresentersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsTownhallsItemPresentersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByVirtualEventPresenterId provides operations to manage the presenters property of the microsoft.graph.virtualEvent entity.
// returns a *VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilder when successful
func (m *VirtualEventsTownhallsItemPresentersRequestBuilder) ByVirtualEventPresenterId(virtualEventPresenterId string)(*VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if virtualEventPresenterId != "" {
        urlTplParams["virtualEventPresenter%2Did"] = virtualEventPresenterId
    }
    return NewVirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualEventsTownhallsItemPresentersRequestBuilderInternal instantiates a new VirtualEventsTownhallsItemPresentersRequestBuilder and sets the default values.
func NewVirtualEventsTownhallsItemPresentersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsTownhallsItemPresentersRequestBuilder) {
    m := &VirtualEventsTownhallsItemPresentersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/townhalls/{virtualEventTownhall%2Did}/presenters{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualEventsTownhallsItemPresentersRequestBuilder instantiates a new VirtualEventsTownhallsItemPresentersRequestBuilder and sets the default values.
func NewVirtualEventsTownhallsItemPresentersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsTownhallsItemPresentersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEventsTownhallsItemPresentersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *VirtualEventsTownhallsItemPresentersCountRequestBuilder when successful
func (m *VirtualEventsTownhallsItemPresentersRequestBuilder) Count()(*VirtualEventsTownhallsItemPresentersCountRequestBuilder) {
    return NewVirtualEventsTownhallsItemPresentersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the list of all virtualEventPresenter objects associated with a virtual event. Currently the supported virtual event types are:- virtualEventTownhall- virtualEventWebinar
// returns a VirtualEventPresenterCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualevent-list-presenters?view=graph-rest-1.0
func (m *VirtualEventsTownhallsItemPresentersRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEventsTownhallsItemPresentersRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.VirtualEventPresenterCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateVirtualEventPresenterCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.VirtualEventPresenterCollectionResponseable), nil
}
// Post create a new virtualEventPresenter object on a virtual event. Currently, the following types of virtual events are supported: - virtualEventTownhall- virtualEventWebinar
// returns a VirtualEventPresenterable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualevent-post-presenters?view=graph-rest-1.0
func (m *VirtualEventsTownhallsItemPresentersRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.VirtualEventPresenterable, requestConfiguration *VirtualEventsTownhallsItemPresentersRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.VirtualEventPresenterable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateVirtualEventPresenterFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.VirtualEventPresenterable), nil
}
// ToGetRequestInformation get the list of all virtualEventPresenter objects associated with a virtual event. Currently the supported virtual event types are:- virtualEventTownhall- virtualEventWebinar
// returns a *RequestInformation when successful
func (m *VirtualEventsTownhallsItemPresentersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsTownhallsItemPresentersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new virtualEventPresenter object on a virtual event. Currently, the following types of virtual events are supported: - virtualEventTownhall- virtualEventWebinar
// returns a *RequestInformation when successful
func (m *VirtualEventsTownhallsItemPresentersRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.VirtualEventPresenterable, requestConfiguration *VirtualEventsTownhallsItemPresentersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualEventsTownhallsItemPresentersRequestBuilder when successful
func (m *VirtualEventsTownhallsItemPresentersRequestBuilder) WithUrl(rawUrl string)(*VirtualEventsTownhallsItemPresentersRequestBuilder) {
    return NewVirtualEventsTownhallsItemPresentersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
