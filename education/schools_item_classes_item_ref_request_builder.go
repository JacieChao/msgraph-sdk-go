package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// SchoolsItemClassesItemRefRequestBuilder provides operations to manage the collection of educationRoot entities.
type SchoolsItemClassesItemRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SchoolsItemClassesItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchoolsItemClassesItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSchoolsItemClassesItemRefRequestBuilderInternal instantiates a new SchoolsItemClassesItemRefRequestBuilder and sets the default values.
func NewSchoolsItemClassesItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchoolsItemClassesItemRefRequestBuilder) {
    m := &SchoolsItemClassesItemRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/schools/{educationSchool%2Did}/classes/{educationClass%2Did}/$ref", pathParameters),
    }
    return m
}
// NewSchoolsItemClassesItemRefRequestBuilder instantiates a new SchoolsItemClassesItemRefRequestBuilder and sets the default values.
func NewSchoolsItemClassesItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchoolsItemClassesItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchoolsItemClassesItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete ref of navigation property classes for education
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SchoolsItemClassesItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *SchoolsItemClassesItemRefRequestBuilderDeleteRequestConfiguration)(error) {
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
// ToDeleteRequestInformation delete ref of navigation property classes for education
// returns a *RequestInformation when successful
func (m *SchoolsItemClassesItemRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *SchoolsItemClassesItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SchoolsItemClassesItemRefRequestBuilder when successful
func (m *SchoolsItemClassesItemRefRequestBuilder) WithUrl(rawUrl string)(*SchoolsItemClassesItemRefRequestBuilder) {
    return NewSchoolsItemClassesItemRefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
