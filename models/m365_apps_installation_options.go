package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type M365AppsInstallationOptions struct {
    Entity
}
// NewM365AppsInstallationOptions instantiates a new M365AppsInstallationOptions and sets the default values.
func NewM365AppsInstallationOptions()(*M365AppsInstallationOptions) {
    m := &M365AppsInstallationOptions{
        Entity: *NewEntity(),
    }
    return m
}
// CreateM365AppsInstallationOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateM365AppsInstallationOptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewM365AppsInstallationOptions(), nil
}
// GetAppsForMac gets the appsForMac property value. The appsForMac property
// returns a AppsInstallationOptionsForMacable when successful
func (m *M365AppsInstallationOptions) GetAppsForMac()(AppsInstallationOptionsForMacable) {
    val, err := m.GetBackingStore().Get("appsForMac")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AppsInstallationOptionsForMacable)
    }
    return nil
}
// GetAppsForWindows gets the appsForWindows property value. The appsForWindows property
// returns a AppsInstallationOptionsForWindowsable when successful
func (m *M365AppsInstallationOptions) GetAppsForWindows()(AppsInstallationOptionsForWindowsable) {
    val, err := m.GetBackingStore().Get("appsForWindows")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AppsInstallationOptionsForWindowsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *M365AppsInstallationOptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appsForMac"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAppsInstallationOptionsForMacFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppsForMac(val.(AppsInstallationOptionsForMacable))
        }
        return nil
    }
    res["appsForWindows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAppsInstallationOptionsForWindowsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppsForWindows(val.(AppsInstallationOptionsForWindowsable))
        }
        return nil
    }
    res["updateChannel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppsUpdateChannelType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateChannel(val.(*AppsUpdateChannelType))
        }
        return nil
    }
    return res
}
// GetUpdateChannel gets the updateChannel property value. The updateChannel property
// returns a *AppsUpdateChannelType when successful
func (m *M365AppsInstallationOptions) GetUpdateChannel()(*AppsUpdateChannelType) {
    val, err := m.GetBackingStore().Get("updateChannel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AppsUpdateChannelType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *M365AppsInstallationOptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("appsForMac", m.GetAppsForMac())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("appsForWindows", m.GetAppsForWindows())
        if err != nil {
            return err
        }
    }
    if m.GetUpdateChannel() != nil {
        cast := (*m.GetUpdateChannel()).String()
        err = writer.WriteStringValue("updateChannel", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppsForMac sets the appsForMac property value. The appsForMac property
func (m *M365AppsInstallationOptions) SetAppsForMac(value AppsInstallationOptionsForMacable)() {
    err := m.GetBackingStore().Set("appsForMac", value)
    if err != nil {
        panic(err)
    }
}
// SetAppsForWindows sets the appsForWindows property value. The appsForWindows property
func (m *M365AppsInstallationOptions) SetAppsForWindows(value AppsInstallationOptionsForWindowsable)() {
    err := m.GetBackingStore().Set("appsForWindows", value)
    if err != nil {
        panic(err)
    }
}
// SetUpdateChannel sets the updateChannel property value. The updateChannel property
func (m *M365AppsInstallationOptions) SetUpdateChannel(value *AppsUpdateChannelType)() {
    err := m.GetBackingStore().Set("updateChannel", value)
    if err != nil {
        panic(err)
    }
}
type M365AppsInstallationOptionsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppsForMac()(AppsInstallationOptionsForMacable)
    GetAppsForWindows()(AppsInstallationOptionsForWindowsable)
    GetUpdateChannel()(*AppsUpdateChannelType)
    SetAppsForMac(value AppsInstallationOptionsForMacable)()
    SetAppsForWindows(value AppsInstallationOptionsForWindowsable)()
    SetUpdateChannel(value *AppsUpdateChannelType)()
}

type AppsInstallationOptionsForWindows struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAppsInstallationOptionsForWindows instantiates a new AppsInstallationOptionsForWindows and sets the default values.
func NewAppsInstallationOptionsForWindows()(*AppsInstallationOptionsForWindows) {
    m := &AppsInstallationOptionsForWindows{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAppsInstallationOptionsForWindowsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAppsInstallationOptionsForWindowsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppsInstallationOptionsForWindows(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AppsInstallationOptionsForWindows) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *AppsInstallationOptionsForWindows) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AppsInstallationOptionsForWindows) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isMicrosoft365AppsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMicrosoft365AppsEnabled(val)
        }
        return nil
    }
    res["isProjectEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsProjectEnabled(val)
        }
        return nil
    }
    res["isSkypeForBusinessEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSkypeForBusinessEnabled(val)
        }
        return nil
    }
    res["isVisioEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsVisioEnabled(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetIsMicrosoft365AppsEnabled gets the isMicrosoft365AppsEnabled property value. Specifies whether users can install Microsoft 365 apps, including Skype for Business, on their Windows devices. The default value is true.
// returns a *bool when successful
func (m *AppsInstallationOptionsForWindows) GetIsMicrosoft365AppsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isMicrosoft365AppsEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsProjectEnabled gets the isProjectEnabled property value. Specifies whether users can install Microsoft Project on their Windows devices. The default value is true.
// returns a *bool when successful
func (m *AppsInstallationOptionsForWindows) GetIsProjectEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isProjectEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsSkypeForBusinessEnabled gets the isSkypeForBusinessEnabled property value. Specifies whether users can install Skype for Business (standalone) on their Windows devices. The default value is true.
// returns a *bool when successful
func (m *AppsInstallationOptionsForWindows) GetIsSkypeForBusinessEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isSkypeForBusinessEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsVisioEnabled gets the isVisioEnabled property value. Specifies whether users can install Visio on their Windows devices. The default value is true.
// returns a *bool when successful
func (m *AppsInstallationOptionsForWindows) GetIsVisioEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isVisioEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AppsInstallationOptionsForWindows) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AppsInstallationOptionsForWindows) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isMicrosoft365AppsEnabled", m.GetIsMicrosoft365AppsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isProjectEnabled", m.GetIsProjectEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSkypeForBusinessEnabled", m.GetIsSkypeForBusinessEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isVisioEnabled", m.GetIsVisioEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppsInstallationOptionsForWindows) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AppsInstallationOptionsForWindows) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetIsMicrosoft365AppsEnabled sets the isMicrosoft365AppsEnabled property value. Specifies whether users can install Microsoft 365 apps, including Skype for Business, on their Windows devices. The default value is true.
func (m *AppsInstallationOptionsForWindows) SetIsMicrosoft365AppsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isMicrosoft365AppsEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsProjectEnabled sets the isProjectEnabled property value. Specifies whether users can install Microsoft Project on their Windows devices. The default value is true.
func (m *AppsInstallationOptionsForWindows) SetIsProjectEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isProjectEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSkypeForBusinessEnabled sets the isSkypeForBusinessEnabled property value. Specifies whether users can install Skype for Business (standalone) on their Windows devices. The default value is true.
func (m *AppsInstallationOptionsForWindows) SetIsSkypeForBusinessEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isSkypeForBusinessEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsVisioEnabled sets the isVisioEnabled property value. Specifies whether users can install Visio on their Windows devices. The default value is true.
func (m *AppsInstallationOptionsForWindows) SetIsVisioEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isVisioEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AppsInstallationOptionsForWindows) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
type AppsInstallationOptionsForWindowsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetIsMicrosoft365AppsEnabled()(*bool)
    GetIsProjectEnabled()(*bool)
    GetIsSkypeForBusinessEnabled()(*bool)
    GetIsVisioEnabled()(*bool)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetIsMicrosoft365AppsEnabled(value *bool)()
    SetIsProjectEnabled(value *bool)()
    SetIsSkypeForBusinessEnabled(value *bool)()
    SetIsVisioEnabled(value *bool)()
    SetOdataType(value *string)()
}
