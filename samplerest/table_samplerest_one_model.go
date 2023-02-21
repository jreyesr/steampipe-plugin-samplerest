package samplerest

import (
	"context"
	"encoding/json"
	"time"

	// Any other dependencies, including packages that wrap calls to your target API, if any exist

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableSampleRESTOneModel() *plugin.Table {
	return &plugin.Table{
		Name:        "samplerest_one_model",
		Description: "Describe the model or data object that is represented by this table",
		List: &plugin.ListConfig{
			Hydrate: listOneModel,
			// Delete if your API doesn't suport searching over all instances
			KeyColumns: plugin.OptionalColumns([]string{"query_string", "query_json"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getOneModel,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromField("ID"), Description: "Primary identifier"},
			{Name: "column_1", Type: proto.ColumnType_STRING, Transform: transform.FromField("Column1"), Description: "A description for the column"},
			{Name: "column_2", Type: proto.ColumnType_INT, Transform: transform.FromField("Column2"), Description: "A description for the column"},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("CreatedAt"), Description: "A description of the column"},
			// Any other columns go here!
			// BEGIN Delete if your API doesn't suport searching over all instances
			{Name: "query_string", Type: proto.ColumnType_STRING, Transform: transform.FromField("QueryString"), Description: "Send a search string here"},
			{Name: "query_json", Type: proto.ColumnType_JSON, Transform: transform.FromField("QueryJson"), Description: "Send a search JSON here"},
			// END Delete if your API doesn't suport searching over all instances
		},
	}
}

type OneModel struct {
	ID          int
	Column1     string
	Column2     int
	CreatedAt   time.Time
	QueryString string
	QueryJson   map[string]any
}

func listOneModel(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Warn("listOneModel")

	// BEGIN Delete if your API doesn't suport searching over all instances
	realQueryString := d.EqualsQuals["query_string"]
	plugin.Logger(ctx).Warn("listOneModel", "realQueryString", realQueryString)

	queryJson, found := d.EqualsQuals["query_json"]
	realQueryJson := make(map[string]any)
	if found && json.Valid([]byte(queryJson.GetJsonbValue())) {
		json.Unmarshal([]byte(queryJson.GetJsonbValue()), &realQueryJson)
	}
	plugin.Logger(ctx).Warn("listOneModel", "realQueryJson", realQueryJson)
	// END Delete if your API doesn't suport searching over all instances

	config := d.Connection.Config.(SampleRESTConfig)
	plugin.Logger(ctx).Warn("listOneModel", "config", config)

	// 1. Initialize your technology plugin, if one exists. use the credentials in the config variable

	// 2. Iterate over all instances of the model by calling the appropriate API
	// It's probably something like GET https://service.com/api/onemodel, with no filters
	// If the API suports searching, use the data in realQueryString and/or realQueryJson to compose the filter

	// 3. Then, for each object, call d.StreamListItem(ctx, your_obj)
	// your_obj may be of any type, but it must have all the fields declared in Columns above
	// (see the Transform property for the correct field names). For example:
	// For example, the OneModel struct declared above can be returned

	// The List function should return nil. Data is returned via d.StreamListItem above
	return nil, nil
}

func getOneModel(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Warn("getOneModel")

	quals := d.EqualsQuals
	plugin.Logger(ctx).Warn("getOneModel", "quals", quals)
	id := quals["id"].GetInt64Value()
	plugin.Logger(ctx).Warn("getOneModel", "id", id)

	config := d.Connection.Config.(SampleRESTConfig)
	plugin.Logger(ctx).Warn("getOneModel", "config", config)

	// 1. Initialize your technology plugin, if one exists. use the credentials in config

	// 2. Get the instance of the model with the passed ID (stored in the id variable)
	// It's probably something like GET https://service.com/api/onemodel/<ID>

	// return result, nil
	// If an error happened, return nil, error
	// As above, result may be of any type, as long as it has the fields specified in Columns
	return nil, nil
}
