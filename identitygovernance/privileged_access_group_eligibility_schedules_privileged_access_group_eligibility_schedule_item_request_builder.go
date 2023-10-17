package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder provides operations to manage the eligibilitySchedules property of the microsoft.graph.privilegedAccessGroup entity.
type PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderGetQueryParameters read the properties and relationships of a privilegedAccessGroupEligibilitySchedule object.
type PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderGetQueryParameters
}
// PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderInternal instantiates a new PrivilegedAccessGroupEligibilityScheduleItemRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) {
    m := &PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/eligibilitySchedules/{privilegedAccessGroupEligibilitySchedule%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewPrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder instantiates a new PrivilegedAccessGroupEligibilityScheduleItemRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property eligibilitySchedules for identityGovernance
func (m *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read the properties and relationships of a privilegedAccessGroupEligibilitySchedule object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/privilegedaccessgroupeligibilityschedule-get?view=graph-rest-1.0
func (m *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrivilegedAccessGroupEligibilityScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleable), nil
}
// Group provides operations to manage the group property of the microsoft.graph.privilegedAccessGroupEligibilitySchedule entity.
func (m *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) Group()(*PrivilegedAccessGroupEligibilitySchedulesItemGroupRequestBuilder) {
    return NewPrivilegedAccessGroupEligibilitySchedulesItemGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property eligibilitySchedules in identityGovernance
func (m *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleable, requestConfiguration *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrivilegedAccessGroupEligibilityScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleable), nil
}
// Principal provides operations to manage the principal property of the microsoft.graph.privilegedAccessGroupEligibilitySchedule entity.
func (m *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) Principal()(*PrivilegedAccessGroupEligibilitySchedulesItemPrincipalRequestBuilder) {
    return NewPrivilegedAccessGroupEligibilitySchedulesItemPrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property eligibilitySchedules for identityGovernance
func (m *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a privilegedAccessGroupEligibilitySchedule object.
func (m *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
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
// ToPatchRequestInformation update the navigation property eligibilitySchedules in identityGovernance
func (m *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleable, requestConfiguration *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) WithUrl(rawUrl string)(*PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) {
    return NewPrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
