package identityprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// RiskyserviceprincipalsConfirmcompromisedConfirmCompromisedRequestBuilder provides operations to call the confirmCompromised method.
type RiskyserviceprincipalsConfirmcompromisedConfirmCompromisedRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RiskyserviceprincipalsConfirmcompromisedConfirmCompromisedRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RiskyserviceprincipalsConfirmcompromisedConfirmCompromisedRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRiskyserviceprincipalsConfirmcompromisedConfirmCompromisedRequestBuilderInternal instantiates a new RiskyserviceprincipalsConfirmcompromisedConfirmCompromisedRequestBuilder and sets the default values.
func NewRiskyserviceprincipalsConfirmcompromisedConfirmCompromisedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyserviceprincipalsConfirmcompromisedConfirmCompromisedRequestBuilder) {
    m := &RiskyserviceprincipalsConfirmcompromisedConfirmCompromisedRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityProtection/riskyServicePrincipals/confirmCompromised", pathParameters),
    }
    return m
}
// NewRiskyserviceprincipalsConfirmcompromisedConfirmCompromisedRequestBuilder instantiates a new RiskyserviceprincipalsConfirmcompromisedConfirmCompromisedRequestBuilder and sets the default values.
func NewRiskyserviceprincipalsConfirmcompromisedConfirmCompromisedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyserviceprincipalsConfirmcompromisedConfirmCompromisedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRiskyserviceprincipalsConfirmcompromisedConfirmCompromisedRequestBuilderInternal(urlParams, requestAdapter)
}
// Post confirm one or more riskyServicePrincipal objects as compromised. This action sets the targeted service principal account's risk level to high.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/riskyserviceprincipal-confirmcompromised?view=graph-rest-1.0
func (m *RiskyserviceprincipalsConfirmcompromisedConfirmCompromisedRequestBuilder) Post(ctx context.Context, body RiskyserviceprincipalsConfirmcompromisedConfirmCompromisedPostRequestBodyable, requestConfiguration *RiskyserviceprincipalsConfirmcompromisedConfirmCompromisedRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation confirm one or more riskyServicePrincipal objects as compromised. This action sets the targeted service principal account's risk level to high.
// returns a *RequestInformation when successful
func (m *RiskyserviceprincipalsConfirmcompromisedConfirmCompromisedRequestBuilder) ToPostRequestInformation(ctx context.Context, body RiskyserviceprincipalsConfirmcompromisedConfirmCompromisedPostRequestBodyable, requestConfiguration *RiskyserviceprincipalsConfirmcompromisedConfirmCompromisedRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RiskyserviceprincipalsConfirmcompromisedConfirmCompromisedRequestBuilder when successful
func (m *RiskyserviceprincipalsConfirmcompromisedConfirmCompromisedRequestBuilder) WithUrl(rawUrl string)(*RiskyserviceprincipalsConfirmcompromisedConfirmCompromisedRequestBuilder) {
    return NewRiskyserviceprincipalsConfirmcompromisedConfirmCompromisedRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
