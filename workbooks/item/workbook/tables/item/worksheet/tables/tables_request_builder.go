package tables

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i73b59b4c820c658765ec34a527c863a9e2283cfbc3cd45b85d501a07d24d4fae "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/tables/add"
    i8c5be1bf92534987886913c0a578608dfa0949167ca84b350759174eccb845a8 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/tables/itematwithindex"
    ib415e5f6df547c00658667b5f14ddfd2a45f2062fac744035a34420985392b32 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/tables/count"
)

// TablesRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\worksheet\tables
type TablesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TablesRequestBuilderGetOptions options for Get
type TablesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TablesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TablesRequestBuilderGetQueryParameters collection of tables that are part of the worksheet. Read-only.
type TablesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select_escaped []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// TablesRequestBuilderPostOptions options for Post
type TablesRequestBuilderPostOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookTable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *TablesRequestBuilder) Add()(*i73b59b4c820c658765ec34a527c863a9e2283cfbc3cd45b85d501a07d24d4fae.AddRequestBuilder) {
    return i73b59b4c820c658765ec34a527c863a9e2283cfbc3cd45b85d501a07d24d4fae.NewAddRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTablesRequestBuilderInternal instantiates a new TablesRequestBuilder and sets the default values.
func NewTablesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TablesRequestBuilder) {
    m := &TablesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/tables/{workbookTable_id}/worksheet/tables{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTablesRequestBuilder instantiates a new TablesRequestBuilder and sets the default values.
func NewTablesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TablesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTablesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\worksheet\tables\microsoft.graph.count()
func (m *TablesRequestBuilder) Count()(*ib415e5f6df547c00658667b5f14ddfd2a45f2062fac744035a34420985392b32.CountRequestBuilder) {
    return ib415e5f6df547c00658667b5f14ddfd2a45f2062fac744035a34420985392b32.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation collection of tables that are part of the worksheet. Read-only.
func (m *TablesRequestBuilder) CreateGetRequestInformation(options *TablesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePostRequestInformation collection of tables that are part of the worksheet. Read-only.
func (m *TablesRequestBuilder) CreatePostRequestInformation(options *TablesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get collection of tables that are part of the worksheet. Read-only.
func (m *TablesRequestBuilder) Get(options *TablesRequestBuilderGetOptions)(*TablesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTablesResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*TablesResponse), nil
}
// ItemAtWithIndex builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\worksheet\tables\microsoft.graph.itemAt(index={index})
func (m *TablesRequestBuilder) ItemAtWithIndex(index *int32)(*i8c5be1bf92534987886913c0a578608dfa0949167ca84b350759174eccb845a8.ItemAtWithIndexRequestBuilder) {
    return i8c5be1bf92534987886913c0a578608dfa0949167ca84b350759174eccb845a8.NewItemAtWithIndexRequestBuilderInternal(m.pathParameters, m.requestAdapter, index);
}
// Post collection of tables that are part of the worksheet. Read-only.
func (m *TablesRequestBuilder) Post(options *TablesRequestBuilderPostOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookTable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewWorkbookTable() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookTable), nil
}
