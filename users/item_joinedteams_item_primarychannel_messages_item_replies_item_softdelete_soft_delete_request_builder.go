package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilder provides operations to call the softDelete method.
type ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilderInternal instantiates a new ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilder and sets the default values.
func NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilder) {
    m := &ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/primaryChannel/messages/{chatMessage%2Did}/replies/{chatMessage%2Did1}/softDelete", pathParameters),
    }
    return m
}
// NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilder instantiates a new ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilder and sets the default values.
func NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post delete a single chatMessage or a chat message reply in a channel or a chat.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chatmessage-softdelete?view=graph-rest-1.0
func (m *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation delete a single chatMessage or a chat message reply in a channel or a chat.
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilder when successful
func (m *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilder) WithUrl(rawUrl string)(*ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilder) {
    return NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemSoftdeleteSoftDeleteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
