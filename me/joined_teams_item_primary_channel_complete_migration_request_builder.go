package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// JoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilder provides operations to call the completeMigration method.
type JoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// JoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type JoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewJoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilderInternal instantiates a new CompleteMigrationRequestBuilder and sets the default values.
func NewJoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*JoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilder) {
    m := &JoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedTeams/{team%2Did}/primaryChannel/microsoft.graph.completeMigration";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewJoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilder instantiates a new CompleteMigrationRequestBuilder and sets the default values.
func NewJoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*JoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewJoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation complete the message migration process by removing `migration mode` from a channel in a team. `Migration mode` is a special state that prevents certain operations, like sending messages and adding members, during the data migration process. After a **completeMigration** request is made, you cannot import additional messages into the team. You can add members to the team after the request returns a successful response.
func (m *JoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *JoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post complete the message migration process by removing `migration mode` from a channel in a team. `Migration mode` is a special state that prevents certain operations, like sending messages and adding members, during the data migration process. After a **completeMigration** request is made, you cannot import additional messages into the team. You can add members to the team after the request returns a successful response.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/channel-completemigration?view=graph-rest-1.0
func (m *JoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilder) Post(ctx context.Context, requestConfiguration *JoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
