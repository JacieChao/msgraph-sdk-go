package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilder provides operations to manage the swapShiftsChangeRequests property of the microsoft.graph.schedule entity.
type ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilderGetQueryParameters the swap requests for shifts in the schedule.
type ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilderGetQueryParameters
}
// ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilderInternal instantiates a new ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilder and sets the default values.
func NewItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilder) {
    m := &ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/team/schedule/swapShiftsChangeRequests/{swapShiftsChangeRequest%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilder instantiates a new ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilder and sets the default values.
func NewItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property swapShiftsChangeRequests for groups
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the swap requests for shifts in the schedule.
// returns a SwapShiftsChangeRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SwapShiftsChangeRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSwapShiftsChangeRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SwapShiftsChangeRequestable), nil
}
// Patch update the navigation property swapShiftsChangeRequests in groups
// returns a SwapShiftsChangeRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SwapShiftsChangeRequestable, requestConfiguration *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SwapShiftsChangeRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSwapShiftsChangeRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SwapShiftsChangeRequestable), nil
}
// ToDeleteRequestInformation delete navigation property swapShiftsChangeRequests for groups
// returns a *RequestInformation when successful
func (m *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the swap requests for shifts in the schedule.
// returns a *RequestInformation when successful
func (m *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property swapShiftsChangeRequests in groups
// returns a *RequestInformation when successful
func (m *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SwapShiftsChangeRequestable, requestConfiguration *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilder when successful
func (m *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilder) WithUrl(rawUrl string)(*ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilder) {
    return NewItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
