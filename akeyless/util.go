package akeyless

import (
	"context"
	"fmt"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func convertDisplayedDate(ctx context.Context, input *transform.TransformData) (interface{}, error) {
	layout := "2006-01-02T15Z"
	const logPrefix = "akeyless_utils.convertDisplayedDate"

	if displayedDate, ok := input.Value.(string); ok {
		if accDate, err := time.Parse(layout, displayedDate); err == nil {
			return accDate, nil
		} else {
			plugin.Logger(ctx).Warn(logPrefix, "failed to parse displayed date", displayedDate, err)
			return nil, nil
		}
	}

	plugin.Logger(ctx).Warn(logPrefix, "failed to cast input value of displayed date", input.Value)

	return nil, nil
}

func listTableTemplate[T any](ctx context.Context, qd *plugin.QueryData, _ *plugin.HydrateData, p apiClientProvider, fetcher apiDataFetcher[T]) (interface{}, error) {
	{
		pluginLogger := plugin.Logger(ctx)
		logPrefix := fmt.Sprintf("%v.list", qd.Table.Name)

		plugin.Logger(ctx).Trace(logPrefix, "start connecting to the api")

		client, token, err := p.GetApiService(ctx, qd)

		plugin.Logger(ctx).Trace(logPrefix, "connected")

		if err != nil {
			plugin.Logger(ctx).Error("faile")
			return nil, err
		}

		var paginationToken *string = nil

		pluginLogger.Trace(logPrefix, "start api list call")

		for {
			paginator, httpResp, err := fetcher.FetchData(ctx, client, token, paginationToken)

			pluginLogger.Trace(logPrefix, "api list call return status ", httpResp.StatusCode)

			if err != nil {
				pluginLogger.Error(logPrefix, "list_error", err)
				return nil, err
			} else if httpResp.StatusCode != 200 {
				return nil, fmt.Errorf("api list call return %v", httpResp.StatusCode)
			}

			pageData := paginator.GetPageItems()
			pluginLogger.Trace(logPrefix, "api list call return  %v rows", len(pageData))

			for _, item := range pageData {
				qd.StreamListItem(ctx, item)
			}

			paginationToken = paginator.NextPageToken()

			if paginationToken == nil {
				pluginLogger.Trace(logPrefix, "no next page, completed")
				break
			}

			pluginLogger.Trace(logPrefix, "next page exists, continue")
		}

		return nil, nil
	}
}
