package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// OnenoteNotebooksItemSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilder provides operations to manage the sectionGroups property of the microsoft.graph.sectionGroup entity.
type OnenoteNotebooksItemSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OnenoteNotebooksItemSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilderGetQueryParameters the section groups in the section. Read-only. Nullable.
type OnenoteNotebooksItemSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OnenoteNotebooksItemSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnenoteNotebooksItemSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnenoteNotebooksItemSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilderGetQueryParameters
}
// NewOnenoteNotebooksItemSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilderInternal instantiates a new SectionGroupItemRequestBuilder and sets the default values.
func NewOnenoteNotebooksItemSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnenoteNotebooksItemSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilder) {
    m := &OnenoteNotebooksItemSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/onenote/notebooks/{notebook%2Did}/sectionGroups/{sectionGroup%2Did}/sectionGroups/{sectionGroup%2Did1}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnenoteNotebooksItemSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilder instantiates a new SectionGroupItemRequestBuilder and sets the default values.
func NewOnenoteNotebooksItemSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnenoteNotebooksItemSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenoteNotebooksItemSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the section groups in the section. Read-only. Nullable.
func (m *OnenoteNotebooksItemSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *OnenoteNotebooksItemSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get the section groups in the section. Read-only. Nullable.
func (m *OnenoteNotebooksItemSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilder) Get(ctx context.Context, requestConfiguration *OnenoteNotebooksItemSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSectionGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable), nil
}
