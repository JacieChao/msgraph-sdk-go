package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody provides operations to call the getMemberObjects method.
type ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The securityEnabledOnly property
    securityEnabledOnly *bool
}
// NewItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody instantiates a new ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody and sets the default values.
func NewItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody()(*ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody) {
    m := &ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["securityEnabledOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityEnabledOnly(val)
        }
        return nil
    }
    return res
}
// GetSecurityEnabledOnly gets the securityEnabledOnly property value. The securityEnabledOnly property
func (m *ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody) GetSecurityEnabledOnly()(*bool) {
    return m.securityEnabledOnly
}
// Serialize serializes information the current object
func (m *ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("securityEnabledOnly", m.GetSecurityEnabledOnly())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetSecurityEnabledOnly sets the securityEnabledOnly property value. The securityEnabledOnly property
func (m *ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody) SetSecurityEnabledOnly(value *bool)() {
    m.securityEnabledOnly = value
}