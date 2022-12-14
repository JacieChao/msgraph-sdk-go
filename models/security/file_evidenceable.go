package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FileEvidenceable 
type FileEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDetectionStatus()(*DetectionStatus)
    GetFileDetails()(FileDetailsable)
    GetMdeDeviceId()(*string)
    SetDetectionStatus(value *DetectionStatus)()
    SetFileDetails(value FileDetailsable)()
    SetMdeDeviceId(value *string)()
}
