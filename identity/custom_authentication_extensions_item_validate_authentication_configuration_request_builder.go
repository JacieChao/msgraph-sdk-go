package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CustomAuthenticationExtensionsItemValidateAuthenticationConfigurationRequestBuilder provides operations to call the validateAuthenticationConfiguration method.
type CustomAuthenticationExtensionsItemValidateAuthenticationConfigurationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CustomAuthenticationExtensionsItemValidateAuthenticationConfigurationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustomAuthenticationExtensionsItemValidateAuthenticationConfigurationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCustomAuthenticationExtensionsItemValidateAuthenticationConfigurationRequestBuilderInternal instantiates a new CustomAuthenticationExtensionsItemValidateAuthenticationConfigurationRequestBuilder and sets the default values.
func NewCustomAuthenticationExtensionsItemValidateAuthenticationConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomAuthenticationExtensionsItemValidateAuthenticationConfigurationRequestBuilder) {
    m := &CustomAuthenticationExtensionsItemValidateAuthenticationConfigurationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/customAuthenticationExtensions/{customAuthenticationExtension%2Did}/validateAuthenticationConfiguration", pathParameters),
    }
    return m
}
// NewCustomAuthenticationExtensionsItemValidateAuthenticationConfigurationRequestBuilder instantiates a new CustomAuthenticationExtensionsItemValidateAuthenticationConfigurationRequestBuilder and sets the default values.
func NewCustomAuthenticationExtensionsItemValidateAuthenticationConfigurationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomAuthenticationExtensionsItemValidateAuthenticationConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCustomAuthenticationExtensionsItemValidateAuthenticationConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post an API to check validity of the endpoint and and authentication configuration for a customAuthenticationExtension object, which can represent one of the following derived types:
// returns a AuthenticationConfigurationValidationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/customauthenticationextension-validateauthenticationconfiguration?view=graph-rest-1.0
func (m *CustomAuthenticationExtensionsItemValidateAuthenticationConfigurationRequestBuilder) Post(ctx context.Context, requestConfiguration *CustomAuthenticationExtensionsItemValidateAuthenticationConfigurationRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationConfigurationValidationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationConfigurationValidationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationConfigurationValidationable), nil
}
// ToPostRequestInformation an API to check validity of the endpoint and and authentication configuration for a customAuthenticationExtension object, which can represent one of the following derived types:
// returns a *RequestInformation when successful
func (m *CustomAuthenticationExtensionsItemValidateAuthenticationConfigurationRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CustomAuthenticationExtensionsItemValidateAuthenticationConfigurationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CustomAuthenticationExtensionsItemValidateAuthenticationConfigurationRequestBuilder when successful
func (m *CustomAuthenticationExtensionsItemValidateAuthenticationConfigurationRequestBuilder) WithUrl(rawUrl string)(*CustomAuthenticationExtensionsItemValidateAuthenticationConfigurationRequestBuilder) {
    return NewCustomAuthenticationExtensionsItemValidateAuthenticationConfigurationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
