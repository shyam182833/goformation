package cloudformation

// AWSEFSFileSystem_ElasticFileSystemTag AWS CloudFormation Resource (AWS::EFS::FileSystem.ElasticFileSystemTag)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-filesystemtags.html
type AWSEFSFileSystem_ElasticFileSystemTag struct {

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-filesystemtags.html#cfn-efs-filesystem-filesystemtags-key
	Key string `json:"Key,omitempty"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-filesystemtags.html#cfn-efs-filesystem-filesystemtags-value
	Value string `json:"Value,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEFSFileSystem_ElasticFileSystemTag) AWSCloudFormationType() string {
	return "AWS::EFS::FileSystem.ElasticFileSystemTag"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEFSFileSystem_ElasticFileSystemTag) AWSCloudFormationSpecificationVersion() string {
	return "1.5.0"
}
