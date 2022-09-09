package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func MlTransforms() *schema.Table {
	return &schema.Table{
		Name:        "aws_glue_ml_transforms",
		Description: "A structure for a machine learning transform",
		Resolver:    fetchGlueMlTransforms,
		Multiplex:   client.ServiceAccountRegionMultiplexer("glue"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) of the workflow.",
				Type:            schema.TypeString,
				Resolver:        resolveGlueMlTransformArn,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "tags",
				Description: "Resource tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveGlueMlTransformTags,
			},
			{
				Name:        "created_on",
				Description: "A timestamp",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "description",
				Description: "A user-defined, long-form description text for the machine learning transform Descriptions are not guaranteed to be unique and can be changed at any time",
				Type:        schema.TypeString,
			},
			{
				Name:     "evaluation_metrics",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EvaluationMetrics"),
			},
			{
				Name:        "glue_version",
				Description: "This value determines which version of Glue this machine learning transform is compatible with",
				Type:        schema.TypeString,
			},
			{
				Name:        "label_count",
				Description: "A count identifier for the labeling files generated by Glue for this transform As you create a better transform, you can iteratively download, label, and upload the labeling file",
				Type:        schema.TypeInt,
			},
			{
				Name:        "last_modified_on",
				Description: "A timestamp",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "max_capacity",
				Description: "The number of Glue data processing units (DPUs) that are allocated to task runs for this transform",
				Type:        schema.TypeFloat,
			},
			{
				Name:        "max_retries",
				Description: "The maximum number of times to retry after an MLTaskRun of the machine learning transform fails",
				Type:        schema.TypeInt,
			},
			{
				Name:        "name",
				Description: "A user-defined name for the machine learning transform",
				Type:        schema.TypeString,
			},
			{
				Name:        "number_of_workers",
				Description: "The number of workers of a defined workerType that are allocated when a task of the transform runs",
				Type:        schema.TypeInt,
			},
			{
				Name:     "parameters",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Parameters"),
			},
			{
				Name:        "role",
				Description: "The name or Amazon Resource Name (ARN) of the IAM role with the required permissions",
				Type:        schema.TypeString,
			},
			{
				Name:        "schema",
				Description: "A map of key-value pairs representing the columns and data types that this transform can run against",
				Type:        schema.TypeJSON,
				Resolver:    resolveMlTransformsSchema,
			},
			{
				Name:        "status",
				Description: "The current status of the machine learning transform",
				Type:        schema.TypeString,
			},
			{
				Name:        "timeout",
				Description: "The timeout in minutes of the machine learning transform",
				Type:        schema.TypeInt,
			},
			{
				Name:     "transform_encryption",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TransformEncryption"),
			},
			{
				Name:        "transform_encryption_ml_user_data_encryption_kms_key_id",
				Description: "The ID for the customer-provided KMS key",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TransformEncryption.MlUserDataEncryption.KmsKeyId"),
			},
			{
				Name:        "id",
				Description: "The unique transform ID that is generated for the machine learning transform The ID is guaranteed to be unique and does not change",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TransformId"),
			},
			{
				Name:        "worker_type",
				Description: "The type of predefined worker that is allocated when a task of this transform runs",
				Type:        schema.TypeString,
			},
			{
				Name:        "input_record_tables",
				Description: "The database and table in the Glue Data Catalog that is used for input or output data",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("InputRecordTables"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_glue_ml_transform_task_runs",
				Description: "The sampling parameters that are associated with the machine learning transform",
				Resolver:    fetchGlueMlTransformTaskRuns,
				Columns: []schema.Column{
					{
						Name:        "ml_transform_cq_id",
						Description: "Unique CloudQuery ID of aws_glue_ml_transforms table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "completed_on",
						Description: "The last point in time that the requested task run was completed",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "error_string",
						Description: "The list of error strings associated with this task run",
						Type:        schema.TypeString,
					},
					{
						Name:        "execution_time",
						Description: "The amount of time (in seconds) that the task run consumed resources",
						Type:        schema.TypeInt,
					},
					{
						Name:        "last_modified_on",
						Description: "The last point in time that the requested task run was updated",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "log_group_name",
						Description: "The names of the log group for secure logging, associated with this task run",
						Type:        schema.TypeString,
					},
					{
						Name:        "export_labels_task_run_properties_output_s3_path",
						Description: "The Amazon Simple Storage Service (Amazon S3) path where you will export the labels",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.ExportLabelsTaskRunProperties.OutputS3Path"),
					},
					{
						Name:     "properties",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("Properties"),
					},
					{
						Name:        "started_on",
						Description: "The date and time that this task run started",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "status",
						Description: "The current status of the requested task run",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "The unique identifier for this task run",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TaskRunId"),
					},
					{
						Name:        "transform_id",
						Description: "The unique identifier for the transform",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchGlueMlTransforms(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.GetMLTransformsInput{}
	for {
		result, err := svc.GetMLTransforms(ctx, &input)
		if err != nil {
			return err
		}
		res <- result.Transforms
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}
func resolveGlueMlTransformArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	r := resource.Item.(types.MLTransform)
	arn := aws.String(mlTransformARN(cl, &r))
	return resource.Set(c.Name, arn)
}
func resolveGlueMlTransformTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	r := resource.Item.(types.MLTransform)
	result, err := svc.GetTags(ctx, &glue.GetTagsInput{
		ResourceArn: aws.String(mlTransformARN(cl, &r)),
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, result.Tags)
}
func resolveMlTransformsEvaluationMetricsFindMatchesMetricsColumnImportances(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.MLTransform)
	j := make(map[string]float64)
	if r.EvaluationMetrics == nil || r.EvaluationMetrics.FindMatchesMetrics == nil {
		return nil
	}
	for _, c := range r.EvaluationMetrics.FindMatchesMetrics.ColumnImportances {
		j[*c.ColumnName] = *c.Importance
	}
	return resource.Set(c.Name, j)
}
func resolveMlTransformsSchema(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.MLTransform)
	j := make(map[string]string)
	for _, c := range r.Schema {
		j[*c.Name] = *c.DataType
	}
	return resource.Set(c.Name, j)
}
func fetchGlueMlTransformTaskRuns(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.MLTransform)
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.GetMLTaskRunsInput{
		TransformId: r.TransformId,
	}
	for {
		result, err := svc.GetMLTaskRuns(ctx, &input)
		if err != nil {
			return err
		}
		res <- result.TaskRuns
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func mlTransformARN(cl *client.Client, tr *types.MLTransform) string {
	return cl.ARN(client.GlueService, "mlTransform", *tr.TransformId)
}
