package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnenoteNotebooksItemSectionsItemPagesItemCopyToSectionPostRequestBodyable 
type OnenoteNotebooksItemSectionsItemPagesItemCopyToSectionPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroupId()(*string)
    GetId()(*string)
    GetSiteCollectionId()(*string)
    GetSiteId()(*string)
    SetGroupId(value *string)()
    SetId(value *string)()
    SetSiteCollectionId(value *string)()
    SetSiteId(value *string)()
}
