package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// Getoffice365groupsactivitygroupcountswithperiodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder provides operations to call the getOffice365GroupsActivityGroupCounts method.
type Getoffice365groupsactivitygroupcountswithperiodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Getoffice365groupsactivitygroupcountswithperiodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Getoffice365groupsactivitygroupcountswithperiodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetoffice365groupsactivitygroupcountswithperiodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderInternal instantiates a new Getoffice365groupsactivitygroupcountswithperiodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder and sets the default values.
func NewGetoffice365groupsactivitygroupcountswithperiodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*Getoffice365groupsactivitygroupcountswithperiodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) {
    m := &Getoffice365groupsactivitygroupcountswithperiodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getOffice365GroupsActivityGroupCounts(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetoffice365groupsactivitygroupcountswithperiodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder instantiates a new Getoffice365groupsactivitygroupcountswithperiodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder and sets the default values.
func NewGetoffice365groupsactivitygroupcountswithperiodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Getoffice365groupsactivitygroupcountswithperiodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetoffice365groupsactivitygroupcountswithperiodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get the daily total number of groups and how many of them were active based on email conversations, Yammer posts, and SharePoint file activities.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getoffice365groupsactivitygroupcounts?view=graph-rest-1.0
func (m *Getoffice365groupsactivitygroupcountswithperiodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *Getoffice365groupsactivitygroupcountswithperiodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation get the daily total number of groups and how many of them were active based on email conversations, Yammer posts, and SharePoint file activities.
// returns a *RequestInformation when successful
func (m *Getoffice365groupsactivitygroupcountswithperiodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Getoffice365groupsactivitygroupcountswithperiodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *Getoffice365groupsactivitygroupcountswithperiodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder when successful
func (m *Getoffice365groupsactivitygroupcountswithperiodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) WithUrl(rawUrl string)(*Getoffice365groupsactivitygroupcountswithperiodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) {
    return NewGetoffice365groupsactivitygroupcountswithperiodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
