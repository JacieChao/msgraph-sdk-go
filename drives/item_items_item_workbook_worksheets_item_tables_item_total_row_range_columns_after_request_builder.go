package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemTablesItemTotalRowRangeColumnsAfterRequestBuilder provides operations to call the columnsAfter method.
type ItemItemsItemWorkbookWorksheetsItemTablesItemTotalRowRangeColumnsAfterRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookWorksheetsItemTablesItemTotalRowRangeColumnsAfterRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemTablesItemTotalRowRangeColumnsAfterRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookWorksheetsItemTablesItemTotalRowRangeColumnsAfterRequestBuilderInternal instantiates a new ItemItemsItemWorkbookWorksheetsItemTablesItemTotalRowRangeColumnsAfterRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesItemTotalRowRangeColumnsAfterRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemTablesItemTotalRowRangeColumnsAfterRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemTablesItemTotalRowRangeColumnsAfterRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/tables/{workbookTable%2Did}/totalRowRange()/columnsAfter()", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemTablesItemTotalRowRangeColumnsAfterRequestBuilder instantiates a new ItemItemsItemWorkbookWorksheetsItemTablesItemTotalRowRangeColumnsAfterRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesItemTotalRowRangeColumnsAfterRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemTablesItemTotalRowRangeColumnsAfterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemTotalRowRangeColumnsAfterRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function columnsAfter
// returns a WorkbookRangeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemTotalRowRangeColumnsAfterRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemTotalRowRangeColumnsAfterRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookRangeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookRangeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookRangeable), nil
}
// ToGetRequestInformation invoke function columnsAfter
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemTotalRowRangeColumnsAfterRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemTotalRowRangeColumnsAfterRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookWorksheetsItemTablesItemTotalRowRangeColumnsAfterRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemTotalRowRangeColumnsAfterRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookWorksheetsItemTablesItemTotalRowRangeColumnsAfterRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemTotalRowRangeColumnsAfterRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
