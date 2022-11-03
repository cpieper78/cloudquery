// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
)

//go:generate mockgen -package=mocks -destination=../mocks/guardduty.go -source=guardduty.go GuarddutyClient
type GuarddutyClient interface {
	DescribeMalwareScans(context.Context, *guardduty.DescribeMalwareScansInput, ...func(*guardduty.Options)) (*guardduty.DescribeMalwareScansOutput, error)
	DescribeOrganizationConfiguration(context.Context, *guardduty.DescribeOrganizationConfigurationInput, ...func(*guardduty.Options)) (*guardduty.DescribeOrganizationConfigurationOutput, error)
	DescribePublishingDestination(context.Context, *guardduty.DescribePublishingDestinationInput, ...func(*guardduty.Options)) (*guardduty.DescribePublishingDestinationOutput, error)
	GetAdministratorAccount(context.Context, *guardduty.GetAdministratorAccountInput, ...func(*guardduty.Options)) (*guardduty.GetAdministratorAccountOutput, error)
	GetDetector(context.Context, *guardduty.GetDetectorInput, ...func(*guardduty.Options)) (*guardduty.GetDetectorOutput, error)
	GetFilter(context.Context, *guardduty.GetFilterInput, ...func(*guardduty.Options)) (*guardduty.GetFilterOutput, error)
	GetFindings(context.Context, *guardduty.GetFindingsInput, ...func(*guardduty.Options)) (*guardduty.GetFindingsOutput, error)
	GetFindingsStatistics(context.Context, *guardduty.GetFindingsStatisticsInput, ...func(*guardduty.Options)) (*guardduty.GetFindingsStatisticsOutput, error)
	GetIPSet(context.Context, *guardduty.GetIPSetInput, ...func(*guardduty.Options)) (*guardduty.GetIPSetOutput, error)
	GetInvitationsCount(context.Context, *guardduty.GetInvitationsCountInput, ...func(*guardduty.Options)) (*guardduty.GetInvitationsCountOutput, error)
	GetMalwareScanSettings(context.Context, *guardduty.GetMalwareScanSettingsInput, ...func(*guardduty.Options)) (*guardduty.GetMalwareScanSettingsOutput, error)
	GetMasterAccount(context.Context, *guardduty.GetMasterAccountInput, ...func(*guardduty.Options)) (*guardduty.GetMasterAccountOutput, error)
	GetMemberDetectors(context.Context, *guardduty.GetMemberDetectorsInput, ...func(*guardduty.Options)) (*guardduty.GetMemberDetectorsOutput, error)
	GetMembers(context.Context, *guardduty.GetMembersInput, ...func(*guardduty.Options)) (*guardduty.GetMembersOutput, error)
	GetRemainingFreeTrialDays(context.Context, *guardduty.GetRemainingFreeTrialDaysInput, ...func(*guardduty.Options)) (*guardduty.GetRemainingFreeTrialDaysOutput, error)
	GetThreatIntelSet(context.Context, *guardduty.GetThreatIntelSetInput, ...func(*guardduty.Options)) (*guardduty.GetThreatIntelSetOutput, error)
	GetUsageStatistics(context.Context, *guardduty.GetUsageStatisticsInput, ...func(*guardduty.Options)) (*guardduty.GetUsageStatisticsOutput, error)
	ListDetectors(context.Context, *guardduty.ListDetectorsInput, ...func(*guardduty.Options)) (*guardduty.ListDetectorsOutput, error)
	ListFilters(context.Context, *guardduty.ListFiltersInput, ...func(*guardduty.Options)) (*guardduty.ListFiltersOutput, error)
	ListFindings(context.Context, *guardduty.ListFindingsInput, ...func(*guardduty.Options)) (*guardduty.ListFindingsOutput, error)
	ListIPSets(context.Context, *guardduty.ListIPSetsInput, ...func(*guardduty.Options)) (*guardduty.ListIPSetsOutput, error)
	ListInvitations(context.Context, *guardduty.ListInvitationsInput, ...func(*guardduty.Options)) (*guardduty.ListInvitationsOutput, error)
	ListMembers(context.Context, *guardduty.ListMembersInput, ...func(*guardduty.Options)) (*guardduty.ListMembersOutput, error)
	ListOrganizationAdminAccounts(context.Context, *guardduty.ListOrganizationAdminAccountsInput, ...func(*guardduty.Options)) (*guardduty.ListOrganizationAdminAccountsOutput, error)
	ListPublishingDestinations(context.Context, *guardduty.ListPublishingDestinationsInput, ...func(*guardduty.Options)) (*guardduty.ListPublishingDestinationsOutput, error)
	ListTagsForResource(context.Context, *guardduty.ListTagsForResourceInput, ...func(*guardduty.Options)) (*guardduty.ListTagsForResourceOutput, error)
	ListThreatIntelSets(context.Context, *guardduty.ListThreatIntelSetsInput, ...func(*guardduty.Options)) (*guardduty.ListThreatIntelSetsOutput, error)
}