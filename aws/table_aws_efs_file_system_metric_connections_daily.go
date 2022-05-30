package aws

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
	"github.com/aws/aws-sdk-go/service/efs"
)

//// TABLE DEFINITION
func tableAwsEfsFileSystemMetricClientConnectionsDaily(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_efs_file_system_metric_client_connections_daily",
		Description: "AWS EFS File System Cloudwatch Metrics - Client Connections (Daily)",
		List: &plugin.ListConfig{
			ParentHydrate: listElasticFileSystem,
			Hydrate:       listEfsFileSystemMetricClientConnectionsDaily,
		},
		GetMatrixItem: BuildRegionList,
		Columns: awsRegionalColumns(cwMetricColumns(
			[]*plugin.Column{
				{
					Name:        "efs_file_system_id",
					Description: "The ID to identify the EFS file system.",
					Type:        proto.ColumnType_STRING,
					Transform:   transform.FromField("DimensionValue"),
				},
			})),
	}
}

func listEfsFileSystemMetricClientConnectionsDaily(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	fileSystem := h.Item.(*efs.FileSystemDescription)
	return listCWMetricStatistics(ctx, d, "DAILY", "AWS/EFS", "ClientConnections", "FileSystemId", *fileSystem.FileSystemId)
}
