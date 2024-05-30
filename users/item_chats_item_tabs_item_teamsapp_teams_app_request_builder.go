package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemChatsItemTabsItemTeamsappTeamsAppRequestBuilder provides operations to manage the teamsApp property of the microsoft.graph.teamsTab entity.
type ItemChatsItemTabsItemTeamsappTeamsAppRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemChatsItemTabsItemTeamsappTeamsAppRequestBuilderGetQueryParameters the application that is linked to the tab. This can't be changed after tab creation.
type ItemChatsItemTabsItemTeamsappTeamsAppRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemChatsItemTabsItemTeamsappTeamsAppRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChatsItemTabsItemTeamsappTeamsAppRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemChatsItemTabsItemTeamsappTeamsAppRequestBuilderGetQueryParameters
}
// NewItemChatsItemTabsItemTeamsappTeamsAppRequestBuilderInternal instantiates a new ItemChatsItemTabsItemTeamsappTeamsAppRequestBuilder and sets the default values.
func NewItemChatsItemTabsItemTeamsappTeamsAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemTabsItemTeamsappTeamsAppRequestBuilder) {
    m := &ItemChatsItemTabsItemTeamsappTeamsAppRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/chats/{chat%2Did}/tabs/{teamsTab%2Did}/teamsApp{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemChatsItemTabsItemTeamsappTeamsAppRequestBuilder instantiates a new ItemChatsItemTabsItemTeamsappTeamsAppRequestBuilder and sets the default values.
func NewItemChatsItemTabsItemTeamsappTeamsAppRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemTabsItemTeamsappTeamsAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemChatsItemTabsItemTeamsappTeamsAppRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the application that is linked to the tab. This can't be changed after tab creation.
// returns a TeamsAppable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemChatsItemTabsItemTeamsappTeamsAppRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemChatsItemTabsItemTeamsappTeamsAppRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamsAppable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamsAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamsAppable), nil
}
// ToGetRequestInformation the application that is linked to the tab. This can't be changed after tab creation.
// returns a *RequestInformation when successful
func (m *ItemChatsItemTabsItemTeamsappTeamsAppRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemChatsItemTabsItemTeamsappTeamsAppRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemChatsItemTabsItemTeamsappTeamsAppRequestBuilder when successful
func (m *ItemChatsItemTabsItemTeamsappTeamsAppRequestBuilder) WithUrl(rawUrl string)(*ItemChatsItemTabsItemTeamsappTeamsAppRequestBuilder) {
    return NewItemChatsItemTabsItemTeamsappTeamsAppRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
