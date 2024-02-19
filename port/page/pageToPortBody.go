package page

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/port-labs/terraform-provider-port-labs/internal/cli"
)

func PageToPortBody(pm *PageModel) (*cli.Page, error) {
	pb := &cli.Page{
		Identifier: pm.Identifier.ValueString(),
		Type:       pm.Type.ValueString(),
		Icon:       pm.Icon.ValueStringPointer(),
		Title:      pm.Title.ValueStringPointer(),
		Locked:     pm.Locked.ValueBoolPointer(),
		Blueprint:  pm.Blueprint.ValueStringPointer(),
		Parent:     pm.Parent.ValueStringPointer(),
		After:      pm.After.ValueStringPointer(),
	}

	widgets, err := widgetsToPortBody(pm.Widgets)
	if err != nil {
		return nil, err
	}
	pb.Widgets = widgets

	return pb, nil
}

func widgetsToPortBody(widgets []types.String) (*[]map[string]any, error) {
	if widgets == nil {
		return nil, nil
	}
	widgetsBody := make([]map[string]any, len(widgets))
	for i, w := range widgets {
		var widgetObject map[string]any
		if err := json.Unmarshal([]byte(w.ValueString()), &widgetObject); err != nil {
			return nil, err
		}
		widgetsBody[i] = widgetObject
	}

	return &widgetsBody, nil
}
